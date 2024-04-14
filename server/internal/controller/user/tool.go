package user

import (
	"library/internal/model/entity"
	"time"
)

func userHandle(user entity.User) map[string]any {
	return map[string]any{
		"id":          user.Id,
		"username":    user.Username,
		"name":        user.Name,
		"phone":       user.Phone,
		"sex":         user.Sex,
		"birth":       user.Birth,
		"isAdmin":     user.IsAdmin,
		"update_time": user.UpdateTime,
		"create_time": user.CreateTime,
	}
}

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
