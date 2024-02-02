package main

import (
	"backend/model"
	"backend/model/seeder"
	"backend/router"
	"flag"
)

func main() {

	var migrate, seed, refresh, refreshSeed bool

	flag.BoolVar(&refreshSeed, "refreshseed", false, "Refresh the database")
	flag.BoolVar(&refresh, "refresh", false, "Refresh the database")
	flag.BoolVar(&migrate, "migrate", false, "Migrate the schemas")
	flag.BoolVar(&seed, "seed", false, "Seed the database")

	flag.Parse()

	if refreshSeed {
		model.Refresh()
		model.Migrate()
		seeder.Seed()
		return
	}

	if refresh {
		model.Refresh()
		return
	}

	if migrate {
		model.Migrate()
		return
	}

	if seed {
		seeder.Seed()
		return
	}

	r := router.SetupRouter()
	r.Run(":4000")
}
