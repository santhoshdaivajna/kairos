package utils

import (
	"bytes"

	"github.com/sirupsen/logrus"

	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func SHInDir(c, dir string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", c)
	logrus.Infof("command %s %s", c, dir)
	logrus.Infof("environ %s", os.Environ())
	cmd.Env = os.Environ()
	cmd.Dir = dir
	o, err := cmd.CombinedOutput()
	return string(o), err
}

func SH(c string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	logrus.Infof("command %s", c)
	logrus.Infof("environ %s", os.Environ())
	o, err := cmd.CombinedOutput()
	return string(o), err
}

func WriteEnv(envFile string, config map[string]string) error {
	content, _ := os.ReadFile(envFile)
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
