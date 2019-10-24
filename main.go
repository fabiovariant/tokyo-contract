package main

import (
	"context"
	"net/http"

	"github.com/fabiovariant/tokyo-commons/database"
	d "github.com/fabiovariant/tokyo-contracts/delivery"
	"github.com/fabiovariant/tokyo-contracts/repository/sql"
	"github.com/fabiovariant/tokyo-contracts/service"
	"github.com/urfave/negroni"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			newMux,
			database.GetPostgresqlConn,
			sql.NewClientContractsSQLRepository,
			service.NewClientContractService,
			d.NewClientContractHTTPDelivery,
		),
		fx.Invoke(server),
	)
	app.Run()
}

func server(mux *http.ServeMux, cd d.ClientContractsDelivery) {
	r := d.GetMuxRoutes(cd)
	mux.Handle("/", r)
}

func newMux(lc fx.Lifecycle) *http.ServeMux {

	mux := http.NewServeMux()
	n := negroni.Classic()
	n.UseHandler(mux)
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return mux
}
