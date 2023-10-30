package initialize

import "go.uber.org/zap"


func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		zap.ReplaceGlobals(logger)
	}
}