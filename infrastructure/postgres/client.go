package postgres

import (
	"context"

	"sqe-otp/config"

	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService interface {
	Insert(context.Context, any) error
	QueryByCondition(context.Context, map[string]any, any) error
	UpdateByCondition(context.Context, map[string]any, any) error
	Shutdown()
}

type client struct {
	instance *gorm.DB
}

func New(cfg config.Database) DatabaseService {
	primary := cfg.PrimaryDsn()
	conn, err := gorm.Open(driver.Open(primary), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	db := &client{
		instance: conn,
	}
	return db
}

func (c *client) Insert(ctx context.Context, table any) error {
	return c.instance.WithContext(ctx).Create(table).Error
}

func (c *client) QueryByCondition(ctx context.Context, conditions map[string]any, table any) error {
	return c.instance.WithContext(ctx).Where(conditions).Take(table).Error
}

func (c *client) UpdateByCondition(ctx context.Context, conditions map[string]any, table any) error {
	return c.instance.WithContext(ctx).Where(conditions).Updates(table).Error
}

func (c *client) Shutdown() {
	conn, _ := c.instance.DB()
	conn.Close()
}
