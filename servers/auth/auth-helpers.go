package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jacksonopp/skuman/internal/constants"
	"github.com/jacksonopp/skuman/internal/jwt"
	"github.com/jacksonopp/skuman/internal/logger"
)

func claimToCookie(claim *jwt.Claims) (*http.Cookie, error) {
	token, err := claim.Sign()
	if err != nil {
		logger.Errorln("error creating jwt ", err)
		return nil, fmt.Errorf("error creating jwt: %v", err)
	}

	expiresAt := time.Until(claim.ExpiresAt.Time).Seconds()

	cookie := http.Cookie{
		Name:     constants.AUTH_COOKIE_NAME,
		Value:    token,
		Path:     "/",
		MaxAge:   int(expiresAt),
		Secure:   true,
		HttpOnly: true,
	}

	return &cookie, nil
}
