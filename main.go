package main

import (
	"log"
	"strconv"

	"github.com/spf13/viper"

	"github.com/i1kondratiuk/kanban/application"
	"github.com/i1kondratiuk/kanban/config"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/infrastructure/persistence"
	"github.com/i1kondratiuk/kanban/interface/web/webjson/handler"
)

func init() {
	// Loading environmental variables
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	db, err := persistence.NewDbConnection(
		configuration.Database.Host,
		strconv.Itoa(int(configuration.Database.Port)),
		configuration.Database.User,
		configuration.Database.Password,
		configuration.Database.Name,
		configuration.Database.Driver,
	)

	repository.InitBoardRepository(persistence.NewBoardRepository(db))

	if err != nil {
		panic(err.Error())
	}
}

func main() {
	application.InitBoardManagerApp(&application.BoardManagerAppImpl{})
	application.InitColumnManagerApp(&application.ColumnManagerAppImpl{})
	application.InitCommentManagerApp(&application.CommentManagerAppImpl{})
	application.InitTaskManagerApp(&application.TaskManagerAppImpl{})

	handler.Run(8080)
}
