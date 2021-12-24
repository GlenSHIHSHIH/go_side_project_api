package backstagedto

type JwtInfoDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type JwtTokenDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type JwtRefTokenDTO struct {
	RefreshToken string `json:"refreshToken"`
}
