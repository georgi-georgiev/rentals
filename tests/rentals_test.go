package tests

import (
	"rentals/internal/app"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

var serverAddress string = "http://localhost:5353"

func TestFindById(t *testing.T) {
	result := app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals/1")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := app.RentalResponse{
		ID:              1,
		Name:            "'Abaco' VW Bay Window: Westfalia Pop-top",
		Description:     "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
		Type:            "camper-van",
		Make:            "Volkswagen",
		Model:           "Bay Window",
		Year:            2021,
		Length:          15,
		Sleeps:          4,
		PrimaryImageURL: "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
		Price: app.PriceResponse{
			Day: 16900,
		},
		Location: app.LocationResponse{
			City:    "Costa Mesa",
			State:   "CA",
			Zip:     "92627",
			Country: "US",
			Lat:     33.64,
			Lng:     -117.93,
		},
		User: app.UserResponse{
			ID:        1,
			FirstName: "John",
			LastName:  "Smith",
		},
	}

	require.EqualValues(t, expectedResult, result, "responses are not equal")

}

func TestSearchByPrice(t *testing.T) {
	result := []app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals?price_min=16800&price_max=16900")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := []app.RentalResponse{
		{
			ID:              1,
			Name:            "'Abaco' VW Bay Window: Westfalia Pop-top",
			Description:     "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
			Type:            "camper-van",
			Make:            "Volkswagen",
			Model:           "Bay Window",
			Year:            2021,
			Length:          15,
			Sleeps:          4,
			PrimaryImageURL: "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
			Price: app.PriceResponse{
				Day: 16900,
			},
			Location: app.LocationResponse{
				City:    "Costa Mesa",
				State:   "CA",
				Zip:     "92627",
				Country: "US",
				Lat:     33.64,
				Lng:     -117.93,
			},
			User: app.UserResponse{
				ID:        1,
				FirstName: "John",
				LastName:  "Smith",
			},
		},
	}

	require.EqualValues(t, expectedResult, result, "responses are not equal")
}

func TestSearchByPriceMissing(t *testing.T) {
	result := []app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals?price_min=0&price_max=1000")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := []app.RentalResponse{}

	require.EqualValues(t, expectedResult, result, "responses are not equal")
}

func TestPaging(t *testing.T) {
	result := []app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals?limit=1&offset=0")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := []app.RentalResponse{
		{
			ID:              1,
			Name:            "'Abaco' VW Bay Window: Westfalia Pop-top",
			Description:     "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
			Type:            "camper-van",
			Make:            "Volkswagen",
			Model:           "Bay Window",
			Year:            2021,
			Length:          15,
			Sleeps:          4,
			PrimaryImageURL: "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
			Price: app.PriceResponse{
				Day: 16900,
			},
			Location: app.LocationResponse{
				City:    "Costa Mesa",
				State:   "CA",
				Zip:     "92627",
				Country: "US",
				Lat:     33.64,
				Lng:     -117.93,
			},
			User: app.UserResponse{
				ID:        1,
				FirstName: "John",
				LastName:  "Smith",
			},
		},
	}

	require.EqualValues(t, expectedResult, result, "responses are not equal")
}

func TestSearchByIds(t *testing.T) {
	result := []app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals?ids=1")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := []app.RentalResponse{
		{
			ID:              1,
			Name:            "'Abaco' VW Bay Window: Westfalia Pop-top",
			Description:     "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
			Type:            "camper-van",
			Make:            "Volkswagen",
			Model:           "Bay Window",
			Year:            2021,
			Length:          15,
			Sleeps:          4,
			PrimaryImageURL: "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
			Price: app.PriceResponse{
				Day: 16900,
			},
			Location: app.LocationResponse{
				City:    "Costa Mesa",
				State:   "CA",
				Zip:     "92627",
				Country: "US",
				Lat:     33.64,
				Lng:     -117.93,
			},
			User: app.UserResponse{
				ID:        1,
				FirstName: "John",
				LastName:  "Smith",
			},
		},
	}

	require.EqualValues(t, expectedResult, result, "responses are not equal")
}

func TestSearchByLocation(t *testing.T) {
	result := []app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals?near=33.64,-117.93")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := app.RentalResponse{
		ID:              1,
		Name:            "'Abaco' VW Bay Window: Westfalia Pop-top",
		Description:     "ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi",
		Type:            "camper-van",
		Make:            "Volkswagen",
		Model:           "Bay Window",
		Year:            2021,
		Length:          15,
		Sleeps:          4,
		PrimaryImageURL: "https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg",
		Price: app.PriceResponse{
			Day: 16900,
		},
		Location: app.LocationResponse{
			City:    "Costa Mesa",
			State:   "CA",
			Zip:     "92627",
			Country: "US",
			Lat:     33.64,
			Lng:     -117.93,
		},
		User: app.UserResponse{
			ID:        1,
			FirstName: "John",
			LastName:  "Smith",
		},
	}

	require.EqualValues(t, expectedResult, result[0], "responses are not equal")
}

func TestSorting(t *testing.T) {
	result := []app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals?sort=description")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := app.RentalResponse{
		ID:              28,
		Name:            "sCAMPer X",
		Description:     "ac tellus phasellus ultrices nostra eros aenean metus ridiculus adipiscing habitant nulla cubilia tortor rhoncus quisque sem ultrices varius massa mollis congue praesent nam ante",
		Type:            "camper-van",
		Make:            "Ram",
		Model:           "Promaster",
		Year:            2021,
		Length:          19,
		Sleeps:          4,
		PrimaryImageURL: "https://res.cloudinary.com/outdoorsy/image/upload/v1589910541/p/rentals/156152/images/jvyvtqoeljadoizjjzag.jpg",
		Price: app.PriceResponse{
			Day: 17500,
		},
		Location: app.LocationResponse{
			City:    "Atlanta",
			State:   "GA",
			Zip:     "30310",
			Country: "US",
			Lat:     33.73,
			Lng:     -84.41,
		},
		User: app.UserResponse{
			ID:        3,
			FirstName: "Barry",
			LastName:  "Martin",
		},
	}

	require.EqualValues(t, expectedResult, result[0], "responses are not equal")
}

func TestSearchCombinationEmptyPage(t *testing.T) {
	result := []app.RentalResponse{}
	error := &app.HTTPErrorResponse{}
	client := resty.New()
	response, err := client.R().
		SetResult(&result).
		SetError(error).
		Get(serverAddress + "/rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price_per_day")

	if err != nil {
		t.FailNow()
	}

	if response.IsError() {
		t.FailNow()
	}

	expectedResult := []app.RentalResponse{}

	require.EqualValues(t, expectedResult, result, "responses are not equal")
}
