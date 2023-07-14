package vkapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"vk-notification-monitor/entity"
	"vk-notification-monitor/entity/vkapi"
)

type WallUsecase struct {
	db    *sql.DB
	vkapi vkapi.VKApi
}

func NewWallUsecase(db *sql.DB, vkapi vkapi.VKApi) WallUsecase {
	return WallUsecase{db: db, vkapi: vkapi}
}

func (w *WallUsecase) GetWallPosts(ownerId string) (*entity.Wall, error) {
	client := http.Client{}
	resource := "/method/wall.get"
	params := url.Values{}
	params.Add("access_token", w.vkapi.AccessToken)
	params.Add("owner_id", "-"+ownerId)
	params.Add("v", w.vkapi.V)

	req, err := http.NewRequest(http.MethodGet, w.vkapi.URL+resource, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var wall entity.Wall

	json.Unmarshal(body, &wall)

	return &wall, err
}

func (w *WallUsecase) GetWallPostsByDomain(domain string) (*entity.Wall, error) {
	client := http.Client{}
	resource := "/method/wall.get"
	params := url.Values{}
	params.Add("access_token", w.vkapi.AccessToken)
	params.Add("domain", domain)
	params.Add("v", w.vkapi.V)

	req, err := http.NewRequest(http.MethodGet, w.vkapi.URL+resource, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var wall entity.Wall

	json.Unmarshal(body, &wall)

	return &wall, err
}

func (w *WallUsecase) GetProfilePhotoByUserId(userId int) (*vkapi.VKPhoto, error) {
	client := http.Client{}
	resource := "/method/photos.get"
	params := url.Values{}
	params.Add("access_token", w.vkapi.AccessToken)
	params.Add("owner_id", strconv.Itoa(userId))
	params.Add("album_id", "profile")
	params.Add("rev", "true")
	params.Add("v", w.vkapi.V)

	req, err := http.NewRequest(http.MethodGet, w.vkapi.URL+resource, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var photos vkapi.VKPhoto

	err = json.Unmarshal(body, &photos)
	if err != nil {
		return nil, err
	}

	return &photos, nil
}

func (w *WallUsecase) GetUserById(userid int) (*vkapi.VKUser, error) {
	client := http.Client{}
	resource := "/method/users.get"
	params := url.Values{}
	params.Add("access_token", w.vkapi.AccessToken)
	params.Add("user_ids", strconv.Itoa(userid))
	params.Add("fields", "photo_100")
	params.Add("v", w.vkapi.V)

	req, err := http.NewRequest(http.MethodGet, w.vkapi.URL+resource, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var user vkapi.VKUser

	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (w *WallUsecase) GetGroupById(id int) (*entity.Group, error) {
	client := http.Client{}
	resource := "/method/groups.getById"
	params := url.Values{}
	params.Add("access_token", w.vkapi.AccessToken)
	params.Add("group_id", strconv.Itoa(int(math.Abs(float64(id)))))
	params.Add("v", w.vkapi.V)

	req, err := http.NewRequest(http.MethodGet, w.vkapi.URL+resource, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	type Resp struct {
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

	var respStruct Resp

	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return nil, err
	}
	if len(respStruct.Response) <= 0 {
		fmt.Println(len(respStruct.Response))
		fmt.Println(respStruct.Response)
		return nil, fmt.Errorf("пустая структура, group id %v", strconv.Itoa(int(math.Abs(float64(id)))))
	}
	group := entity.Group{
		Id:       strconv.Itoa(respStruct.Response[0].ID),
		FullUrl:  "https://vk.com/" + respStruct.Response[0].ScreenName,
		Name:     respStruct.Response[0].Name,
		PhotoUrl: respStruct.Response[0].Photo100,
		Address:  "asd",
	}
	return &group, nil
}

func (w *WallUsecase) GetPostAuthor(authorId int) (*entity.Author, error) {
	var a entity.Author
	if authorId > 0 {
		usr, err := w.GetUserById(authorId)
		if err != nil || len(usr.Response) <= 0 {
			return nil, err
		}
		a.Id = strconv.Itoa(authorId)
		a.Name = usr.Response[0].FirstName + " " + usr.Response[0].LastName
		a.Photo = usr.Response[0].Photo100
		a.FullUrl = "https://vk.com/id" + a.Id
	}

	return &a, nil
}
