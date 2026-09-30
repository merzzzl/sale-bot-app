package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestAppHandler(t *testing.T) {
	dir := t.TempDir()
	for name, content := range map[string]string{"index.html": "mini app", "app.js": "asset"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
	}
	api := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("api"))
	})
	handler := appHandler(api, dir)
	for _, tc := range []struct {
		method, path string
		status       int
		body         string
	}{
		{"GET", "/", 200, "mini app"},
		{"GET", "/bots/123", 200, "mini app"},
		{"GET", "/app.js", 200, "asset"},
		{"GET", "/missing.js", 404, ""},
		{"GET", "/api/v1/bots", 202, "api"},
		{"POST", "/webhook/token", 202, "api"},
		{"GET", "/tg/bot123/getMe", 202, "api"},
		{"POST", "/bots/123", 404, ""},
		{"HEAD", "/bots/123", 200, ""},
	} {
		t.Run(tc.method+tc.path, func(t *testing.T) {
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
			if w.Code != tc.status || (tc.body != "" && w.Body.String() != tc.body) {
				t.Fatalf("got %d %q, want %d %q", w.Code, w.Body.String(), tc.status, tc.body)
			}
		})
	}
}
