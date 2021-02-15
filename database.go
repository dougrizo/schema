package schema

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

const (
	devModeEnv      = "DEVELOPMENT_MODE"
	uidEnv          = "SQL_USERNAME"
	pwdEnv          = "SQL_PASSWORD"
	hostEnv         = "SQL_HOSTNAME"
	portEnv         = "SQL_PORT"
	dbEnv           = "SQL_DATABASE"
	dsnFmt          = "host=%s user=%s password=%s dbname=%s port=%s"
	filenameEnv     = "SQLITE_FILENAME"
	defaultFilename = "file::memory:?cache=shared"
)

func envError(env string) string {
	return fmt.Sprintf("environment variable %s was empty", env)
}

func NewSQLite() (*gorm.DB, error) {
	filename := os.Getenv(filenameEnv)
	if filename == "" {
		filename = defaultFilename
	}

	return gorm.Open(sqlite.Open(filename), &gorm.Config{})
}

func NewPostgres(l hclog.Logger) (*gorm.DB, error) {
	var err error

	uid := os.Getenv(uidEnv)
	if uid == "" {
		err = errors.New(envError(uidEnv))
		l.Error(err.Error())
		return nil, err
	}

	pwd := os.Getenv(pwdEnv)
	if pwd == "" {
		err = errors.New(envError(pwdEnv))
		l.Error(err.Error())
		return nil, err
	}

	hostname := os.Getenv(hostEnv)
	if hostname == "" {
		err = errors.New(envError(hostEnv))
		l.Error(err.Error())
		return nil, err
	}

	port := os.Getenv(portEnv)
	if port == "" {
		err = errors.New(envError(portEnv))
		l.Error(err.Error())
		return nil, err
	}

	db := os.Getenv(dbEnv)
	if db == "" {
		err = errors.New(envError(dbEnv))
		l.Error(err.Error())
		return nil, err
	}

	dsn := fmt.Sprintf(dsnFmt, hostname, uid, pwd, db, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func New(l hclog.Logger) (*gorm.DB, error) {
	if os.Getenv(devModeEnv) == "true" {
		return NewSQLite()
	} else {
		return NewPostgres(l)
	}
}
