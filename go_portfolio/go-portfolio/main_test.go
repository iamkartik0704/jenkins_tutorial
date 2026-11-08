package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupTestServer(t *testing.T) *Server {
	t.Helper()
	server, err := NewServer(content)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	return server
}

func TestIndexRoute(t *testing.T) {
	server := setupTestServer(t)
	handler := server.Routes()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", rec.Code)
	}

	body := rec.Body.String()
	expectedStrings := []string{
		"Kartik Chawla",
		"IIT Patna",
		"Software Engineer",
		"Backend Developer",
		"https://www.linkedin.com/in/kartik-chawla-189430203/",
		"https://github.com/iamkartik0704",
		"https://www.kartikchawla.in/",
		"Compile",
		"ToadCode",
		"Celesta",
		"Anwesha",
		"Jenkins Pipeline",
		"Docker Mastery",
		"Git and GitHub",
		"AWS Cloud",
		"/jenkins",
		"/docker",
		"/github",
		"/aws",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(body, expected) {
			t.Errorf("Index response body does not contain expected substring %q", expected)
		}
	}
}

func TestPlaylistRoutes(t *testing.T) {
	server := setupTestServer(t)
	handler := server.Routes()

	tests := []struct {
		route          string
		expectedTitle  string
		expectedListID string
	}{
		{
			route:          "/jenkins",
			expectedTitle:  "🧩 Jenkins Pipeline",
			expectedListID: "PLBr8obKbpkYvJEaPmrzhHhwx8uPj8WYbg",
		},
		{
			route:          "/docker",
			expectedTitle:  "🐳 Docker Mastery",
			expectedListID: "PLBr8obKbpkYsFtSF1XY9rM_3LH8LKRwSw",
		},
		{
			route:          "/github",
			expectedTitle:  "💻 Git and GitHub",
			expectedListID: "PLBr8obKbpkYt679NgO1KqZOoY_QYYLBl2",
		},
		{
			route:          "/aws",
			expectedTitle:  "☁️ AWS Cloud",
			expectedListID: "PLBr8obKbpkYtn5eIK3iTgiV7MSgFwO-tr",
		},
	}

	for _, tc := range tests {
		t.Run(tc.route, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.route, nil)
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d for %s", rec.Code, tc.route)
			}

			body := rec.Body.String()
			if !strings.Contains(body, tc.expectedTitle) {
				t.Errorf("Response body missing expected title %q for %s", tc.expectedTitle, tc.route)
			}
			if !strings.Contains(body, tc.expectedListID) {
				t.Errorf("Response body missing expected list ID %q for %s", tc.expectedListID, tc.route)
			}
			if !strings.Contains(body, "https://www.youtube.com/embed/videoseries?list=") {
				t.Errorf("Response body missing YouTube embed iframe for %s", tc.route)
			}
		})
	}
}

func TestStaticAssets(t *testing.T) {
	server := setupTestServer(t)
	handler := server.Routes()

	t.Run("CSS stylesheet", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/static/style.css", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200 for style.css, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "--bg:#0b0f17") {
			t.Errorf("style.css does not contain expected CSS variable --bg")
		}
	})

	t.Run("Hero image", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/static/img/hero.jpg", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200 for hero.jpg, got %d", rec.Code)
		}
		if rec.Body.Len() == 0 {
			t.Errorf("hero.jpg body is empty")
		}
	})
}

func TestHealthCheck(t *testing.T) {
	server := setupTestServer(t)
	handler := server.Routes()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK for /healthz, got %d", rec.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("Failed to parse json health response: %v", err)
	}

	if res["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got %v", res["status"])
	}
	if res["author"] != "Kartik Chawla" {
		t.Errorf("Expected author 'Kartik Chawla', got %v", res["author"])
	}
	if res["portfolio"] != "https://www.kartikchawla.in/" {
		t.Errorf("Expected portfolio URL 'https://www.kartikchawla.in/', got %v", res["portfolio"])
	}
}

func TestNotFound(t *testing.T) {
	server := setupTestServer(t)
	handler := server.Routes()

	req := httptest.NewRequest(http.MethodGet, "/unknown-page", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("Expected 404 Not Found, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "404 - Page Not Found") {
		t.Errorf("Expected 404 page body to contain '404 - Page Not Found'")
	}
}
