package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"kredi-plus.com/be/model"
	"kredi-plus.com/be/repository"
	"testing"
)

var (
	userMockRepo = repository.NewUserMock()
	userService  = user{
		userRepo: userMockRepo,
	}
)

func TestUserService_FindById(t *testing.T) {
	user1 := model.User{
		Id:       1,
		Username: "superadmin",
		Password: "$2a$04$mVVUGfk3n86YIgAI9hxxGeB/MNBuIyW05PutPhK8zFxtATpM2n0Na",
		IsAdmin:  true,
	}

	ctx := context.Background()

	// set mock
	userMockRepo.Mock.On("GetDetailById", int64(1)).Return(user1, nil)
	userMockRepo.Mock.On("GetDetailById", int64(2)).Return(model.User{}, gorm.ErrRecordNotFound)

	// not found
	res, err := userService.GetDetailById(ctx, 2)
	assert.NotNil(t, err)
	assert.Equal(t, res, toUserResponse(model.User{}))

	// success
	res, err = userService.GetDetailById(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, res, toUserResponse(user1))
}
