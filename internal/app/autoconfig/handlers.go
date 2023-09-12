package autoconfig

import (
	"automation-hub-nginxconfigmanager/internal/app/dto"
	"encoding/json"
	"net/http"
)

func AddConfigHandler(w http.ResponseWriter, r *http.Request) {
	var auto dto.Automation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&auto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := auto.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := manageConfig(Add, auto); err != nil {
		http.Error(w, "Failed to add config", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func RemoveConfigHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	if err := manageConfig(Remove, dto.Automation{Name: name}); err != nil {
		http.Error(w, "Failed to remove config", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	var auto dto.Automation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&auto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := auto.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := manageConfig(Update, auto); err != nil {
		http.Error(w, "Failed to update config", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
