// Copyright 2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// policyGetCmd represents the policy_get command
var policyGetCmd = &cobra.Command{
	Use:   "get [<labels>]",
	Short: "Display policy node information",
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := client.PolicyGet(args); err != nil {
			Fatalf("Cannot get policy: %s\n", err)
		} else if resp != "[]" {
			fmt.Printf("%s\n", resp)
		}
	},
}

func init() {
	policyCmd.AddCommand(policyGetCmd)
}
