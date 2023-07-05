package main

import (
	"os"

	"github.com/PikeEcosystem/cosmos-sdk/server"
	svrcmd "github.com/PikeEcosystem/cosmos-sdk/server/cmd"

	"github.com/PikeEcosystem/ibc-go/v3/testing/simapp"
	"github.com/PikeEcosystem/ibc-go/v3/testing/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
