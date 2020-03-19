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
	"fmt"
	"github.com/madjlzz/madprobectl/internal/endpoints"
	"github.com/madjlzz/madprobectl/internal/service"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var (
	// Convenient way to print a table to Stdout.
	PrettyTab = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "GetWithParam the detail of a specific probe or all probes",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var probesDetails []service.ProbeDetails
		if len(args) == 1 {
			pd, err := service.GetWithParam(endpoints.GetProbe, map[string]string{"name": args[0]})
			if err != nil {
				cmd.PrintErr(err)
				return
			}
			probesDetails = append(probesDetails, pd)
		} else {
			pds, err := service.GetAll(endpoints.FindAllProbe)
			if err != nil {
				cmd.PrintErr(err)
				return
			}
			probesDetails = append(probesDetails, pds...)
		}
		prettyPrintProbesDetails(probesDetails)
	},
}

func prettyPrintProbesDetails(probesDetails []service.ProbeDetails) {
	_, _ = fmt.Fprintln(PrettyTab, "NAME\tURL\tDELAY")
	for _, pd := range probesDetails {
		_, _ = fmt.Fprintf(PrettyTab, "%s\t%s\t%d\n", pd.Name, pd.URL, pd.Delay)
	}
	_ = PrettyTab.Flush()
}

func init() {
	probeCmd.AddCommand(getCmd)
}