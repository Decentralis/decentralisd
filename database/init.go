package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type SqliteOptions string
type PostgresqlOptions struct {
	Username string
	Password string
	Ip       string
	Port     uint
}

func (p *PostgresqlOptions) Dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s port=%d sslmode=disable",
		p.Ip,
		p.Username,
		p.Password,
		p.Port)
}

type Options struct {
	Sqlite   *SqliteOptions
	Postgres *PostgresqlOptions
}

// InitDB creates (if not created) and connects to the database
func InitDB(options *Options) error {
	var err error

	if options.Sqlite != nil {
		if db, err = gorm.Open(sqlite.Open(string(*options.Sqlite))); err != nil {
			return err
		}
	} else if options.Postgres != nil {
		if db, err = gorm.Open(postgres.Open(options.Postgres.Dsn())); err != nil {
			return err
		}
	}

	db.AutoMigrate(&UserModel{}, &AccessTokenModel{}, &KeysModel{}, &MessageModel{})
	if err = GenerateKeys(); err != nil {
		return err
	}

	return nil
}
