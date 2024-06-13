package Routes

import (
	"context"
	"log"
	"main/Database"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World")
	})

	// app.Get("/ram", func(ctx *fiber.Ctx) error {
	// 	nameCol := "ram"
	// 	collection := Instance.Mg.Db.Collection(nameCol)

	// 	err := collection.Drop(context.TODO())
	// 	if err != nil {
	// 		return err
	// 	}

	// 	cmd := exec.Command("sh", "-c", "cat /proc/ram_201800722")
	// 	output, err := cmd.CombinedOutput()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	out := string(output[:])

	// 	doc := Model.Ram{Libre: out}

	// 	_, err = collection.InsertOne(context.TODO(), doc)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	return ctx.JSON(fiber.Map{
	// 		"status":  200,
	// 		"percent": out,
	// 	})
	// })

	app.Get("/cpu", func(ctx *fiber.Ctx) error {
		log.Println("Insertando proceso")

		cmd := exec.Command("sleep", "infinity")
		err := cmd.Start()
		if err != nil {
			log.Printf("Error starting process: %v\n", err)
			return ctx.Status(500).SendString("Error starting process")
		}

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"pid":     cmd.Process.Pid,
		})
	})

	app.Get("/delProceso", func(ctx *fiber.Ctx) error {
		pid := ctx.Query("pid")
		pidInt, err := strconv.Atoi(pid)
		if err != nil {
			log.Printf("Error converting pid: %v\n", err)
			return ctx.Status(400).SendString("Invalid pid")
		}

		cmd := exec.Command("kill", "-9", strconv.Itoa(pidInt))
		err = cmd.Run()
		if err != nil {
			log.Printf("Error killing process: %v\n", err)
			return ctx.Status(500).SendString("Error killing process")
		}

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
		})
	})

	app.Get("/dbstatus", func(ctx *fiber.Ctx) error {
		// Comprobar el estado de la conexión a la base de datos
		err := Database.Mg.Ping(context.TODO(), nil)
		if err != nil {
			log.Printf("Database connection is not alive: %v\n", err)
			return ctx.Status(500).SendString("Database connection is not alive")
		}
		return ctx.Status(200).SendString("Database connection is alive")
	})
}

// package Routes

// import (
// 	"context"
// 	"log"
// 	"main/Instance"
// 	"main/Model"
// 	"math/rand"
// 	"os/exec"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// )

// func Setup(app *fiber.App) {
// 	app.Get("/", func(ctx *fiber.Ctx) error {
// 		return ctx.JSON("Hello World")
// 	})

// 	app.Get("/ram", func(ctx *fiber.Ctx) error {
// 		nameCol := "ram"
// 		collection := Instance.Mg.Db.Collection(nameCol)

// 		err := collection.Drop(context.TODO())
// 		if err != nil {
// 			return err
// 		}

// 		dataParam := strconv.Itoa(rand.Intn(100))

// 		// collection = Instance.Mg.Db.Collection(nameCol)
// 		doc := Model.Ram{Libre: dataParam}

// 		_, err = collection.InsertOne(context.TODO(), doc)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		return ctx.Status(201).JSON(dataParam)
// 	})

// 	app.Get("/cpu", func(ctx *fiber.Ctx) error {
// 		log.Println("Insertando proceso")

// 		cmd := exec.Command("sleep", "infinity")
// 		err := cmd.Start()
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		return ctx.Status(200).JSON(fiber.Map{
// 			"success": true,
// 			"pid":     cmd.Process.Pid,
// 		})
// 	})

// 	app.Get("/delProceso", func(ctx *fiber.Ctx) error {
// 		pid := ctx.Query("pid")
// 		pidInt, err := strconv.Atoi(pid)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		cmd := exec.Command("kill", "-9", strconv.Itoa(pidInt))
// 		err = cmd.Run()
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		return ctx.Status(200).JSON(fiber.Map{
// 			"success": true,
// 		})
// 	})
// }
