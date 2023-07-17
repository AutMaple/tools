package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

var (
	logger zerolog.Logger
)

func init() {
	zerolog.ErrorFieldName = "message"
	zerolog.ErrorMarshalFunc = func(err error) interface{} {
		return fmt.Sprintf("%+v", err)
	}
	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	}
	cw.FormatLevel = func(data interface{}) string {
		level := strings.ToUpper(data.(string))
		var cl string
		switch level {
		case "TRACE", "DEBUG":
			cl = color.WhiteString(level)
		case "INFO":
			cl = color.GreenString(level)
		case "WARN":
			cl = color.YellowString(level)
		case "ERROR", "FATAL", "PANIC":
			cl = color.RedString(level)
		}
		return cl
	}
	logger = zerolog.New(cw).With().
		Timestamp().
		Caller().
		CallerWithSkipFrameCount(3).
		Logger()
}

func Trace(msg string) {
	logger.Trace().Msg(msg)
}

func Debug(msg string) {
	logger.Debug().Msg(msg)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}
func Warn(msg string) {
	logger.Warn().Msg(msg)
}

func Error(err error) {
	logger.Error().Err(err).Send()
}

func ErrorMsg(msg string, err error) {
	logger.Error().Msg(msg)
	logger.Error().Err(err).Send()
}

func Fatal(err error) {
	logger.Fatal().Err(err).Send()
}

func FatalMsg(msg string, err error) {
	logger.Fatal().Msg(msg)
	logger.Fatal().Err(err).Send()
}

func Panic(err error) {
	logger.Panic().Err(err).Send()
}

func PanicMsg(msg string, err error) {
	logger.Panic().Msg(msg)
	logger.Panic().Err(err).Send()
}
