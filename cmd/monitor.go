package cmd

import (
	"nats/internal/monitor"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type TemplateVars struct {
	natsServer  string
	natsSubject string
}

var templateVars TemplateVars

// publishCmd represents the publish command
var monitorCmd = &cobra.Command{
	Use:     "monitor",
	Aliases: []string{"mon"},
	Short:   "Monitoring subjects",
	Args:    cobra.MinimumNArgs(0),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := monitor.Run(
			templateVars.natsServer,
			templateVars.natsSubject); err != nil {
			logrus.Error(err)
			return
		}
	},
}

func init() {
	monitorCmd.Flags().StringVarP(&templateVars.natsServer, "addr", "a", "", "NATS server addr")
	monitorCmd.Flags().StringVarP(&templateVars.natsSubject, "subject", "s", "", "subject name")

	// if err := monitorCmd.MarkFlagRequired("addr"); err != nil {
	// 	logrus.Fatal(err)
	// }
	// if err := monitorCmd.MarkFlagRequired("subject"); err != nil {
	// 	logrus.Fatal(err)
	// }

	rootCmd.AddCommand(monitorCmd)
}
