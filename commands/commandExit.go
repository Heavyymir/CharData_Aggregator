package commands

import (
	"fmt"
	"os"

	"github.com/Heavyymir/CharData_Aggregator/config"
)

func commandExit(cfg *config.Config, args ...string) error {
	fmt.Println("Closing... Goodbye!")
	os.Exit(0)
	return nil
}
