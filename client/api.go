package restapi

import (
	"net/http"
	"github.com/Coop25/the-meme-index-api/client/controllers"
	"github.com/Coop25/the-meme-index-api/config"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type RestAPI struct {
	Config *config.Config
	Router *chi.Mux
}

func New(config *config.Config) *RestAPI {
	api := &RestAPI{
		Config: config,
	}
	api.Router = api.newRouter()
	return api
}

func (api *RestAPI) newRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Serve Swagger UI
	router.Get("/swagger/*", api.serveSwaggerUI)

	// Serve OpenAPI specification
	router.Get("/swagger.yaml", api.serveOpenAPISpec)

	controller := controllers.New(api.Config)

	handler := restapi.HandlerWithOptions(controller, restapi.ChiServerOptions{})
	router.Mount("/", handler)
	return router
}

func (api *RestAPI) serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/swagger/", http.FileServer(http.Dir("swaggerui"))).ServeHTTP(w, r)
}

func (api *RestAPI) serveOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger/swagger.yaml")
}
