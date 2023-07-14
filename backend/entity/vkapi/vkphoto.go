package vkapi

type VKPhoto struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			AlbumID int `json:"album_id"`
			Date    int `json:"date"`
			ID      int `json:"id"`
			OwnerID int `json:"owner_id"`
			PostID  int `json:"post_id"`
			Sizes   []struct {
				Height int    `json:"height"`
				Type   string `json:"type"`
				Width  int    `json:"width"`
				URL    string `json:"url"`
			} `json:"sizes"`
			SquareCrop string `json:"square_crop"`
			Text       string `json:"text"`
			HasTags    bool   `json:"has_tags"`
		} `json:"items"`
	} `json:"response"`
}
