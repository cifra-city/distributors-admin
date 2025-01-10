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
						r.Post("/add", handlers.DistributorEmployeeAdd)
						r.Put("/update/{user_id}", handlers.DistributorEmployeeUpdate)
						r.Delete("/delete/{user_id}", handlers.DistributorEmployeeDelete)
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
								r.Get("/{user_id}", handlers.GetPlaceEmployee)
								r.Patch("/{user_id}", handlers.UpdatePlaceEmployee)
								r.Delete("/{user_id}", handlers.DeletePlaceEmployee)
							})

							r.Post("/add", handlers.AddPlaceEmployee)
							r.Get("/", handlers.GetPlacesEmployees)
						})
					})

					r.Post("/create", nil)
				})
			})
		})

	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
