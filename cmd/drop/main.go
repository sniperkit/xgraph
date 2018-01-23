package drop

import (
	//"fmt"
	"github.com/bmeg/arachne/aql"
	"github.com/spf13/cobra"
)

var host string = "localhost:9090"

var Cmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop Graph on Arachne Server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := aql.Connect(host, true)
		if err != nil {
			return err
		}
		conn.DeleteGraph(args[0])
		return nil
	},
}

func init() {
	flags := Cmd.Flags()
	flags.StringVar(&host, "host", "localhost:9090", "Host Server")
}
