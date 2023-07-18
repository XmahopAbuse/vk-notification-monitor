package usecase

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
	"strings"
	"vk-notification-monitor/entity"
)

type PostRepository interface {
	GetAll() (*[]entity.Post, error)
	GetPost(map[string]string) (*[]entity.Post, error)
	Add(post *entity.Post) error
}

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

func (w *PostUsecase) GetPostsByKeywords(wall *entity.Wall, keywords []entity.Keyword) {
	for _, item := range wall.Response.Items {
		for _, keyword := range keywords {
			if strings.ToLower(item.Text) == strings.ToLower(string(keyword)) {
				fmt.Println(item.Text)
			}
		}
	}
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
