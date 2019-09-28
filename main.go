package main

import (
	"golang_demo/routers"
)

func main() {
	router := routers.InitRouter()
	_ = router.Run(":80")
}
