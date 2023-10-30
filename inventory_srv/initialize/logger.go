package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = append(config.OutputPaths, "./log/inventory.log")
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	logger, err := config.Build()
	if err != nil {
		zap.S().Panic(err)
		panic("init logger failed")
	}
	zap.ReplaceGlobals(logger)
}