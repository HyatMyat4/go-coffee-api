package services

import (
	"context"
	"time"
)

type Coffee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roast     string    `json:"roast"`
	Image     string    `json:"image"`
	Region    string    `json:"region"`
	Price     float32   `json:"price"`
	GrindUnit int16     `json:"grind_unit"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (c *Coffee) GetCoffeeById(id string) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, name, roast, image, region, price, grind_unit, create_at, update_at FROM coffees WHERE id = $1`

	var coffee Coffee

	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&coffee.ID,
		&coffee.Name,
		&coffee.Roast,
		&coffee.Image,
		&coffee.Region,
		&coffee.Price,
		&coffee.GrindUnit,
		&coffee.CreateAt,
		&coffee.UpdateAt,
	)

	if err != nil {
		return nil, err
	}

	return &coffee, nil

}

func (c *Coffee) GetAllCoffees() ([]*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, name, roast, image, region, price, grind_unit, create_at, update_at FROM coffees`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var coffees []*Coffee

	for rows.Next() {
		var coffee Coffee
		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Roast,
			&coffee.Image,
			&coffee.Region,
			&coffee.Price,
			&coffee.GrindUnit,
			&coffee.CreateAt,
			&coffee.UpdateAt,
		)
		if err != nil {
			return nil, err
		}

		coffees = append(coffees, &coffee)

	}
	return coffees, nil
}

func (c *Coffee) CreateCoffee(coffee Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `INSERT INTO coffees (name, roast, image, region, price, grind_unit, create_at, update_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning *`

	_, err := db.ExecContext(
		ctx,
		query,
		coffee.Name,
		coffee.Roast,
		coffee.Image,
		coffee.Region,
		coffee.Price,
		coffee.GrindUnit,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return &coffee, nil
}

func (c *Coffee) UpdateCoffee(id string, body Coffee) (*Coffee, error) {

}
