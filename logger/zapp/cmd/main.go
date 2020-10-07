package main

import (
	"fmt"

	"mugdha-adhav/go-for-go/logger/zapp/pkg/logger"
)

func printDebugLog() error {
	logger.Log.Debug("Testing debug message")
	return nil
}

func printInfoLog() error {
	logger.Log.Info("Testing info message")
	return nil
}

func printErrorLog() error {
	logger.Log.Error("Testing error message")
	return nil
}

func printPanicLog() error {
	logger.Log.Panic("Testing panic message")
	return nil
}

func printFatalLog() error {
	logger.Log.Fatal("Testing fatal message")
	return nil
}

func printSugaredLog() error {
	logger.Log.Infof("Testing sugared message: %s", "Here is some sample content")
	return nil
}

func main() {
	if err := logger.New(); err != nil {
		fmt.Printf("Logger initialization failed with error: %s", err.Error())
	}

	// Print log messages of all levels
	if err := printDebugLog(); err != nil {
		fmt.Printf("Debug level logging failed with error: %s", err.Error())
	}

	if err := printInfoLog(); err != nil {
		fmt.Printf("Info level logging failed with error: %s", err.Error())
	}

	if err := printSugaredLog(); err != nil {
		fmt.Printf("Sugared logging failed with error: %s", err.Error())
	}

	if err := printErrorLog(); err != nil {
		fmt.Printf("Error level logging failed with error: %s", err.Error())
	}

	// *******
	// TODO: Need to handle panic logging
	// *******

	// if err := printPanicLog(); err != nil {
	// 	fmt.Printf("Panic level logging failed with error: %s", err.Error())
	// }

	if err := printFatalLog(); err != nil {
		fmt.Printf("Fatal level logging failed with error: %s", err.Error())
	}
}
