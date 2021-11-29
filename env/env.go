package env

import (
	"fmt"
	"log"
	"strings"
	"syscall"
)

var panickingMessages []string
var warningMessages []string

func Require(envName string, description ...string) string {
	env, found := syscall.Getenv(envName)
	if !found {
		message := fmt.Sprintf("%s env is required.", envName)
		message = prependDescription(message, description)
		prePanic(message)
	}
	return env
}

func WarnIfEmpty(envName string, description ...string) string {
	env, found := syscall.Getenv(envName)
	if !found {
		message := fmt.Sprintf("%s env is empty, it may be needed.", envName)
		message = prependDescription(message, description)
		preWarn(message)
	}
	return env
}

func Default(envName string, defaultValue string) string {
	env, found := syscall.Getenv(envName)
	if !found {
		return defaultValue
	}
	return env
}

func Assert() {
	if len(warningMessages) > 0 {
		log.Println(strings.Join(warningMessages, "\n"))
	}

	if len(panickingMessages) > 0 {
		log.Panic(strings.Join(panickingMessages, "\n"))
	}

	resetState()
}

func prependDescription(message string, description []string) string {
	if len(description) > 0 {
		message = fmt.Sprintf("%s (%s)", message, description[0])
	}
	return message
}

func prePanic(message string) {
	panickingMessages = append(panickingMessages, message)
}

func preWarn(message string) {
	warningMessages = append(warningMessages, message)
}

func resetState() {
	panickingMessages = nil
	warningMessages = nil
}
