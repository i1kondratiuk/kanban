package repository

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// CommentRepository represents a storage of all existing comments
type CommentRepository interface {
	GetOrderedByCreatedDateTimeBy(parentId common.Id) ([]*entity.Comment, error)
	Insert(newComment *entity.Comment) (*entity.Comment, error)
	Update(storedCommentId common.Id, newBodyText value.BodyText) error
	DeleteAllBy(parentId common.Id) error
	Delete(storedCommentId common.Id) error
}

var commentRepository CommentRepository

// GetCommentRepository returns the CommentRepository
func GetCommentRepository() CommentRepository {
	return commentRepository
}

// InitCommentRepository injects CommentRepository with its implementation
func InitCommentRepository(r CommentRepository) {
	commentRepository = r
}
