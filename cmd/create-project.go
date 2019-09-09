package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var project_name string
var tmpl_type string

var CreateProject = &cobra.Command{
	Use:   "create-project",
	Short: "创建DogeGo项目.",
	Run:   createProject,
}

func replaceProjectName() {
	err := filepath.Walk(project_name, func(path string, f os.FileInfo, err error) error {
		var rb string

		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		buffer := make([]byte, f.Size())

		fp, _ := os.Open(path)

		fp.Read(buffer)
		fp.Close()

		if tmpl_type == "mini" {
			rb = strings.Replace(string(buffer), "dogego-mini", project_name, -1)
		} else {
			rb = strings.Replace(string(buffer), "dogego", project_name, -1)
		}

		fp, _ = os.OpenFile(path, os.O_WRONLY, 666)
		fp.WriteString(rb)
		fp.Close()
		log.Printf("处理: %s", path)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func createProject(cmd *cobra.Command, args []string) {
	var cmdx *exec.Cmd

	if tmpl_type == "" {
		tmpl_type = "full"
	}

	log.Println("创建项目中.........")

	if tmpl_type == "mini" {
		cmdx = exec.Command(
			"git",
			"clone",
			"https://github.com/Xiangrui2019/dogego",
			project_name)
	} else {
		cmdx = exec.Command(
			"git",
			"clone",
			"https://github.com/Xiangrui2019/dogego-mini",
			project_name)
	}
	wd, _ := os.Getwd()

	cmdx.Stdout = os.Stdout
	cmdx.Stderr = os.Stdout
	cmdx.Stdin = os.Stdin
	cmdx.Dir = wd
	err := cmdx.Run()

	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("rm", "-rf", project_name+"/.git").Run()

	if err != nil {
		log.Fatal(err)
	}

	replaceProjectName()
}

func init() {
	CreateProject.Flags().StringVarP(&project_name, "name", "n", "", "项目名称")
	CreateProject.MarkFlagRequired("name")
	CreateProject.Flags().StringVarP(&tmpl_type, "type", "t", "", "项目类型(full or mini)")
}
