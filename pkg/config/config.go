package config

import (
	"github.com/harpy-wings/fibonacci-kenshi/pkg/constants"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitConfig initializes the viper and loads the configuration.
func InitConfig(v *viper.Viper) error {
	var err error

	// ─── CONFIG THE PATH AND FILE NAME ──────────────────────────────────────────────
	v.SetConfigName(constants.ConfigFileName) // name of config file (without extension)
	v.AddConfigPath(constants.PathWorkingDirectory)
	v.AutomaticEnv()

	// ─── READ THE CONFIG ────────────────────────────────────────────────────────────
	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	// ─── BINDING FLAGS ──────────────────────────────────────────────────────────────
	pflag.Parse()
	err = v.BindPFlags(pflag.CommandLine)
	if err != nil {
		return err
	}

	return nil
}
