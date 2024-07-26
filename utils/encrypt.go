package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

// CBCEncrypted AES CBC 加密
func CBCEncrypted(buf []byte, keyStr string) ([]byte, error) {
	key := sha256.Sum256([]byte(keyStr))
	plaintext := standardizeDataEn(buf)

	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key[:sha256.Size])
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func standardizeDataEn(data []byte) []byte {
	appendingLen := aes.BlockSize - (len(data) % aes.BlockSize)
	sd := make([]byte, len(data)+appendingLen)
	copy(sd, data)
	for i := 0; i < appendingLen; i++ {
		sd[i+len(data)] = byte(appendingLen)
	}
	return sd
}
