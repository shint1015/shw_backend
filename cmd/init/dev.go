package main

import (
	"shwgrpc/db"
	"shwgrpc/model"
)

func main() {
	model.Init()
	db.DoSeed()
}
