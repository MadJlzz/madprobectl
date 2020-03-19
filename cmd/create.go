/*
Copyright © 2020 Julien KLAER <klaer.julien@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/madjlzz/madprobectl/internal/endpoints"
	"github.com/madjlzz/madprobectl/internal/parser"
	"github.com/madjlzz/madprobectl/internal/service"
	"github.com/spf13/cobra"
)

var (
	File string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new probe and run it",
	Run: func(cmd *cobra.Command, args []string) {
		var yamlProbe parser.HttpProbe
		err := parser.ParseYAML(File, &yamlProbe)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		err = service.Post(endpoints.CreateProbe, yamlProbe)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
	},
}

func init() {
	probeCmd.AddCommand(createCmd)
	createCmd.Flags().
		StringVarP(&File, "file", "f", "", "Configuration file describing a probe")
	_ = createCmd.MarkFlagRequired("file")
}
