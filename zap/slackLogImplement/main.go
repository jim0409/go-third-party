package main

import (
	"go.uber.org/zap"
)

var (
	// Please rewrite it with your webhook URL
	slackWebHookURL = "https://hooks.slack.com/services/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"
)

func main() {
	logger, _ := zap.NewProduction()

	// Send a notification to slack at only error, fatal, panic level
	logger = logger.WithOptions(
		zap.Hooks(NewSlackHook(slackWebHookURL, zap.ErrorLevel).GetHook()),
	)

	logger.Debug("don't need to send a message")
	logger.Error("an error happened!")
}
