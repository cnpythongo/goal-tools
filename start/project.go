package start

import (
	"fmt"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/urfave/cli"
	"os/exec"
	"path/filepath"
)

const GoalHelperRepo = "https://github.com/cnpythongo/goal-helper.git"

func NewProject(c *cli.Context) error {
	args := c.Args()
	projectName := args.First()
	if projectName == "" {
		projectName = "goal-demo"
	}

	abs, err := filepath.Abs(projectName)
	if err != nil {
		return err
	}

	err = utils.MkdirIfNotExist(abs)
	if err != nil {
		return err
	}

	err = clone(abs)
	if err != nil {
		return err
	}

	appName := c.String("app")
	if appName == "" {
		appName = projectName
	}

	return nil
}

func clone(targetDir string) error {
	fmt.Printf("Create project %s, please wait...\n", targetDir)
	cmd := exec.Command("git", "clone", GoalHelperRepo, targetDir)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("cd %s\n", targetDir)
	cmd = exec.Command("/bin/sh", "-c", "cd", targetDir, "")
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Print("rm -rf .git\n")
	cmd = exec.Command("/bin/sh","-c", "rm", "-rf", ".git")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
