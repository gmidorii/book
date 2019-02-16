//+build wireinject

package injector

import (
	"github.com/google/wire"

	"github.com/gmidorii/book/server/app"
	"github.com/gmidorii/book/server/infra/db"
)

func InitBookApp() (app.Book, error) {
	wire.Build(db.NewBook, app.NewBook)
	return app.Book{}, nil
}
