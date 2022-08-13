// This file is safe to edit. Once it exists it will not be overwritten

package apis

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"hulhay-mall/internal/apis/operations"
	"hulhay-mall/internal/apis/operations/health"
	"hulhay-mall/internal/apis/operations/store"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/handlers"
	"hulhay-mall/internal/models"
)

//go:generate swagger generate server --target ../../../hulhay-mall-be --name HulhayMall --spec ../../api/swagger.yml --model-package internal/models --server-package internal/apis --principal interface{}

func configureFlags(api *operations.HulhayMallAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HulhayMallAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// GET HEALTH
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
		result, err := handlers.NewHandler().GetHealtcheck()
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return health.NewGetHealthDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return health.NewGetHealthOK().WithPayload(&health.GetHealthOKBody{
			Data: &models.Health{
				Status: result.Status,
			},
			Message: "Success",
		})
	})

	// POST STORE
	api.StorePostStoresHandler = store.PostStoresHandlerFunc(func(params store.PostStoresParams) middleware.Responder {
		err := handlers.NewHandler().CreateStore(context.Background(), params.Identifier, params.Body)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return store.NewPostStoresDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return store.NewPostStoresCreated().WithPayload(&store.PostStoresCreatedBody{
			Message: "Create data successfully",
		})
	})

	// GET STORE
	api.StoreGetStoresHandler = store.GetStoresHandlerFunc(func(params store.GetStoresParams) middleware.Responder {
		result, err := handlers.NewHandler().GetStores(context.Background(), params.Identifier)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return store.NewGetStoresDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return store.NewGetStoresOK().WithPayload(&store.GetStoresOKBody{
			Message: "Success",
			Data:    result,
		})
	})

	// GET STORE BY ID
	api.StoreGetStoresIDHandler = store.GetStoresIDHandlerFunc(func(params store.GetStoresIDParams) middleware.Responder {
		result, err := handlers.NewHandler().GetStoreByID(context.Background(), params.Identifier, params.ID)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return store.NewGetStoresIDDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return store.NewGetStoresIDOK().WithPayload(&store.GetStoresIDOKBody{
			Message: "Success",
			Data:    result,
		})
	})

	// DELETE STORE BY ID
	api.StoreDeleteStoresIDHandler = store.DeleteStoresIDHandlerFunc(func(params store.DeleteStoresIDParams) middleware.Responder {
		err := handlers.NewHandler().DeleteStoreByID(context.Background(), params.Identifier, params.ID)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return store.NewDeleteStoresIDDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return store.NewDeleteStoresIDOK().WithPayload(&store.DeleteStoresIDOKBody{
			Message: "Delete data successfully",
		})
	})

	// UPDATE STORE BY ID
	api.StorePatchStoresIDHandler = store.PatchStoresIDHandlerFunc(func(params store.PatchStoresIDParams) middleware.Responder {
		err := handlers.NewHandler().UpdateStoreByID(context.Background(), params.Identifier, params.Body, params.ID)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return store.NewPatchStoresIDDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return store.NewPatchStoresIDOK().WithPayload(&store.PatchStoresIDOKBody{
			Message: "Update data successfully",
		})
	})

	// REGISTER
	api.UserPostRegisterHandler = user.PostRegisterHandlerFunc(func(params user.PostRegisterParams) middleware.Responder {
		err := handlers.NewHandler().Register(context.Background(), params.Body)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return user.NewPostRegisterDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return user.NewPostRegisterCreated().WithPayload(&user.PostRegisterCreatedBody{
			Message: "Success Register New Account",
		})
	})

	// LOGIN
	api.UserPatchLoginHandler = user.PatchLoginHandlerFunc(func(params user.PatchLoginParams) middleware.Responder {
		result, err := handlers.NewHandler().Login(context.Background(), params.Body)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return user.NewPatchLoginDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return user.NewPatchLoginOK().WithPayload(&user.PatchLoginOKBody{
			Message: "Login Succcessfully",
			Data:    result,
		})
	})

	// LOGOUT
	api.UserPatchLogoutHandler = user.PatchLogoutHandlerFunc(func(params user.PatchLogoutParams) middleware.Responder {
		err := handlers.NewHandler().Logout(context.Background(), &params)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return user.NewPatchLogoutDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return user.NewPatchLogoutOK().WithPayload(&user.PatchLogoutOKBody{
			Message: "Logout Successfully",
		})
	})

	// GET PROFILE BY IDENTIFIER
	api.UserGetProfileHandler = user.GetProfileHandlerFunc(func(params user.GetProfileParams) middleware.Responder {
		result, err := handlers.NewHandler().GetProfile(context.Background(), &params)
		if err != nil {
			var errorMessage = new(string)
			*errorMessage = err.Error()
			return user.NewGetProfileDefault(400).WithPayload(&models.Error{Code: "400", Message: *errorMessage})
		}

		return user.NewGetProfileOK().WithPayload(&user.GetProfileOKBody{
			Message: "Success get profile",
			Data:    result,
		})
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
