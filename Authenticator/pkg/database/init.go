package database

import (
	"embed"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"

	"authenticator/pkg/log"
)

func InitDatabase(cfg Config, logger log.Logger, fs embed.FS) *sqlx.DB {
	dbConn, err := New(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	// UP migrations
	if cfg.AutoMigrate {
		n, err := dbAutoMigrate(cfg, dbConn, fs)
		if err != nil {
			logger.Fatalf("failed: sql migrations: %v", err)
		}

		logger.Infof("database migration complete: apply %d", n)
	}

	return dbConn
}

func dbAutoMigrate(cfg Config, db *sqlx.DB, fs embed.FS) (int, error) {
	migrate.SetTable("gorp_migrations")

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: fs,
		Root:       cfg.MigrationDirectory,
	}

	return migrate.Exec(db.DB, cfg.Dialect, migrations, migrate.Up)
}
