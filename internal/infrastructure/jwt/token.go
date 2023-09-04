package jwt

import (
	"monolith/internal/domain"
	"monolith/pkg/byteconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	expirationAccessTime  = time.Hour * 24
	expirationRefreshTime = expirationAccessTime * 30
)

type _TokenClaims struct {
	jwt.Claims
	CustomerId int   `json:"customer_id"`
	Expiration int64 `json:"expiration"`
}

type _TokenManager struct {
	TokenSalt []byte
}

func NewTokenManager(tokenSalt string) domain.TokenManager {
	return &_TokenManager{
		TokenSalt: byteconv.Bytes(tokenSalt),
	}
}

func (tokenManager *_TokenManager) generateToken(customerId int, expiration time.Duration) (string, error) {
	claims := _TokenClaims{
		CustomerId: customerId,
		Expiration: time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(tokenManager.TokenSalt)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (tokenManager *_TokenManager) GenerateAccess(customerId int) (string, error) {
	return tokenManager.generateToken(customerId, expirationAccessTime)
}

func (tokenManager *_TokenManager) GenerateRefresh(customerId int) (string, error) {
	return tokenManager.generateToken(customerId, expirationRefreshTime)
}

func (tokenManager *_TokenManager) Parse(token string) (domain.TokenClaims, error) {
	signedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrorInvalidType
		}

		return tokenManager.TokenSalt, nil
	})
	if err != nil {
		return domain.TokenClaims{}, err
	}

	if !signedToken.Valid {
		return domain.TokenClaims{}, domain.ErrorInvaledToken
	}

	claims, ok := signedToken.Claims.(_TokenClaims)
	if !ok {
		return domain.TokenClaims{}, domain.ErrorInvaledToken
	}

	return domain.TokenClaims{
		CustomerId: claims.CustomerId,
		Expiration: time.Unix(claims.Expiration, 0),
	}, nil
}
