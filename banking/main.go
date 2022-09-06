package main

import (
	"banking-lib/logger"
	"banking/app"
)

func main() {
	logger.Info("Start banking application...")
	app.Start()
}
