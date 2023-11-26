package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"time"
)

func GenerateFilename(originalFilename string) string {
	extension := filepath.Ext(originalFilename)
	filename := fmt.Sprintf("image_%d%s", time.Now().UnixNano(), extension)
	return filename
}

func GenerateImageID() (string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
