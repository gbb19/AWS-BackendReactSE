package services

import (
	"onez19/config"
	"onez19/models"
)

func GetAllUsers() ([]models.UserResponse, error) {
	var users []models.UserResponse

	// ดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := config.DB.Query("SELECT username, first_name, last_name FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.Username, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}