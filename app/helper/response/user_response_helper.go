package response

import (
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/user"
)

func ToUserResponse(userDomain domain.User) user.UserResponse {
	return user.UserResponse{
		Id:    userDomain.Id,
		Name:  userDomain.Name,
		Email: userDomain.Email,
	}
}

func ToUserResponses(users []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToUserLoginResponse(token string) string {
	return token
}
