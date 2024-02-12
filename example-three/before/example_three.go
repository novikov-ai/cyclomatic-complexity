package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
}

type UserService struct {
	Users []*User
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

	us.Users = append(us.Users, user)
	fmt.Println("User added successfully")
}

func (us *UserService) UpdateUser(id int, newUser *User) {
	if newUser == nil {
		fmt.Println("Error: New user is nil")
		return
	}

	if newUser.Username == "" {
		fmt.Println("Error: Username is required")
		return
	}

	if newUser.Email == "" {
		fmt.Println("Error: Email is required")
		return
	}

	for i, user := range us.Users {
		if user == nil {
			continue
		}

		if user.ID == id {
			us.Users[i] = newUser
			fmt.Println("User updated successfully")
			return
		}
	}

	fmt.Println("Error: User not found")
}

func (us *UserService) DeleteUser(id int) {
	for i, user := range us.Users {
		if user == nil {
			continue
		}

		if user.ID == id {
			us.Users = append(us.Users[:i], us.Users[i+1:]...)
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
