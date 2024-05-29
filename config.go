package main

type Config struct {
	Port           int
	DownloadSuffix string
}

var config = Config{
	Port:           8080,
	DownloadSuffix: "download",
}
