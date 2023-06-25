package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger
var logChan chan func()

func init() {
	// logger setup
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	// For file name
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	// For function name
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	Logger = zerolog.New(output).With().Timestamp().Logger()
	logChan = make(chan func())

	// for thread safe
	go func() {
		for log := range logChan {
			log()
		}
	}()
}

func pushLog(log func()) {
	logChan <- log
}

func Info(msg any, fileName string, funcName string) {
	message := fmt.Sprint(msg)
	log := func() {
		Logger.Info().Str(fileName, funcName).Msg(message)
	}
	pushLog(log)
}

func Error(msg any, fileName string, funcName string) {
	message := fmt.Sprint(msg)
	log := func() {
		Logger.Error().Str(fileName, funcName).Msg(message)
	}
	pushLog(log)
}

func Debug(msg any, fileName string, funcName string) {
	message := fmt.Sprint(msg)
	log := func() {
		Logger.Debug().Str(fileName, funcName).Msg(message)
	}
	pushLog(log)
}
