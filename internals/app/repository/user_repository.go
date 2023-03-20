package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"security-service/internals/app/entity"
)

// UserRepository Слой для работы с таблицой пользователя entity.User
type UserRepository struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

var log = logrus.New()

// NewUserRepository Создание нового репозитория
func NewUserRepository(pool *pgxpool.Pool, ctx context.Context) *UserRepository {
	return &UserRepository{pool: pool, ctx: ctx}
}

// Save Сохранение пользователя
func (r *UserRepository) Save(user entity.User) error {
	tx, err := r.pool.Begin(r.ctx)

	query := "INSERT INTO users(uuid, name, password, number) VALUES($1, $2, $3, $4)"

	_, err = tx.Exec(r.ctx, query, user.Uuid, user.Name, user.Password, user.Number)
	if err != nil {
		log.Errorln("Error save user. Error: ", err)

		go func() {
			err = tx.Rollback(r.ctx)
			if err != nil {
				log.Errorln("Error rollback transaction. Error: ", err)
			}
		}()

		return err
	}

	err = tx.Commit(r.ctx)
	if err != nil {
		log.Errorln("Error commit transaction. Error: ", err)
	}

	return err
}

// FindUserByNumber Поиск пользователя по номеру телефона
func (r *UserRepository) FindUserByNumber(number string) (entity.User, error) {
	var user entity.User

	query := "SELECT uuid, password FROM users WHERE number = $1"

	err := pgxscan.Select(r.ctx, r.pool, &user, query, number)
	if err != nil {
		log.Errorln("Error find user password. Error: ", err)
		return entity.User{}, err
	}

	return user, nil
}
