package backstagedto

type JwtInfoDTO struct {
	Id   int    `validate:"min=6" json:"id"`
	Name string `validate:"min=4" json:"name"`
}

type JwtTokenDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type JwtRefTokenDTO struct {
	RefreshToken string `json:"refreshToken"`
}
