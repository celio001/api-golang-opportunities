package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("fake erro")
}

func GetLogger(p string) *Logger {
	//Inicialize Logger
	logger = NewLogger(p)
	return logger
}
