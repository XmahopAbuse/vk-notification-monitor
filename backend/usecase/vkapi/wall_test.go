package vkapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"vk-notification-monitor/entity/vkapi"
)

type MockWall struct {
}

func TestGetWallPostsByDomain(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/method/wall.get" {
			t.Errorf("unexpected path, expected /method/wall.get, got %s", r.URL.Path)
		}

		response := `{"response":{"count":1,"items":[{"text":"Test"}]}}`
		w.Write([]byte(response))
	}))
	defer mockServer.Close()

	vkapi := vkapi.VKApi{
		URL:         mockServer.URL,
		AccessToken: "mock_access_token",
		V:           "5.0",
	}
	wallUsecase := &WallUsecase{
		vkapi: vkapi,
	}

	domain := "test_domain"
	wall, err := wallUsecase.GetWallPostsByDomain(domain)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if wall.Response.Count != 1 {
		t.Errorf("expected count = 1, got %d", wall.Response.Count)
	}
	if len(wall.Response.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(wall.Response.Items))
	}
	if wall.Response.Items[0].Text != "Test" {
		t.Errorf("unexpected text, expected 'Test', got '%s'", wall.Response.Items[0].Text)
	}
}
