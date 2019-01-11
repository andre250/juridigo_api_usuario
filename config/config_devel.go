package config

import (
	"github.com/juridigo/juridigo_api_usuario/models"
)

/*
Devel - Responsável pode difinir confirgurações de ambiente de desenvolvimento
*/
func devel() {
	globaConfig = models.Config{
		App: models.App{
			Port:   "3030",
			Secret: "JUR1d1G00S3cr377",
		},
		Version: "0.0.1",
		Database: models.Database{
			Path:     "mongodb://<dbuser>:<dbpassword>@ds229701.mlab.com:29701/mica-develop",
			User:     "juridigo",
			Password: "jur1digo",
			Database: "mica-develop",
		},
		Amazon: models.Amazon{
			Prefix:  "juri_",
			Profile: "default",
			Bucket:  "juridigo",
			Region:  "us-east-1",
		},
		Facebook: models.Facebook{
			AppToken: "199329427332694|PluJMN9CqiFsZCVmgNIOA_Z6H08",
		},
	}
}
