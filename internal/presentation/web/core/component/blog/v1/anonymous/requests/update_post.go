package requests

import (
	"app/internal/core/component/activity/domain"

	optional "github.com/aarondl/null/v9"
)

type UpdatePostRequest struct {
	Name    optional.String `json:"name" binding:"notnull"`
	Content optional.String `json:"content" binding:"notnull"`
}

func (r *UpdatePostRequest) Populate(p *domain.Activity) {
	if r.Name.IsSet() {
		p.Name = r.Name.String
	}
	if r.Content.IsSet() {
		p.Content = r.Content.String
	}
}
