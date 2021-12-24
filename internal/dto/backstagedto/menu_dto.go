package backstagedto

type MenuDTO struct {
	Menu []*MenuData `json:"menu"`
}

type MenuData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
	Parent  int    `json:"parent"`
	// Child   *[]MenuDTO `json:"child"`
}
