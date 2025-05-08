package main

type DBConfig struct {
}

type Config struct {
	Port     int      `json:"port" bson:"port"`
	AppName  string   `json:"appName" bson:"app_name"`
	DBConfig DBConfig `json:"dbConfig" bson:"db_config"`
}

type Dependencies struct {
	TodoRepo    ITodoRepo
	TodoService ITodoService
}

type App struct {
	Server Server
	Config Config
}
