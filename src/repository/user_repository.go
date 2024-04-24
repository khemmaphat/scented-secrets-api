package repository

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/khemmaphat/scented-secrets-api/src/entities"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
)

type UserRepository struct {
	client *firestore.Client
}

func MakeUserRepository(client *firestore.Client) infRepo.IUserRepository {
	return &UserRepository{client: client}
}

func (r UserRepository) GetUserById(ctx context.Context, id string) (entities.User, error) {
	user := entities.User{}
	userDoc, err := r.client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return user, err
	}

	if err := userDoc.DataTo(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (r UserRepository) CrateUser(ctx context.Context, user entities.User) error {
	query, err := r.client.Collection("users").Where("Username", "==", user.Username).Documents(ctx).GetAll()

	if err != nil {
		return err
	}

	if len(query) != 0 {
		return errors.New("username is exists")
	}

	_, _, err = r.client.Collection("users").Add(context.Background(), user)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func (r UserRepository) LoginUser(ctx context.Context, user entities.User) (string, error) {
	doc, err := r.client.Collection("users").Where("Username", "==", user.Username).Documents(ctx).Next()

	if err != nil {
		return "", errors.New("invalid username or password")
	}

	password, ok := doc.Data()["Password"].(string)
	if !ok {
		return "", errors.New("password field not found in user document")
	}

	if password != user.Password {
		return "", errors.New("invalid username or password")
	}

	return doc.Ref.ID, err
}

func (r UserRepository) EditUser(ctx context.Context, id string, user entities.User) error {
	_, err := r.client.Collection("users").Doc(id).Set(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) UpdateNameUser(ctx context.Context, id string, name string) error {
	docRef := r.client.Collection("users").Doc(id)

	_, err := docRef.Update(ctx, []firestore.Update{
		{Path: "FirstName", Value: name},
	})

	if err != nil {
		return err
	}

	return nil
}
