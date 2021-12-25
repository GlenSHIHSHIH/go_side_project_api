package backstagedto

type MenuData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
	Parent  int    `json:"parent"`
}

type MenuDTO struct {
	Menu []*MenuNestData `json:"menu"`
}

type MenuNestData struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Key     string          `json:"key"`
	Url     string          `json:"url"`
	Feature string          `json:"feature"`
	Parent  int             `json:"parent"`
	Child   []*MenuNestData `json:"child"`
}
