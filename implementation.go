package gotest

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2"
	proto "gotest/service"
	"time"
)



func (api *Api) CreateNewUser(ctx context.Context,req *proto.UserInput,res *proto.UserResponse) error {
	user := &User{
		ID:           "",
		Profile:      Profile{
			Username: req.Username,
			Password: req.Password,
			Email:    "",
			Phone:    req.Phone,
		},
		PersonalInfo: PersonalInfo{
			Name:     req.Name,
			LastName: req.LastName,
			Address:  req.Address,
			Age:      "",
			DNI:      req.Dni,
		},
		CreateAt:     time.Now(),
	}

	if err := api.repo.CreateNewUser(user); err != nil {
		return err
	}

	u,err := json.Marshal(user)
	if err != nil{
		return err
	}


	newUser := new(proto.UserInput)
	if err = json.Unmarshal(u,newUser); err != nil {
		return err
	}

	res.User = newUser
	return nil
}



func (api *Api) DeleteUser(ctx context.Context,req *proto.UserID,res *proto.Response) error {
	if err := api.repo.DeleteUserByID(req.Id); err != nil{
		return err
	}
	res.Message = "User with "+ req.Id+ "deleted succesfully"
	return nil
}

func (api *Api) UpdateUser(ctx context.Context, req *proto.UserInput,res *proto.Response) error {
	user := &User{
		ID:           "",
		Profile:      Profile{
			Username: req.Username,
			Password: req.Password,
			Email:    "",
			Phone:    req.Phone,
		},
		PersonalInfo: PersonalInfo{
			Name:     req.Name,
			LastName: req.LastName,
			Address:  req.Address,
			Age:      "",
			DNI:      req.Dni,
		},
		CreateAt:     time.Now(),
	}

	if err := api.repo.UpdateUserByID(user); err != nil {
		return err
	}


	res.Message = "Update user"
	return nil
}

func (api *Api) GetUser(ctx context.Context,req *proto.UserID,res *proto.UserResponse) error {

	user, err :=  api.repo.GetUserByID(req.Id)
	if err != nil{
		return err
	}

	u,err := json.Marshal(user)
	if err != nil{
		return err
	}

	newUser := new(proto.UserInput)
	if err = json.Unmarshal(u,newUser); err != nil {
		return err
	}

	res.User = newUser

	return nil
}




func (api *Api) ConnectWithGRPC() error {
	service :=  micro.NewService(micro.Name("proto"))
	service.Init()

	if err := proto.RegisterCrudUserHandler(service.Server(), new(Api)); err != nil {
		return err
	}

	if err := service.Run(); err != nil {
		return err
	}

	return nil
}