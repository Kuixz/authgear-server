package accountmanagement

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	goredis "github.com/go-redis/redis/v8"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis/appredis"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/duration"
)

type RedisStore struct {
	Context context.Context
	AppID   config.AppID
	Redis   *appredis.Handle
	Clock   clock.Clock
}

type GenerateOAuthTokenOptions struct {
	UserID      string
	Alias       string
	MaybeState  string
	RedirectURI string
}

type GenerateTOTPTokenOptions struct {
	UserID     string
	TOTPSecret string
	OTPAuthURI string
}

type GenerateOOBTokenOptions struct {
	UserID     string
	OOBChannel string
	OOBTarget  string
	State      string
}

func (s *RedisStore) GenerateOAuthToken(options GenerateOAuthTokenOptions) (string, error) {
	tokenString := GenerateToken()
	tokenHash := HashToken(tokenString)

	now := s.Clock.NowUTC()
	ttl := duration.UserInteraction
	expireAt := now.Add(ttl)

	token := &Token{
		AppID:       string(s.AppID),
		UserID:      options.UserID,
		Alias:       options.Alias,
		State:       options.MaybeState,
		RedirectURI: options.RedirectURI,
		TokenHash:   tokenHash,
		CreatedAt:   &now,
		ExpireAt:    &expireAt,
	}

	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	tokenKey := tokenKey(token.AppID, token.TokenHash)

	err = s.Redis.WithConnContext(s.Context, func(conn *goredis.Conn) error {
		_, err = conn.SetNX(s.Context, tokenKey, tokenBytes, ttl).Result()
		if errors.Is(err, goredis.Nil) {
			return errors.New("account management token collision")
		} else if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *RedisStore) GenerateTOTPToken(options GenerateTOTPTokenOptions) (string, error) {
	tokenString := GenerateToken()
	tokenHash := HashToken(tokenString)

	now := s.Clock.NowUTC()
	ttl := duration.UserInteraction
	expireAt := now.Add(ttl)

	token := &Token{
		AppID:      string(s.AppID),
		UserID:     options.UserID,
		TOTPSecret: options.TOTPSecret,
		OTPAuthURI: options.OTPAuthURI,
		TokenHash:  tokenHash,
		CreatedAt:  &now,
		ExpireAt:   &expireAt,
	}

	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	tokenKey := tokenKey(token.AppID, token.TokenHash)

	err = s.Redis.WithConnContext(s.Context, func(conn *goredis.Conn) error {
		_, err = conn.SetNX(s.Context, tokenKey, tokenBytes, ttl).Result()
		if errors.Is(err, goredis.Nil) {
			return errors.New("account management token collision")
		} else if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *RedisStore) GenerateOOBToken(options GenerateOOBTokenOptions) (string, error) {
	tokenString := GenerateToken()
	tokenHash := HashToken(tokenString)

	now := s.Clock.NowUTC()
	ttl := duration.UserInteraction
	expireAt := now.Add(ttl)

	token := &Token{
		AppID:      string(s.AppID),
		UserID:     options.UserID,
		OOBChannel: options.OOBChannel,
		OOBTarget:  options.OOBTarget,
		State:      options.State,
		TokenHash:  tokenHash,
		CreatedAt:  &now,
		ExpireAt:   &expireAt,
	}

	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	tokenKey := tokenKey(token.AppID, token.TokenHash)

	err = s.Redis.WithConnContext(s.Context, func(conn *goredis.Conn) error {
		_, err = conn.SetNX(s.Context, tokenKey, tokenBytes, ttl).Result()
		if errors.Is(err, goredis.Nil) {
			return errors.New("account management token collision")
		} else if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *RedisStore) ConsumeToken(tokenStr string) (*Token, error) {
	tokenHash := HashToken(tokenStr)

	tokenKey := tokenKey(string(s.AppID), tokenHash)

	var tokenBytes []byte
	err := s.Redis.WithConnContext(s.Context, func(conn *goredis.Conn) error {
		var err error
		tokenBytes, err = conn.GetDel(s.Context, tokenKey).Bytes()
		if errors.Is(err, goredis.Nil) {
			// Token Invalid
			return ErrOAuthTokenInvalid
		} else if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var token Token
	err = json.Unmarshal(tokenBytes, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *RedisStore) GetToken(tokenStr string) (*Token, error) {
	tokenHash := HashToken(tokenStr)

	tokenKey := tokenKey(string(s.AppID), tokenHash)

	var tokenBytes []byte
	err := s.Redis.WithConnContext(s.Context, func(conn *goredis.Conn) error {
		var err error
		tokenBytes, err = conn.Get(s.Context, tokenKey).Bytes()
		if errors.Is(err, goredis.Nil) {
			// Token Invalid
			return ErrOAuthTokenInvalid
		} else if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var token Token
	err = json.Unmarshal(tokenBytes, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func tokenKey(appID string, tokenHash string) string {
	return fmt.Sprintf("app:%s:account-management-token:%s", appID, tokenHash)
}
