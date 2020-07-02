package apidto

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/entity"
)

func NewComments(ces []*entity.Comment) []*apimodel.Comment {
	var cms = make([]*apimodel.Comment, len(ces))

	for i, ce := range ces {
		cms[i] = NewComment(ce)
	}

	return cms
}

func NewComment(ce *entity.Comment) *apimodel.Comment {
	return &apimodel.Comment{
		Id:              ce.Id,
		CreatedDateTime: ce.CreatedDateTime,
		Comment:         ce.Comment,
	}
}
