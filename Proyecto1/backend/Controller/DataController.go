package Controller

import (
	"context"
	"fmt"
	"log"
	"main/Instance"
	"main/Model"
	//"github.com/gofiber/fiber/v2"
)

func InsertRam(nameCol string, Total int, Enuso int, Libre int, Porcentaje int) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Ram{Total: Total, En_uso: Enuso, Libre: Libre, Porcentaje: Porcentaje}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertData2(nameCol string, PID int, Name string, State int, PidPadre int) {

	collection := Instance.Mg.Db.Collection(nameCol)

	doc := Model.Hijos{PID: PID, Name: Name, State: State, PIDPadre: PidPadre}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertData1(nameCol string) error {
	collection := Instance.Mg.Db.Collection(nameCol)

	err := collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("en poceso de guardar datos")
	return nil
}
