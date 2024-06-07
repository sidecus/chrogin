package main

import (
	"os"
	"strconv"
)

const (
	ServerPortEnv                 = "Port"
	DefaultServerPort             = "8080"
	ChromiumTimeoutSecondsEnv     = "Timeout"
	DefaultChromiumTimeoutSeconds = "5000"
	EmbededReportUriFmtString     = "http://localhost:%d/reports/%s"
)

func getEnv(envName string, defaultValue string) string {
	val, ok := os.LookupEnv(envName)
	if !ok {
		val = defaultValue
	}

	return val
}

type Config struct {
	Port                   int
	ChromiumTimeoutSeconds int
}

func newConfig() Config {
	var config = Config{}
	var val string

	val = getEnv(ServerPortEnv, DefaultServerPort)
	port, err := strconv.Atoi(val)
	if err != nil || port < 1024 || port > 65535 {
		panic(ServerPortEnv + "is not valid")
	}
	config.Port = port

	val = getEnv(ChromiumTimeoutSecondsEnv, DefaultChromiumTimeoutSeconds)
	timeout, err := strconv.Atoi(val)
	if err != nil || timeout < 1000 || timeout > 10000 {
		panic(ChromiumTimeoutSecondsEnv + "is not valid")
	}
	config.ChromiumTimeoutSeconds = timeout

	return config
}
