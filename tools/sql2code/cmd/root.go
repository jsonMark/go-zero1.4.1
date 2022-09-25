package cmd

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/sql2code/api"
	"github.com/zeromicro/go-zero/tools/sql2code/gen"
	"github.com/zeromicro/go-zero/tools/sql2code/model"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	codeFailure = 1
	//dash        = "-"
	//doubleDash  = "--"
	//assign      = "="
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sql2code",
	Short: "基于goctl的代码生成工具",
	Long:  `代码生成工具`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(gen.Cmd)
	rootCmd.AddCommand(model.Cmd)
	rootCmd.AddCommand(api.Cmd)
	//RootCmd.AddCommand(gen.Cmd)
}
