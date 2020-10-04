package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mycontroller-org/backend/v2/pkg/api/field"
	"github.com/mycontroller-org/backend/v2/pkg/model"
	mtrml "github.com/mycontroller-org/backend/v2/pkg/model/metric"
	svc "github.com/mycontroller-org/backend/v2/pkg/service"
)

// global constants
const (
	QuickID = "quick_id"
)

func registerMetricRoutes(router *mux.Router) {
	router.HandleFunc("/api/metric", getMetricList).Methods(http.MethodPost)
	router.HandleFunc("/api/metric", getMetric).Methods(http.MethodGet)
}

func getMetric(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params, err := ReceivedQueryMap(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// returns values from parameters / user input
	getValue := func(key string) string {
		if values, ok := params[key]; ok {
			if len(values) > 0 {
				return values[0]
			}
		}
		return ""
	}
	// returns all the values
	getValues := func(key string) []string {
		if values, ok := params[key]; ok {
			if len(values) > 0 {
				return values
			}
		}
		return nil
	}

	queryConfig := &mtrml.QueryConfig{}
	queryConfig.Individual = []mtrml.Query{{Name: QuickID, Tags: map[string]string{}}}

	if quickID, ok := params[QuickID]; ok {
		if len(quickID) > 0 {
			rt, kvMap, err := model.ResourceKeyValueMap(quickID[0])
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			// get resource details
			switch rt {
			case model.QuickIDSensorField:
				// get field details
				field, err := field.GetByIDs(kvMap[model.KeyGatewayID], kvMap[model.KeyNodeID], kvMap[model.KeySensorID], kvMap[model.KeyFieldID])
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				queryConfig.Individual[0].Tags[model.KeyID] = field.ID
				queryConfig.Individual[0].MetricType = field.MetricType

			default:
				http.Error(w, fmt.Sprintf("resource type not supported in metric. ResourceType:%s", rt), 500)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("%s not supplied", QuickID), 500)
			return
		}
	} else {
		http.Error(w, fmt.Sprintf("%s not supplied", QuickID), 500)
		return
	}

	// update optional parameters
	if value := getValue(mtrml.QueryKeyStart); value != "" {
		queryConfig.Global.Start = value
	}
	if value := getValue(mtrml.QueryKeyStop); value != "" {
		queryConfig.Global.Stop = value
	}
	if value := getValue(mtrml.QueryKeyWindow); value != "" {
		queryConfig.Global.Window = value
	}
	if values := getValues(mtrml.QueryKeyFunctions); values != nil {
		queryConfig.Global.Functions = values
	}

	result, err := svc.MTS.Query(queryConfig)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	od, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}

func getMetricList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	queryConfig := &mtrml.QueryConfig{}

	err = json.Unmarshal(d, queryConfig)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result, err := svc.MTS.Query(queryConfig)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	od, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}