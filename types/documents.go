package types

import "go.mongodb.org/mongo-driver/bson/primitive"

//	type Documents struct {
//		// The ID of the document.
//		Path      string     `bson:"path"`
//		Documents []Document `bson:"documents"`
//	}
//
//	type Document struct {
//		Id  primitive.ObjectID `bson:"_id, omitempty"`
//		Doc string             `bson:"doc"`
//	}

type Documents struct {
	Path      string     `bson:"path"      json:"path"`
	Documents []Document `bson:"documents" json:"documents"`
}

type Document struct {
	Id  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Doc string             `bson:"doc"           json:"doc"`
}
