package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
Hints:
1. use `mapstructure` as if it is a struct tag
*/
type Config struct {
	//USELESS
	Env string `mapstructure:"env"`
	//USELESS
	ValidateConfig bool `mapstructure:"validate_config"`
	Logger         struct {
		Level            string   `mapstructure:"level"`
		Encoding         string   `mapstructure:"encoding"`
		OutputPaths      []string `mapstructure:"output_paths"`
		ErrorOutputPaths []string `mapstructure:"error_output_paths"`
		EncoderConfig    struct {
			LevelEncoder string `mapstructure:"level_encoder"`
		} `mapstructure:"encoder_config"`
	} `mapstructure:"logger"`
}

var C Config

func InitConfig() {
	setConfigOptions()
	setConfigDefaults()
	setConfigEnvAndCommandLine()
	setConfigElse()
	pflag.Parse()
	bindFlags()
	readConfig()
	setConfigValidate()
}
func bindFlags() {
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}
func readConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}

}
func setConfigOptions() {
	/* === Config file === */
	viper.AddConfigPath("../config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
}
func setConfigDefaults() {
	/*  === Default values === */
	viper.SetDefault("workdir", "../")
	//	/tmpdir now used to store
	//		- .data files
	//		- log file
	viper.SetDefault("tmpdir", "./tmp")
	viper.SetDefault("options_size", 8)
	viper.SetDefault("query_size", 8)
	viper.SetDefault("data_rows_size", 16)
	viper.SetDefault("methods", []string{"start", "stop", "insert"})
	viper.SetDefault("validate_config", true)
}

func setConfigEnvAndCommandLine() {
	/* === Environment variables === */
	/* === Command line arguments === */
	pflag.String("file", "", "File to codegen visualise")

}
func setConfigElse() {
	/* === Watch config file changes === */
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

}
func setConfigValidate() {
	if viper.GetString("file") == "" {
		panic("Missing Input File")
	}
}
