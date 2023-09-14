package config

const (
	ADMIN = "admin"
	GUEST = "guest"
	USER  = "user"
)

const (
	GetInit      = "/init"
	PostLogin    = "auth/login"
	PostRegister = "auth/register"
	GetUserList  = "/users"
	GetUser      = "/users/:id"
	PostUser     = "/users"
	PutUser      = "/users"
	DeleteUser   = "/users/:id"
	GetTodoList  = "/todos"
	GetTodo      = "/todos/:id"
	PostTodo     = "/todos"
	PutTodo      = "/todos"
	DeleteTodo   = "/todos/:id"
)
