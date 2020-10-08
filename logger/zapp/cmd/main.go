package main

import (
	"fmt"

	"mugdha-adhav/go-for-go/logger/zapp/pkg/log"
)

func printDebugLog() error {
	log.Debug("Testing debug message")
	return nil
}

func printInfoLog() error {
	log.Info("Testing info message")
	return nil
}

func printErrorLog() error {
	log.Error("Testing error message")
	return nil
}

func printPanicLog() error {
	log.Panic("Testing panic message")
	return nil
}

func printFatalLog() error {
	log.Fatal("Testing fatal message")
	return nil
}

func printSugaredLog() error {
	log.Infof("Testing sugared message: %s. And then another content: %s", "Hey", "Hello")
	return nil
}

func main() {
	if err := log.New(); err != nil {
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
