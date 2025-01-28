package router

import (
	"github.com/D-Sorrow/meli-frescos/internal/repository"
	"github.com/D-Sorrow/meli-frescos/internal/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/D-Sorrow/meli-frescos/pkg/models"
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
	})
}
