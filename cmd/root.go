package cmd
import (
	"github.com/spf13/cobra"
	"fmt"
	"imageview/server"
	"os"
)

var Verbose bool
var Port int
var DownloadDir string

var rootCmd = &cobra.Command{
	Use: "image",
	Short: "image view is a simple http server",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
    Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hello word %v\n", args)
    	server.Start(Port)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8080, "port")
	rootCmd.PersistentFlags().StringVarP(&DownloadDir, "dir", "d", "download", "download dir")
}

