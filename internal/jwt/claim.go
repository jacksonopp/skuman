package jwt

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jacksonopp/skuman/db/db"
)

type Claims struct {
	SessionId string `json:"sessionId"`
	UserId    int64  `json:"userId"`
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}

func NewClaims(data db.GetSessionBySessionIdRow) Claims {
	s := fmt.Sprintf("%s%s", data.SessionID, data.ExpiresAt.String())
	h := sha256.New()

	h.Write([]byte(s))
	id := h.Sum(nil)

	return Claims{
		SessionId: data.SessionID,
		UserId:    data.ID,
		UserEmail: data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(data.ExpiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        string(id),
		},
	}
}

func (c *Claims) Sign() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	ss, err := token.SignedString([]byte("areallysecurekeyfromenv"))
	if err != nil {
		return ss, err
	}

	ss = fmt.Sprintf("Bearer %s", ss)

	return ss, nil
}

func (c *Claims) Parse(tokenString string) (*Claims, error) {
	tokenSlc := strings.Split(tokenString, " ")
	token, err := jwt.ParseWithClaims(tokenSlc[1], c, func(t *jwt.Token) (interface{}, error) {
		return []byte("areallysecurekeyfromenv"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		c = claims
		return claims, nil
	} else {
		return nil, fmt.Errorf("error parsing claims")
	}
}
