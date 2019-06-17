package dto

/*User in request*/
type UserFromRequest struct {
	Username string `json="username" bind="required"`
	Password string `json="password" bind="required"`
}

/*User in stored file*/
type UserStored struct {
	Username string `json="username"`
	Password []byte `json="password"`
}
