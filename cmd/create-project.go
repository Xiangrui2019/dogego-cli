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
	Use:   "create",
	Short: "创建新的DogeGo项目.",
	Run:   createProject,
}

func defaultValue() {
	if tmpl_type == "" {
		tmpl_type = "full"
	}
}

func ProjectTypeGit(workdir string) *exec.Cmd {
	var gitCommand *exec.Cmd

	switch tmpl_type {
	case "mini":
		gitCommand = exec.Command(
			"git",
			"clone",
			"https://github.com/Xiangrui2019/dogego-mini",
			project_name)
	case "jrpc":
		gitCommand = exec.Command(
			"git",
			"clone",
			"https://github.com/Xiangrui2019/jrpc",
			project_name)
	case "jrpc-orm":
		gitCommand = exec.Command(
			"git",
			"clone",
			"https://github.com/Xiangrui2019/jrpc-orm",
			project_name)
	default:
		gitCommand = exec.Command(
			"git",
			"clone",
			"https://github.com/Xiangrui2019/dogego",
			project_name)
	}

	gitCommand.Stdout = os.Stdout
	gitCommand.Stderr = os.Stdout
	gitCommand.Stdin = os.Stdin
	gitCommand.Dir = workdir

	return gitCommand
}

func replacer(s string) string {
	var result string

	switch tmpl_type {
	case "mini":
		result = strings.Replace(s, "dogego-mini", project_name, -1)
	case "jrpc":
		result = strings.Replace(s, "jrpc", project_name, -1)
	case "jrpc-orm":
		result = strings.Replace(s, "jrpc-orm", project_name, -1)
	default:
		result = strings.Replace(s, "dogego", project_name, -1)
	}

	return result
}

func replaceProjectName() {
	err := filepath.Walk(project_name, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		buffer := make([]byte, f.Size())

		fp, err := os.Open(path)
		fp.Read(buffer)
		fp.Close()

		result := replacer(string(buffer))

		fp, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 666)
		if err != nil {
			log.Fatal(err)
		} else {
			fp.WriteString(result)
			log.Printf(path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func createProject(cmd *cobra.Command, args []string) {
	defaultValue()

	workdir, _ := os.Getwd()
	command := ProjectTypeGit(workdir)

	err := command.Run()
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
	CreateProject.Flags().StringVarP(&tmpl_type, "type", "t", "", "项目类型(full or mini or jrpc or jrpc-orm)")
}
