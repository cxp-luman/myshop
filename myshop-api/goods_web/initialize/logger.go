package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = append(config.OutputPaths, "../log/goods.log")
	config.Level = zap.NewAtomicLevel()
	logger, err := config.Build()
	if err != nil {
		panic("init logger failed!")
	}
	zap.ReplaceGlobals(logger)
}
