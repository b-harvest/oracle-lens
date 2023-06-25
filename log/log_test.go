package log_test

import (
	"testing"
	"time"

	"bharvest.io/oracle-lens/log"
)

func TestInfo(t *testing.T) {
	log.Info("Info log test", "log_test.go", "TestInfo()")

	// For not terminate test
	time.Sleep(time.Second*1)
}

func TestError(t *testing.T) {
	log.Error("Error log test", "log_test.go", "ErrorInfo()")

	// For not terminate test
	time.Sleep(time.Second*1)
}

func TestDebug(t *testing.T) {
	log.Debug("Debug log test", "log_test.go", "DebugInfo()")

	// For not terminate test
	time.Sleep(time.Second*1)
}
