package configuration

import "flag"

type Application struct {
	Host     string
	Port     uint
	Postgres Postgres
}

func NewApplicationConfiguration() *Application {
	return &Application{
		Postgres: Postgres{},
	}
}

func (a *Application) ParseFlags() {
	flag.StringVar(&a.Host, "ip", "127.0.0.1", "HTTP bind IP")
	flag.UintVar(&a.Port, "port", 8080, "HTTP port to listen")
	flag.StringVar(&a.Postgres.Address, "pg.addr", "127.0.0.1:5432", "Postgres address. <IP>:<PORT>")
	flag.StringVar(&a.Postgres.DBName, "pg.db", "ki_test", "Postgres DB name")
	flag.StringVar(&a.Postgres.Username, "pg.username", "ki_test_user", "MongoDB username")
	flag.StringVar(&a.Postgres.Password, "pg.password", "ki_test_pass", "MongoDB password")

	flag.Parse()
}
