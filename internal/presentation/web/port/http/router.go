package http

type Router interface {
	// POST method handling on a relative path
	POST(string, ...Handler)

	// GET method handling on a relative path
	GET(string, ...Handler)

	// DELETE method handling on a relative path
	DELETE(string, ...Handler)

	// PATCH method handling on a relative path
	PATCH(string, ...Handler)

	// PUT method handling on a relative path
	PUT(string, ...Handler)

	// Group routes under a common relative path
	Group(string) Router

	// Use global handler
	Use(...Handler)
}

type Handler func(Context)
