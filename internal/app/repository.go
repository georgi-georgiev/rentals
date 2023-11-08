package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(dbpool *pgxpool.Pool) *Repository {
	return &Repository{db: dbpool}
}

func (r *Repository) SelectRentalById(ctx context.Context, rentalId int) (*Rental, error) {
	rental := &Rental{}
	rental.User = &User{}
	id := rentalId

	err := r.db.QueryRow(ctx, `SELECT r.id, r.user_id, r.name, r.type, r.description, r.sleeps, r.price_per_day,
		r.home_city, r.home_state, r.home_zip, r.home_country, r.vehicle_make, r.vehicle_model, r.vehicle_year,
		r.vehicle_length, r.created, r.updated, r.lat, r.lng, r.primary_image_url, u.id, u.first_name, u.last_name FROM rentals r
		LEFT JOIN users u ON r.user_id = u.id
		WHERE r.id = $1`, id).Scan(
		&rental.ID, &rental.UserID, &rental.Name, &rental.Type, &rental.Description,
		&rental.Sleeps, &rental.PricePerDay, &rental.HomeCity, &rental.HomeState,
		&rental.HomeZIP, &rental.HomeCountry, &rental.VehicleMake, &rental.VehicleModel,
		&rental.VehicleYear, &rental.VehicleLength, &rental.Created, &rental.Updated,
		&rental.Lat, &rental.Lng, &rental.PrimaryImageURL, &rental.User.ID, &rental.User.FirstName, &rental.User.LastName,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return rental, nil
}

func (r *Repository) SelectRentals(ctx context.Context, queryParams RentalQueryParams) ([]*Rental, error) {

	dq := NewDynamicQuery(`SELECT r.id, r.user_id, r.name, r.type, r.description, r.sleeps, r.price_per_day,
	r.home_city, r.home_state, r.home_zip, r.home_country, r.vehicle_make, r.vehicle_model, r.vehicle_year,
	r.vehicle_length, r.created, r.updated, r.lat, r.lng, r.primary_image_url, u.id, u.first_name, u.last_name FROM rentals r
	LEFT JOIN users u ON r.user_id = u.id`)

	if queryParams.PriceMin > 0 {
		dq.Where("r.price_per_day", ">=", queryParams.PriceMin)
	}

	if queryParams.PriceMax > 0 {
		dq.Where("r.price_per_day", "<=", queryParams.PriceMax)
	}

	if queryParams.Ids != "" {
		ids := strings.Split(queryParams.Ids, ",")
		dq.Contains("r.id", ids)
	}

	if queryParams.Near != "" {
		near := strings.Split(queryParams.Near, ",")
		if len(near) == 2 {
			latPlaceholder := dq.AddArgument(near[0])
			lngPlaceholder := dq.AddArgument(near[1])
			dq.Where(fmt.Sprintf(`3959 * acos (
				cos ( radians(%s) )
				* cos( radians( r.lat ) )
				* cos( radians( r.lng ) - radians(%s) )
				+ sin ( radians(%s) )
				* sin( radians( r.lat ) )
			  )`, latPlaceholder, lngPlaceholder, latPlaceholder), "<", "100")
		}
	}

	if queryParams.Sort != "" {
		dq.Order(fmt.Sprintf("r.%s", queryParams.Sort))
	}

	if queryParams.Limit > 0 {
		dq.Limit(queryParams.Limit)
	}

	if queryParams.Offset > 0 {
		dq.Offset(queryParams.Offset)
	}

	fmt.Println("query", dq.query.String())
	fmt.Println("args", dq.args)

	rows, err := r.db.Query(ctx, dq.query.String(), dq.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentals []*Rental
	for rows.Next() {
		rental := &Rental{}
		rental.User = &User{}
		err := rows.Scan(
			&rental.ID, &rental.UserID, &rental.Name, &rental.Type, &rental.Description, &rental.Sleeps,
			&rental.PricePerDay, &rental.HomeCity, &rental.HomeState, &rental.HomeZIP, &rental.HomeCountry,
			&rental.VehicleMake, &rental.VehicleModel, &rental.VehicleYear, &rental.VehicleLength, &rental.Created,
			&rental.Updated, &rental.Lat, &rental.Lng, &rental.PrimaryImageURL, &rental.User.ID, &rental.User.FirstName, &rental.User.LastName,
		)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rentals, nil
}
