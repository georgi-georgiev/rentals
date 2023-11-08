package app

type RentalQueryParams struct {
	PriceMin int    `form:"price_min" example:"9000"`
	PriceMax int    `form:"price_max" example:"75000"`
	Limit    int    `form:"limit" minimum:"1" maximum:"100" default:"200"`
	Offset   int    `form:"offset" default:"0"`
	Ids      string `form:"ids" example:"1,2,3"`
	Near     string `form:"near" example:"33.64,-117.93"`
	Sort     string `form:"sort" default:"id"`
}
