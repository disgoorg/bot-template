package types

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/disgoorg/bot-template/internal/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func (b *Bot) SetupDatabase(shouldSyncDBTables bool) error {
	sqlDB := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", b.Config.Database.Host, b.Config.Database.Port)),
		pgdriver.WithUser(b.Config.Database.User),
		pgdriver.WithPassword(b.Config.Database.Password),
		pgdriver.WithDatabase(b.Config.Database.DBName),
		pgdriver.WithInsecure(true),
	))
	b.DB = bun.NewDB(sqlDB, pgdialect.New())

	if shouldSyncDBTables {
		if err := b.DB.ResetModel(context.TODO(), (*models.Idk)(nil)); err != nil {
			return err
		}
	}
	return nil
}
