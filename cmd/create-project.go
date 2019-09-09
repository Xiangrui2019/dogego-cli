package cmd

import "github.com/spf13/cobra"

var CreateProject = &cobra.Command{
	Use:   "create-project",
	Short: "创建DogeGo项目.",
	Run:   createProject,
}

func createProject(cmd *cobra.Command, args []string) {

}
