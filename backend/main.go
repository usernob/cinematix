package main

import (
	"backend/model"
	"backend/router"
)

func main() {
  model.SeedRefresh()

	r := router.SetupRouter()
	r.Run(":4000")
}
