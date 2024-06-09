package logger

import (
	"backend-visualiser/cli-codegen/internal/config"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.SugaredLogger {
	// For some users, the presets offered by the NewProduction, NewDevelopment,
	// and NewExample constructors won't be appropriate. For most of those
	// users, the bundled Config struct offers the right balance of flexibility
	// and convenience. (For more complex needs, see the AdvancedConfiguration
	// example.)
	//
	// See the documentation for Config and zapcore.EncoderConfig for all the
	// available options.
	/* 	rawJSON := []byte(`{
	   	  "level": "debug",
	   	  "encoding": "json",
	   	  "outputPaths": ["stdout", "./tmp/logs"],
	   	  "errorOutputPaths": ["stderr"],
	   	  "encoderConfig": {
	   	    "levelEncoder": "lowercase"
	   	  }
	   	}`)
	   	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
	   		panic(err)
	   	} */
	fmt.Println(config.C)
	loglevel, err := zap.ParseAtomicLevel(config.C.Logger.Level)
	if err != nil {
		panic(err)
	}

	encodigConfig := zapcore.EncoderConfig{
		EncodeLevel: levelEncoderFromConfig(config.C.Logger.EncoderConfig.LevelEncoder),
	}
	cfg := zap.Config{
		Level: loglevel,
		/* Development:       false, */
		/* DisableCaller:     false, */
		/* DisableStacktrace: false, */
		/* Sampling:          &zap.SamplingConfig{}, */
		Encoding:         config.C.Logger.Encoding,
		EncoderConfig:    encodigConfig,
		OutputPaths:      config.C.Logger.OutputPaths,
		ErrorOutputPaths: config.C.Logger.ErrorOutputPaths,
		/*
			InitialFields:    map[string]interface{}{},
			"initialFields": {"foo": "bar"},
			just key:value that will be added to every log record
		*/
	}

	logger, err := cfg.Build()
	sugerlogger := logger.Sugar()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugerlogger.Info("logger construction succeeded")
	return sugerlogger
}
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
