package model

import (
	"encoding/json"
	"os"
)

type UserFile struct {
	Users map[string]*User
}

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	IsDeleted bool   `json:"is_deleted"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ImageDetail struct {
	FileName string `json:"file_name"`
	Labels   string `json:"labels"`
}

type EditImageRequest struct {
	FileName string `form:"file_name"`
	Labels   string `form:"labels"`
}

func LoadUsersFromFile(filename string) (*UserFile, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return &UserFile{Users: make(map[string]*User)}, nil
	}

	var users map[string]*User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}

	return &UserFile{Users: users}, nil
}

func SaveUsersToFile(filename string, userDB *UserFile) error {
	file, err := json.MarshalIndent(userDB.Users, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
