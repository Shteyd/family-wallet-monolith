package crypto

import (
	"crypto/sha256"
	"hash"
	"monolith/internal/domain"
	"monolith/pkg/byteconv"
	"sync"
)

type _CryptoManager struct {
	*sync.Mutex
	Hash         hash.Hash
	PasswordSalt []byte
}

func NewCryptoManager(passwordSalt string) domain.CryptoManager {
	return &_CryptoManager{
		Mutex:        new(sync.Mutex),
		Hash:         sha256.New(),
		PasswordSalt: byteconv.Bytes(passwordSalt),
	}
}

func (crypto *_CryptoManager) Encrypt(value string) (string, error) {
	crypto.Lock()
	defer crypto.Unlock()

	crypto.Hash.Write(byteconv.Bytes(value))
	encryptValue := crypto.Hash.Sum(crypto.PasswordSalt)

	return byteconv.String(encryptValue), nil
}
