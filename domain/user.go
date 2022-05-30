package domain

import (
	"context"
	"douyin-app/repository"
)

// UserInfo 返回用户信息
type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// FillUserInfo 将user对象转换为UserInfo对象
func FillUserInfo(user *repository.User) *UserInfo {
	userInfo := &UserInfo{
		Id:            user.UserId,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      true,
	}
	return userInfo
}

// GetToUserIdsFromFollowList 从返回的follow对象中，获得关注id列表
func GetToUserIdsFromFollowList(followList []*repository.Follow) []int64 {
	ids := make([]int64, len(followList))
	for i := range followList {
		ids[i] = followList[i].ToUserId
	}
	return ids
}

// GetFromUserIdsFromFollowList 从返回的follow对象中，获得粉丝id列表
func GetFromUserIdsFromFollowList(followList []*repository.Follow) []int64 {
	ids := make([]int64, len(followList))
	for i := range followList {
		ids[i] = followList[i].FromUserId
	}
	return ids
}

// GetUserInfosFromIds 根据获得的id列表去User表中查询
func GetUserInfosFromIds(ctx context.Context, userIds []int64) ([]*UserInfo, error) {
	// 根据id列表去查询users列表
	users, err := repository.GetUserRepository().FindByUserIds(ctx, userIds)
	if err != nil {
		return nil, err
	}
	userInfos := make([]*UserInfo, len(userIds))
	for i := range users {
		userInfos[i] = FillUserInfo(users[i])
	}
	return userInfos, nil
}
