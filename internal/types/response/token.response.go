package response

import (
	"time"

	tokenprovider "github.com/online_marketplace/helper/token_provider"
)

type TokenResponse struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func TokenMapToResponse(token tokenprovider.Token) *TokenResponse {
	return &TokenResponse{
		Token:   token.GetToken(),
		Created: token.GetCreated(),
		Expiry:  token.GetExpiry(),
	}
}
