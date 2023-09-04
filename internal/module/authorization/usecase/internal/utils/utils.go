package utils

import token "monolith/internal/module/token/core"

func NewToken(access, refresh string) token.Token {
	return token.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}
