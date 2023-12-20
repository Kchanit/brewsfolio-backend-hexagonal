package services

import (
	"log"

	"github.com/Kchanit/brewsfolio-backend/models"
	"github.com/Kchanit/brewsfolio-backend/repositories"
)

// BeerService provides methods for handling beer-related business logic.
type BeerService struct {
	beerRepository repositories.BeerRepository
}

// NewBeerService creates a new instance of BeerService.
func NewBeerService(beerRepository repositories.BeerRepository) *BeerService {
	return &BeerService{beerRepository: beerRepository}
}

// GetBeerByID retrieves a beer by its ID.
func (s *BeerService) GetBeerByID(id string) (models.Beer, error) {

	beer, err := s.beerRepository.GetBeerByID(id)
	if err != nil {
		log.Println(err)
		return models.Beer{}, err
	}
	return beer, nil
}

// GetBeersByIDs retrieves a list of beers by their IDs.
func (s *BeerService) GetBeersByIDs(ids []uint) ([]models.Beer, error) {

	beers, err := s.beerRepository.GetBeersByIDs(ids)
	if err != nil {
		log.Println(err)
		return []models.Beer{}, err
	}

	return beers, nil
}

// GetBeers retrieves all beers.
func (s *BeerService) GetBeers() ([]models.Beer, error) {

	beers, err := s.beerRepository.GetBeers()
	if err != nil {
		log.Println(err)
		return []models.Beer{}, err
	}
	return beers, nil
}

// CreateBeer creates a new beer.
func (s *BeerService) CreateBeer(beer *models.Beer) (*models.Beer, error) {
	beer, err := s.beerRepository.CreateBeer(beer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return beer, nil
}

// UpdateBeer updates a beer.
func (s *BeerService) UpdateBeer(id string, beer *models.Beer) (*models.Beer, error) {
	beer, err := s.beerRepository.UpdateBeer(id, beer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return beer, nil
}

// DeleteBeer deletes a beer.
func (s *BeerService) DeleteBeer(id string) error {
	err := s.beerRepository.DeleteBeer(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetFavoritesByUserID retrieves favorites by user ID.
func (s *BeerService) GetFavoritesByUserID(id string) ([]models.Beer, error) {

	favorites, err := s.beerRepository.GetFavoritesByUserID(id)
	if err != nil {
		log.Println(err)
		return []models.Beer{}, err
	}
	return favorites, nil
}

// AddFavoriteBeer adds a beer to favorites.
func (s *BeerService) AddFavoriteBeer(userID uint, beerID string) (*models.Beer, error) {

	beer, err := s.beerRepository.GetBeerByID(beerID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := s.beerRepository.AddFavoriteBeer(userID, &beer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

// RemoveFavoriteBeer removes a beer from favorites.
func (s *BeerService) RemoveFavoriteBeer(userID uint, beerID string) (*models.Beer, error) {
	beer, err := s.beerRepository.GetBeerByID(beerID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := s.beerRepository.RemoveFavoriteBeer(userID, &beer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
