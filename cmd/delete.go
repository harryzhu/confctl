package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete --name=KEY",
	Long:  `-`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if Name == "" {
			log.Fatal("--name cannot be empty")
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("delete ...")
		config.Delete(Name)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVar(&Name, "name", "", "key name in conf datebase")

	deleteCmd.MarkFlagRequired("name")
}
