package main

import "context"

func (app *PlaylistService) GetPlaylistInfo(ctx context.Context, playlistID string) ([]Dish, []Restaurant, error) {

	dishIDs, err := app.DBConnection.GetDishesByPlaylistID(ctx, playlistID)

	if err != nil {
		return nil, nil, err
	}

	dishes := []Dish{}
	restaurants := []Restaurant{}

	for _, dishID := range dishIDs {
		dish, err := app.DBConnection.GetDishByID(ctx, dishID)
		restaurant, err := app.DBConnection.GetRestaurantsByID(ctx, dish.RestaurantID)

		if err != nil {
			return nil, nil, err
		}
		dishes = append(dishes, dish)
		restaurants = append(restaurants, restaurant)
	}

	return dishes, restaurants, nil

}

func (app *PlaylistService) GetRestaurantInfo(ctx context.Context, restaurantID string) (*RestaurantResponseDataDTO, error) {

	restaurant, err := app.DBConnection.GetRestaurantsByID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	dishes, err := app.DBConnection.GetDishesByRestaurantID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	DTO := RestaurantResponseDataDTO{
		RestaurantInfo: restaurant,
		DishIncluded:   dishes,
	}
	return &DTO, nil

}

func (app *PlaylistService) GetAllRestaurantInfo(ctx context.Context) (*[]RestaurantResponseDataDTO, error) {

	restaurants, err := app.DBConnection.GetAllRestaurants(ctx)
	if err != nil {
		return nil, err
	}

	restaurantDTOs := []RestaurantResponseDataDTO{}
	for _, restaurant := range restaurants {
		dishes, err := app.DBConnection.GetDishesByRestaurantID(ctx, restaurant.ID)

		if err != nil {
			return nil, err
		}
		DTO := RestaurantResponseDataDTO{
			RestaurantInfo: restaurant,
			DishIncluded:   dishes,
		}
		restaurantDTOs = append(restaurantDTOs, DTO)
	}

	return &restaurantDTOs, nil

}
