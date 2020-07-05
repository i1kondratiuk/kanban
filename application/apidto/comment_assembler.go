package apidto

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/entity"
)

func NewComments(ces []*entity.Comment) []*apimodel.Comment {
	var cms = make([]*apimodel.Comment, 0, len(ces))

	for _, ce := range ces {
		cms = append(cms, NewComment(ce))
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
