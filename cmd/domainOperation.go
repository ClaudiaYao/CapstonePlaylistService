package main

import "context"

func (app *PlaylistService) GetPlaylistInfo(ctx context.Context, playlistID string) ([]Dish, []Restaurant, error) {
	dishes, err := app.DBConnection.GetDishesByPlaylistID(ctx, playlistID)

	if err != nil {
		return nil, nil, err
	}

	restaurants := []Restaurant{}

	for _, dish := range dishes {
		restaurant, err := app.DBConnection.GetRestaurantsByID(ctx, dish.RestaurantID)

		if err != nil {
			return nil, nil, err
		}
		restaurants = append(restaurants, restaurant)
	}

	return dishes, restaurants, nil

}
