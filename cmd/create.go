package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vlad-marlo/enrollment/pkg/httpclient"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a record in server",
	Long:  `Creates asks to enter user and message type. Then application will connect to server and store provided data.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := httpclient.New(cmd.Flag("server").Value.String())
		if err != nil {
			fmt.Printf("ERROR WHILE CREATING RECORDS\n\tError: %v\n", err)
			return
		}
		var user, msgType string
		fmt.Print("Enter User: ")
		if _, err := fmt.Scanln(&user); err != nil {
			fmt.Printf("ERROR WHILE GETTING USERNAME\n\tError: %v\n", err)
			return
		}
		fmt.Print("Enter MsgType: ")
		if _, err := fmt.Scanln(&msgType); err != nil {
			fmt.Printf("ERROR WHILE GETTING MSGTYPE\n\tError: %v\n", err)
			return
		}
		resp, err := cli.CreateRecord(user, msgType)
		if err != nil {
			fmt.Printf("ERROR WHILE CREATING RECORD\n\tError: %v\n", err)
			return
		}
		printResult(resp)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().String("server", "http://localhost:8080", "Address of running application.")
}
