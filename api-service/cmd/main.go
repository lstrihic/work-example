package main

import "github.com/lstrihic/work-example/core/log"

var version = "develop"

func main() {
	log.InitLogger("dev", version)
	logger := log.GetLogger()
	logger.Info().Msg("Sample")
}
