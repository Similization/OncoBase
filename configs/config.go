package configs

type ConfigDatabase struct {
	Port     string `yml:"port" env:"PORT" env-default:"5432"`
	Host     string `yml:"host" env:"HOST" env-default:"localhost"`
	Name     string `yml:"name" env:"NAME" env-default:"postgres"`
	User     string `yml:"user" env:"USER" env-default:"user"`
	Password string `yml:"password" env:"PASSWORD"`
}

type ConfigServer struct {
	Port string `yml:"port" env:"PORT" env-default:"8080"`
	Host string `yml:"host" env:"HOST" env-default:"localhost"`
}

type ConfigApp struct {
	Database ConfigDatabase
	Server   ConfigServer
}
