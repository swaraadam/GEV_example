package auth

type Tuple[A, B, C any] struct {
	Route       string
	RequireAuth bool
	Roles       []string
}

var Endpoints = map[string]Tuple[string, bool, []string]{
	"Register": {
		Route:       "/m_user.UserService/Register",
		RequireAuth: false,
	},
	"Login": {
		Route:       "/m_user.UserService/Login",
		RequireAuth: false,
	},
}
