package main

import (
	"context"
	"fmt"
	"github.com/Sal1mbek/letsgofurther/internal/assert"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestReadIDParam(t *testing.T) {
	app := newTestApp()

	tests := []struct {
		name      string
		requestID int64
		wantID    int64
	}{
		{
			name:      "123 ID test",
			requestID: 123,
			wantID:    123,
		},
		{
			name:      "Negative ID test",
			requestID: -2048,
			wantID:    0,
		},
		{
			name:      "int64 max number ID test",
			requestID: 1<<63 - 1,
			wantID:    1<<63 - 1,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/movies"+strconv.FormatInt(testCase.requestID, 10), nil)
			params := httprouter.Params{{Key: "id", Value: strconv.FormatInt(testCase.requestID, 10)}}
			ctx := req.Context()
			ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
			req = req.WithContext(ctx)

			id, _ := app.readIDParam(req)
			assert.Equal(t, id, testCase.wantID)
			fmt.Printf("Test %v results\n", testCase.name)
			fmt.Printf("\trequested: %v\n", testCase.requestID)
			fmt.Printf("\twant: %v\n", testCase.wantID)
			fmt.Printf("\tresult: %v\n", id)
		})
	}
}

func TestIntStr(t *testing.T) {
	tests := []struct {
		name string
		give int64
		want string
	}{
		{
			name: "1 number test",
			give: 1,
			want: "1",
		},
		{
			name: "0 number test",
			give: 0,
			want: "0",
		},
		{
			name: "negative number",
			give: -100,
			want: "-100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, strconv.FormatInt(tt.give, 10), tt.want)
		})
	}

}
