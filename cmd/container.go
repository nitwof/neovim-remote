package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:     "container",
	Aliases: []string{"c"},
	Short:   "Neovim remote containers development",
	Long: `
		Lets you use a Docker container as a full-featured development environment
	`,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("docker-compose", "-f", "./.devcontainer/docker-compose.yml", "build")
		fmt.Println(c.String())
		cobra.CheckErr(c.Run())
	},
}

func init() {
	rootCmd.AddCommand(containerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
