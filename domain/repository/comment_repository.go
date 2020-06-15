package repository

import (
	"github.com/i1kondratiuk/kanban/domain/value"
)

// CommentRepository represents a storage of all existing comments
type CommentRepository interface {
	GetAllCommentsBy(parentId value.Id) (*[]value.Comment, error)
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
