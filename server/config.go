package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Constant Constant
	BookDB   *sql.DB
}

type Constant struct {
	PORT   string
	BookDB DBConstant
}

type DBConstant struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

const (
	defaultPort = "8088"
)

func initViper() (Constant, error) {
	viper.SetConfigName("book.config")
	viper.AddConfigPath("/app")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return Constant{}, errors.Wrap(err, "failed read config")
	}

	viper.SetDefault("PORT", defaultPort)

	var constant Constant
	if err := viper.Unmarshal(&constant); err != nil {
		return Constant{}, errors.Wrap(err, "failed unmarshal config file")
	}
	return constant, nil
}

func New() (Config, error) {
	constant, err := initViper()
	if err != nil {
		return Config{}, errors.Wrap(err, "faild create constant value")
	}

	bDB, err := sql.Open("mysql", createMysqlDSN(constant.BookDB))
	if err != nil {
		return Config{}, errors.Wrap(err, "failed create book db connection")
	}

	return Config{
		Constant: constant,
		BookDB: bDB,
	}, nil
}

func createMysqlDSN(c DBConstant) string {
	return fmt.Sprintf("%v:%v@%v/%v", c.User, c.Pass, c.Host, c.DBName)
}
