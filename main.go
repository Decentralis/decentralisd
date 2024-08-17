package main

import (
	"decentralisd/database"
	"decentralisd/endpoints"
	"flag"
	"log"
)

func main() {
	// Parse flags
	// server flags
	var (
		port uint
		ip   string
	)
	flag.UintVar(&port, "P", 42069, "port to listen on")
	flag.StringVar(&ip, "A", "0.0.0.0", "ip to listen on")
	// database settings
	var (
		db         string
		dbFile     string
		dbIp       string
		dbPort     uint
		dbUser     string
		dbPassword string
	)
	flag.StringVar(&db, "db", "", "Database server (possible: sqlite, postgresql)")
	flag.StringVar(&dbFile, "dfile", "decentralisd.db", "path to database file")
	flag.StringVar(&dbIp, "dip", "0.0.0.0", "database server ip")
	flag.UintVar(&dbPort, "dport", 5432, "database server port")
	flag.StringVar(&dbUser, "duser", "postgres", "database user")
	flag.StringVar(&dbPassword, "dpassword", "postgres", "database password")
	flag.Parse()

	// Select database
	var databaseOptions = &database.Options{}
	switch db {
	case "sqlite":
		databaseOptions.Sqlite = (*database.SqliteOptions)(&dbFile)
	case "postgresql":
		if dbPort < 0 || dbPort > 65535 {
			log.Fatalln("invalid database port number")
		}

		databaseOptions.Postgres = &database.PostgresqlOptions{
			Username: dbUser,
			Password: dbPassword,
			Ip:       dbIp,
			Port:     dbPort,
		}
	default:
		log.Fatalln("invalid database server")
	}

	// Check if port is valid
	if port < 0 || port > 65535 {
		log.Fatalln("port must be between 0 and 65535")
	}

	// Setup database
	if err := database.InitDB(databaseOptions); err != nil {
		panic(err)
	}

	// Run server
	endpoints.RunServer(ip, port)
}
