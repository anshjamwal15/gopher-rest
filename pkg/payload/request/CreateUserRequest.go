package request

type CreateUserRequest struct {
	Username string
	Password string
	OrgId    int
	AdminId  int
}
