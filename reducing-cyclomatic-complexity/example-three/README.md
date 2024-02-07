# Before (Cyclomatic Complexity of UpdateUser = 7)

~~~
type UserService struct {
	Users []*User
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
~~~

# After (Cyclomatic Complexity of UpdateUser = 3 )

~~~
type UserService struct {
	Users map[int]*User
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
~~~

## Methods used

- removed nil check
- changed another data structure for Users
- removed for loop
- encapsulated processing logic into separated functions