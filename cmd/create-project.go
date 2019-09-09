package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var project_name string
var tmpl_type string

var CreateProject = &cobra.Command{
	Use:   "create-project",
	Short: "创建DogeGo项目.",
	Run:   createProject,
}

func createProject(cmd *cobra.Command, args []string) {
	log.Println(project_name)
}

func init() {
	CreateProject.Flags().StringVarP(&project_name, "name", "n", "", "项目名称")
	CreateProject.MarkFlagRequired("name")
	CreateProject.Flags().StringVarP(&tmpl_type, "type", "t", "", "项目类型(full or mini)")
}
