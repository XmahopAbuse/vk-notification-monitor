package usecase

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudflare/ahocorasick"
	"github.com/mymmrac/telego"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
	"vk-notification-monitor/entity"
	"vk-notification-monitor/usecase/vkapi"
)

type PostUsecase struct {
	db *sql.DB
}

func NewPostUsecase(db *sql.DB) PostUsecase {
	return PostUsecase{db: db}
}

func (p *PostUsecase) Add(post *entity.Post) error {
	sql, query, err := squirrel.Insert("posts").
		Columns("vk_id", "author", "text", "author_id", "group_id", "hash", "from_id", "date", "post_url").
		Values(
			post.Id,
			post.Author,
			post.Text,
			post.AuthorId,
			post.GroupId,
			post.Hash,
			post.FromId,
			post.Date,
			post.PostUrl,
		).PlaceholderFormat(squirrel.Dollar).
		//Suffix("ON CONFLICT (hash) DO NOTHING").
		ToSql()

	if err != nil {
		return err
	}
	_, err = p.db.Exec(sql, query...)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostUsecase) GetAll() (*[]entity.Post, error) {
	query, args, err := squirrel.Select("vk_id", "author", "text", "author_id", "group_id", "hash", "from_id", "post_url", "date").From("posts").OrderBy("date DESC").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []entity.Post

	for rows.Next() {
		var post entity.Post
		err := rows.Scan(
			&post.Id,
			&post.Author,
			&post.Text,
			&post.AuthorId,
			&post.GroupId,
			&post.Hash,
			&post.FromId,
			&post.PostUrl,
			&post.Date,
		)

		if err != nil {
			log.Println(err)
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

func (p *PostUsecase) ChangeState(post entity.Post) {

}

func (w *PostUsecase) GetPostsByKeywords(wall *entity.Wall, keywords []entity.Keyword) {
	for _, item := range wall.Response.Items {
		for _, keyword := range keywords {
			if strings.ToLower(item.Text) == strings.ToLower(string(keyword)) {
				fmt.Println(item.Text)
			}
		}
	}
}

func (w *PostUsecase) SyncPostsByKeywords(g *GroupUsecase, k *KeywordUsecase, wa *vkapi.WallUsecase, p *PostUsecase, n *NotificationUsecase, tgbot *telego.Bot, notify int) ([]entity.Post, error) {
	groups, err := g.GetAllGroups()
	keywords, _ := k.GetAll()
	ac := ahocorasick.NewStringMatcher(keywords)
	errCh := make(chan error, len(*groups))
	posts := []entity.Post{}
	if err != nil {
		return nil, err
	}
	notifyUsers, err := n.GetAll()
	if err != nil {
		log.Println(err)
	}
	wg := sync.WaitGroup{}
	for _, group := range *groups {
		wg.Add(1)

		go func(group entity.Group) {
			defer wg.Done()
			wall, err := wa.GetWallPostsByDomain(group.Address)
			if err != nil {
				errCh <- err
			}
			for _, item := range wall.Response.Items {
				text := strings.ToLower(item.Text)
				hits := ac.Match([]byte(text))
				if len(hits) > 0 {
					post := entity.Post{
						Id:       item.ID,
						Author:   strconv.Itoa(item.FromID),
						Text:     item.Text,
						AuthorId: strconv.Itoa(item.FromID),
						GroupId:  group.Id,
						Hash:     item.Hash,
						Status:   false,
						Date:     time.Unix(int64(item.Date), 0),
						FromId:   item.FromID,
						PostUrl:  "https://vk.com/" + group.Address + "?w=wall-" + group.Id + "_" + strconv.Itoa(item.ID),
					}

					postExist, _ := p.GetPost(map[string]string{"hash": item.Hash})
					if postExist == nil || len(*postExist) != 0 {
						continue
					} else {
						err = p.Add(&post)
						if err != nil {
							log.Println(err)
							continue
						}
						posts = append(posts, post)
						fmt.Println("NOTIFY", notify)
						if notify == 1 {
							// Отправка сообщения по источникам уведомлений
							fmt.Println("отправляем сообщения")
							fmt.Println(notifyUsers)
							for _, user := range notifyUsers {
								switch user.Type {
								case "telegram":
									chatId, _ := strconv.Atoi(user.Value)
									c := telego.ChatID{ID: int64(chatId)}
									text := fmt.Sprintf("Новое сообщение с упоминанием\n%s", post.PostUrl)
									msg := telego.SendMessageParams{ChatID: c, Text: text}
									_, err = tgbot.SendMessage(&msg)
									if err != nil {
										log.Println(err)
									}
								}
							}
						}
					}
				}
			}
		}(group)
	}
	wg.Wait()
	close(errCh)

	return posts, nil
}

func (p *PostUsecase) GetPost(q map[string]string) (*[]entity.Post, error) {

	queryBuilder := squirrel.Select("vk_id", "author", "text", "author_id", "group_id", "hash", "status", "from_id", "post_url", "date").From("posts")

	for k, v := range q {
		condition := squirrel.Eq{k: v}
		queryBuilder = queryBuilder.Where(condition)
	}

	sql, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []entity.Post{}

	for rows.Next() {
		post := entity.Post{}
		err = rows.Scan(&post.VKId, &post.Author, &post.Text, &post.AuthorId, &post.GroupId, &post.Hash, &post.Status, &post.FromId, &post.PostUrl, &post.Date)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

func (p *PostUsecase) Update(data map[string]interface{}, hash string) error {
	queryBuilder := squirrel.Update("posts").Where(squirrel.Eq{"hash": hash})

	for key, value := range data {
		queryBuilder = queryBuilder.Set(key, value)
	}

	sql, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = p.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
