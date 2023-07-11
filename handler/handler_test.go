package handler_test

import (
	"belajar-golang-mock/constant"
	"belajar-golang-mock/handler"
	service_mocks "belajar-golang-mock/service/mocks"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCat(t *testing.T) {

	mockService := new(service_mocks.Service)
	handler := handler.New(mockService)

	payload := constant.Request{
		Name: "dummy-name",
	}

	// CreateNewCat(ctx context.Context, request constant.Request) (int64, error)
	mockService.On("CreateNewCat", mock.Anything, payload).Return(int64(1), nil).Once()

	payloadJson, err := json.Marshal(payload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/cats", strings.NewReader(string(payloadJson)))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	c.Request = req
	handler.CreateCat(c)

	res, err := io.ReadAll(rec.Result().Body)
	assert.Nil(t, err)

	actualResult := make(map[string]interface{}, 0)
	json.Unmarshal(res, &actualResult)

	expectedResult := map[string]interface{} {
		"success": true,
		"data": map[string]interface{} {
			"id": float64(1),
		},
	}

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Equal(t, expectedResult, actualResult)
}