package Controller

import (
	"context"
	"log"
	"main/Instance"
	"main/Model"
)

func InsertarProcesos(nameCol string, dataParam string) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Ram{Libre: dataParam}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

}
