package cmd

import (
	"fmt"
	"log"

	"github.com/diazharizky/go-cmd-embed/pkg/mariadb"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate-up",
	Short: "Run db migration up.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running db migration up...")

		if err := mariadb.MigrateUp(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Success!")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate-down",
	Short: "Run db migration down.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running db migration down...")

		if err := mariadb.MigrateDown(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Success!")
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd, migrateDownCmd)
}
