package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/marioharper/commuter/cmd/utils"
	"github.com/marioharper/commuter/directions"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your commuter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var addresses []directions.Location

		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(fmt.Sprintf("%s/commuter-config.json", usr.HomeDir))
		utils.Check(err)

		workReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter work address: ")

		// Get work address
		workAddress, _ := workReader.ReadString('\n')
		workAddress = strings.TrimSpace(workAddress)

		work := directions.Location{
			Name:    "work",
			Address: workAddress,
		}
		addresses = append(addresses, work)

		// Get home address
		homeReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter home address: ")
		homeAddress, _ := homeReader.ReadString('\n')
		homeAddress = strings.TrimSpace(homeAddress)
		home := directions.Location{
			Name:    "home",
			Address: homeAddress,
		}
		addresses = append(addresses, home)
		addressesJSON, _ := json.Marshal(addresses)

		// Write to config file
		w := bufio.NewWriter(f)
		w.WriteString(string(addressesJSON))
		w.Flush()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
