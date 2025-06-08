package main

import (
	"log"

	"github.com/erodriguezg/gueshop/internal/app"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env (puede no ser necesario en producción)")
	}

	fx.New(
		app.Modules, // app/fx.go
	).Run()
}
