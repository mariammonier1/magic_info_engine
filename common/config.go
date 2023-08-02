package common

import (
	"github.com/JeremyLoy/config"
	"log"
)

type EnvConfig struct {
	Debug                  string `config:"DEBUG"`
	ServerPort             int    `config:"SERVER_PORT"`
	RepositoryHost         string `config:"REPOSITORY_HOST"`
	RepositoryPort         int    `config:"REPOSITORY_PORT"`
	RepositoryUsername     string `config:"REPOSITORY_USERNAME"`
	RepositoryPassword     string `config:"REPOSITORY_PASSWORD"`
	RepositoryDatabaseName string `config:"REPOSITORY_DATABASE_NAME"`
	AuthHost               string `config:"AUTH_HOST"`
	AuthRealm              string `config:"AUTH_REALM"`
	AuthClientId           string `config:"AUTH_CLIENT_ID"`
	AuthClientSecret       string `config:"AUTH_CLIENT_SECRET"`
	MessagingHost          string `config:"MESSAGING_HOST"`
	MessagingUsername      string `config:"MESSAGING_Username"`
	MessagingPassword      string `config:"MESSAGING_Password"`
	NatsCreateProject      string `config:"NATS_CREATE_PROJECT"`
	NatsListProject        string `config:"NATS_List_PROJECT"`
}

var ec EnvConfig

func GetEnvConfig() *EnvConfig {

	err := config.FromEnv().To(&ec)
	if err != nil {
		log.Fatal(err)
	}
	return &ec

}
