package main

import (
	"leonardoramosantos/tfc_plan_notifier/controller"
	"os"

	"github.com/op/go-logging"
)

func configLog() {
	var log_level = logging.INFO
	switch os.Getenv("LOG_LEVEL") {
	case "ERROR":
		log_level = logging.ERROR
	case "CRITICAL":
		log_level = logging.CRITICAL
	case "DEBUG":
		log_level = logging.DEBUG
	case "INFO":
		log_level = logging.INFO
	}
	logging.SetLevel(log_level, "")
}

// Main method of the application
func main() {
	configLog()

	var scanController = controller.GetController()
	scanController.StartPlans()
}
