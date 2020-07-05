package persistence

import (
	"database/sql"
	"errors"
	"strconv"

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
	if c.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := c.db.Query(`
		SELECT
		       id,
		       body,
		       created_at
		FROM comments
		WHERE parent_id = $1
		ORDER BY created_at DESC`,
		parentId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("there is no record with id " + strconv.Itoa(int(parentId)))
		}

		return nil, err
	}

	defer rows.Close()

	comments := make([]*entity.Comment, 0)

	for rows.Next() {
		var (
			id        sql.NullInt64
			body      sql.NullString
			createdAt sql.NullTime
		)

		err = rows.Scan(
			&id,
			&body,
			&createdAt,
		)

		if err != nil {
			return nil, err
		}

		if id.Valid {
			comment := &entity.Comment{
				Id:       common.Id(id.Int64),
				ParentId: parentId,
			}

			if body.Valid {
				comment.Comment = value.Comment{BodyText: value.BodyText(body.String)}
			}

			if createdAt.Valid {
				comment.CreatedDateTime = createdAt.Time
			}

			comments = append(comments, comment)
		}
	}

	return comments, nil
}

func (c CommentRepositoryImpl) Insert(newComment *entity.Comment) (*entity.Comment, error) {
	var insertedCommentId int64

	if err := c.db.QueryRow(
		`INSERT INTO comments (body) VALUES ($1) RETURNING id`,
		string(newComment.Comment.BodyText),
	).Scan(&insertedCommentId); err != nil {
		return nil, err
	}

	newComment.Id = common.Id(insertedCommentId)

	return newComment, nil
}

func (c CommentRepositoryImpl) Update(storedCommentId common.Id, newBodyText value.BodyText) (err error) {
	_, err = c.db.Exec(
		`UPDATE comments SET body = $1 WHERE id = $2`,
		string(newBodyText),
		int(storedCommentId),
	)

	return
}

func (c CommentRepositoryImpl) DeleteAllBy(parentId common.Id) (err error) {
	_, err = c.db.Exec(`DELETE FROM comments WHERE parent_id = $1`, parentId)
	return
}

func (c CommentRepositoryImpl) Delete(storedCommentId common.Id) error {
	res, err := c.db.Exec(`DELETE FROM tasks WHERE id = $1`, storedCommentId)

	if err == nil {
		count, err := res.RowsAffected()
		if err != nil {
			return err
		} else if count != 1 {
			return errors.New("the record cannot be found, thus it is not deleted")
		}
	} else {
		return err
	}

	return nil
}
