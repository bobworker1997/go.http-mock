package http_mock_demo_test

import (
	"github.com/wiremock/go-wiremock"
	"net/http"
	"testing"
)

func TestDemo(t *testing.T) {
	// 建立 WireMock 客戶端，連接到你的 WireMock server（預設是 localhost:8080）
	client := wiremock.NewClient("http://localhost:8080")

	// 清空之前的 stub
	defer client.Reset()

	// 建立一個 stub for GET /api/example
	stub := wiremock.Get(wiremock.URLPathEqualTo("/api/example")).
		WillReturnResponse(
			wiremock.NewResponse().
				WithStatus(http.StatusOK).
				WithHeader("Content-Type", "application/json").
				WithBody(`{"UserId": 123}`),
		)

	// 註冊 stub
	err := client.StubFor(stub)
	if err != nil {
		t.Fatalf("failed to stub: %v", err)
	}

}
