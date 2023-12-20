package repositories

import (
	"github.com/Kchanit/brewsfolio-backend/models"
	"gorm.io/gorm"
)

type CollectionRepository interface {
	GetCollections() ([]models.Collection, error)
	GetCollectionByID(id string) (models.Collection, error)
	GetCollectionsByUserID(id string) ([]models.Collection, error)
	GetPublicCollections() ([]models.Collection, error)
	CreateCollection(collection *models.Collection) (models.Collection, error)
	DeleteCollection(id string) error
	UpdateCollection(collection *models.Collection) (models.Collection, error)
	AddBeerToCollection(collection models.Collection, beer models.Beer) (models.Collection, error)
	AddBeersToCollection(collection models.Collection, beers []models.Beer) (models.Collection, error)
	RemoveBeerFromCollection(collection models.Collection, beer models.Beer) (models.Collection, error)
}

type collectionRepository struct {
	db *gorm.DB
}

func NewCollectionRepository(db *gorm.DB) CollectionRepository {
	return &collectionRepository{db: db}
}

func (r *collectionRepository) GetCollections() ([]models.Collection, error) {
	var collections []models.Collection
	if err := r.db.Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

func (r *collectionRepository) GetCollectionByID(id string) (models.Collection, error) {
	var collection models.Collection
	if err := r.db.Preload("Beers").First(&collection, id).Error; err != nil {
		return models.Collection{}, err
	}
	return collection, nil
}

func (r *collectionRepository) GetCollectionsByUserID(id string) ([]models.Collection, error) {
	var collections []models.Collection
	if err := r.db.Preload("Beers").Preload("User").Where("user_id = ?", id).Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

func (r *collectionRepository) GetPublicCollections() ([]models.Collection, error) {
	var collections []models.Collection
	if err := r.db.Preload("Beers").Preload("User").Where("is_private = ?", false).Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

// Create Collection
func (r *collectionRepository) CreateCollection(collection *models.Collection) (models.Collection, error) {
	err := r.db.Create(collection).Error
	return *collection, err
}

// Delete Collection
func (r *collectionRepository) DeleteCollection(id string) error {
	err := r.db.Delete(&models.Collection{}, id).Error
	return err
}

// Update Collection
func (r *collectionRepository) UpdateCollection(collection *models.Collection) (models.Collection, error) {
	err := r.db.Save(collection).Error
	return *collection, err
}

// Add beer to collection
func (r *collectionRepository) AddBeerToCollection(collection models.Collection, beer models.Beer) (models.Collection, error) {
	collection.Beers = append(collection.Beers, beer)
	err := r.db.Save(&collection).Error
	return collection, err
}

// Add beers to collection
func (r *collectionRepository) AddBeersToCollection(collection models.Collection, beers []models.Beer) (models.Collection, error) {
	collection.Beers = append(collection.Beers, beers...)
	err := r.db.Save(&collection).Error
	return collection, err
}

// Remove beer from collection
func (r *collectionRepository) RemoveBeerFromCollection(collection models.Collection, beer models.Beer) (models.Collection, error) {
	for i := 0; i < len(collection.Beers); i++ {
		if collection.Beers[i].ID == beer.ID {
			collection.Beers = append(collection.Beers[:i], collection.Beers[i+1:]...)
			break
		}
	}
	err := r.db.Save(&collection).Error
	return collection, err
}
