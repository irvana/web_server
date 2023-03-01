package logger

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(path string, name string) (*viper.Viper, error) {
	log.Info("initiating config")
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("sinarmas")
	err := viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Error("error loading config")
		return nil, err
	}
	log.Info("initiating custom log")
	initLog()
	return viper.GetViper(), nil
}

func initLog() {
	logFormat := viper.GetString("log.format")
	switch logFormat {
	case "text":
		log.SetFormatter(&log.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	case "json":
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	default:
		log.SetFormatter(&log.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	}

	if logLvl, err := log.ParseLevel(viper.GetString("log.level")); err != nil {
		log.SetLevel(log.DebugLevel)
		log.Warn("Cant parsing log level, using ", log.DebugLevel, " level")
	} else {
		log.SetLevel(logLvl)
	}
}
