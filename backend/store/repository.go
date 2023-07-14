package store

import (
	"vk-notification-monitor/entity/vkapi"
	"vk-notification-monitor/usecase"
	vkapi2 "vk-notification-monitor/usecase/vkapi"
)

type Repository struct {
	Wall         vkapi2.WallUsecase
	Group        usecase.GroupUsecase
	Keyword      usecase.KeywordUsecase
	Post         usecase.PostUsecase
	Notification usecase.NotificationUsecase
}

func NewRepository(st *store, vkapi vkapi.VKApi) Repository {
	Wall := vkapi2.NewWallUsecase(st.db, vkapi)
	Group := usecase.NewGroupUsecase(st.db, vkapi)
	Keyword := usecase.NewKeywordUsecase(st.db)
	Post := usecase.NewPostUsecase(st.db)
	Notification := usecase.NewNotificationUsecase(st.db)
	return Repository{
		Wall,
		Group,
		Keyword,
		Post,
		Notification,
	}
}
