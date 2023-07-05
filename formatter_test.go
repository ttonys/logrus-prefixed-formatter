package prefixed

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

var Logger *log.Logger

func TestFormatter(t *testing.T) {
	Logger = log.New()

	var stdFormatter *TextFormatter
	// todo: 配置读取
	Logger.SetLevel(log.InfoLevel)

	stdFormatter = &TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		ForceFormatting:  true,
		ForceColors:      true,
		DisableColors:    false,
		QuoteEmptyFields: true,
		SpacePadding:     60,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			funcName := strings.Split(path.Base(frame.Function), ".")[len(strings.Split(path.Base(frame.Function), "."))-1]
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			return funcName, fileName
		},
	}

	Logger.SetFormatter(stdFormatter)
	Logger.SetOutput(os.Stdout)
	Logger.SetReportCaller(true)
	Logger.Infof("test")
}
