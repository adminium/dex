package memory

import (
	"os"
	"testing"
	
	"github.com/sirupsen/logrus"
	
	"github.com/adminium/dex/storage"
	"github.com/adminium/dex/storage/conformance"
)

func TestStorage(t *testing.T) {
	logger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.TextFormatter{DisableColors: true},
		Level:     logrus.DebugLevel,
	}
	
	newStorage := func() storage.Storage {
		return New(logger)
	}
	conformance.RunTests(t, newStorage)
}
