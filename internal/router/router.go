package router

import (
	"context"
	"log"

	"github.com/aliirsyaadn/kodein/grpc/grader"
	"google.golang.org/grpc"

	"github.com/aliirsyaadn/kodein/handlers"
	"github.com/aliirsyaadn/kodein/model"
	"github.com/aliirsyaadn/kodein/services/member"
	fiber "github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App, model *model.Queries){
	api := app.Group("/api")

	// Member
	memberService := member.NewService(model)
	handlers.NewMemberHandler(memberService).Register(api)

	api.Post("/grade", Grade)
}

func Grade(c *fiber.Ctx) error {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	service := grader.NewGraderServiceClient(conn)

	message := grader.Message{
		Body: "Hello from the client!",
	}

	res, err := service.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", res.Body)
	return nil
}