package Controller

import (
	"context"
	"log"
	"main/Instance"
	"main/Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// func InsertCpu(nameCol string, Porcentaje int) {
// 	collection := Instance.Mg.Db.Collection(nameCol)
// 	doc := Model.Cpu{Porcentaje: Porcentaje}
// 	_, err := collection.InsertOne(context.TODO(), doc)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func DeleteOldRecords(limit int) error {
	// Definir las colecciones a limpiar
	collections := []string{"cpu", "%cpu", "ram"}

	for _, collectionName := range collections {
		collection := Instance.Mg.Db.Collection(collectionName)

		// Establecer la opción de eliminación para limitar el número de registros a eliminar
		opts := options.Find().SetSort(bson.D{{"_id", 1}}).SetLimit(int64(limit))
		cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
		if err != nil {
			return err
		}

		var ids []interface{}
		for cursor.Next(context.TODO()) {
			var result struct {
				ID interface{} `bson:"_id"`
			}
			if err := cursor.Decode(&result); err != nil {
				return err
			}
			ids = append(ids, result.ID)
		}

		// Cerrar el cursor
		if err := cursor.Close(context.TODO()); err != nil {
			return err
		}

		// Si no hay registros para eliminar, continuar con la siguiente colección
		if len(ids) == 0 {
			continue
		}

		// Eliminar los registros encontrados
		_, err = collection.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
		if err != nil {
			return err
		}

		log.Printf("Eliminados %d registros de la colección %s\n", len(ids), collectionName)
	}

	return nil
}
