package usecase_test

import (
	"context"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/repository/mock"
	"glamgrove/pkg/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock repository
	mockRepo := mock.NewMockUserRepository(ctrl)

	// Create a user use case with the mock repository
	userUseCase := usecase.NewUserUseCase(mockRepo)

	// Define a test user
	testUser := domain.User{
		ID: 1, UserName: "Ajin", FirstName: "Ajin", LastName: "A", Age: 22, Email: "ajin@gmail.com", Phone: "9074386600", Password: "ajin123", BlockStatus: false,
	}

	// Set up expectations for the mock repository
	mockRepo.EXPECT().
		FindUser(gomock.Any(), testUser).
		Return(testUser, nil).
		Times(1) // Expect the FindUser method to be called once with the test user

	// Call the FindUser method
	found, err := userUseCase.FindUser(context.Background(), testUser)

	// Check if user is found and error is nil
	assert.True(t, found, "expected user to be found")
	assert.NoError(t, err, "expected no error")
}
