package system

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
)

// CheckSum calculates the SHA-256 hex digest of a multipart file,
// then resets the file pointer to the beginning.
func CheckSum(file multipart.File) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
