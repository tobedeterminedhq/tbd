package lib

import "embed"

//go:embed init/*
var initContent embed.FS

// Init returns the content of the init directory.
func Init() embed.FS {
	return initContent
}
