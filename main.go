package main

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

func main() {
	var err = errors.New("Some error")
	err = errors.Wrap(err, "Wrapping")
	err = errors.Wrap(err, "Wrapping 2")

	log := zap.NewExample().Sugar()
	defer log.Sync()

	testExample(log)
	log.Info(err.Error())
	log.Debug("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second)
}

func testExample(log *zap.SugaredLogger) {
	log.Info("in func")
}
