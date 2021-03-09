package tests

import (
	"context"
	"fmt"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/lyouthzzz/go-web-layout/internal/repo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser(t *testing.T) {

	userRepository := repo.NewUserRepository(db)

	t.Run("create", func(t *testing.T) {
		user, err := userRepository.Create(context.TODO(), &domain.User{
			Username: "阳",
			Nickname: "youn",
			Password: "youn",
			Email:    "111111312@qq.com",
		})

		require.NoError(t, err)
		require.NotEmpty(t, user)
		require.Equal(t, user.Username, "阳")
		require.Equal(t, user.Nickname, "youn")
	})

	t.Run("get", func(t *testing.T) {
		user, err := userRepository.Get(context.TODO(), 1)

		require.NoError(t, err)
		require.NotEmpty(t, user)
		fmt.Printf("%+v\n", user)
	})

	t.Run("update", func(t *testing.T) {
		err := userRepository.Update(context.TODO(), 1, &domain.User{Username: "younupdated"})
		require.NoError(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		err := userRepository.Delete(context.TODO(), 1)
		require.NoError(t, err)
	})
}
