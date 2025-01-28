package router

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitEmployeeRouter(rt *chi.Mux) {

	db := make(map[int]models.Employee)
	db[1] = models.Employee{
		Id:           1,
		CardNumberId: 123,
		FirstName:    "Alejandro",
		LastName:     "Salazar",
		WarehouseId:  1000,
	}
	repositoryImp := repository.NewEmployeeRepository(db)

	serviceImp := service.NewEmployeeService(repositoryImp)

	handler := handlers.NewEmployeeHandler(serviceImp)

	rt.Route("/employees", func(rt chi.Router) {
		rt.Get("/", handler.GetEmployees())
		rt.Get("/{id}", handler.GetEmployeeById())
	})
}
