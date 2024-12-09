package log

import (
	"fin_api_gateway/internal/monitoring"
	"fmt"
	"log/slog"
)

func Error(additionalMessage string, err error) {
	if err != nil {
		msg := fmt.Sprintf("%s: %s", additionalMessage, err.Error())
		monitoring.ProcessingErrorCount.WithLabelValues(msg).Inc()
		slog.Error(msg)
	}
}

func Info(message string) {
	monitoring.ProcessingSuccessCount.WithLabelValues(message).Inc()
	slog.Info(message)
}
