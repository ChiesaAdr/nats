package cmd

import (
	"nats/internal/publish"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var publishVars struct {
	natsServer     string
	natsSubject    string
	natsQuantity   uint
	startDeltaFrom time.Duration
	message        string
}

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"pub"},
	Short:   "Publish to subject",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := publish.Run(
			publishVars.natsServer,
			publishVars.natsSubject,
			publishVars.natsQuantity,
			publishVars.startDeltaFrom,
			[]byte(publishVars.message)); err != nil {
			logrus.Error(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.Flags().StringVarP(&publishVars.natsServer, "addr", "a", "", "NATS server addr")
	publishCmd.Flags().StringVarP(&publishVars.natsSubject, "subject", "s", "", "subject name")
	publishCmd.Flags().UintVarP(&publishVars.natsQuantity, "quantity", "q", 0, "quantity of messages")
	publishCmd.Flags().DurationVarP(&publishVars.startDeltaFrom, "delta-time", "d", 0, "delta-time")
	// publishCmd.Flags().StringVarP(&publishVars.message, "message", "m", "", "JSON message")

	// if err := publishCmd.MarkFlagRequired("addr"); err != nil {
	// 	logrus.Fatal(err)
	// }

	// if err := publishCmd.MarkFlagRequired("subject"); err != nil {
	// 	logrus.Fatal(err)
	// }

	// if err := publishCmd.MarkFlagRequired("message"); err != nil {
	// 	logrus.Fatal(err)
	// }

}
