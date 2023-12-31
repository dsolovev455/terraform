// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build linux || darwin || openbsd || netbsd || solaris || dragonfly
// +build linux darwin openbsd netbsd solaris dragonfly

package configload

import (
	"fmt"
	"os"
	"syscall"
)

// lookup the inode of a file on posix systems
func inode(path string) (uint64, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	if st, ok := stat.Sys().(*syscall.Stat_t); ok {
		return st.Ino, nil
	}
	return 0, fmt.Errorf("could not determine file inode")
}
