package service

import (
	"fmt"
	"path/filepath"
	"time"
)

func GenerateFilename(originalFilename string) string {
	extension := filepath.Ext(originalFilename)
	filename := fmt.Sprintf("image_%d%s", time.Now().UnixNano(), extension)
	return filename
}
