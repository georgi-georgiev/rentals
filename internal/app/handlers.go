package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handlers struct {
	repository *Repository
	logger     *zap.Logger
}

func NewHandlers(repository *Repository, logger *zap.Logger) *Handlers {
	return &Handlers{repository: repository, logger: logger}
}

func MapRentalModelToResponse(rental *Rental) RentalResponse {
	response := RentalResponse{
		ID:              rental.ID,
		Name:            rental.Name,
		Description:     rental.Description,
		Type:            rental.Type,
		Make:            rental.VehicleMake,
		Model:           rental.VehicleModel,
		Year:            rental.Created.Year(),
		Length:          rental.VehicleLength,
		Sleeps:          rental.Sleeps,
		PrimaryImageURL: rental.PrimaryImageURL,
		Price: PriceResponse{
			Day: rental.PricePerDay,
		},
		Location: LocationResponse{
			City:    rental.HomeCity,
			State:   rental.HomeState,
			Zip:     rental.HomeZIP,
			Country: rental.HomeCountry,
			Lat:     rental.Lat,
			Lng:     rental.Lng,
		},
	}

	if rental.User != nil {
		response.User = UserResponse{
			ID:        rental.User.ID,
			FirstName: rental.User.FirstName,
			LastName:  rental.User.LastName,
		}
	}

	return response
}

// GetRentalById godoc
// @Summary Get rental
// @Description by id
// @Accept  json
// @Produce  json
// @Param rental_id query int true "1"
// @Success 200 {object} RentalResponse
// @Failure 400  {object} HTTPErrorResponse
// @Failure 404  {object} HTTPErrorResponse
// @Failure 500  {object} HTTPErrorResponse
// @Router /rentals/{rental_id} [get]
func (h *Handlers) GetRentalById(c *gin.Context) {
	id := c.Param("rental_id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		h.logger.With(zap.Error(err)).Error("Field rental_id is invalid")
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest())
		return
	}

	rental, err := h.repository.SelectRentalById(c.Request.Context(), aid)
	if err != nil {
		h.logger.With(zap.Error(err)).Error("Could not select rental by id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError())
		return
	}

	if rental == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, NotFound())
		return
	}

	response := MapRentalModelToResponse(rental)
	c.JSON(http.StatusOK, response)
}

const MinimumLimit int = 1
const MaximumLimit int = 100

func DefaultRentalQueryParams() RentalQueryParams {
	return RentalQueryParams{
		Limit: 20,
		Sort:  "id",
	}
}

// GetRentals godoc
// @Summary Get all rentals
// @Description by query
// @Accept  json
// @Produce  json
// @Param price_min query int false "100"
// @Param price_max query int false "500"
// @Param limit query int false "10"
// @Param offset query int false "0"
// @Param ids query string false "1,2,3"
// @Param near query string false "40.7128,-74.0060"
// @Param sort query string false "name"
// @Success 200 {array} RentalResponse
// @Failure 400  {object} HTTPErrorResponse
// @Failure 404  {object} HTTPErrorResponse
// @Failure 500  {object} HTTPErrorResponse
// @Router /rentals [get]
func (h *Handlers) GetRentals(c *gin.Context) {
	queryParams := DefaultRentalQueryParams()
	err := c.BindQuery(&queryParams)
	if err != nil {
		h.logger.With(zap.Error(err)).Error("Could not bind query params")
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest())
		return
	}

	if queryParams.Limit < MinimumLimit || queryParams.Limit > MaximumLimit {
		httpError := NewHTTPErrorResponse(http.StatusBadRequest, "Limit is not in the allowed range [1,100]")
		c.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}

	validSortColumns := map[string]bool{
		"id":            true,
		"user_id":       true,
		"name":          true,
		"type":          true,
		"description":   true,
		"price_per_day": true,
	}

	if _, valid := validSortColumns[queryParams.Sort]; !valid {
		httpError := NewHTTPErrorResponse(http.StatusBadRequest, "Column name in sort parameter is not supported")
		c.AbortWithStatusJSON(http.StatusBadRequest, httpError)
		return
	}

	fmt.Println("params", queryParams)

	rentals, err := h.repository.SelectRentals(c.Request.Context(), queryParams)
	if err != nil {
		h.logger.With(zap.Error(err)).Error("Could not select rentals")
		c.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError())
		return
	}

	response := make([]RentalResponse, 0)

	for _, rental := range rentals {
		response = append(response, MapRentalModelToResponse(rental))
	}

	c.JSON(http.StatusOK, response)
}
