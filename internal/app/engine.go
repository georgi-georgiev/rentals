package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab1.dev.cns.bg/topup/brm/docs"
	"go.uber.org/zap"
)

func NewGinEngine(conf *Config, logger *zap.Logger, handlers *Handlers) *gin.Engine {
	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(ErrorHandler(logger))
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(NoMethod)
	engine.NoRoute(NoRoute)

	engine.Use(cors.Default())

	docs.SwaggerInfo.Host = conf.Swagger.Host

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := engine.Group("")
	{
		group.GET("rentals/:rental_id", handlers.GetRentalById)
		group.GET("rentals", handlers.GetRentals)
	}

	return engine
}

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		containsError := false

		for _, err := range c.Errors {
			logger.Error(err.Error())
			containsError = true
		}

		if containsError {
			c.Header("Content-Type", "application/problem+json")
			InternalError(c)
		}
	}
}

func NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, NotFound())
}

func NoMethod(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, MethodNotAllowed())
}

func InternalError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError())
}

func NewHTTPErrorResponse(status int, message string) HTTPErrorResponse {
	return HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Title: message,
			},
		},
		Status: status,
	}
}

func BadRequest() HTTPErrorResponse {
	return NewHTTPErrorResponse(http.StatusBadRequest, "The request failed because it contained an invalid value or missing required value. The value could be a parameter value, a header value, or a property value.")
}

func MethodNotAllowed() HTTPErrorResponse {
	return NewHTTPErrorResponse(http.StatusMethodNotAllowed, "The HTTP method associated with the request is not supported.")
}

func NotFound() HTTPErrorResponse {
	return NewHTTPErrorResponse(http.StatusNotFound, "The requested operation failed because a resource associated with the request could not be found.")
}

func InternalServerError() HTTPErrorResponse {
	return NewHTTPErrorResponse(http.StatusInternalServerError, "The request failed due to an internal error.")
}

func ServiceUnavailable() HTTPErrorResponse {
	return NewHTTPErrorResponse(http.StatusServiceUnavailable, "Service Unavailable.")
}
