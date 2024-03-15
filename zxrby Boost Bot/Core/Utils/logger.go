package Utils

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
)

func LogError(message string, key string, value string) {
	log.ErrorLevelStyle = lipgloss.NewStyle().SetString("ERROR [⚠ ] ").Foreground(lipgloss.Color("#FFA500")).Background(lipgloss.Color("#00000000"))
	Logger.Error(message, key, value)
}

func LogSuccess(message string, key string, value string) {
	log.InfoLevelStyle = lipgloss.NewStyle().SetString("SUCCESS [✔ ] ").Foreground(lipgloss.Color("#008000")).Background(lipgloss.Color("#00000000"))
	Logger.Info(message, key, value)
}

func LogInfo(message string, key string, value string) {
	log.WarnLevelStyle = lipgloss.NewStyle().SetString("INFO [❔ ] ").Foreground(lipgloss.Color("#0000FF")).Background(lipgloss.Color("#00000000"))
	Logger.Warn(message, key, value)
}

func LogPanic(message string, key string, value string) {
	log.ErrorLevelStyle = lipgloss.NewStyle().SetString("PANIC [☠ ] ").Foreground(lipgloss.Color("#800000")).Background(lipgloss.Color("#00000000"))
	Logger.Fatal(message, key, value)
}

