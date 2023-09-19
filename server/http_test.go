package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/REST-API-Test/mocks"
	"github.com/REST-API-Test/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setup() (*StoreServer, *gin.Context, *httptest.ResponseRecorder, *mocks.IUsecase) {
	usecase := &mocks.IUsecase{}
	server := NewHttpServer(usecase)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return server, c, w, usecase
}

func TestPing(t *testing.T) {
	server, c, w, _ := setup()
	server.Ping(c)
	assert.Equal(t, 200, w.Code, "Received 200!")
}

func TestOrderHistorySuccess(t *testing.T) {
	server, c, w, m := setup()

	c.Request = &http.Request{}
	c.Request.URL = &url.URL{
		RawQuery: "start=2021-05-01&end=2021-06-01&type=day",
	}

	start, _ := time.Parse("2006-01-02", "2021-05-01")
	end, _ := time.Parse("2006-01-02", "2021-06-01")
	dr := types.DateRange{
		Start: start,
		End:   end,
		Type:  "day",
	}
	m.On("OrderHistory", dr).Return([]types.DateBucket{}, nil)
	server.OrderHistory(c)
	assert.Equal(t, 200, w.Code, "Order History failed")
}

func TestOrderHistoryFailOrder(t *testing.T) {
	server, c, w, m := setup()

	c.Request = &http.Request{}
	c.Request.URL = &url.URL{
		RawQuery: "start=2021-05-01&end=2021-06-01&type=day",
	}

	start, _ := time.Parse("2006-01-02", "2021-05-01")
	end, _ := time.Parse("2006-01-02", "2021-06-01")
	dr := types.DateRange{
		Start: start,
		End:   end,
		Type:  "day",
	}
	m.On("OrderHistory", dr).Return([]types.DateBucket{}, errors.New("Failed"))
	server.OrderHistory(c)
	assert.Equal(t, 404, w.Code, "Order History failed")
}

func TestOrderHistoryFailureBadQuery(t *testing.T) {
	server, c, w, _ := setup()
	c.Request = &http.Request{}
	c.Request.URL = &url.URL{
		RawQuery: "start=2021-05-01&type=day",
	}
	server.OrderHistory(c)
	assert.Equal(t, 400, w.Code)
}
