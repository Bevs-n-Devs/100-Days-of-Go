package main

/*
UserRepository defines a GetUser method that retrieves a user by ID.
UserService depends on UserRepository to get user information, allowing us to pass any type that implements UserRepository.
*/

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

/*
To make the UserService even more flexible, you can split interfaces into smaller ones when needed.

Example: Adding a SaveUser Method
Suppose you also want to add a method to save a user in the repository. Instead of modifying the existing interface, create a new one

With two interfaces (UserRepository and UserSaver), you can mock them separately and test specific parts of UserService without affecting other functionalities.
*/
// type UserSaver interface {
// 	SaveUser(user *User) error
// }

// type UserService2 struct {
// 	repo  UserRepository
// 	saver UserSaver
// }

// func NewUserService2(repo UserRepository, saver UserSaver) *UserService2 {
// 	return &UserService2{repo: repo, saver: saver}
// }

// func (s *UserService2) SaveNewUser(id int, name string) error {
// 	user := &User{ID: id, Name: name}
// 	return s.saver.SaveUser(user)
// }
