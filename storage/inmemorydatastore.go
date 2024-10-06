package storage

import (
	"maps"
	"sync"

	"github.com/chinmay211997/RiskManagerApp/models"
	"github.com/chinmay211997/RiskManagerApp/utils"
	"github.com/google/uuid"
)

type InMemoryDataStore struct {
	storageType DataStorageType
	mu          sync.RWMutex
	riskMap     map[uuid.UUID]models.Risk
}

var lock = &sync.Mutex{}

var inMemoryDataStoreInstance *InMemoryDataStore

func GetinMemoryDataStoreInstance() *InMemoryDataStore {
	if inMemoryDataStoreInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if inMemoryDataStoreInstance == nil {
			utils.Logger.Debug("Creating inMemoryDataStore Instance")
			inMemoryDataStoreInstance = &InMemoryDataStore{
				storageType: IN_MEMORY,
				riskMap:     make(map[uuid.UUID]models.Risk),
			}
		} else {
			utils.Logger.Debug("inMemoryDataStore Instance already created")
		}
	} else {
		utils.Logger.Debug("inMemoryDataStore Instance already created.")
	}
	return inMemoryDataStoreInstance
}

func (ds *InMemoryDataStore) PutKeyInDataStore(key uuid.UUID, value models.Risk) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.riskMap[key] = value
}

func (ds *InMemoryDataStore) GetKeyFromDataStore(key uuid.UUID) (*models.Risk, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	value, ok := ds.riskMap[key]
	if ok {
		return &value, true
	}
	return nil, false
}

func (ds *InMemoryDataStore) GetAllKeysFromDataStore() []models.Risk {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	var riskListData = make([]models.Risk, 0)
	for val := range maps.Values(ds.riskMap) {
		riskListData = append(riskListData, val)
	}
	return riskListData
}
