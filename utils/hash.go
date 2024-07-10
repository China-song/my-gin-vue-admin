package utils

import "golang.org/x/crypto/bcrypt"

// BcryptCheck 对比明文密码和数据库的哈希密码
func BcryptCheck(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
