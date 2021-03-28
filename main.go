package main

import (
	"errors"
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/aristat/http-server/internal/app/router"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "start",
				Aliases:     []string{"s"},
				Usage:       "run server",
				Description: "http server",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "envFile", Value: ".env.development"},
				},
				Action: func(c *cli.Context) error {
					fileName := c.String("envFile")

					err := godotenv.Load(fileName)
					if err != nil {
						return errors.New("Error loading .env file")
					}

					r, cleanup, err := router.Build()
					defer cleanup()
					if err != nil {
						return err
					}

					s := &http.Server{
						Handler: r,
						Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
					}

					err = s.ListenAndServe()
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("cli error ", err)
		os.Exit(1)
	}
}
