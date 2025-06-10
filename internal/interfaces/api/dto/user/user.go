package user

import (
	domainUser "jollej/db-scout/internal/domain/user"
)

type ListUsersResponse struct {
	Users []User `json:"users"`
}

type User struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Age        int    `json:"Age"`
	Created_at string `json:"Created_at"`
	Updated_at string `json:"Updated_at"`
	Is_active  bool   `json:"Is_active"`
}

func ToUserDto(u domainUser.User) User {
	return User{
		Id:         u.Id,
		Name:       u.Name,
		Email:      u.Email,
		Age:        u.Age,
		Created_at: u.Created_at,
		Updated_at: u.Updated_at,
		Is_active:  u.Is_active,
	}
}

func ToUserDomain(u User) domainUser.User {
	return domainUser.User{
		Id:         u.Id,
		Name:       u.Name,
		Email:      u.Email,
		Age:        u.Age,
		Created_at: u.Created_at,
		Updated_at: u.Updated_at,
		Is_active:  u.Is_active,
	}
}
