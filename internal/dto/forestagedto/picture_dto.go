package forestagedto

type PictureData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Alt    string `json:"alt"`
	Url    string `json:"url"`
	Weight int    `json:"weight"`
}

type PictureListData struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PictureUrl string `json:"pictureUrl"`
	Alt        string `json:"alt"`
	Url        string `json:"url"`
	Weight     int    `json:"weight"`
	Status     bool   `json:"status"`
}
