package model

import (
	"fmt"
	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var Cmd = &cobra.Command{
	Use:   "model",
	Short: "A brief description of your command",
	Long:  `模型`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
	},
}

func init() {
	//rootCmd.AddCommand(modelCmd)
	//cmd.RootCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
