package run

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"libs.altipla.consulting/errors"
)

func Interactive(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return errors.Trace(cmd.Run())
}

func Shell(script string, vars ...map[string]string) error {
	script = "#!/bin/bash\nset -eu\n" + script

	for _, replace := range vars {
		for k, v := range replace {
			script = strings.Replace(script, "$"+k, v, -1)
		}
	}

	f, err := ioutil.TempFile("", "cdm")
	if err != nil {
		return errors.Trace(err)
	}
	defer os.Remove(f.Name())
	fmt.Fprint(f, script)
	if err := f.Close(); err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(Interactive("bash", f.Name()))
}

func InteractiveCaptureOutput(name string, args ...string) (string, error) {
	var buf bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", errors.Trace(err)
	}

	return strings.TrimSpace(buf.String()), nil
}

func WriteEnv(key, value string) error {
	if os.Getenv(env) != "" {
		return
	}

	os.Setenv(key, value)

	zsh, err := hasZshInstalled()
	if err != nil {
		return errors.Trace(err)
	}
	if !zsh {
		return nil
	}

	f, err := os.OpenFile("~/.zshrc", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Trace(err)
	}
	defer f.Close()

	if _, err = fmt.Fprintln(f, fmt.Sprintf("export %s=%s", key, value)); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func WriteAlias(env, key, value string) error {
	env := fmt.Sprintf("CONFMACHINE_%s", strings.ToUpper(env))
	if os.Getenv(env) != "" {
		return nil
	}

	if err := WriteEnv(env, "true"); err != nil {
		return errors.Trace(err)
	}

	fbash, err := os.OpenFile("~/.bashrc", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Trace(err)
	}
	defer fbash.Close()

	if _, err = fmt.Fprintln(fbash, fmt.Sprintf("alias %s='%s'", key, value)); err != nil {
		return errors.Trace(err)
	}

	zsh, err := hasZshInstalled()
	if err != nil {
		return errors.Trace(err)
	}
	if !zsh {
		return nil
	}

	fzsh, err := os.OpenFile("~/.zshrc", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Trace(err)
	}
	defer fzsh.Close()

	if _, err = fmt.Fprintln(fzsh, fmt.Sprintf("alias %s='%s'", key, value)); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func hasZshInstalled() (bool, error) {
	if _, err := os.Stat("~/.zshrc"); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, errors.Trace(err)
	}

	return true, nil
}
