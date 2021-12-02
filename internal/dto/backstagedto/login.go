package backstagedto

type LoginDTO struct {
	LoginName string `validate:"min=4" json:"loginName"`
	Password  string `validate:"min=6" json:"password"`
}
