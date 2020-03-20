package services

import (
	"context"
	"errors"

	"github.com/casbin/casbin/v2/persist"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v2"

	// import sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// supportDriverNames contains the supported driver name
	// Note here `[...]` is used instead of `[]` to get a (fixed size) array instead of a slice
	supportDriverNames = [...]string{"file", "mysql", "postgres", "sqlite3"}
)

var (
	// ErrDriverNotSupported is an error indicating that driver is not supported
	ErrDriverNotSupported = errors.New("driver not supported")
)

func (s *service) getAdapter(handle AdapterHandler) (persist.Adapter, error) {
	if _, ok := s.adapterMap[handle]; ok {
		return s.adapterMap[handle], nil
	}
	return nil, ErrAdapterNotFound
}

func (s *service) addAdapter(a persist.Adapter) AdapterHandler {
	cnt := AdapterHandler(len(s.adapterMap))
	s.adapterMap[cnt] = a
	return cnt
}

func newAdapter(config config) (persist.Adapter, error) {
	var a persist.Adapter
	switch config.Driver {
	case "file":
		// file adapter
		a = fileadapter.NewAdapter(config.Connection)
	default:
		// Database adapter

		// check if dirver is supported
		if !supportedDriver(config.Driver) {
			return nil, ErrDriverNotSupported
		}

		var err error
		a, err = gormadapter.NewAdapter(config.Driver, config.Connection, config.DbSpecified)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}

// supportedDriver checks whether the driver is supported
func supportedDriver(driver string) bool {
	for _, n := range supportDriverNames {
		if driver == n {
			return true
		}
	}
	return false
}

func (s *service) NewAdapter(ctx context.Context, driver, connection string, dbSpecified bool) (AdapterHandler, error) {
	// new an adapter
	a, err := newAdapter(config{
		Driver:      driver,
		Connection:  connection,
		DbSpecified: dbSpecified,
	})

	if err != nil {
		return -1, err
	}

	h := s.addAdapter(a)
	return h, nil
}

// config is a private struct, representing the configurations for adapter
type config struct {
	Driver      string
	Connection  string
	DbSpecified bool
}
