package syncpost

import (
	"github.com/cloudflare/ahocorasick"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
	"vk-notification-monitor/entity"
	"vk-notification-monitor/usecase"
	"vk-notification-monitor/usecase/vkapi"
)

func SyncPostsByKeywords(groups []entity.Group, keywords []string, wa vkapi.WallRepository, p usecase.PostRepository, n *usecase.NotificationUsecase, notify int) ([]entity.Post, error) {
	ac := ahocorasick.NewStringMatcher(keywords)
	errCh := make(chan error, len(groups))
	posts := []entity.Post{}
	wg := sync.WaitGroup{}
	for _, group := range groups {
		wg.Add(1)

		go func(group entity.Group) {
			defer wg.Done()
			wall, err := wa.GetWallPostsByDomain(group.Address)
			if err != nil {
				errCh <- err
			}
			processedPosts := processWallPosts(wall, ac, group, p)
			posts = append(posts, processedPosts...)
		}(group)
	}
	wg.Wait()
	close(errCh)

	return posts, nil
}

func processWallPosts(wall *entity.Wall, matcher *ahocorasick.Matcher, group entity.Group, p usecase.PostRepository) []entity.Post {
	processedPosts := []entity.Post{}
	for _, item := range wall.Response.Items {
		text := strings.ToLower(item.Text)
		hits := matcher.Match([]byte(text))
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
			if postExist == nil {
				continue
			} else {
				err := p.Add(&post)
				if err != nil {
					log.Println(err)
					continue
				}
				processedPosts = append(processedPosts, post)
			}
		}
	}
	return processedPosts
}
