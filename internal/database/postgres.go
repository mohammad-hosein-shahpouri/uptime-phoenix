package database

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mohammad-hosein-shahpouri/uptime-phoenix/internal/config"
	"github.com/mohammad-hosein-shahpouri/uptime-phoenix/internal/models"
)

type DB struct {
	db *gorm.DB
	mu sync.RWMutex
}

// New initializes the PostgreSQL database connection using GORM, automatically creates
// the target database if it does not exist, and runs table migrations.
func New(config config.PostgresConfig) (*DB, error) {
	sslmode := "disable"
	if config.SslEnabled {
		sslmode = "require"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		config.Host, config.Port, config.User, config.Password, config.Database, sslmode)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres database with gorm: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic sql.DB from gorm: %w", err)
	}

	// Recommended connection pooling for PostgreSQL in production
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	instance := &DB{db: gormDB}
	if err := instance.migrate(); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to run postgres database migrations: %w", err)
	}

	slog.Info("Connected to PostgreSQL successfully via GORM", "database", config.Database)
	return instance, nil
}

func (d *DB) migrate() error {
	return d.db.AutoMigrate(
		&models.Monitor{},
		&models.Heartbeat{},
		&models.NotificationChannel{},
		&models.MonitorChannel{},
		&models.StatusPage{},
		&models.StatusPageMonitor{},
		&models.Incident{},
		&models.Setting{},
	)
}
