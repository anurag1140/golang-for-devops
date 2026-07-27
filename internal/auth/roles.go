package auth

const (
	RoleAdmin     = "admin"
	RoleLibrarian = "librarian"
	RoleMember    = "member"
)

// func RequireRole(roles ...string) func(http.Handler) http.Handler
