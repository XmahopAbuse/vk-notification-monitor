package entity

type wallPhoto struct {
	Type  string `json:"type"`
	Photo struct {
		AlbumID   int    `json:"album_id"`
		Date      int    `json:"date"`
		ID        int    `json:"id"`
		OwnerID   int    `json:"owner_id"`
		AccessKey string `json:"access_key"`
		PostID    int    `json:"post_id"`
		Sizes     []struct {
			Height int    `json:"height"`
			Type   string `json:"type"`
			Width  int    `json:"width"`
			URL    string `json:"url"`
		} `json:"sizes"`
		Text    string `json:"text"`
		UserID  int    `json:"user_id"`
		HasTags bool   `json:"has_tags"`
	} `json:"photo"`
}

type WallItem struct {
	CopyHistory []struct {
		Type        string      `json:"type"`
		Attachments []wallPhoto `json:"attachments"`
		Date        int         `json:"date"`
		FromID      int         `json:"from_id"`
		ID          int         `json:"id"`
		OwnerID     int         `json:"owner_id"`
		PostSource  struct {
			Type string `json:"type"`
		} `json:"post_source"`
		PostType string `json:"post_type"`
		Text     string `json:"text"`
	} `json:"copy_history,omitempty"`
	Donut struct {
		IsDonut bool `json:"is_donut"`
	} `json:"donut"`
	IsPinned int `json:"is_pinned,omitempty"`
	Comments struct {
		CanPost       int  `json:"can_post"`
		Count         int  `json:"count"`
		GroupsCanPost bool `json:"groups_can_post"`
	} `json:"comments"`
	MarkedAsAds   int           `json:"marked_as_ads"`
	ShortTextRate float64       `json:"short_text_rate"`
	Hash          string        `json:"hash"`
	Type          string        `json:"type"`
	Attachments   []interface{} `json:"attachments"`
	Date          int           `json:"date"`
	FromID        int           `json:"from_id"`
	ID            int           `json:"id"`
	IsFavorite    bool          `json:"is_favorite"`
	Likes         struct {
		CanLike        int  `json:"can_like"`
		Count          int  `json:"count"`
		UserLikes      int  `json:"user_likes"`
		CanPublish     int  `json:"can_publish"`
		RepostDisabled bool `json:"repost_disabled"`
	} `json:"likes"`
	OwnerID    int `json:"owner_id"`
	PostSource struct {
		Platform string `json:"platform"`
		Type     string `json:"type"`
	} `json:"post_source,omitempty"`
	PostType string `json:"post_type"`
	Reposts  struct {
		Count        int `json:"count"`
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Text  string `json:"text"`
	Views struct {
		Count int `json:"count"`
	} `json:"views,omitempty"`
	CarouselOffset int  `json:"carousel_offset,omitempty"`
	ZoomText       bool `json:"zoom_text,omitempty"`
	PostSource0    struct {
		Type string `json:"type"`
	} `json:"post_source,omitempty"`
	PostSource1 struct {
		Type string `json:"type"`
	} `json:"post_source,omitempty"`
}

type Wall struct {
	Response struct {
		Count int        `json:"count"`
		Items []WallItem `json:"items"`
	} `json:"response"`
}
