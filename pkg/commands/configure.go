package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "configure",
	Short:                 "Configure fgh with an interactive prompt",
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#%EF%B8%8F-fgh-configure",
	Run: func(cmd *cobra.Command, args []string) {
		config := configure.AskQuestions()
		configure.WriteConfig(config)
		statuser.Success("Wrote to config")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
