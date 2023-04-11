package dao

import (
	"context"
	"fmt"
	"user/database"
	"user/internal/model"
)

type userDao struct {
	conn *database.DBConn
}

func NewUserDao(conn *database.DBConn) *userDao {
	return &userDao{
		conn: conn,
	}
}

func (d *userDao) Save(ctx context.Context, u *model.User) error {
	sql := fmt.Sprintf("insert into %s (name,gender) values ( ?, ?)", u.TableName())

	r, _ := d.conn.Conn.ExecCtx(ctx, sql, u.Name, u.Gender)

	u.Id, _ = r.LastInsertId()

	return nil
}
