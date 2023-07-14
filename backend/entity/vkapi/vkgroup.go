package vkapi

type VKGroup struct {
	Response []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		ScreenName   string `json:"screen_name"`
		IsClosed     int    `json:"is_closed"`
		Type         string `json:"type"`
		IsAdmin      int    `json:"is_admin"`
		IsMember     int    `json:"is_member"`
		IsAdvertiser int    `json:"is_advertiser"`
		Photo50      string `json:"photo_50"`
		Photo100     string `json:"photo_100"`
		Photo200     string `json:"photo_200"`
	} `json:"response"`
}
