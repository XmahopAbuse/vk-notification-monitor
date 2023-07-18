package usecase

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/mymmrac/telego"
	"log"
	"vk-notification-monitor/entity"
)

type NotificationUsecase struct {
	db *sql.DB
}

func NewNotificationUsecase(db *sql.DB) NotificationUsecase {
	return NotificationUsecase{db: db}
}

func (notificator *NotificationUsecase) AddNotification(t, value string) error {
	sql, query, err := squirrel.Insert("notification").Columns("type", "value").Values(t, value).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}
	_, err = notificator.db.Exec(sql, query...)

	if err != nil {
		return err
	}
	return nil
}

func (notificator *NotificationUsecase) GetByType(t string) ([]entity.Notification, error) {
	sql, query, err := squirrel.Select("type", "value").From("notification").Where(squirrel.Eq{"type": t}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := notificator.db.Query(sql, query...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notifications := []entity.Notification{}

	for rows.Next() {
		notification := entity.Notification{}
		err := rows.Scan(&notification.Type, &notification.Value)
		if err != nil {
			log.Println(err)
			continue
		}

	}

	return notifications, nil

}

func (notificator *NotificationUsecase) GetAll() ([]entity.Notification, error) {
	sql, query, err := squirrel.Select("type", "value").From("notification").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := notificator.db.Query(sql, query...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notifications := []entity.Notification{}

	for rows.Next() {
		notification := entity.Notification{}
		err = rows.Scan(&notification.Type, &notification.Value)
		if err != nil {
			log.Println(err)
			continue
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil

}

func (notificator *NotificationUsecase) sendNotifications(post entity.Post, notifyUsers []entity.Notification, tgbot *telego.Bot) {
	for _, user := range notifyUsers {
		switch user.Type {
		case "telegram":
			log.Println("TELEGRAM SEND")
		}
	}
}
