package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func SH(c string) (string, error) {
	fmt.Printf("Command %s\n", c)
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	o, err := cmd.CombinedOutput()
	fmt.Printf("env %+v\n", cmd.Env)
	fmt.Printf("Command %s %s\n", string(o), err.Error())
	return string(o), err
}

func WriteEnv(envFile string, config map[string]string) error {
	content, _ := ioutil.ReadFile(envFile)
	env, _ := godotenv.Unmarshal(string(content))

	for key, val := range config {
		env[key] = val
	}

	return godotenv.Write(env, envFile)
}

func Shell() *exec.Cmd {
	cmd := exec.Command("/bin/sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd
}

func ShellSTDIN(s, c string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = bytes.NewBuffer([]byte(s))
	o, err := cmd.CombinedOutput()
	return string(o), err
}
