package restaurantservice

import (
	"context"
	"errors"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantService struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantService (store UpdateRestaurantStore) *updateRestaurantService {
	return &updateRestaurantService{ store: store }
}

func (service *updateRestaurantService) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) (error) {

	oldData, err := service.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := service.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}