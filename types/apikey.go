package types

import "time"

type ApiKey struct {
	Id        string    `bson:"_id, omitempty"`
	CreatedBy string    `bson:"created"`
	CreatedAt time.Time `bson:"created_at"`
	Disabled  bool      `bson:"disabled"`
}

func CreateNewApiKey(id string) ApiKey {
	return ApiKey{
		Id:        id,
		CreatedBy: "cli", // TODO: change this to the user's email address
		CreatedAt: time.Now(),
		Disabled:  false,
	}
}
