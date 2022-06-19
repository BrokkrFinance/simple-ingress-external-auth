// Code generated by mockery v2.10.4. DO NOT EDIT.

package authmock

import (
	context "context"

	model "github.com/slok/simple-ingress-external-auth/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// TokenGetter is an autogenerated mock type for the TokenGetter type
type TokenGetter struct {
	mock.Mock
}

// GetToken provides a mock function with given fields: ctx, tokenValue
func (_m *TokenGetter) GetToken(ctx context.Context, tokenValue string) (*model.Token, error) {
	ret := _m.Called(ctx, tokenValue)

	var r0 *model.Token
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Token); ok {
		r0 = rf(ctx, tokenValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
