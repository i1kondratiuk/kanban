package api

import (
	"github.com/i1kondratiuk/kanban/application/apidto"
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// BoardManagerApp represents BoardManagerApp application to be called by interface layer
type CommentManagerApp interface {
	GetAllParentCommentsGroupedByCreatedDateTime(parentId common.Id) ([]*apimodel.Comment, error)
	Create(newComment *entity.Comment) (*apimodel.Comment, error)                                    // TODO bulk create
	UpdateBodyText(storedCommentId common.Id, newBodyText value.BodyText) (*apimodel.Comment, error) // TODO bulk update
	Delete(storedCommentId common.Id) error                                                          // TODO bulk delete
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

func (c CommentManagerAppImpl) GetAllParentCommentsGroupedByCreatedDateTime(parentId common.Id) ([]*apimodel.Comment, error) {
	storedComments, err := repository.GetCommentRepository().GetOrderedByCreatedDateTimeBy(parentId)

	if err != nil {
		return nil, err
	}

	return apidto.NewComments(storedComments), nil
}

func (c CommentManagerAppImpl) Create(newComment *entity.Comment) (*apimodel.Comment, error) {
	insertedComment, err := repository.GetCommentRepository().Insert(newComment)

	if err != nil {
		return nil, err
	}

	return apidto.NewComment(insertedComment), nil
}

func (c CommentManagerAppImpl) UpdateBodyText(storedCommentId common.Id, newBodyText value.BodyText) (*apimodel.Comment, error) {
	updatedComment, err := repository.GetCommentRepository().Update(storedCommentId, newBodyText)

	if err != nil {
		return nil, err
	}

	return apidto.NewComment(updatedComment), nil
}

func (c CommentManagerAppImpl) Delete(storedCommentId common.Id) error {
	if err := repository.GetCommentRepository().Delete(storedCommentId); err != nil {
		return err
	}

	return nil
}
