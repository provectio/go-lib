package command

import "os/exec"

func Exec(command string, environments map[string]string, args ...string) (res []byte, err error) {

	cmd := exec.Command(command, args...)

	for key, value := range environments {
		cmd.Env = append(cmd.Env, key+"="+value)
	}

	res, err = cmd.CombinedOutput()

	return
}
