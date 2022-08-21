package cmd

import (
	"data-engine/utils"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "data-engine",
	Short:        "Provide dataset and model engine",
	SilenceUsage: true,
	Long: `
 ███████       ██     ██████████     ██                 ████████ ████     ██   ████████  ██ ████     ██ ████████
░██░░░░██     ████   ░░░░░██░░░     ████               ░██░░░░░ ░██░██   ░██  ██░░░░░░██░██░██░██   ░██░██░░░░░ 
░██    ░██   ██░░██      ░██       ██░░██              ░██      ░██░░██  ░██ ██      ░░ ░██░██░░██  ░██░██      
░██    ░██  ██  ░░██     ░██      ██  ░░██  █████ █████░███████ ░██ ░░██ ░██░██         ░██░██ ░░██ ░██░███████ 
░██    ░██ ██████████    ░██     ██████████░░░░░ ░░░░░ ░██░░░░  ░██  ░░██░██░██    █████░██░██  ░░██░██░██░░░░  
░██    ██ ░██░░░░░░██    ░██    ░██░░░░░░██            ░██      ░██   ░░████░░██  ░░░░██░██░██   ░░████░██      
░███████  ░██     ░██    ░██    ░██     ░██            ░████████░██    ░░███ ░░████████ ░██░██    ░░███░████████
░░░░░░░   ░░      ░░     ░░     ░░      ░░             ░░░░░░░░ ░░      ░░░   ░░░░░░░░  ░░ ░░      ░░░ ░░░░░░░░ 
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(utils.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + utils.Green(`data-engine`) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(ServeCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
