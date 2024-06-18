package logger

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// I have better template now
// TODO: Rewtite it
func InitLogger() *zap.SugaredLogger {

	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	ws, err := os.OpenFile(viper.GetString("tmpdir")+"/logs", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(enc, ws, zapcore.InfoLevel)

	logger := zap.New(core)
	sugerlogger := logger.Sugar()

	defer sugerlogger.Sync()

	sugerlogger.Info("logger construction succeeded")

	return sugerlogger
}

// UNUSED
func levelEncoderFromConfig(levelEncoder string) zapcore.LevelEncoder {
	switch levelEncoder {
	case "lowercase":
		return zapcore.LowercaseLevelEncoder
	case "uppercase", "capital":
		return zapcore.CapitalLevelEncoder
	default:
		panic("unknown levelEncoder: " + levelEncoder)
	}
}
