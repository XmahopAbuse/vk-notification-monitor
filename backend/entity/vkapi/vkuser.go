package vkapi

type VKUser struct {
	Response []struct {
		ID              int    `json:"id"`
		Photo100        string `json:"photo_100"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		CanAccessClosed bool   `json:"can_access_closed"`
		IsClosed        bool   `json:"is_closed"`
	} `json:"response"`
}
