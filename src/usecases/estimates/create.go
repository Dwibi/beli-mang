package estimatesusecase

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/Dwibi/beli-mang/src/helpers"
	"github.com/gofiber/fiber/v2"
)

type CreateEstimateParams struct {
	UserId        int
	EstimatesBody entities.CreateEstimateParams
}

func (i sEstimatesUseCase) Create(p CreateEstimateParams) (*entities.ResultEstimate, int, error) {
	// validate user doing the request is user
	userRole, err := i.userRepository.FindUserRole(p.UserId)

	if err != nil {
		fmt.Println("error 1")
		return nil, fiber.StatusInternalServerError, err
	}

	if userRole != "user" {
		fmt.Println("error 2")
		return nil, fiber.StatusUnauthorized, errors.New("only user can use this route")
	}

	// Get all merchantId and itemId
	merchantIds := []int{}
	itemIds := []int{}
	var startingPointMerchant string

	for _, merchant := range p.EstimatesBody.Orders {
		if merchant.IsStartingPoint {
			startingPointMerchant = merchant.MerchantId
		} else {
			merchantIdInt, _ := strconv.Atoi(merchant.MerchantId)
			merchantIds = append(merchantIds, merchantIdInt)
		}
		for _, item := range merchant.Items {
			itemIdInt, _ := strconv.Atoi(item.ItemId)
			itemIds = append(itemIds, itemIdInt)
		}
	}

	// Validate if all the merchant id is exist
	missingMerchantId, err := i.merchantrepository.GetMissingIDs(merchantIds)

	if err != nil {
		fmt.Println("error 3")
		return nil, fiber.StatusInternalServerError, err
	}

	if len(missingMerchantId) > 0 {
		fmt.Println("error 4")
		return nil, fiber.StatusNotFound, fmt.Errorf("the following IDs do not exist in the merchants table: %v", missingMerchantId)
	}

	// Validate if all the item id is exist
	missingItemId, err := i.itemsrepository.GetMissingIDs(itemIds)

	if err != nil {
		fmt.Println("error 5")
		return nil, fiber.StatusInternalServerError, err
	}

	if len(missingItemId) > 0 {
		fmt.Println("error 6")
		return nil, fiber.StatusNotFound, fmt.Errorf("the following IDs do not exist in the items table: %v", missingItemId)
	}

	// Get first startingPointMerchant lat and long
	merchant, err := i.merchantrepository.FindMany(&entities.SearchMerchantParams{
		MerchantId: startingPointMerchant,
		Limit:      1,
	})

	if err != nil {
		fmt.Println("error 7")
		return nil, fiber.StatusInternalServerError, err
	}

	startingPointLat := merchant.Data[0].Location.Lat
	startingPointLong := merchant.Data[0].Location.Long

	// Get total distance tsp
	totalDistance := 0.0

	for len(merchantIds) != 0 {
		var merchant entities.FindDistanceResult

		tempMerchant, err := i.merchantrepository.FindDistance(startingPointLat, startingPointLong, merchantIds)

		merchant = *tempMerchant

		fmt.Println("exec 2")

		if err != nil {
			fmt.Println("error 9")
			return nil, fiber.StatusInternalServerError, err
		}

		startingPointLat = merchant.Lat
		startingPointLong = merchant.Long
		totalDistance += merchant.Distance
		valueToDelete := merchant.Id

		// Deleted used merchant in merchantIds
		indexToDelete := -1
		for i, v := range merchantIds {
			if v == valueToDelete {
				indexToDelete = i
				break
			}
		}

		if indexToDelete != -1 {
			merchantIds = append(merchantIds[:indexToDelete], merchantIds[indexToDelete+1:]...)
		}
	}

	if len(merchantIds) == 0 {
		totalDistance += helpers.Haversine(startingPointLat, startingPointLong, p.EstimatesBody.UserLocation.Lat, p.EstimatesBody.UserLocation.Long)
	}

	fmt.Println("distance : ", totalDistance)

	// Validate that total tsp is > 3 km
	if totalDistance > 3 {
		fmt.Println("error 10")
		return nil, fiber.StatusBadRequest, errors.New("the coordinates is too far > 3km")
	}

	// Calculate estimate time delivery
	estimatedTimeInHours := totalDistance / float64(40) // 40km/h
	estimatedTimeInMinutes := estimatedTimeInHours * 60

	fmt.Println("estimatedTimeInMinutes : ", math.Ceil(estimatedTimeInMinutes))

	// Create estimates and estimates_item
	estimateId, err := i.estimatesRepository.Create(p.UserId, int(math.Ceil(estimatedTimeInMinutes)))

	if err != nil {
		fmt.Println("error 11")
		return nil, fiber.StatusInternalServerError, err
	}

	err = i.estimateitemsrepository.Create(estimateId, p.EstimatesBody.Orders)

	if err != nil {
		fmt.Println("error 12")
		return nil, fiber.StatusInternalServerError, err
	}

	// Calculate Total Price
	totalPrice, err := i.estimateitemsrepository.FindTotalPrice(estimateId)
	if err != nil {
		fmt.Println("error 13")
		return nil, fiber.StatusInternalServerError, err
	}

	result, err := i.estimatesRepository.Update(totalPrice)
	if err != nil {
		fmt.Println("error 14")
		return nil, fiber.StatusInternalServerError, err
	}

	return result, fiber.StatusOK, nil
}
