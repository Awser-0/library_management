package user

import "library/internal/model/entity"

func userHandle(user entity.User) map[string]any {
	return map[string]any{
		"id":          user.Id,
		"username":    user.Username,
		"phone":       user.Phone,
		"sex":         user.Sex,
		"birth":       user.Birth,
		"isAdmin":     user.IsAdmin,
		"update_time": user.UpdateTime,
		"create_time": user.CreateTime,
	}
}
