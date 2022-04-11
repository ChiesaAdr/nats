package cmd

import (
	"github.com/spf13/cobra"

	"nats/internal/consumer"
)

var subscribeVars struct {
	natsServer  string
	natsSubject string
}

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:     "consumer",
	Aliases: []string{"cs"},
	Short:   "Consumer the subject",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		consumer.Run(
			subscribeVars.natsServer,
			subscribeVars.natsSubject)
	},
}

func init() {

	subscribeCmd.Flags().StringVarP(&subscribeVars.natsServer, "addr", "a", "", "NATS server addr")
	subscribeCmd.Flags().StringVarP(&subscribeVars.natsSubject, "subject", "s", "", "subject name")

	// if err := subscribeCmd.MarkFlagRequired("addr"); err != nil {
	// 	logrus.Fatal(err)
	// }

	// if err := subscribeCmd.MarkFlagRequired("subject"); err != nil {
	// 	logrus.Fatal(err)
	// }

	rootCmd.AddCommand(subscribeCmd)
}
