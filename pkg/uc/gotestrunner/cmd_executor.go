package gotestrunner

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type command struct {
	Name string
	Args []string
	Env  []string
}

func Command(name string, env []string, arg ...string) command {
	return command{Name: name, Args: arg, Env: env}
}

func (c command) Execute() (int, error) {
	return executeCommand(c, false)
}

func (c command) ExecuteWithLog() (int, error) {
	return executeCommand(c, true)
}

func (c command) ExecuteWithOutput() ([]byte, error) {
	return exec.Command(c.Name, c.Args...).CombinedOutput()
}

func executeCommand(c command, logs bool) (int, error) {
	if logs {
		commandStr := c.Name + " " + strings.Join(c.Args, " ")
		fmt.Println(commandStr)
	}

	cmd := exec.Command(c.Name, c.Args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "TEST_RUNNER_CONF_DEFAULT="+`{ "db": {   "databasetype": "postgresql",   "host": "localhost",   "port": 5432,   "database": "postgres",   "username": "postgres",   "password": "postgres",   "sslmode": "disable",   "scheme": "postgres",   "dynamic_port_name": "5432/tcp" }, "db_schema":{"path":"/schemas/postgres/test_schema.sql"} }`)
	cmd.Env = append(cmd.Env, c.Env...)

	stdOutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdOutPipe)
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = cmd.Wait()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode(), nil
		}
		return 0, err
	}

	return 0, nil
}
