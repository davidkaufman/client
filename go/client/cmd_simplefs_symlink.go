// Copyright 2018 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
)

// CmdSimpleFSSymlink is the 'fs ln' command.
type CmdSimpleFSSymlink struct {
	libkb.Contextified
	src  string
	dest keybase1.Path
}

// NewCmdSimpleFSSymlink creates a new cli.Command.
func NewCmdSimpleFSSymlink(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:         "ln",
		ArgumentHelp: "<target> <link>",
		Usage:        "create a symlink from link to target",
		Action: func(c *cli.Context) {
			cl.ChooseCommand(&CmdSimpleFSSymlink{Contextified: libkb.NewContextified(g)}, "ln", c)
			cl.SetNoStandalone()
		},
	}
}

// Run runs the command in client/server mode.
func (c *CmdSimpleFSSymlink) Run() error {
	cli, err := GetSimpleFSClient(c.G())
	if err != nil {
		return err
	}

	ctx := context.TODO()
	arg := keybase1.SimpleFSSymlinkArg{
		Src:  c.src,
		Dest: c.dest,
	}

	return cli.SimpleFSSymlink(ctx, arg)
}

// ParseArgv gets the required path argument for this command.
func (c *CmdSimpleFSSymlink) ParseArgv(ctx *cli.Context) error {
	nargs := len(ctx.Args())
	if nargs != 2 {
		return errors.New("ln requires exactly 2 arguments")
	}
	srcStr := ctx.Args()[0]
	destStr := ctx.Args()[1]

	rev := int64(0)
	timeString := ""
	relTimeString := ""
	destPath, err := makeSimpleFSPathWithArchiveParams(
		destStr, rev, timeString, relTimeString)
	if err != nil {
		return err
	}

	destPathType, err := destPath.PathType()
	if err != nil {
		return err
	}
	if destPathType != keybase1.PathType_KBFS {
		return errors.New("keybase fs ln: link must be a KBFS path")
	}
	c.dest = destPath

	c.src = srcStr

	return err
}

// GetUsage says what this command needs to operate.
func (c *CmdSimpleFSSymlink) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config:    true,
		KbKeyring: true,
		API:       true,
	}
}
