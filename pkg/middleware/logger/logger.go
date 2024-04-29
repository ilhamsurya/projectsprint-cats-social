package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = log.Output(output)
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		statusColor = params.StatusCodeColor()
		methodColor = params.MethodColor()
		resetColor = params.ResetColor()

		if params.Latency > time.Minute {
			params.Latency = params.Latency.Truncate(time.Second)
		}

		return fmt.Sprintf("%v%s %s %s %s %s %3d %s Latency:%5v [GIN]\n",
			params.TimeStamp.Format("2006-01-02T15:04:05+07:00"),
			methodColor, params.Method, resetColor,
			params.Path,
			statusColor, params.StatusCode, resetColor,
			params.Latency,
		)
	})
}

func GetLog() *os.File {
	logData := fmt.Sprintf("gin_" + time.Now().Format("20060102") + ".log")
	f, err := os.OpenFile(logData, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("error opening file: %v", err))
	}
	return f
}
