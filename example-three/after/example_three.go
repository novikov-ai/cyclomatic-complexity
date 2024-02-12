package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
}

type UserService struct {
	Users map[int]*User
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AddUser(user *User) {
	if user == nil {
		fmt.Println("Error: User is nil")
		return
	}

	if user.Username == "" {
		fmt.Println("Error: Username is required")
		return
	}

	if user.Email == "" {
		fmt.Println("Error: Email is required")
		return
	}

	us.Users[user.ID] = user
	fmt.Println("User added successfully")
}

func (us *UserService) UpdateUser(id int, newUser *User) {
	if !validUser(newUser) {
		fmt.Println("User has invalid data")
		return
	}

	_, found := us.Users[id]
	if !found {
		fmt.Println("User not found")
		return
	}

	us.Users[id] = newUser
	fmt.Println("User updated successfully")
}

func validUser(newUser *User) bool {
	return newUser != nil && newUser.Username != "" && newUser.Email != ""
}

func (us *UserService) DeleteUser(id int) {
	for _, user := range us.Users {
		if user == nil {
			continue
		}

		if user.ID == id {
			delete(us.Users, user.ID)
			fmt.Println("User deleted successfully")
			return
		}
	}

	fmt.Println("Error: User not found")
}

func main() {
	userService := NewUserService()

	// Add user
	user1 := &User{ID: 1, Username: "user1", Email: "user1@example.com"}
	userService.AddUser(user1)

	// Update user
	user2 := &User{ID: 2, Username: "user2", Email: "user2@example.com"}
	userService.UpdateUser(2, user2)

	// Delete user
	userService.DeleteUser(3)
}
