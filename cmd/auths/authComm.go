package auths

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	datacfg string
	AuthCmd = &cobra.Command{
		Use:     "authcommand",
		Short:   "import batch commands from json.file",
		Example: "zeus authcommand -a config/auth.json",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
			run()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	AuthCmd.PersistenFlags().StringVarP(&datacfg, "datacfg", "d", "./data/auth.json")
}

func usage() {
	fmt.Println("insert auth command")
}

func setup() {
	fmt.Println("setup")
}

func run() error {
	_, err := importAuth()

	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	return nil
}

func importAuth()
