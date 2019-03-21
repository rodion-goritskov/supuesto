package apiv1

import (
	"github.com/go-chi/chi"
	"net/http"
	"encoding/json"
	"github.com/rodion-goritskov/supuesto/backend/scenario"
	"log"
	"github.com/rodion-goritskov/supuesto/backend/dgraph"
	"context"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux{
	router := chi.NewRouter()
	router.Post("/scenario", UpdateScenario)
	router.Get("/scenario/{scenarioId}", FindScenarioById)
	router.Delete("/scenario/{scenarioId}", DeleteScenarioById)

	return router
}

func UpdateScenario(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body scenario.Scenario
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	body.DatatypeScenario = ""

	client, conn := dgraph.CreateClient()
	tr := client.NewTxn()
	sc, err :=  scenario.CreateOrUpdateScenario(body)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		conn.Close()
		return
	}

	res, err := tr.Mutate(context.Background(), sc)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		conn.Close()
		return
	}

	conn.Close()
	render.JSON(w, r, res)
}

func FindScenarioById(w http.ResponseWriter, r *http.Request) {
	scenarioId := chi.URLParam(r, "scenarioId")
	client, conn := dgraph.CreateClient()

	type Root struct {
		 ScenariosById []scenario.Scenario `json:"scenarios_by_id"`
	}

	response, err := client.NewTxn().QueryWithVars(context.Background(), scenario.GetByID, map[string]string{"$id": scenarioId})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		conn.Close()
		return
	}

	var root Root
	err = json.Unmarshal(response.Json, &root)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		conn.Close()
		return
	}

	conn.Close()
	render.JSON(w, r, root.ScenariosById)
}

func DeleteScenarioById(w http.ResponseWriter, r *http.Request) {
	scenarioId := chi.URLParam(r, "scenarioId")

	client, conn := dgraph.CreateClient()
	sc, err := scenario.DeleteScenario(scenario.Scenario{
		Uid:scenarioId,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		conn.Close()
		return
	}

	result, err := client.NewTxn().Mutate(context.Background(), sc)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		conn.Close()
		return
	}

	conn.Close()
	render.JSON(w, r, result)
}
