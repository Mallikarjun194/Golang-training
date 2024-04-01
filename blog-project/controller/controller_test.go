package controller

import (
	"blog-project/database"
	"blog-project/mocks"
	"blog-project/model"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var rPath = "/test"

func TestBlogController_AddBlog(t *testing.T) {
	tests := []struct {
		name string
		body model.Blog
		mock func() database.DatabaseI
		want int
	}{
		{
			name: "positive test",
			body: model.Blog{
				Title:   "Golang",
				Content: "Let's learn golang!!",
				Author:  "Zayn",
			},
			mock: func() database.DatabaseI {
				r := mocks.DatabaseI{}
				r.On("Create", mock.Anything).Return(nil)
				return &r
			},
			want: http.StatusCreated,
		},
		{
			name: "negative test",
			body: model.Blog{
				Title:  "Golang",
				Author: "Zayn",
			},
			mock: func() database.DatabaseI {
				r := mocks.DatabaseI{}
				r.On("Create", mock.Anything).Return(nil)
				return &r
			},
			want: http.StatusBadRequest,
		},
	}
	router := gin.Default()
	m := &BlogController{}
	router.POST(rPath, m.AddBlog)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Database = tt.mock()
			byteArray, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", rPath, bytes.NewReader(byteArray))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status : %d", w.Code)
			t.Logf("response: %s ", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}

}

func TestBlogController_GetBlogByIDNeg(t *testing.T) {
	tests := []struct {
		name string
		ID   string
		body model.Blog
		mock func() database.DatabaseI
		want int
	}{
		{
			name: "Negative test",
			ID:   "123",
			body: model.Blog{},
			mock: func() database.DatabaseI {
				r := mocks.DatabaseI{}
				r.On("FetchByID", mock.Anything).Return(errors.New("db call failed"))
				return &r
			},
			want: http.StatusInternalServerError,
		},
	}
	router := gin.Default()
	m := &BlogController{}
	router.GET(rPath+"/:id", m.GetBlogByID)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Database = tt.mock()
			byteArray, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath+"/"+tt.ID, bytes.NewReader(byteArray))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status : %d", w.Code)
			t.Logf("response: %s ", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}
