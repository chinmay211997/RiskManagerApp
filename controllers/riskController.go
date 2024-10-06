package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chinmay211997/RiskManagerApp/models"
	"github.com/chinmay211997/RiskManagerApp/storage"
	"github.com/chinmay211997/RiskManagerApp/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var DataStore = storage.GetinMemoryDataStoreInstance()

func GetRisksById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	riskId, err := uuid.Parse(params["Id"])
	if err != nil {
		utils.Logger.Error(err.Error(), "Id", params["Id"])
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GetLongError(models.BAD_REQUEST, http.StatusBadRequest,
			models.MALFORMED_RISKID))
		return
	}
	fmt.Println(DataStore)
	riskObj, ok := storage.GetKeyFromDataStorage(DataStore, riskId)
	if ok {
		utils.Logger.Info(fmt.Sprintf("Returned Risk ID:%s", riskId))
		json.NewEncoder(w).Encode(riskObj)
	} else {
		utils.Logger.Error(fmt.Sprintf("%s ID:%s", models.RISKID_NOT_EXIST, riskId))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.GetLongError(models.NOT_FOUND, http.StatusNotFound,
			models.RISKID_NOT_EXIST))
	}
}

func GetRisksList(w http.ResponseWriter, req *http.Request) {

	var riskListData = storage.GetAllKeysDataStorage(DataStore)
	riskList := models.GetRisksListResponse{
		RiskListData: riskListData,
		TotalLength:  len(riskListData),
		Code:         http.StatusOK,
	}

	utils.Logger.Info(fmt.Sprintf("Total Risks Returned:%d", len(riskListData)))
	json.NewEncoder(w).Encode(riskList)
}

func CreateRisk(w http.ResponseWriter, req *http.Request) {
	var riskReq models.CreateRiskRequest

	err := json.NewDecoder(req.Body).Decode(&riskReq)
	if err != nil {
		utils.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			models.GetLongError(models.INTERNAL_SERVER_ERROR, http.StatusInternalServerError,
				models.MALFORMED_REQUEST_BODY))
		return
	}
	if !riskReq.ValidateCreateRiskRequest() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.GetLongError(models.BAD_REQUEST, http.StatusBadRequest,
			models.INVALID_REQUEST_BODY))
		return
	}

	risk := models.Risk{
		ID:          uuid.New(),
		Title:       riskReq.Title,
		State:       riskReq.State,
		Description: riskReq.Description,
	}

	storage.PutKeyInDataStorage(DataStore, risk.ID, risk)

	response := models.CreateRiskResponse{
		RiskID:  risk.ID,
		Message: models.RISK_CREATE_SUCCESS,
		Code:    http.StatusOK,
	}

	utils.Logger.Info(fmt.Sprintf("%s with ID:%s", models.RISK_CREATE_SUCCESS, risk.ID))
	json.NewEncoder(w).Encode(response)
}
