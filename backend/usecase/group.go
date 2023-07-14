package usecase

import (
	"database/sql"
	"encoding/json"
	"github.com/Masterminds/squirrel"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"vk-notification-monitor/entity"
	"vk-notification-monitor/entity/vkapi"
)

type GroupUsecase struct {
	db    *sql.DB
	vkapi vkapi.VKApi
}

func NewGroupUsecase(db *sql.DB, vkapi vkapi.VKApi) GroupUsecase {
	return GroupUsecase{db: db, vkapi: vkapi}
}

func (g *GroupUsecase) Add(url string) (*entity.Group, error) {

	group, err := g.GetInfoByAddressVKAPI(url)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO "groups" ("group_id", "group_address", "photo_url", "full_address", "name")
				VALUES ($1, $2, $3, $4, $5);`

	_, err = g.db.Exec(query, group.Id, url, group.PhotoUrl, group.FullUrl, group.Name)

	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *GroupUsecase) GetInfoByAddressVKAPI(address string) (*entity.Group, error) {
	client := http.Client{}
	resource := "/method/groups.getById"
	params := url.Values{}
	params.Add("access_token", g.vkapi.AccessToken)
	params.Add("group_id", address)
	params.Add("v", g.vkapi.V)

	req, err := http.NewRequest(http.MethodGet, g.vkapi.URL+resource, nil)
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

	respStruct := struct {
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
	}{}

	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return nil, err
	}

	group := entity.Group{
		Id:       strconv.Itoa(respStruct.Response[0].ID),
		FullUrl:  "https://vk.com/" + respStruct.Response[0].ScreenName,
		Name:     respStruct.Response[0].Name,
		PhotoUrl: respStruct.Response[0].Photo100,
		Address:  address,
	}
	return &group, nil
}

func (g *GroupUsecase) GetAllGroups() (*[]entity.Group, error) {
	rows, err := g.db.Query("SELECT group_id, group_address, photo_url, name, full_address FROM groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []entity.Group

	for rows.Next() {
		var group entity.Group
		err = rows.Scan(&group.Id, &group.Address, &group.PhotoUrl, &group.Name, &group.FullUrl)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	return &groups, nil
}

func (g *GroupUsecase) GetGroupAsAuthor(id int) (*entity.Author, error) {
	query := `SELECT group_id, name, photo_url, full_address FROM groups WHERE group_id = $1`

	id = int(math.Abs(float64(id)))

	a := entity.Author{}

	err := g.db.QueryRow(query, id).Scan(&a.Id, &a.Name, &a.Photo, &a.FullUrl)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (g *GroupUsecase) GetById(id string) (*entity.Group, error) {
	group := entity.Group{
		Id:       id,
		FullUrl:  "",
		Name:     "",
		PhotoUrl: "",
		Address:  "",
	}
	query := `SELECT full_address, name, photo_url, group_address FROM groups WHERE group_id = $1`

	strId, _ := strconv.Atoi(id)

	as := int(math.Abs(float64(strId)))

	err := g.db.QueryRow(query, as).Scan(&group.FullUrl, &group.Name, &group.PhotoUrl, &group.Address)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (g *GroupUsecase) DeleteByAddress(address string) error {
	sql, query, err := squirrel.Delete("groups").Where(squirrel.Eq{"full_address": address}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = g.db.Exec(sql, query...)
	if err != nil {
		return err
	}

	return nil
}
