package webhook

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/rs/zerolog/log"
)

type Router struct {
	prefix      string
	mu          sync.RWMutex
	paths       map[string]http.Handler
	secretToken string
}

func NewRouter(prefix, secretToken string) *Router {
	if prefix[len(prefix)-1] != '/' {
		prefix += "/"
	}

	return &Router{
		prefix:      prefix,
		paths:       make(map[string]http.Handler),
		secretToken: secretToken,
	}
}

func (r *Router) add(b *bot.Bot) string {
	r.mu.Lock()
	r.paths[strconv.FormatInt(b.ID(), 10)] = b.WebhookHandler()
	r.mu.Unlock()

	return r.prefix + "webhook/" + strconv.FormatInt(b.ID(), 10)
}

func (r *Router) remove(b *bot.Bot) {
	r.mu.Lock()
	delete(r.paths, strconv.FormatInt(b.ID(), 10))
	r.mu.Unlock()
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Telegram-Bot-Api-Secret-Token") != r.secretToken {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	puri, err := url.Parse(r.prefix)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	if !strings.HasPrefix(req.URL.Path, puri.Path+"webhook/") {
		http.NotFound(w, req)

		return
	}

	id := strings.TrimPrefix(req.URL.Path, puri.Path+"webhook/")

	r.mu.RLock()
	h := r.paths[id]
	r.mu.RUnlock()

	if h == nil {
		http.NotFound(w, req)

		return
	}

	log.Info().Str("bot_id", id).Msg("handle webhook request")

	h.ServeHTTP(w, req)
}
