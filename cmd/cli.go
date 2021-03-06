/*
 * Bottlenet (C) 2020 MinIO, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	clientMode = false
	serverMode = false
)

var bottlenetCmd = &cobra.Command{
	Use: fmt.Sprintf("%s [IP...] [-a]", os.Args[0]),
	RunE: func(c *cobra.Command, args []string) error {
		return bottlenetEntrypoint(context.Background(), args)
	},
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	SilenceErrors:         true,
	Long: `
Bottlenet finds bottlenecks in your cluster network

Steps to find bottlenecks in network using bottlenet:

1. Run one instance of bottlenet on control node, where output will be collected:

  $>_ bottlenet

2. Run one instance of bottlenet on each of the peer nodes:

  $>_ bottlenet CONTROL-SERVER-IP:PORT

Once all the peer nodes have been added, press 'y' on the prompt (on control node) to start the tests.

In order to bind bottlenet to specific interface and port

  $>_ bottlenet --adddress IP:PORT

Note: --address should be applied to both control and peer nodes

  $>_ bottlenet --address IP:PORT CONTROL-SERVER-IP:PORT
`,
}

var (
	address = ":7007"
)

func init() {
	bottlenetCmd.PersistentFlags().StringVarP(&address, "address", "a", address, "listen address")
	// Turned-off for now
	// bottlenetCmd.PersistentFlags().BoolVarP(&clientMode, "client", "c", clientMode, "run in client mode")
	// bottlenetCmd.PersistentFlags().BoolVarP(&serverMode, "server", "s", serverMode, "run in server mode")
}

// Execute runs the binary
func Execute() error {
	return bottlenetCmd.Execute()
}
