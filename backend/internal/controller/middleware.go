package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type contextKey string

const initDataKey contextKey = "init_data"

var (
	ErrInitDataNotFound   = errors.New("init data not found in context")
	ErrMetadataNotFound   = errors.New("metadata not found")
	ErrInvalidToken       = errors.New("invalid tg token")
	ErrInvalidUserData    = errors.New("invalid user data")
	ErrFailedToCreateUser = errors.New("failed to create user")
)

func (c *ControllerService) InitDataMiddleware() runtime.Interceptor {
	return func(ctx context.Context, req *http.Request) (context.Context, error) {
		var token string

		token = req.Header.Get("tg-token")
		if token == "" {
			return ctx, runtime.Error(http.StatusUnauthorized, ErrInvalidToken.Error())
		}

		if err := initdata.Validate(token, c.botToken, 24*time.Hour); err != nil {
			return ctx, runtime.Error(http.StatusUnauthorized, ErrInvalidToken.Error())
		}

		user, err := initdata.Parse(token)
		if err != nil {
			return ctx, runtime.Error(http.StatusUnauthorized, ErrInvalidUserData.Error())
		}

		if err := c.service.UserCreate(ctx, user.User.ID, user.ChatInstance, user.User.LanguageCode); err != nil {
			return ctx, runtime.Error(http.StatusInternalServerError, ErrFailedToCreateUser.Error())
		}

		ctx = context.WithValue(ctx, initDataKey, user)

		return ctx, nil
	}
}

func GetInitDataFromContext(ctx context.Context) (initdata.InitData, error) {
	user, ok := ctx.Value(initDataKey).(initdata.InitData)
	if !ok {
		return initdata.InitData{}, runtime.Error(http.StatusUnauthorized, ErrInitDataNotFound.Error())
	}

	return user, nil
}

func (c *ControllerService) LogMiddleware() runtime.Interceptor {
	return func(ctx context.Context, req *http.Request) (context.Context, error) {
		initData, _ := GetInitDataFromContext(ctx)

		log.Info().Str("method", req.Method).Str("path", req.URL.Path).Str("user_id", strconv.FormatInt(initData.User.ID, 10)).Msg("handle controller request")

		return ctx, nil
	}
}
