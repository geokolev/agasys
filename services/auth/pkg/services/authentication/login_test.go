package authentication

import (
	"errors"
	"testing"
	"time"

	"github.com/Kamva/mgm"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	mocks "github.com/jonsch318/royalafg/mocks/services/auth/pkg/services/user"
	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/models"
	"github.com/jonsch318/royalafg/services/auth/pkg/security"
	"github.com/jonsch318/royalafg/services/auth/pkg/serviceconfig"
)

func TestLogin(t *testing.T) {
	viper.SetDefault(config.JWTExpiresAt, time.Minute*60)
	viper.SetDefault(config.JWTIssuer, "example.com")
	viper.SetDefault(config.JWTSigningKey, "testkey")
	viper.SetDefault(serviceconfig.Pepper, "")
	mockUserService := &mocks.IUser{}
	auth := NewAuthentication(mockUserService)
	hash, err := security.HashPassword("testPassword", "")
	assert.Nil(t, err)

	user := &models.User{
		DefaultModel: mgm.DefaultModel{},
		Username:     "test",
		Email:        "test@test.com",
		Hash:         hash,
		FullName:     "test test",
	}

	t.Run("CredentialsValid", func(t *testing.T) {
		mockUserService.On("GetUserByUsernameOrEmail", "testUser").Return(user, nil)
		user, token, err := auth.Login("testUser", "testPassword")

		assert.Nil(t, err)
		assert.Equal(t, user, user)

		//Could test the
		assert.NotNil(t, token)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockUserService.On("GetUserByUsernameOrEmail", "testUser2").Return(nil, errors.New("user not found"))
		user, token, err := auth.Login("testUser2", "testPassword")

		assert.NotNil(t, err)
		assert.Equal(t, token, "")
		assert.Nil(t, user)
	})

}
