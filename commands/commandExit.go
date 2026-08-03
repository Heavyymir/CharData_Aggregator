package commands

import(
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
        fmt.Println("Closing... Goodbye!")
        os.Exit(0)
        return nil
}
