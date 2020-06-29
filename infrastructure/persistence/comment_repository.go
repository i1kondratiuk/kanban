package persistence

import (
	"database/sql"

	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// CommentRepositoryImpl is the implementation of CommentRepository
type CommentRepositoryImpl struct {
	db *sql.DB
}

// CCommentRepositoryImpl implements the domain CommentRepository interface
var _ repository.CommentRepository = &CommentRepositoryImpl{}

// CommentRepository returns initialized CommentRepositoryImpl
func NewCommentRepository(db *sql.DB) repository.CommentRepository {
	return &CommentRepositoryImpl{db: db}
}

func (c CommentRepositoryImpl) GetOrderedByCreatedDateTimeBy(parentId common.Id) ([]*entity.Comment, error) {
	panic("implement me")
}

func (c CommentRepositoryImpl) GetAllBy(parentId common.Id) ([]*entity.Comment, error) {
	panic("implement me")
}

func (c CommentRepositoryImpl) Insert(newComment *entity.Comment) (*entity.Comment, error) {
	panic("implement me")
}

func (c CommentRepositoryImpl) Update(storedCommentId common.Id, newBodyText value.BodyText) (*entity.Comment, error) {
	panic("implement me")
}

func (c CommentRepositoryImpl) DeleteBulk(storedCommentIds []common.Id) error {
	panic("implement me")
}

func (c CommentRepositoryImpl) Delete(storedCommentId common.Id) error {
	panic("implement me")
}
