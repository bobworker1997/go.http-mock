package http_mock_demo

import (
	"github.com/h2non/gock"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	defer gock.Off() // 清除 mock

	// 設定 mock
	gock.New("http://example.com").
		Get("/api/GetUserInfo").
		Reply(200).
		JSON(map[string]interface{}{
			"UserId": 1,
			"Name":   "bob",
		})

	// 執行測試
	user, err := GetUserInfo("http://example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// 驗證結果
	if user.UserId != 1 {
		t.Errorf("Expected UserId 1, got %d", user.UserId)
	}
	if user.Name != "bob" {
		t.Errorf("Expected Name 'bob', got %s", user.Name)
	}

	// 確認 mock 被使用
	if !gock.IsDone() {
		t.Error("Expected mock to be called")
	}
}
