package response

import "github.com/online_marketplace/internal/types/entity"

type UserResponse struct {
	Id        int32  `json:"id"`
	UId       string `json:"uid"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Phone     string `json:"phone"`
}

func UserMapToResponse(data *entity.User) *UserResponse {
	return &UserResponse{
		Id:        data.Id,
		UId:       data.UId,
		Email:     data.Email,
		LastName:  data.LastName,
		FirstName: data.FirstName,
		Phone:     data.Phone,
	}
}

func UserMapToResponses(items []*entity.User) []*UserResponse {
	var results []*UserResponse
	for _, val := range items {
		results = append(results, UserMapToResponse(val))
	}
	return results
}

type RegisterResponse struct {
	AccessToken  *TokenResponse `json:"access_token"`
	RefreshToken *TokenResponse `json:"refresh_token"`
	User         *UserResponse  `json:"user"`
}

type LoginResponse struct {
	AccessToken  *TokenResponse `json:"access_token"`
	RefreshToken *TokenResponse `json:"refresh_token"`
	User         *UserResponse  `json:"user"`
}
