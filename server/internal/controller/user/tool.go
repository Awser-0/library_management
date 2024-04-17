package user

import (
	"library/internal/model/entity"
	"time"
)

type UserInfo struct {
	Id         int64      `json:"id"`
	Username   string     `json:"username"`
	Name       string     `json:"name"`
	Sex        string     `json:"sex"`
	Phone      string     `json:"phone"`
	Birth      *time.Time `json:"birth"`
	IsAdmin    bool       `json:"isAdmin"`
	UpdateTime time.Time  `json:"update_time"`
	CreateTime time.Time  `json:"create_time"`
}

func userHandle(user entity.User) UserInfo {
	return UserInfo{
		Id:         user.Id,
		Username:   user.Username,
		Name:       user.Name,
		Sex:        user.Sex,
		Phone:      user.Phone,
		Birth:      user.Birth,
		IsAdmin:    user.IsAdmin,
		UpdateTime: user.UpdateTime,
		CreateTime: user.CreateTime,
	}
}

func usersHandle(users []entity.User) []UserInfo {
	var userInfos = make([]UserInfo, len(users))
	for i, user := range users {
		userInfos[i].Id = user.Id
		userInfos[i].Username = user.Username
		userInfos[i].Name = user.Name
		userInfos[i].Sex = user.Sex
		userInfos[i].Phone = user.Phone
		userInfos[i].Birth = user.Birth
		userInfos[i].IsAdmin = user.IsAdmin
		userInfos[i].UpdateTime = user.UpdateTime
		userInfos[i].CreateTime = user.CreateTime
	}
	return userInfos
}
