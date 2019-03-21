package scenario

import (
	"time"
	"github.com/dgraph-io/dgo/protos/api"
	"encoding/json"
)

type Scenario struct {
	Uid string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	CreatedDate *time.Time `json:"created.date,omitempty"`
	Steps string `json:"scenario.steps,omitempty"`
	DatatypeScenario string `json:"datatype.scenario"`
}

var SearchByName = `query Name($name: string){
	scenarios_by_regexp(func: has(datatype.scenario)) @filter(regexp(name,/.*$name.*/i)) {
    expand(_all_)
	}
}`

var GetByID = `query scenarios_by_id($id: string){
	scenarios_by_id(func: uid($id)) @filter(has(datatype.scenario)) {
    expand(_all_)
	}
}`

func CreateOrUpdateScenario(class Scenario) (*api.Mutation, error) {
	mutation := &api.Mutation{
		CommitNow:true,
	}

	sj, err := json.Marshal(class)
	if err != nil {
		return nil, err
	}

	mutation.SetJson = sj
	return mutation, nil
}

func DeleteScenario(class Scenario) (*api.Mutation, error) {
	mutation := &api.Mutation{
		CommitNow:true,
	}

	sj, err := json.Marshal(class)
	if err != nil {
		return nil, err
	}

	mutation.DeleteJson = sj
	return mutation, nil
}
