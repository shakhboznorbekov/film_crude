package models

type Actor struct {
	ActorId    string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	LastUpdate string `json:"last_update"`
}

type UserData struct {
	Id   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}
