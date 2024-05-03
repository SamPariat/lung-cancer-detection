/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package predict_result

import (
	"fmt"

	"github.com/spf13/cobra"
)

// predictResultCmd represents the predictResult command
var PredictResultCmd = &cobra.Command{
	Use:   "predictResult",
	Short: "The command used for predicting lung cancer detection result",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("predictResult called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// predictResultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// predictResultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
