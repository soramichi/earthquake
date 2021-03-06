// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package searchtools

import (
	// "encoding/gob"
	"fmt"
	"os"

	// . "../equtils"

	"github.com/mitchellh/cli"
)

func run(args []string) {
	if len(args) != 1 {
		fmt.Printf("specify <storage dir path>\n")
		os.Exit(1)
	}

	storage := args[0]
	conf := storage + "/" + storageConfigPath

	cf, cerr := os.Open(conf)
	if cerr != nil {
		fmt.Printf("failed to open config file %s (%s)\n", conf, cerr)
		os.Exit(1)
	}

	fmt.Printf("ok: %s\n", cf)

}

type runCmd struct {
}

func (cmd runCmd) Help() string {
	return "run help (todo)"
}

func (cmd runCmd) Run(args []string) int {
	run(args)
	return 0
}

func (cmd runCmd) Synopsis() string {
	return "run subcommand"
}

func RunCommandFactory() (cli.Command, error) {
	return runCmd{}, nil
}
