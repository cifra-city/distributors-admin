package service

import (
	"context"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/service/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	authMW := service.TokenManager.AuthMiddleware(service.Config.JWT.AccessToken.SecretKey)

	r.Route("/distributors-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Post("/create", handlers.DistributorCreate)

				r.Route("/distributors/{distributor_id}", func(r chi.Router) {
					r.Route("/update", func(r chi.Router) {
						r.Put("/name", handlers.DistributorUpdate)
					})
					r.Route("/employees", func(r chi.Router) {
						r.Post("/add", handlers.EmployeeAdd)
						r.Put("/update", handlers.EmployeeUpdate)
						r.Delete("/delete", handlers.EmployeeDelete)
					})
				})
			})
		})

	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
