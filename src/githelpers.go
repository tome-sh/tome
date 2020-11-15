package tome

import (
	"os/exec"
	"strings"
)

func GetGitConfigSetting(key string) (string, error) {
	app := "git"

	arg0 := "config"
	arg1 := "--includes"
	arg2 := "--get"

	cmd := exec.Command(app, arg0, arg1, arg2, key)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}
