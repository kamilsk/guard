package storage

import (
	"context"
	"database/sql"
	"errors"

	"go.octolab.org/ecosystem/guard/internal/config"
	"go.octolab.org/ecosystem/guard/internal/storage/internal"
)

// Must returns a new instance of the Storage or panics if it cannot configure it.
func Must(configs ...Configurator) *Storage {
	instance, err := New(configs...)
	if err != nil {
		panic(err)
	}
	return instance
}

// New returns a new instance of the Storage or an error if it cannot configure it.
func New(configs ...Configurator) (*Storage, error) {
	instance := &Storage{}
	for _, configure := range configs {
		if err := configure(instance); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Database returns Database Configurator.
func Database(cnf config.DatabaseConfig) Configurator {
	return func(instance *Storage) (err error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case error:
					err = v
				default:
					err = errors.New("unknown panic")
				}
			}
		}()
		instance.exec = internal.New(cnf.DriverName())
		instance.db, err = sql.Open(cnf.DriverName(), string(cnf.DSN))
		if err == nil {
			instance.db.SetMaxOpenConns(cnf.MaxOpenConns)
			instance.db.SetMaxIdleConns(cnf.MaxIdleConns)
			instance.db.SetConnMaxLifetime(cnf.ConnMaxLifetime)
		}
		return
	}
}

// Configurator defines a function which can use to configure the Storage.
type Configurator func(*Storage) error

// Storage is an implementation of Data Access Object.
type Storage struct {
	db   *sql.DB
	exec internal.Executor
}

// Database returns the current database handle.
func (storage *Storage) Database() *sql.DB {
	return storage.db
}

// Dialect returns a supported database dialect.
func (storage *Storage) Dialect() string {
	return storage.exec.Dialect()
}

func (storage *Storage) connection(ctx context.Context) (*sql.Conn, func() error, error) {
	conn, err := storage.db.Conn(ctx)
	if err != nil {
		return conn, nil, errors.New("trying to get connection")
	}
	return conn, conn.Close, nil
}
