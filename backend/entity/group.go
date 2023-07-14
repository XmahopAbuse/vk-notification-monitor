package entity

type Group struct {
	Id       string `json:"id"`
	FullUrl  string `json:"full_url"`
	Name     string `json:"name"`
	PhotoUrl string `json:"photo_url"`
	Address  string `json:"address"`
}
