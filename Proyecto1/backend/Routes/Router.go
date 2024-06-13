package Routes

import (
	"context"
	"log"
	"main/Instance"
	"main/Model"
	"math/rand"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World")
	})

	app.Get("/ram", func(ctx *fiber.Ctx) error {
		nameCol := "ram"
		collection := Instance.Mg.Db.Collection(nameCol)

		err := collection.Drop(context.TODO())
		if err != nil {
			return err
		}

		dataParam := strconv.Itoa(rand.Intn(100))

		// collection = Instance.Mg.Db.Collection(nameCol)
		doc := Model.Ram{Libre: dataParam}

		_, err = collection.InsertOne(context.TODO(), doc)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(201).JSON(dataParam)
	})

	app.Get("/cpu", func(ctx *fiber.Ctx) error {
		log.Println("Insertando proceso")

		cmd := exec.Command("sleep", "infinity")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}

		cmd := exec.Command("kill", "-9", strconv.Itoa(pidInt))
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
		})
	})
}
