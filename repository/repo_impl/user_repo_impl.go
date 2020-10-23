package repo_impl

import "C"
import (
	"backend-github-trending/banana"
	"backend-github-trending/db"
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"backend-github-trending/repository"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		//log ra
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}
	return user, nil
}

func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req.RequestSignIn) (model.User, error) {
	var user = model.User{}

	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", loginReq.Email)

	if err != nil {
		//log err
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		return user, err
	}
	return user, nil
}
func (u *UserRepoImpl) SelectUserById(context context.Context, userId string) (model.User, error){
	var user model.User
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE user_id= $1", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		//log error
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	sqlStatement := `
		UPDATE users
		SET
			full_name = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			updated_at = COALESCE (:updated_at, updated_at)
		WHERE user_id = :user_id
	`
	user.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		//log err
		return user, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		//log err
		return user, banana.UserNotUpdated
	}
	if count == 0 {
		return user, banana.UserNotUpdated
	}

	return user, nil

}