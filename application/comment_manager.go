package application

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// BoardManagerApp represents BoardManagerApp application to be called by interface layer
type CommentManagerApp interface {
	GetAllParentCommentsGroupedByCreatedDateTime(parentId *common.Id) ([]*entity.Comment, error)
	Create(newComment *entity.Comment) (*entity.Comment, error)
	UpdateBodyText(storedCommentId common.Id, newBodyText value.BodyText) (*entity.Comment, error)
	Delete(storedCommentId common.Id) error
}

// CommentManagerAppImpl is the implementation of CommentManagerApp
type CommentManagerAppImpl struct{}

var commentManagerApp CommentManagerApp

// InitBoardManagerApp injects implementation for CommentManagerApp application
func InitCommentManagerApp(a CommentManagerApp) {
	commentManagerApp = a
}

// GetCommentManagerApp returns CommentManagerApp application
func GetCommentManagerApp() CommentManagerApp {
	return commentManagerApp
}

// CommentManagerAppImpl implements the CommentManagerApp interface
var _ CommentManagerApp = &CommentManagerAppImpl{}

func (c CommentManagerAppImpl) GetAllParentCommentsGroupedByCreatedDateTime(parentId *common.Id) ([]*entity.Comment, error) {
	storedComments, err := repository.GetCommentRepository().GetAllBy(*parentId)

	if err != nil {
		panic(err)
	}

	return storedComments, nil
}

func (c CommentManagerAppImpl) Create(newComment *entity.Comment) (*entity.Comment, error) {
	panic("implement me")
}

func (c CommentManagerAppImpl) UpdateBodyText(storedCommentId common.Id, newBodyText value.BodyText) (*entity.Comment, error) {
	panic("implement me")
}

func (c CommentManagerAppImpl) Delete(storedCommentId common.Id) error {
	panic("implement me")
}
