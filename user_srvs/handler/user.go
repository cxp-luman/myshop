package handler

import (
	"context"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"myshop/user_srvs/global"
	"myshop/user_srvs/model"
	"myshop/user_srvs/proto"
	"strings"
	"time"
)

func UserModelToResp(user model.User) *proto.UserInfoResp {
	UserInfoResp := proto.UserInfoResp{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.NickName,
		Mobile:   user.Mobile,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		UserInfoResp.BirthDay = uint64(user.Birthday.Unix())
	}
	return &UserInfoResp
}

type UserServer struct {
	proto.UnimplementedUserServer
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func (*UserServer) GetUserList(ctx context.Context, res *proto.PageInfo) (*proto.UserListResp, error) {
	var users []model.User
	DB := global.DB
	result := DB.Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	resp := &proto.UserListResp{}
	resp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(res.Pn), int(res.Pn))).Find(&users)
	for _, user := range users {
		userInfoResp := UserModelToResp(user)
		resp.Data = append(resp.Data, userInfoResp)
	}
	return resp, nil
}

// GetUserById find user info by telephone
func (*UserServer) GetUserById(ctx context.Context, res *proto.IdRequest) (*proto.UserInfoResp, error) {
	var user model.User
	id := res.Id
	DB := global.DB
	result := DB.First(&user, int(id))
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return UserModelToResp(user), nil
}

func (*UserServer) GetUserByMobile(ctx context.Context, res *proto.MobileRequest) (*proto.UserInfoResp, error) {
	var user model.User
	mobile := res.Mobile
	DB := global.DB
	result := DB.Where("mobile = ?", mobile).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	resp := UserModelToResp(user)
	return resp, nil
}

func (*UserServer) CreateUser(ctx context.Context, res *proto.CreateUserInfo) (*proto.UserInfoResp, error) {
	var user model.User
	DB := global.DB
	result := DB.Where(&model.User{Mobile: res.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Mobile = res.Mobile
	user.NickName = res.NickName
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(res.PassWord, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	result = DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	resp := UserModelToResp(user)
	return resp, nil
}

func (*UserServer) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResp, error) {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[1], passwordInfo[2], options)
	return &proto.CheckResp{Success: check}, nil
}

func (*UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	var user model.User
	DB := global.DB
	result := DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user.NickName = req.NickName
	user.Gender = req.Gender
	birthday := time.Unix(int64(req.BirthDay), 0)
	user.Birthday = &birthday
	result = DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}
