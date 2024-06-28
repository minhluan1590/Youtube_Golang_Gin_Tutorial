package services

import (
    "errors"
)

var users = map[string]string{
    "user1": "password1",
    "user2": "password2",
}

func Login(username, password string) (string, error) {
    // Check if the user exists and the password is correct
    if storedPassword, exists := users[username]; exists && storedPassword == password {
        // Generate JWT token
        token, err := GenerateToken(username)
        if err != nil {
            return "", err
        }
        return token, nil
    }

    return "", errors.New("invalid credentials")
}
