package services

import (
	"errors"
	"log"
	"strconv"

	"github.com/Kchanit/brewsfolio-backend/models"
	"github.com/Kchanit/brewsfolio-backend/repositories"
)

type CollectionService struct {
	collectionRepository repositories.CollectionRepository
	userRepository       repositories.UserRepository
	beerRepository       repositories.BeerRepository
}

func NewCollectionService(collectionRepository repositories.CollectionRepository, userRepository repositories.UserRepository, beerRepository repositories.BeerRepository) *CollectionService {
	return &CollectionService{
		collectionRepository: collectionRepository,
		userRepository:       userRepository,
		beerRepository:       beerRepository,
	}
}

func (s *CollectionService) GetCollections() ([]models.Collection, error) {
	collections, err := s.collectionRepository.GetCollections()
	if err != nil {
		log.Println(err)
		return []models.Collection{}, err
	}
	return collections, nil
}

func (s *CollectionService) GetCollectionByID(id string) (models.Collection, error) {
	collection, err := s.collectionRepository.GetCollectionByID(id)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}
	return collection, nil
}

func (s *CollectionService) GetCollectionsByUserID(id string) ([]models.Collection, error) {
	collections, err := s.collectionRepository.GetCollectionsByUserID(id)
	if err != nil {
		log.Println(err)
		return []models.Collection{}, err
	}
	return collections, nil
}

func (s *CollectionService) GetPublicCollections() ([]models.Collection, error) {
	collections, err := s.collectionRepository.GetPublicCollections()
	if err != nil {
		log.Println(err)
		return []models.Collection{}, err
	}
	return collections, nil
}

func (s *CollectionService) CreateCollection(userID uint, title string, description string, isPrivate bool) (*models.Collection, error) {
	// format userID uint to string
	userIDString := strconv.Itoa(int(userID))
	user, err := s.userRepository.GetUserByID(userIDString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	collection := models.Collection{
		UserID:      user.ID,
		User:        user,
		Title:       title,
		Description: description,
		IsPrivate:   isPrivate,
	}

	collection, err = s.collectionRepository.CreateCollection(&collection)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &collection, nil
}

func (s *CollectionService) DeleteCollection(userID uint, collectionID string) error {
	collection, err := s.collectionRepository.GetCollectionByID(collectionID)
	if err != nil {
		log.Println(err)
		return err
	}

	if collection.UserID != userID {
		return errors.New("Unauthorized")
	}

	err = s.collectionRepository.DeleteCollection(collectionID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *CollectionService) UpdateCollection(collectionID string, title string, description string, isPrivate bool) (collection models.Collection, err error) {
	collection, err = s.collectionRepository.GetCollectionByID(collectionID)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}

	if collection.UserID != collection.UserID {
		return models.Collection{}, errors.New("Unauthorized")
	}

	collection.Title = title
	collection.Description = description
	collection.IsPrivate = isPrivate

	collection, err = s.collectionRepository.UpdateCollection(&collection)
	if err != nil {
		log.Println(err)
		return
	}

	return collection, nil
}

// AddBeerToCollection
func (s *CollectionService) AddBeerToCollection(userID uint, collectionID string, beerID string) (models.Collection, error) {
	collection, err := s.collectionRepository.GetCollectionByID(collectionID)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}

	if userID != collection.UserID {
		return models.Collection{}, errors.New("Unauthorized")
	}

	beer, err := s.beerRepository.GetBeerByID(beerID)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}

	collection, err = s.collectionRepository.AddBeerToCollection(collection, beer)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}
	return collection, nil

}

func (s *CollectionService) AddBeersToCollection(userID uint, collectionID string, beerIDs []uint) (models.Collection, error) {
	collection, err := s.collectionRepository.GetCollectionByID(collectionID)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}

	if userID != collection.UserID {
		return models.Collection{}, errors.New("Unauthorized")
	}

	beers, err := s.beerRepository.GetBeersByIDs(beerIDs)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}
	collection, err = s.collectionRepository.AddBeersToCollection(collection, beers)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}
	return collection, nil

}

func (s *CollectionService) RemoveBeerFromCollection(userID uint, collectionID string, beerID string) (models.Collection, error) {
	collection, err := s.collectionRepository.GetCollectionByID(collectionID)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}

	if userID != collection.UserID {
		return models.Collection{}, errors.New("Unauthorized")
	}

	beer, err := s.beerRepository.GetBeerByID(beerID)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}

	collection, err = s.collectionRepository.RemoveBeerFromCollection(collection, beer)
	if err != nil {
		log.Println(err)
		return models.Collection{}, err
	}

	return collection, nil
}
