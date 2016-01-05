// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build !windows

package minterm

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func (m *MinTerm) open() error {
	f, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	m.termFile = f
	fd := int(m.termFile.Fd())
	w, h, err := terminal.GetSize(fd)
	if err != nil {
		return err
	}
	m.width, m.height = w, h
	return nil
}
