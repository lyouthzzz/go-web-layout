package authenticator

import (
	"context"
	"fmt"
	"github.com/lyouthzzz/framework/pkg/auth/authn"
	"github.com/lyouthzzz/framework/pkg/auth/user"
	"github.com/lyouthzzz/go-web-layout/internal/auth/store"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type UserInfo struct {
	Id string
}

func (u *UserInfo) GetUID() string {
	return u.Id
}

type TokenAuthN struct {
	store store.Store
}

func NewTokenAuthN(store store.Store) authn.Authenticator {
	return &TokenAuthN{store: store}
}

func (t *TokenAuthN) GetAuthentication(ctx context.Context, req *http.Request) (interface{}, error) {
	ah := req.Header.Get("Authorization")
	if ah == "" {
		return nil, errors.New("authorization not found")
	}
	return &authn.TokenAuthentication{Token: ah}, nil
}

func (t *TokenAuthN) DeleteAuthentication(ctx context.Context, authentication interface{}) error {
	key := fmt.Sprintf("session:%s", authentication.(*authn.TokenAuthentication).Token)
	return t.store.Delete(ctx, key)
}

func (t *TokenAuthN) Authenticate(ctx context.Context, authentication interface{}) (user.Info, error) {
	key := fmt.Sprintf("session:%s", authentication.(*authn.TokenAuthentication).Token)
	userId, err := t.store.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	_ = t.store.ExpireKey(ctx, key, time.Hour)
	return &UserInfo{Id: userId}, nil
}

func (t *TokenAuthN) AddAuthentication(ctx context.Context, req *http.Request, authentication interface{}) error {
	req.Header.Add("Authorization", authentication.(*authn.TokenAuthentication).Token)
	return nil
}

func (t *TokenAuthN) WriteAuthentication(ctx context.Context, authentication interface{}, userInfo user.Info) error {
	key := fmt.Sprintf("session:%s", authentication.(*authn.TokenAuthentication).Token)
	return t.store.Write(ctx, key, userInfo.GetUID())
}

func (t *TokenAuthN) AuthenticateFailedCB(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusForbidden)
	_, _ = w.Write([]byte("403 forbidden"))
}
