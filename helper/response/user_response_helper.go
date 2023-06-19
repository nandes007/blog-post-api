package response

import (
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/user"
)

func ToUserResponse(userDomain domain.User) user.Response {
	return user.Response{
		Id:    userDomain.Id,
		Name:  userDomain.Name,
		Email: userDomain.Email,
	}
}

func ToUserResponses(users []domain.User) []user.Response {
	var userResponses []user.Response
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToUserLoginResponse(token string) string {
	return token
}
