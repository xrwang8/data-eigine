package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "data-eigine",
	Short: "Provide dataset and model engine",
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
	TraverseChildren: true,
}

// App root of application
type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	cmd.AddCommand(GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
