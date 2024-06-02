package config

import "flag"

type Config struct {
	FilesDirectory string
}

func Parse() Config {
	directory := flag.String("directory", "", "the directory to serve files from")
	flag.Parse()
	return Config{
		FilesDirectory: *directory,
	}
}
