package core

import "context"

type (
	PasswordUsecase interface {
		ChangePassword(context.Context, Password) error
	}

	PasswordRepository interface {
		GeneratePassword(context.Context, string) (string, error)
	}
)
