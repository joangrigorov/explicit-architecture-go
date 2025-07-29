package responses

import "app/internal/core/component/blog/domain/post"

type PostResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func OnePostResponse(p *post.Post) *PostResponse {
	return &PostResponse{
		Id:      p.Id,
		Name:    p.Name,
		Content: p.Content,
	}
}

func MultiPostResponse(posts []*post.Post) []*PostResponse {
	responses := make([]*PostResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, OnePostResponse(p))
	}
	return responses
}
