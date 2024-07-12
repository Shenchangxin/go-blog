package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-blog/global"
	"go-blog/model"
	"go-blog/proto"
	"go-blog/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UserServer struct {
}

func Paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (*UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(req.PageNo), int(req.PageSize))).Find(&users)

	for _, user := range users {
		userInfoRes := utils.UserModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRes)
	}
	return rsp, nil
}

func (*UserServer) GetUserByUserName(ctx context.Context, req *proto.UserNameRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{
		UserName: req.Username,
	}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	userInfoRes := utils.UserModelToResponse(user)
	return &userInfoRes, nil
}

func (*UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&user, req.Id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	userInfoRes := utils.UserModelToResponse(user)
	return &userInfoRes, nil
}

func (*UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{
		UserName: req.UserName,
	}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.UserName = req.UserName
	user.Phone = req.Phone
	user.NickName = req.NickName
	user.Sex = req.Sex

	//salt, _ := bcrypt.GenerateFromPassword([]byte("shenchangxin"), bcrypt.DefaultCost)
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	user.Password = string(encryptedPassword)
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	userInfoRes := utils.UserModelToResponse(user)
	return &userInfoRes, nil
}

func (*UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	var user model.User
	result := global.DB.Where(&user, req.Id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user.UserName = req.UserName
	user.NickName = req.NickName
	user.Sex = req.Sex
	user.Phone = req.Phone
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	user.Password = string(encryptedPassword)
	result = global.DB.Updates(&user)
	if result.Error != nil {
		return &empty.Empty{}, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (*UserServer) CheckPassword(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	err := bcrypt.CompareHashAndPassword([]byte(req.EncryptedPassword), []byte(req.Password))
	if err != nil {
		return &proto.CheckResponse{
			Success: false,
		}, err
	} else {
		return &proto.CheckResponse{
			Success: true,
		}, nil
	}

}

func (*UserServer) mustEmbedUnimplementedUserServer() {

}
