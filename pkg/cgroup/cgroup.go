// SPDX-License-Identifier: BSD-3-Clause

package cgroup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	CgroupPath = "/sys/fs/cgroup"
	SlicePath  = "/sys/fs/cgroup/u-bmc.slice"
)

type Cgroup struct {
	path string
}

func New(name string) (*Cgroup, error) {
	cg := Cgroup{
		path: filepath.Join(SlicePath, name),
	}

	stat, err := os.Stat(CgroupPath)
	if err != nil {
		return nil, fmt.Errorf("cgroup mount doesn't exist: %w", err)
	} else if !stat.IsDir() {
		return nil, fmt.Errorf("%q: cgroup mount path is not a dir", CgroupPath)
	}

	if _, err = os.Stat(cg.path); !os.IsNotExist(err) {
		return nil, fmt.Errorf("cgroup with the same name %q already exists", cg.path)
	}

	if err := os.MkdirAll(cg.path, 0o644); err != nil {
		return nil, fmt.Errorf("cannot create cgroup at %q: %w", cg.path, err)
	}

	return &cg, nil
}

func (cg *Cgroup) Run() error {
	if err := unix.Prctl(unix.PR_SET_CHILD_SUBREAPER, 1, 0, 0, 0); err != nil {
		return fmt.Errorf("cannot set child subreaper: %w", err)
	}

	cmdName := filepath.Base(cg.path)
	c, err := exec.LookPath(cmdName)
	if err != nil {
		return fmt.Errorf("cannot find command: %w", err)
	}

	procsHandle, err := os.OpenFile(filepath.Join(cg.path, "cgroup.procs"), os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("cannot open handle to cgroup.procs: %w", err)
	}
	defer procsHandle.Close()

	if _, err := procsHandle.Write([]byte("0")); err != nil {
		return fmt.Errorf("cannot write to cgroup.procs: %w", err)
	}

	cmd := exec.Cmd{
		Path: c,
		Env:  os.Environ(),

		ExtraFiles: []*os.File{
			procsHandle,
		},

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("error on child process %q: %w", cmdName, err)
	}

	return nil
}

func (cg *Cgroup) Teardown() error {
	pids, err := os.ReadFile(filepath.Join(cg.path, "cgroup.procs"))
	if err != nil {
		return fmt.Errorf("cannot get pids in cgroup: %w", err)
	}

	pidNums, err := atoiAll(strings.Split(string(pids), "\n"))
	if err != nil {
		return fmt.Errorf("error parsing pids: %w", err)
	}

	wait := make([]*os.Process, 0, len(pidNums))

	for _, pid := range pidNums {
		proc, _ := os.FindProcess(pid)
		if err := proc.Signal(unix.SIGTERM); err == nil {
			wait = append(wait, proc)
		}
	}

	for _, proc := range wait {
		if _, err := proc.Wait(); err != nil {
			return fmt.Errorf("error waiting for process %d: %w", proc.Pid, err)
		}
	}

	if err := os.Remove(cg.path); err != nil {
		return fmt.Errorf("cannot remove cgroup at %q: %w", cg.path, err)
	}

	return nil
}

func ExecShim() error {
	// FD 0-2 are reserved for stdio, next FD is 3
	handle := os.NewFile(3, "cgroup.procs")
	if handle == nil {
		return fmt.Errorf("failed to open cgroup")
	}

	if _, err := handle.Write([]byte("0")); err != nil {
		return fmt.Errorf("failed to write to cgroup.procs: %w", err)
	}

	_ = handle.Close() // We can ignore this error

	cmd, err := exec.LookPath(os.Args[0])
	if err != nil {
		return fmt.Errorf("cannot find command: %w", err)
	}

	return unix.Exec(cmd, os.Args, os.Environ())
}

func atoiAll(strs []string) ([]int, error) {
	out := make([]int, 0, len(strs))

	for _, str := range strs {
		if str == "" {
			continue
		}

		num, err := strconv.Atoi(str)
		if err != nil {
			return out, fmt.Errorf("cannot convert %q to int: %w", str, err)
		}

		out = append(out, num)
	}

	return out, nil
}
