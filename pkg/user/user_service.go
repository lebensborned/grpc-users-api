package user

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lebensborned/grpc-crud/config"
	"github.com/lebensborned/grpc-crud/pkg/api"
	"github.com/lebensborned/grpc-crud/pkg/cache"
	"github.com/lebensborned/grpc-crud/pkg/kafka"
	"github.com/lebensborned/grpc-crud/storage"
)

type UserServer struct {
	DB     *gorm.DB
	Cache  *cache.RedisClient
	Config config.Config
}

func (s *UserServer) CreateUserProfile(ctx context.Context, req *api.CreateUserProfileRequest) (*api.UserProfile, error) {
	u := storage.UserProfile{
		Name: req.UserProfile.Name,
		Age:  req.UserProfile.Age,
	}
	err := s.DB.Create(&u).Error
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("Added user %s [ID:%d]", u.Name, u.Id)
	err = kafka.PushCommentToQueue("users", []byte(msg))
	if err != nil {
		return nil, err
	}
	return &api.UserProfile{
		Name: u.Name,
		Id:   u.Id,
		Age:  u.Age,
	}, nil
}
func (s *UserServer) DeleteUserProfile(ctx context.Context, req *api.DeleteUserProfileRequest) (*api.EmptyReq, error) {
	u := storage.UserProfile{}
	if err := s.DB.Where("id = ?", req.Id).First(&u).Error; err != nil {
		return &api.EmptyReq{}, err
	}
	s.DB.Delete(&u)
	return &api.EmptyReq{}, nil
}
func (s *UserServer) ListUserProfiles(ctx context.Context, req *api.EmptyReq) (*api.ListUserProfilesResponse, error) {
	/*userlist := []*api.UserProfile{}
	err := s.DB.Model(&api.UserProfile{}).Limit(100).Find(&userlist).Error
	if err != nil {
		return &api.ListUserProfilesResponse{}, err
	}*/
	u, err := s.Cache.GetUsers()
	if err != nil {
		return nil, err
	}
	return &api.ListUserProfilesResponse{Profiles: u}, nil
}
