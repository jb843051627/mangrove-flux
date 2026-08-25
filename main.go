package main

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/handler"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"github.com/jb843051627/mangrove-flux/internal/store"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--smoke-test" {
		if err := smoke(); err != nil {
			log.Fatal(err)
		}
		log.Println("mangrove-flux smoke test passed")
		return
	}
	path := os.Getenv("MANGROVE_FLUX_DB")
	if path == "" {
		path = filepath.Join("data", "mangrove-flux.db")
	}
	db, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app := service.NewLab(db)
	defer app.Close()
	addr := os.Getenv("MANGROVE_FLUX_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("mangrove-flux listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler.New(app)))
}

func smoke() error {
	dir, err := os.MkdirTemp("", "mangrove-flux-smoke-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	db, err := store.Open(filepath.Join(dir, "smoke.db"))
	if err != nil {
		return err
	}
	defer db.Close()
	app := service.NewLab(db)
	defer app.Close()
	ctx := context.Background()
	station := model.FieldStation{ID: "smoke-station", Name: "潮间带样地", Region: "delta", Timezone: "Asia/Shanghai", TideDatum: "local", Active: true}
	if err := app.CreateStation(ctx, station); err != nil {
		return err
	}
	plot := model.Plot{ID: "smoke-plot", StationID: station.ID, Name: "红树林边缘", Habitat: "fringe", AreaM2: 20, TargetDepthM: 0.4}
	if err := app.CreatePlot(ctx, plot); err != nil {
		return err
	}
	return nil
}
