package main

import (
	"github.com/i1kondratiuk/kanban/application"
	"github.com/i1kondratiuk/kanban/interface/web/webjson/handler"
)

func main() {
	application.InitBoardManagerApp(&application.BoardManagerAppImpl{})
	application.InitColumnManagerApp(&application.ColumnManagerAppImpl{})
	application.InitCommentManagerApp(&application.CommentManagerAppImpl{})
	application.InitTaskManagerApp(&application.TaskManagerAppImpl{})

	handler.Run(8080)
}
