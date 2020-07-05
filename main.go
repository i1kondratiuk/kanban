package main

import (
	"log"
	"strconv"

	"github.com/spf13/viper"

	"github.com/i1kondratiuk/kanban/application/api"
	"github.com/i1kondratiuk/kanban/config"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/service"
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
	repository.InitColumnRepository(persistence.NewColumnRepository(db))
	repository.InitTaskRepository(persistence.NewTaskRepository(db))
	repository.InitCommentRepository(persistence.NewCommentRepository(db))

	if err != nil {
		panic(err.Error())
	}
}

func main() {
	service.InitColumnService(&service.ColumnServiceImpl{})

	api.InitBoardManagerApp(&api.BoardManagerAppImpl{})
	api.InitColumnManagerApp(&api.ColumnManagerAppImpl{})
	api.InitCommentManagerApp(&api.CommentManagerAppImpl{})
	api.InitTaskManagerApp(&api.TaskManagerAppImpl{})

	handler.Run(8080)
}
