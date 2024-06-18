package Controller

import (
	"context"
	"log"
	"main/Instance"
	"main/Model"
)

func InsertRam(nameCol string, Total int, Enuso int, Libre int, Porcentaje int) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Ram{Total: Total, En_uso: Enuso, Libre: Libre, Porcentaje: Porcentaje}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}

func InserProcess(nameCol string, PID int, Name string, State int, PidPadre int) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Hijos{PID: PID, Name: Name, State: State, PIDPadre: PidPadre}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertCpu(nameCol string, Porcentaje int) error {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Cpu{Porcentaje: Porcentaje}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil
}

func ResetCollection(nameCol string) error {
	collection := Instance.Mg.Db.Collection(nameCol)
	err := collection.Drop(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
