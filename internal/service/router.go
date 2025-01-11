package service

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/distributors-admin/internal/config"
	"github.com/recovery-flow/distributors-admin/internal/service/handlers"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	authMW := service.TokenManager.AuthMdl(service.Config.JWT.AccessToken.SecretKey)

	r.Route("/distributors-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Post("/create", handlers.DistributorCreate)

				r.Route("/distributors", func(r chi.Router) {
					r.Route("/{distributor_id}", func(r chi.Router) {
						r.Route("/update", func(r chi.Router) {
							r.Put("/name", handlers.DistributorUpdate)
						})
						r.Route("/employees", func(r chi.Router) {
							r.Route("/{user_id}", func(r chi.Router) {
								r.Patch("/", handlers.DistributorEmployeeUpdate)
								r.Delete("/", handlers.DistributorEmployeeDelete)
							})
							r.Post("/add", handlers.DistributorEmployeeAdd)
						})
					})
				})

				r.Route("places", func(r chi.Router) {
					r.Route("/{place_id}", func(r chi.Router) {
						r.Route("/employees", func(r chi.Router) {
							r.Route("/{user_id}", func(r chi.Router) {
								r.Patch("/", handlers.PlaceEmployeeUpdate)
								r.Delete("/", handlers.PlaceEmployeeDelete)
							})
							r.Post("/add", handlers.PlaceEmployeeAdd)
						})
					})
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/distributors", func(r chi.Router) {
					r.Route("/{distributor_id}", func(r chi.Router) {
						r.Route("/employees", func(r chi.Router) {
							r.Get("/", handlers.GetDistributorEmployees)
							r.Get("/{user_id}", handlers.GetDistributorEmployeesData)
						})
					})
				})

				r.Route("places", func(r chi.Router) {
					r.Route("/{place_id}", func(r chi.Router) {
						r.Route("/employees", func(r chi.Router) {
							r.Route("/{user_id}", func(r chi.Router) {
								r.Get("/", handlers.GetPlaceEmployee)
							})
							r.Get("/", handlers.GetPlacesEmployees)
						})
					})
				})
			})
		})

	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
