package app

import "encoding/json"

type HTTPError struct {
	Title      string `json:"title,omitempty" example:"required parameters are missing"`
	Detail     string `json:"detail,omitempty" example:"The parameters: limit, date were not provided"`
	ReasonCode int    `json:"reason_code,omitempty" example:"150"`
	Reason     string `json:"reason,omitempty" example:"invalidParameter"`
	Placement  string `json:"placement,omitempty" example:"field"`
	Field      string `json:"field,omitempty" example:"email"`
	Expression string `json:"expression,omitempty" example:"greater"`
	Argument   string `json:"argument,omitempty" example:"number"`
}

func (e HTTPError) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(bytes)
}

type HTTPErrorResponse struct {
	Errors  []HTTPError `json:"errors"`
	Status  int         `json:"status,omitempty" example:"400"`
	TraceId string      `json:"trace_id,omitempty" example:"EJplcsCHuLu"`
}

type PriceResponse struct {
	Day int64 `json:"day"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LocationResponse struct {
	City    string  `json:"city"`
	State   string  `json:"state"`
	Zip     string  `json:"zip"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type RentalResponse struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Type            string           `json:"type"`
	Make            string           `json:"make"`
	Model           string           `json:"model"`
	Year            int              `json:"year"`
	Length          float64          `json:"length"`
	Sleeps          int              `json:"sleeps"`
	PrimaryImageURL string           `json:"primary_image_url"`
	Price           PriceResponse    `json:"price"`
	Location        LocationResponse `json:"location"`
	User            UserResponse     `json:"user"`
}
