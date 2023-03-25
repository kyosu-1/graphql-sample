package services

import (
	"context"

	"github.com/kyosu-1/graphql-sample/graph/db"
	"github.com/kyosu-1/graphql-sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *userService) GetUsers(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name), // where name = {引数のname}
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}