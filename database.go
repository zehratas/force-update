package main

import (
	"os"
)

type Database interface {
	Get(key string) string
}
type database struct {
}

func (d *database) Get(key string) string {
	return os.Getenv(key)
}
