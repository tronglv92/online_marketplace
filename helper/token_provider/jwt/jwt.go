package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	tokenprovider "github.com/online_marketplace/helper/token_provider"
	"github.com/online_marketplace/internal/config"
)

type TokenPayloadImp struct {
	UId string
}

func (p TokenPayloadImp) UserId() string {
	return p.UId
}

type TokenImp struct {
	Token   string
	Created time.Time
	Expiry  int
}

func (p TokenImp) GetToken() string {
	return p.Token
}
func (p TokenImp) GetCreated() time.Time {
	return p.Created
}
func (p TokenImp) GetExpiry() int {
	return p.Expiry
}

type jwtProvider struct {
	cf config.JWTConfig
}

func NewTokenJWTProvider(cf config.JWTConfig) tokenprovider.Provider {
	return &jwtProvider{cf: cf}
}

// MyClaims represents the custom claims structure
type MyClaims struct {
	Payload TokenPayloadImp `json:"payload"`
	jwt.RegisteredClaims
}

func (j *jwtProvider) SecretKey() string {
	return j.cf.HashSecret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	// generate the JWT
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		TokenPayloadImp{
			UId: data.UserId(),
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Local().Add(time.Minute * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(now.Local()),
			ID:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.cf.HashSecret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &TokenImp{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.cf.HashSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// validate the token
	if !res.Valid {
		// return nil, tokenprovider.ErrInvalidToken
		return nil, nil
	}

	claims, ok := res.Claims.(*MyClaims)

	if !ok {
		// return nil, tokenprovider.ErrInvalidToken
		return nil, nil
	}

	// return the token
	return claims.Payload, nil
}
