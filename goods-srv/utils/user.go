package utils

import (
	"go-blog/goods-srv/model"
	"go-blog/goods-srv/proto"
)

func UserModelToResponse(user model.User) proto.UserInfoResponse {
	//在grpc中的message中的字段有默认值的不能随便赋值nil进去，容易出错
	userInfoRes := proto.UserInfoResponse{
		Id:       user.ID,
		NickName: user.NickName,
		Sex:      user.Sex,
		UserName: user.UserName,
		Phone:    user.Phone,
	}
	return userInfoRes
}
