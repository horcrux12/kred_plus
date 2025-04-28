package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/model"
)

type userMock struct {
	Mock mock.Mock
}

func NewUserMock() *userMock {
	return &userMock{}
}

func (repo userMock) Create(ctx context.Context, inputModel *model.User) (err error) {
	//TODO implement me
	panic("implement me")
}

func (repo userMock) GetList(ctx context.Context, query in.UserRequest) (res []model.User, totalData int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo userMock) GetDetailById(ctx context.Context, userId int64) (res model.User, err error) {
	args := repo.Mock.Called(userId)
	if args.Get(1) != nil {
		err = args.Get(1).(error)
		return
	}

	res = args.Get(0).(model.User)
	return
}

func (repo userMock) Update(ctx context.Context, inputModel *model.User) (err error) {
	//TODO implement me
	panic("implement me")
}

func (repo userMock) DeleteById(ctx context.Context, userId int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (repo userMock) FindByUsername(ctx context.Context, username string) (res model.User, err error) {
	//TODO implement me
	panic("implement me")
}
