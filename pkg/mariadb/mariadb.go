//go:generate go run -v github.com/kevinburke/go-bindata/go-bindata -pkg mariadb -o migration_bindata.go ../../internal/mariadb/migrations

package mariadb

import (
	"errors"
	"fmt"
	"path"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

const (
	migrationDir = "../../internal/mariadb/migrations"
)

var (
	dsn = getDSN("root", "root", "localhost", 3306, "foo")
)

func MigrateUp() (err error) {
	listDir, err := AssetDir(migrationDir)
	if err != nil {
		return
	}

	asset := bindata.Resource(listDir, func(filename string) ([]byte, error) {
		return Asset(path.Join(migrationDir, filename))
	})

	dbDriver, err := bindata.WithInstance(asset)
	if err != nil {
		return
	}

	mig, err := migrate.NewWithSourceInstance("go-bindata", dbDriver, dsn)
	if err != nil {
		return
	}

	if err = mig.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return
	}

	return nil
}

func MigrateDown() (err error) {
	listDir, err := AssetDir(migrationDir)
	if err != nil {
		return
	}

	asset := bindata.Resource(listDir, func(filename string) ([]byte, error) {
		return Asset(path.Join(migrationDir, filename))
	})

	dbDriver, err := bindata.WithInstance(asset)
	if err != nil {
		return
	}

	mig, err := migrate.NewWithSourceInstance("go-bindata", dbDriver, dsn)
	if err != nil {
		return
	}

	if err = mig.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return
	}

	return nil
}

func getDSN(username string, password string, host string, port int, dbName string) string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", username, password, host, strconv.Itoa(port), dbName)
}
