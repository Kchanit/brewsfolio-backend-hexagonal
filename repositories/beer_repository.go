package repositories

import (
	"fmt"

	"github.com/Kchanit/brewsfolio-backend/models"
	"gorm.io/gorm"
)

type BeerRepository interface {
	GetBeerByID(id string) (models.Beer, error)
	GetBeers() ([]models.Beer, error)
	GetBeersByIDs(ids []uint) ([]models.Beer, error)
	CreateBeer(beer *models.Beer) (*models.Beer, error)
	UpdateBeer(id string, beer *models.Beer) (*models.Beer, error)
	DeleteBeer(id string) error
	GetFavoritesByUserID(id string) ([]models.Beer, error)
	AddFavoriteBeer(id uint, beer *models.Beer) (*models.Beer, error)
	RemoveFavoriteBeer(id uint, beer *models.Beer) (*models.Beer, error)
}

// beerRepository is an implementation of the BeerRepository interface.
type beerRepository struct {
	db *gorm.DB
}

// NewBeerRepository creates a new instance of the beerRepository.
func NewBeerRepository(db *gorm.DB) BeerRepository {
	return &beerRepository{db: db}
}

// GetBeerByID retrieves a beer by its ID.
func (r *beerRepository) GetBeerByID(id string) (models.Beer, error) {
	var beer models.Beer
	err := r.db.Preload("Reviews").Where("id = ?", id).First(&beer).Error
	return beer, err
}

// GetBeers retrieves all beers.
func (r *beerRepository) GetBeers() ([]models.Beer, error) {
	var beers []models.Beer
	err := r.db.Find(&beers).Error
	return beers, err
}

// GetBeersByIDs retrieves a list of beers by their IDs.
func (r *beerRepository) GetBeersByIDs(ids []uint) ([]models.Beer, error) {
	var beers []models.Beer
	err := r.db.Where("id IN ?", ids).Find(&beers).Error
	return beers, err
}

// CreateBeer creates a new beer.
func (r *beerRepository) CreateBeer(beer *models.Beer) (*models.Beer, error) {
	err := r.db.Create(&beer).Error
	return beer, err
}

// UpdateBeer updates a beer.
func (r *beerRepository) UpdateBeer(id string, beer *models.Beer) (*models.Beer, error) {
	err := r.db.Model(&beer).Where("id = ?", id).Updates(beer).Error
	return beer, err
}

// DeleteBeer deletes a beer.
func (r *beerRepository) DeleteBeer(id string) error {
	err := r.db.Where("id = ?", id).Delete(&models.Beer{}).Error
	return err
}

// GetFavoritesByUserID retrieves favorites by user ID.
func (r *beerRepository) GetFavoritesByUserID(id string) ([]models.Beer, error) {
	var user models.User
	err := r.db.Preload("Favorites").Find(&user, id).Error
	return user.Favorites, err
}

// AddFavoriteBeer adds a beer to favorites.
func (r *beerRepository) AddFavoriteBeer(userID uint, beer *models.Beer) (*models.Beer, error) {
	var user models.User
	err := r.db.Model(&user).First(&user, userID).Association("Favorites").Append(beer)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return beer, err
}

// RemoveFavoriteBeer removes a beer from favorites.
func (r *beerRepository) RemoveFavoriteBeer(userID uint, beer *models.Beer) (*models.Beer, error) {
	var user models.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = r.db.Model(&user).Association("Favorites").Delete(beer)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return beer, err
}
