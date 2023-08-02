package entity

type Message struct {
	Id              string   `json:"id"`
	InstanceId      string   `json:"instanceId"`
	Entity          Obj      `json:"entity"`
	ApplicationsIds []string `json:"applicationsIds"`
	ThingsProjectId string   `json:"firstProjectID"`
}

type Obj struct {
	Name         string `json:"name"`
	InstanceName string `json:"instanceName"`
}

