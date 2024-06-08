package config

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
	hints

1. use `mapstructure` as if it is a struct tag
2. viper can map time not only to string but also to time.Duration
*/
type Config struct {
	Env         string `mapstructure:"env"`
	StoragePath string `mapstructure:"storage_path"`
	HTTPServer  struct {
		Addres      string        `mapstructure:"addres"`
		Timeout     time.Duration `mapstructure:"timeout"`
		IdleTimeout string        `mapstructure:"idle_timeout"`
	} `mapstructure:"http_server"`
}

var C Config

// go run *.go --flagname 444 // flagname=444
// go run *.go  //  flagname=1234
// __TODO: Logic for config initialiser
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
	viper.SetDefault("tmpdir", "./tmp")
	/*
		TODO:
		tmp dir should be created in workdir or filedir
		not in ../ from workdir
	*/
	viper.SetDefault("env", "release")
	viper.SetDefault("options_size", 8)
	viper.SetDefault("query_size", 8)
	viper.SetDefault("data_rows_size", 16)
	viper.SetDefault("methods", []string{"start", "stop", "insert"})
}

func setConfigEnvAndCommandLine() {
	/* === Environment variables === */
	viper.MustBindEnv("GOVERSION", "GOVERSION")
	viper.BindEnv("ZZGOSRC", "ZZGOSRC", "MYGOSRC", "ANYOTHERALIAS")
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
		panic("file must be set")
	}

}