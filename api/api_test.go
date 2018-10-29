package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Testing the Redirect func

// Since mux basically handle methods and path used
func Test_Redirect(t *testing.T) {
	tt := []struct {
		method string
		url    string
		statusCode   int
	}{
		{method: "GET", url: "http://localhost:8080/paragliding/", statusCode: 301},
		{method: "GET", url: "http://localhost:8080/paragliding/rubish", statusCode: 404},
		{method: "POST", url: "http://localhost:8080/paragliding/123", statusCode: 404},
	}
	for _, tc := range tt {
		fmt.Println(tc.statusCode)
		req, err := http.NewRequest(tc.method, tc.url, nil)
		if err != nil {
			t.Errorf("Unexpected error, %d", err)
		}

		rec := httptest.NewRecorder()
		Redirect(rec,req)

		res := rec.Result()
		if req.URL.Path != "/paragliding/" {
			res.StatusCode = http.StatusNotFound
		}
		defer res.Body.Close()
		if res.StatusCode != tc.statusCode {
			t.Errorf("Expected status %v: got %v", tc.statusCode, res.Status)
		}
	}
}