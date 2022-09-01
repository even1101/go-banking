package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	logger.Info("Start banking application...")
	app.Start()
}
