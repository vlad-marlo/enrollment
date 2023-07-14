package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vlad-marlo/enrollment/pkg/httpclient"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get records from server",
	Long: `Command has 3 modes:
1 - Get record by id. User has to enter integer identifier to get record by it.
2 - Get all records associated to user. User must enter username after it program will return all associated
records.
3 - Get all records that are stored in service.
Records are sorted by creation date in descending order - newer records must be first in response then older.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := httpclient.New(cmd.Flag("server").Value.String())
		if err != nil {
			fmt.Printf("ERROR WHILE CREATING RECORDS\n\tError: %v\n", err)
			return
		}

		var mode int

		for {
			fmt.Print("Select mode:\n\t1. Get by ID.\n\t2. Get all by user.\n\t3. Get all records.\n")
			_, err := fmt.Scanln(&mode)
			if err == nil && 1 <= mode && mode <= 3 {
				break
			}
		}

		switch mode {
		case 1:
			getByID(cli)
		case 2:
			getByUser(cli)
		case 3:
			getAll(cli)
		}
	},
}

func getAll(cli *httpclient.Client) {
	resp, err := cli.GetAll()
	if err != nil {
		fmt.Printf("ERROR WHILE GETTING RECORDS\n\tError: %v\n", err)
		return
	}
	printResult(resp)
}

func getByUser(cli *httpclient.Client) {
	var user string
	fmt.Println("Type user identifier")
	if _, err := fmt.Scanln(&user); err != nil {
		fmt.Printf("ERROR WHILE GETTING USER IDENTIFIER\nError: %v\n", err)
		return
	}
	resp, err := cli.GetRecordsByUser(user)
	if err != nil {
		fmt.Printf("ERROR WHILE GETTING RECORD\n\tError: %v\n", err)
		return
	}
	printResult(resp)
}

func getByID(cli *httpclient.Client) {
	var id int64
	for {
		fmt.Println("Type integer identifier")
		_, err := fmt.Scanln(&id)
		if err == nil {
			break
		}
	}
	resp, err := cli.GetRecordByID(id)
	if err != nil {
		fmt.Printf("ERROR WHILE GETTING RECORD\n\tError: %v\n", err)
		return
	}
	printResult(resp)
}

func printResult(result any) {
	got, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Printf("ERROR WHILE PRINITING RESULT\n\tError: %v\n", err)
		return
	}
	fmt.Printf("%s\n", got)
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().String("server", "http://localhost:8080", "Address of running application.")
}
