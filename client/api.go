package restapi

import (
	"net/http"

	"github.com/Coop25/the-meme-index-api/client/controllers"
	"github.com/Coop25/the-meme-index-api/config"
	"github.com/Coop25/the-meme-index-api/managers"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type RestAPI struct {
	Config   *config.Config
	Router   *chi.Mux
	managers *managers.Managers
}

func New(config *config.Config, managers *managers.Managers) (*RestAPI, error) {
	var err error
	api := &RestAPI{
		Config:   config,
		managers: managers,
	}
	api.Router, err = api.newRouter()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (api *RestAPI) newRouter() (*chi.Mux, error) {
	router := chi.NewRouter()
	// Basic CORS settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{api.Config.Origin}, // Allow specific origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(c.Handler)

	router.Use(middleware.Logger)

	swagger, err := restapi.GetSwagger()
	if err != nil {
		return nil, err
	}

	// Serve Swagger UI
	router.Get("/swaggerui/*", api.serveSwaggerUI)

	swaggerJSON, err := swagger.MarshalJSON()
	if err != nil {
		return nil, err
	}

	// Serve OpenAPI specification
	router.Get("/swagger.json", api.serveOpenAPISpec(swaggerJSON))

	controller := controllers.New(api.Config, *api.managers)

	handler := restapi.HandlerWithOptions(controller, restapi.ChiServerOptions{})
	router.Mount("/", handler)
	return router, nil
}

func (api *RestAPI) serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui"))).ServeHTTP(w, r)
}

func (api *RestAPI) serveOpenAPISpec(swaggerJSON []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Write(swaggerJSON)
	}
}
