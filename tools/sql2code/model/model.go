package model

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/command"
)

// modelCmd represents the model command
var Cmd = &cobra.Command{
	Use:   "model",
	Short: "A brief description of your command",
	Long:  `模型`,
	/*	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("我收到回复看见啊师傅model called")
	},*/
	RunE: MysqlDDL,
}

func init() {
	Cmd.Flags().StringVarP(&command.VarStringSrc, "src", "s", "", "The path or path globbing patterns of the ddl")
	Cmd.Flags().StringVarP(&command.VarStringDir, "dir", "d", "", "The target dir")
	Cmd.Flags().StringVar(&command.VarStringStyle, "style", "", "The file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]")
	//Cmd.Flags().BoolVarP(&command.VarBoolCache, "cache", "c", false, "Generate code with cache [optional]")
	//Cmd.Flags().BoolVar(&command.VarBoolIdea, "idea", false, "For idea plugin [optional]")
	Cmd.Flags().StringVar(&command.VarStringDatabase, "database", "", "The name of database [optional]")
	//Cmd.Flags().StringVar(&command.VarStringHome, "home", "", "The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority")
	//Cmd.Flags().StringVar(&command.VarStringRemote, "remote", "", "The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority\nThe git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure")
	//Cmd.Flags().StringVar(&command.VarStringBranch, "branch", "", "The branch of the remote repo, it does work with --remote")

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
