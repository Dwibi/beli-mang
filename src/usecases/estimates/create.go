package estimatesusecase

import (
	"errors"
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

type LocationList struct {
	Latitude  float64
	Longitude float64
}

func (i sEstimatesUseCase) Create(p CreateEstimateParams) (*entities.ResultEstimate, int, error) {
	// validate user doing the request is user
	userRole, err := i.userRepository.FindUserRole(p.UserId)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	if userRole != "user" {
		return nil, fiber.StatusUnauthorized, errors.New("only user can use this route")
	}

	var locationList []LocationList
	totalItemsPrice := 0
	totalDistance := 0.0
	estimatedTimeInMinutes := 0.0
	const area = 3.0 // km²
	var radius = math.Sqrt(area)

	for _, order := range p.EstimatesBody.Orders {
		if order.MerchantId == "" {
			return nil, fiber.StatusBadRequest, errors.New("merchantId can't be empty")
		}

		// Convert to index for merchant Findone repository
		merchantIdInt, err := strconv.Atoi(order.MerchantId)
		if err != nil || merchantIdInt < 0 {
			return nil, fiber.StatusNotFound, errors.New("invalid merchantId")
		}

		// get merchant location
		merchant, err := i.merchantrepository.FindOne(merchantIdInt)
		if merchant == nil {
			return nil, fiber.StatusNotFound, errors.New("merchantId not found")
		}
		if err != nil {
			return nil, fiber.StatusInternalServerError, err
		}

		location := LocationList{
			Latitude:  merchant.Location.Lat,
			Longitude: merchant.Location.Long,
		}

		// if merchant is starting point append to first index
		if order.IsStartingPoint {
			locationList = append([]LocationList{location}, locationList...)
		} else {
			locationList = append(locationList, location)
		}

		for _, item := range order.Items {
			if item.ItemId == "" {
				return nil, fiber.StatusBadRequest, errors.New("ItemId can't be empty")
			}
			// Convert to index for merchant Findone repository
			itemIdInt, err := strconv.Atoi(item.ItemId)
			if err != nil || itemIdInt < 0 {
				return nil, fiber.StatusNotFound, errors.New("invalid ItemId")
			}

			// Get item price
			price, err := i.itemsrepository.FindItemPrice(itemIdInt)
			if price < 0 {
				return nil, fiber.StatusNotFound, errors.New("ItemId not found")
			}
			if err != nil {
				return nil, fiber.StatusInternalServerError, err
			}

			// added to totalItemsPrice
			// fmt.Println("ItemsPrice:", price*item.Quantity)
			totalItemsPrice += price * item.Quantity
		}
	}

	// Add user location in the last of the locationList
	locationList = append(locationList, LocationList{
		Latitude:  p.EstimatesBody.UserLocation.Lat,
		Longitude: p.EstimatesBody.UserLocation.Long,
	})

	// Validate if all the merchant is not far too user (not > 3 KM)
	for _, merchatLocation := range locationList {
		distance := helpers.Haversine(merchatLocation.Latitude, merchatLocation.Longitude, p.EstimatesBody.UserLocation.Lat, p.EstimatesBody.UserLocation.Long)
		// fmt.Println("distance :", distance)
		// fmt.Println("radius :", radius)

		if distance > radius {
			return nil, fiber.StatusBadRequest, errors.New("the merchant are too far > 3km from the user")
		}
	}

	// calculate for totalDistance
	visited := make(map[int]bool)
	currentLocation := locationList[0]
	for len(visited) < len(locationList)-1 {
		nearestDistance := math.MaxFloat64
		nearestIndex := -1
		for i := 1; i < len(locationList)-1; i++ {
			if visited[i] {
				continue
			}
			distance := helpers.Haversine(currentLocation.Latitude, currentLocation.Longitude, locationList[i].Latitude, locationList[i].Longitude)
			if distance < nearestDistance {
				nearestDistance = distance
				nearestIndex = i
			}
		}
		if nearestIndex == -1 {
			break
		}
		totalDistance += nearestDistance
		visited[nearestIndex] = true
		currentLocation = locationList[nearestIndex]
	}

	// Add distance from the last merchant to the user's location
	totalDistance += helpers.Haversine(currentLocation.Latitude, currentLocation.Longitude, p.EstimatesBody.UserLocation.Lat, p.EstimatesBody.UserLocation.Long)

	// fmt.Println("totalDistance :", totalDistance)

	// Calculate estimateTime
	estimatedTimeInHours := totalDistance / 40.0 // 40km/h
	estimatedTimeInMinutes = estimatedTimeInHours * 60

	// fmt.Println("estimatedTimeInMinutes :", estimatedTimeInMinutes)

	estimate, err := i.estimatesRepository.Create(p.UserId, float64(totalItemsPrice), estimatedTimeInMinutes)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	estimateId, _ := strconv.Atoi(estimate.CalculatedEstimateId)

	err = i.estimateitemsrepository.Create(estimateId, p.EstimatesBody.Orders)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	// fmt.Println("TotalItemsPrice :", totalItemsPrice)

	return estimate, fiber.StatusOK, nil
}

// func (i sEstimatesUseCase) Create(p CreateEstimateParams) (*entities.ResultEstimate, int, error) {
// 	// validate user doing the request is user
// 	userRole, err := i.userRepository.FindUserRole(p.UserId)
// 	if err != nil {
// 		return nil, fiber.StatusInternalServerError, err
// 	}

// 	if userRole != "user" {
// 		return nil, fiber.StatusUnauthorized, errors.New("only user can use this route")
// 	}

// 	// Get all merchantId and itemId
// 	merchantIds := []int{}
// 	itemIds := []int{}
// 	var startingPointMerchant string

// 	for _, merchant := range p.EstimatesBody.Orders {
// 		if merchant.MerchantId == "" {
// 			return nil, fiber.StatusBadRequest, errors.New("merchantId can't be empty")
// 		}
// 		if merchant.IsStartingPoint {
// 			startingPointMerchant = merchant.MerchantId
// 		} else {
// 			merchantIdInt, err := strconv.Atoi(merchant.MerchantId)
// 			if err != nil {
// 				return nil, fiber.StatusNotFound, errors.New("invalid merchantId")
// 			}
// 			merchantIds = append(merchantIds, merchantIdInt)
// 			if merchantIdInt == 0 {
// 				return nil, fiber.StatusNotFound, errors.New("merchantId can't be 0")
// 			}
// 		}
// 		for _, item := range merchant.Items {
// 			if item.ItemId == "" {
// 				return nil, fiber.StatusBadRequest, errors.New("itemId can't be empty")
// 			}
// 			itemIdInt, err := strconv.Atoi(item.ItemId)
// 			if err != nil {
// 				return nil, fiber.StatusNotFound, errors.New("invalid itemId")
// 			}
// 			itemIds = append(itemIds, itemIdInt)
// 			if itemIdInt == 0 {
// 				return nil, fiber.StatusNotFound, errors.New("itemId can't be 0")
// 			}
// 		}
// 	}

// 	// Validate if all the merchant id is exist
// 	missingMerchantId, err := i.merchantrepository.GetMissingIDs(merchantIds)
// 	if err != nil {
// 		return nil, fiber.StatusInternalServerError, err
// 	}

// 	if len(missingMerchantId) > 0 {
// 		return nil, fiber.StatusBadRequest, fmt.Errorf("the following IDs do not exist in the merchants table: %v", missingMerchantId)
// 	}

// 	// Validate if all the item id is exist
// 	missingItemId, err := i.itemsrepository.GetMissingIDs(itemIds)
// 	if err != nil {
// 		return nil, fiber.StatusInternalServerError, err
// 	}

// 	if len(missingItemId) > 0 {
// 		return nil, fiber.StatusNotFound, fmt.Errorf("the following IDs do not exist in the items table: %v", missingItemId)
// 	}

// 	// Get first startingPointMerchant lat and long
// 	startingPointMerchantInt, err := strconv.Atoi(startingPointMerchant)
// 	if err != nil {
// 		return nil, fiber.StatusBadRequest, errors.New("invalid startingPointMerchantId")
// 	}
// 	startingMerchant, err := i.merchantrepository.FindOne(startingPointMerchantInt)
// 	if err != nil {
// 		return nil, fiber.StatusNotFound, err
// 	}

// 	startingPointLat := startingMerchant.Location.Lat
// 	startingPointLong := startingMerchant.Location.Long

// 	// Get total distance tsp
// 	totalDistance := 0.0

// 	for len(merchantIds) != 0 {
// 		fmt.Println("Current merchant IDs:", merchantIds) // Debugging

// 		// Find the nearest merchant from the starting point
// 		nearestMerchant := -1
// 		nearestDistance := math.MaxFloat64
// 		for _, id := range merchantIds {
// 			merchant, err := i.merchantrepository.FindOne(id)
// 			if err != nil {
// 				return nil, fiber.StatusInternalServerError, err
// 			}

// 			distance := helpers.Haversine(startingPointLat, startingPointLong, merchant.Location.Lat, merchant.Location.Long)
// 			if distance < nearestDistance {
// 				nearestDistance = distance
// 				nearestMerchant = id
// 			}
// 		}

// 		if nearestMerchant == -1 {
// 			return nil, fiber.StatusInternalServerError, errors.New("unable to find nearest merchant")
// 		}

// 		merchant, err := i.merchantrepository.FindOne(nearestMerchant)
// 		if err != nil {
// 			return nil, fiber.StatusInternalServerError, err
// 		}

// 		if nearestDistance > 3 {
// 			return nil, fiber.StatusBadRequest, errors.New("the coordinates are too far > 3km")
// 		}

// 		totalDistance += nearestDistance
// 		startingPointLat = merchant.Location.Lat
// 		startingPointLong = merchant.Location.Long

// 		// Remove the nearest merchant from the list
// 		indexToDelete := -1
// 		for i, v := range merchantIds {
// 			if v == nearestMerchant {
// 				indexToDelete = i
// 				break
// 			}
// 		}

// 		if indexToDelete != -1 {
// 			merchantIds = append(merchantIds[:indexToDelete], merchantIds[indexToDelete+1:]...)
// 		} else {
// 			return nil, fiber.StatusInternalServerError, errors.New("failed to delete merchantId from the list")
// 		}
// 	}

// 	if len(merchantIds) == 0 {
// 		totalDistance += helpers.Haversine(startingPointLat, startingPointLong, p.EstimatesBody.UserLocation.Lat, p.EstimatesBody.UserLocation.Long)
// 	}

// 	// Calculate estimate time delivery
// 	estimatedTimeInHours := totalDistance / 40.0 // 40km/h
// 	estimatedTimeInMinutes := estimatedTimeInHours * 60

// 	// Create estimates and estimates_item
// estimateId, err := i.estimatesRepository.Create(p.UserId, int(math.Ceil(estimatedTimeInMinutes)))
// if err != nil {
// 	return nil, fiber.StatusInternalServerError, err
// }

// err = i.estimateitemsrepository.Create(estimateId, p.EstimatesBody.Orders)
// if err != nil {
// 	return nil, fiber.StatusInternalServerError, err
// }

// 	// Calculate Total Price
// 	totalPrice, err := i.estimateitemsrepository.FindTotalPrice(estimateId)
// 	if err != nil {
// 		return nil, fiber.StatusInternalServerError, err
// 	}

// 	result, err := i.estimatesRepository.Update(totalPrice)
// 	if err != nil {
// 		return nil, fiber.StatusInternalServerError, err
// 	}

// 	return result, fiber.StatusOK, nil
// }
