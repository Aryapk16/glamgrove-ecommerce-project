package repository

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/mock"
	"glamgrove/pkg/utils/response"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockUserRepository(ctrl)
	expectedUser := domain.User{ID: 1, UserName: "Ajin", FirstName: "Ajin", LastName: "A", Age: 22, Email: "ajin@gmail.com", Phone: "9074386600", Password: "ajin123", BlockStatus: false}
	expectedResponse := response.UserSignUp{ID: 1,
		UserName: "Ajin"}
	mockRepo.EXPECT().SaveUser(gomock.Any(), expectedUser).Return(expectedResponse, nil)
	ctx := context.Background()
	resp, err := mockRepo.SaveUser(ctx, expectedUser)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resp)
}
