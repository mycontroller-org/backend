package logger

import (
	cfg "github.com/mycontroller-org/backend/v2/pkg/service/configuration"
	"github.com/mycontroller-org/backend/v2/pkg/utils"
	"github.com/mycontroller-org/backend/v2/pkg/version"
	"go.uber.org/zap"
)

// Init logger
func Init() {
	logger := utils.GetLogger(cfg.CFG.Logger.Mode, cfg.CFG.Logger.Level.Core, cfg.CFG.Logger.Encoding, false, 0)
	zap.ReplaceGlobals(logger)
	zap.L().Info("Welcome to MyController.org server :)")
	zap.L().Info("Server detail", zap.Any("version", version.Get()), zap.Any("loggerConfig", cfg.CFG.Logger))
}
