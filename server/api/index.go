package handler

import (
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/urlvet/urlvet/pkg/docs" // swagger docs registration
	"github.com/urlvet/urlvet/pkg/handler"
	"github.com/urlvet/urlvet/pkg/logger"
	"github.com/urlvet/urlvet/pkg/service/rank"
	"github.com/urlvet/urlvet/pkg/service/typosquat"
)

var (
	engine http.Handler
	once   sync.Once
)

func initApp() {
	_ = godotenv.Load(".env", "/app/.env")
	if os.Getenv("LOG_DIR") == "" {
		os.Setenv("LOG_DIR", "/tmp/logs")
	}
	logger.Init()

	if err := rank.LoadDomainRanks(); err != nil {
		logger.Warn("rank load warning", "err", err)
	}

	if err := typosquat.LoadTopDomains(); err != nil {
		logger.Warn("typosquat load warning", "err", err)
	}

	engine = handler.SetupRouter()
}

// Handler is the Vercel serverless entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initApp)

	// Fix path when Vercel serverless rewrites incoming URL
	if r.URL.Path == "/api" || r.URL.Path == "/api/" || r.URL.Path == "/api/index" || r.URL.Path == "/api/index.go" {
		if matched := r.Header.Get("x-matched-path"); matched != "" {
			// strip query if present
			if idx := strings.Index(matched, "?"); idx != -1 {
				r.URL.Path = matched[:idx]
			} else {
				r.URL.Path = matched
			}
		} else if raw := r.Header.Get("x-forwarded-uri"); raw != "" {
			if idx := strings.Index(raw, "?"); idx != -1 {
				r.URL.Path = raw[:idx]
			} else {
				r.URL.Path = raw
			}
		}
	}

	engine.ServeHTTP(w, r)
}

