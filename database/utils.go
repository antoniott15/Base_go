package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"gotest"
	"github.com/rs/xid"
	"time"
)

func (r *Repository) CreateNewUser(user *gotest.User) error {
	if user.ID == "" {
		user.ID = xid.New().String()
	}

	user.CreateAt = time.Now()

	_, err := r.user.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}


func (r *Repository) GetUserByID(id string) (*gotest.User, error) {
	user := new (gotest.User)

	if err := r.user.FindOne(context.Background(), bson.M{"id": id}).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}


func (r *Repository) AllUsers() ([]*gotest.User, error) {

	iter, err := r.user.Find(context.Background(), bson.M{})
	if err != nil {
		 return nil,err
	}

	users := make([]*gotest.User, 0)
	for iter.Next(context.Background()) {
		user := new (gotest.User)
		if err := iter.Decode(user); err != nil {
			m := bson.M{}
			if err := iter.Decode(&m); err != nil {
				continue
			}
			user.ID, _ = m["id"].(string)
		}

		users = append(users, user)
	}

	return users,nil
}




func (r *Repository) UpdateUserByID(user *gotest.User) error {

	_, err := r.user.UpdateOne(context.Background(),bson.M{"id": user.ID}, user)
	if err != nil {
		return  err
	}

	return nil
}


func (r *Repository) DeleteUserByID(id string) error {

	_, err := r.user.DeleteOne(context.Background(),bson.M{"id": id})

	if err != nil {
		return err
	}

	return nil
}