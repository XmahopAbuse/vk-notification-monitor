package syncpost

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"vk-notification-monitor/entity"
)

type MockWallRepository struct{}

type MockPostRepository struct{}

func (m MockPostRepository) GetAll() (*[]entity.Post, error) {
	panic("implement me")
}

func (m MockPostRepository) GetPost(m2 map[string]string) (*[]entity.Post, error) {
	post := entity.Post{
		Id:       0,
		VKId:     "",
		Author:   "",
		Text:     "тест",
		AuthorId: "",
		GroupId:  "",
		Hash:     "",
		Status:   false,
		Date:     time.Time{},
		FromId:   0,
		PostUrl:  "",
	}

	posts := []entity.Post{post}

	return &posts, nil
}

func (m MockPostRepository) Add(post *entity.Post) error {
	return nil
}

func (w *MockWallRepository) GetWallPostsByDomain(domain string) (*entity.Wall, error) {
	item := entity.WallItem{
		Text: "тестовая строка",
	}

	wall := entity.Wall{
		Response: struct {
			Count int               `json:"count"`
			Items []entity.WallItem `json:"items"`
		}(struct {
			Count int
			Items []entity.WallItem
		}{
			Count: 1,
			Items: []entity.WallItem{item},
		}),
	}

	return &wall, nil
}

func TestSyncPostsByKeywords(t *testing.T) {
	group := entity.Group{
		Id:       "",
		FullUrl:  "",
		Name:     "",
		PhotoUrl: "",
		Address:  "",
	}

	groups := []entity.Group{group}

	keywords := []string{"тест"}

	wa := MockWallRepository{}
	p := MockPostRepository{}

	posts, err := SyncPostsByKeywords(groups, keywords, &wa, p, nil, 0)

	// Проверка результата
	require.NotEmpty(t, posts, "Expected non-empty posts array")
	assert.NoError(t, err, "Unexpected error")

	// Вывод результата с подсветкой цветом
	t.Log("\033[32mТест успешно пройден!\033[0m")
}
