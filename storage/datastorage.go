package storage

import (
	"github.com/chinmay211997/RiskManagerApp/models"
	"github.com/google/uuid"
)

type DataStorageInterface interface {
	GetAllKeysFromDataStore() []models.Risk
	GetKeyFromDataStore(key uuid.UUID) (*models.Risk, bool)
	PutKeyInDataStore(key uuid.UUID, value models.Risk)
}

type DataStorageType string

const (
	IN_MEMORY DataStorageType = "IN_MEMORY"
)

func PutKeyInDataStorage(i DataStorageInterface, key uuid.UUID, value models.Risk) {
	i.PutKeyInDataStore(key, value)
}

func GetAllKeysDataStorage(i DataStorageInterface) []models.Risk {
	return i.GetAllKeysFromDataStore()
}

func GetKeyFromDataStorage(i DataStorageInterface, key uuid.UUID) (*models.Risk, bool) {
	return i.GetKeyFromDataStore(key)
}
