package data

import (
	"database/sql"

	"github.com/spf13/viper"
)

func openDb() (*sql.DB, error){
	data, err := sql.Open("mysql", viper.GetString("data.url"))
	if err != nil {
		return nil,err
	}

	if err = data.Ping(); err != nil{
		return nil,err
	}
	return data, nil
}