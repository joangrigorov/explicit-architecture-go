package http

type RouterGroup interface {
	// POST method handling on a relative path
	POST(relativePath string, handler Handler)

	// GET method handling on a relative path
	GET(relativePath string, handler Handler)

	// DELETE method handling on a relative path
	DELETE(relativePath string, handler Handler)

	// PATCH method handling on a relative path
	PATCH(relativePath string, handler Handler)

	// PUT method handling on a relative path
	PUT(relativePath string, handler Handler)

	// Group routes under a common relative path
	Group(relativePath string) RouterGroup
}

type Router interface {
	RouterGroup
}
