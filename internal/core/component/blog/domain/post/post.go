package post

type Post struct {
	Id      int
	Name    string
	Content string
}

func Map(id int, name string, content string) *Post {
	return &Post{
		Id:      id,
		Name:    name,
		Content: content,
	}
}
