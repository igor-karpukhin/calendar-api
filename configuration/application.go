package configuration

import "flag"

type Application struct {
	Host  string
	Port  uint
	Mongo Mongo
}

func NewApplicationConfiguration() *Application {
	return &Application{
		Mongo: Mongo{},
	}
}

func (a *Application) ParseFlags() {
	flag.StringVar(&a.Host, "ip", "127.0.0.1", "HTTP bind IP")
	flag.UintVar(&a.Port, "port", 8080, "HTTP port to listen")
	flag.StringVar(&a.Mongo.Hosts, "mongo.hosts", "127.0.0.1,127.0.0.1", "list of mongoDB hosts")
	flag.StringVar(&a.Mongo.DbName, "mongo.dbname", "ki_test_db", "Mongo DB name")
	flag.StringVar(&a.Mongo.Username, "mongo.username", "ki_test_user", "MongoDB username")
	flag.StringVar(&a.Mongo.Password, "mongo.password", "ki_test_pass", "MongoDB password")

	flag.Parse()
}
