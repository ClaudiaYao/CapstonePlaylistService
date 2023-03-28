package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func convertRestaurantToDTO(res *Restaurant) (*RestaurantDTO, error) {
	var restDTO RestaurantDTO
	var operateHours [][]string
	err := json.Unmarshal([]byte(res.OperateHours), &operateHours)
	if err != nil {
		return nil, err
	}
	fmt.Println("unmarshal operation hours", operateHours)

	restDTO = RestaurantDTO{
		ID:           res.ID,
		Name:         res.Name,
		UnitNumber:   res.UnitNumber,
		AddressLine1: res.AddressLine1,
		AddressLine2: res.AddressLine2,
		PostalCode:   res.PostalCode,
		OperateHours: operateHours,
		Tag:          res.Tag,
		LogoUrl:      res.LogoUrl,
		HeaderUrl:    res.HeaderUrl,
	}
	return &restDTO, nil
}

func convertDishToDTO(dish *Dish) (*DishDTO, error) {
	var dishDTO DishDTO
	var dishOptions [][]string
	err := json.Unmarshal([]byte(dish.DishOptions), &dishOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("unmarshal operation hours", dishOptions)

	dishDTO = DishDTO{
		ID:           dish.ID,
		Name:         dish.Name,
		RestaurantID: dish.RestaurantID,
		Price:        dish.Price,
		CuisineStyle: dish.CuisineStyle,
		Ingredient:   dish.Ingredient,
		Comment:      dish.Comment,
		ServeTime:    dish.ServeTime,
		DishOptions:  dishOptions,
		ImageUrl:     dish.ImageUrl,
	}
	return &dishDTO, nil
}

func (app *PlaylistService) GetPlaylistInfo(ctx context.Context, playlistID string) ([]DishDTO, []RestaurantDTO, error) {

	dishIDs, err := app.DBConnection.GetDishesByPlaylistID(ctx, playlistID)

	if err != nil {
		return nil, nil, err
	}

	dishesDTO := []DishDTO{}
	restaurantsDTO := []RestaurantDTO{}

	for _, dishID := range dishIDs {
		dish, err := app.DBConnection.GetDishByID(ctx, dishID)
		if err != nil {
			return nil, nil, err
		}

		restaurant, err := app.DBConnection.GetRestaurantsByID(ctx, dish.RestaurantID)
		if err != nil {
			return nil, nil, err
		}

		dishDTO, err := convertDishToDTO(&dish)
		if err != nil {
			return nil, nil, err
		}

		resDTO, err := convertRestaurantToDTO(&restaurant)
		if err != nil {
			return nil, nil, err
		}
		dishesDTO = append(dishesDTO, *dishDTO)
		restaurantsDTO = append(restaurantsDTO, *resDTO)
	}

	return dishesDTO, restaurantsDTO, nil

}

func (app *PlaylistService) GetRestaurantInfo(ctx context.Context, restaurantID string) (*RestaurantResponseDataDTO, error) {

	restaurant, err := app.DBConnection.GetRestaurantsByID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	restDTO, err := convertRestaurantToDTO(&restaurant)

	dishes, err := app.DBConnection.GetDishesByRestaurantID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	dishesDTO := []DishDTO{}
	for _, dish := range dishes {
		dishDTO, err := convertDishToDTO(&dish)
		if err != nil {
			return nil, err
		}
		dishesDTO = append(dishesDTO, *dishDTO)
	}

	DTO := RestaurantResponseDataDTO{
		RestaurantInfo: *restDTO,
		DishIncluded:   dishesDTO,
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
		restDTO, err := convertRestaurantToDTO(&restaurant)
		if err != nil {
			return nil, err
		}

		dishes, err := app.DBConnection.GetDishesByRestaurantID(ctx, restaurant.ID)
		if err != nil {
			return nil, err
		}
		dishesDTO := []DishDTO{}
		for _, dish := range dishes {
			dishDTO, err := convertDishToDTO(&dish)
			if err != nil {
				return nil, err
			}

			dishesDTO = append(dishesDTO, *dishDTO)
		}
		DTO := RestaurantResponseDataDTO{
			RestaurantInfo: *restDTO,
			DishIncluded:   dishesDTO,
		}
		restaurantDTOs = append(restaurantDTOs, DTO)
	}

	return &restaurantDTOs, nil

}
