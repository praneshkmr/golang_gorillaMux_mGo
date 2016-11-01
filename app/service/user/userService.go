package user

import (
	CounterDAO "./../../dao/mongo/counter"
	UserDAO "./../../dao/mongo/user"
	"./../../model"
)

func AddUser(user *model.User) (*model.User, error) {
	id, err := CounterDAO.GetNextCounter("user")
	if err != nil {
		return nil, err
	}
	user.ID = id
	user, err = UserDAO.AddUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(id int) (*model.User, error) {
	user, err := UserDAO.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserByID(id int, user *model.User) (*model.User, error) {
	user, err := UserDAO.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserByID(id int) (bool, error) {
	isDeleted, err := UserDAO.DeleteUser(id)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}
