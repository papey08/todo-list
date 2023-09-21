package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"log"
	"time"
	"todo-list/internal/app"
	"todo-list/internal/ports/httpserver"
	"todo-list/internal/repo"
)

// InitConfig initializes configuration file
func InitConfig() error {
	viper.SetConfigFile("configs/config.yml")
	return viper.ReadInConfig()
}

// TaskRepoConfig initializes connection to database
func TaskRepoConfig(ctx context.Context, dbURL string) *pgx.Conn {
	// connecting to a database in the loop with delay 1 sec for correct starting in docker container
	for {
		conn, err := pgx.Connect(ctx, dbURL)
		if err != nil { // database haven't initialized in docker container yet
			log.Printf("taskRepo connection error: %s\n", err.Error())
			time.Sleep(time.Second)
		} else { // database already initialized
			return conn
		}
	}
}

func main() {
	ctx := context.Background()
	if err := InitConfig(); err != nil {
		log.Fatalf("configs error: %s", err.Error())
	}

	taskRepoURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("task_repo.username"),
		viper.GetString("task_repo.password"),
		viper.GetString("task_repo.host"),
		viper.GetInt("task_repo.port"),
		viper.GetString("task_repo.dbname"),
		viper.GetString("task_repo.sslmode"))
	taskRepoConn := TaskRepoConfig(ctx, taskRepoURL)
	defer func(ctx context.Context, conn *pgx.Conn) {
		if err := conn.Close(ctx); err != nil {
			log.Fatalf("taskRepo disconnect error: %s", err.Error())
		}
	}(ctx, taskRepoConn)

	a := app.New(repo.New(taskRepoConn))

	srv := httpserver.New(fmt.Sprintf("%s:%d", viper.GetString("http_server.host"), viper.GetInt("http_server.port")), a)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("unable to listen and serve http server: %s", err.Error())
	}
}
