package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/luvkrai/mytestablecodes/services"
)

type fakePingService struct {
}

func ( pingServer *fakePingService) HandlePing() (string, error) {
	fmt.Println("faking complicated things...")
	return "pong", nil
}

func TestPingController(t *testing.T) {
	fakeService := &fakePingService{}
	services.PingService = fakeService
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	PingController.Ping(context)
	if response.Code != http.StatusOK {
		t.Error("response code should be 200")
	}

	if response.Body.String() != "pong" {
		t.Error("response body should say 'pong'")
	}

}