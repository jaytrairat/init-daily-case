package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "init-daily-report",
		Short: "Initialize the daily report",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			folder_list := []string{"Data", "Recon", "Report"}
			for _, folder := range folder_list {
				err := os.MkdirAll(fmt.Sprintf("%s_daily_report/%s", time.Now().Format("2006-01-02"), folder), 0755)
				if err != nil {
					fmt.Println("Error creating folder:", err)
					return
				}

			}

		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
