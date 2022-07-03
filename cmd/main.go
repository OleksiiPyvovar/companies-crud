package main

import (
	"context"
	"fmt"
	"os"

	"github.com/OleksiiPyvovar/companies-crud/api"
	"github.com/OleksiiPyvovar/companies-crud/pkg/app"
	"github.com/OleksiiPyvovar/companies-crud/pkg/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	conf := &api.Config{
		ServerAddr: getOrDefault("SERVER_ADDR", ":8080"),

		DBUser:    getOrDefault("DB_USER", "postgres"),
		DBPwd:     getOrDefault("DB_PASSWORD", "*"),
		DBPort:    getOrDefault("DB_PORT", "5432"),
		DBTCPHost: getOrDefault("DB_TCP_HOST", "35.238.184.102"),
		DBName:    getOrDefault("DB_NAME", "companies"),
	}

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", conf.DBTCPHost, conf.DBUser, conf.DBPwd, conf.DBPort, conf.DBName)
	conn, err := pgxpool.Connect(context.Background(), dbURI)
	if err != nil {
		fmt.Println("failed to connect db: ", err)
		return
	}

	service := app.NewCompaniesService(db.NewCompaniesPostgresRepository(conn))
	api := api.NewAPI(conf, service)

	api.Run()
}

func getOrDefault(key, defval string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defval
	}
	return val
}
