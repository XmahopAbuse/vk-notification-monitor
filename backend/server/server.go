package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mymmrac/telego"
	"net/http"
	"strings"
	"time"
	"vk-notification-monitor/config"
	"vk-notification-monitor/entity/vkapi"
	"vk-notification-monitor/pkg/sender"
	"vk-notification-monitor/pkg/syncpost"
	"vk-notification-monitor/server/handlers"
	"vk-notification-monitor/store"
)

type Server struct {
	router       *gin.Engine
	Repository   store.Repository
	vkapi        vkapi.VKApi
	EnableNotify int
	senders      []sender.Sender
	config       *config.Config
}

func NewServer(config *config.Config) (*Server, error) {
	var server Server
	st, err := store.NewStore(fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable password=%s user=%s",
		config.DB_HOST, config.DB_PORT, config.DB_NAME, config.DB_PASSWORD, config.DB_USER), "postgres")
	err = st.Ping()
	if err != nil {
		return nil, err
	}

	server.config = config
	server.vkapi = vkapi.VKApi{
		AccessToken: config.VK_TOKEN,
		V:           config.VK_VERSION,
		URL:         config.VK_API_URL,
	}
	server.Repository = store.NewRepository(st, server.vkapi)
	server.router = gin.Default()

	return &server, nil
}

func (s *Server) Run() {

	// Инициализируем роуты
	s.initRoutes()

	// Запускаем тикер для фоновых задач
	s.RunTicker(s.config.SYNC_INTERVAL)

	// Инициализируем сервисы уведомлений
	s.initSenders()

	// Запускаем роутер
	s.router.Run(s.config.SERVER_ADDRESS)
}

func (s *Server) RunTicker(interval int) {

	ticker := time.Tick(time.Duration(interval) * time.Second)

	go func() {
		for range ticker {
			groups, _ := s.Repository.Group.GetAllGroups()
			keywords, _ := s.Repository.Keyword.GetAll()
			syncpost.SyncPostsByKeywords(groups, keywords, &s.Repository.Wall, &s.Repository.Post, s.config.ENABLE_NOTIFICATIONS, s.senders)
		}
	}()

}

func (s *Server) initSenders() error {
	if s.config.ENABLE_NOTIFICATIONS == 1 {
		senders := strings.Split(s.config.SENDERS, ",")
		for _, name := range senders {
			switch name {
			case "telegram":
				bot, err := telego.NewBot(s.config.TELEGRAM_BOT_TOKEN, telego.WithDefaultLogger(false, true))
				if err != nil {
					return err
				}
				sn := sender.NewTelegramSender(bot)
				s.senders = append(s.senders, sn)
			default:
				return fmt.Errorf("Unknown sender type")
				// here you can add some of you senders
			}
		}
	}

	return nil
}

func (s *Server) initRoutes() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowMethods("OPTIONS, GET")
	s.router.Use(cors.New(config))

	// Middleware для обработки CORS
	s.router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	})

	// Отдача статики

	v1 := s.router.Group("/v1")

	//	v1.GET("/ping", handlers.AddPingRoutes())

	//Posts group
	postsGroup := v1.Group("/posts")
	// Ищет посты по ключевым словам
	postsGroup.POST("/sync", handlers.SyncPostsByKeywords(&s.Repository, s.config, s.senders))
	// Отдаем слова из локальной БД
	postsGroup.GET("/get", handlers.GetPost(&s.Repository))
	postsGroup.PATCH("/update/:hash", handlers.UpdatePost(&s.Repository))

	// Groups
	groupGroup := v1.Group("/group")
	groupGroup.POST("/add", handlers.AddGroup(&s.Repository))
	groupGroup.GET("/get", handlers.GetAllGroups(&s.Repository))
	groupGroup.POST("/delete", handlers.DeleteGroupByAddress(&s.Repository))
	groupGroup.GET("/get/:id", handlers.GetGroupById(&s.Repository))
	// Keywords
	keywordGroup := v1.Group("/keyword")
	keywordGroup.POST("/add", handlers.AddKeyword(&s.Repository))
	keywordGroup.GET("/get", handlers.GetAllKeywords(&s.Repository))
	keywordGroup.DELETE("/:name/delete", handlers.DeleteKeyword(&s.Repository))

	//Author
	authorGroup := v1.Group("/author")
	authorGroup.POST("/get", handlers.GetAuthor(&s.Repository))

	// Notification
	notificationGroup := v1.Group("/notification")
	notificationGroup.POST("/add", handlers.AddNotification(&s.Repository))
	notificationGroup.GET("/get", handlers.GetNotificationByType(&s.Repository))

	// Static
	s.router.Use(static.Serve("/", static.LocalFile("./frontend", true)))

}
