// Copyright © 2016 Arjen Schwarz <developer@arjen.eu>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/ArjenSchwarz/aqua/builder"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/spf13/cobra"
)

// apikeyCmd represents the apikey command
var apikeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "List and create API keys",
	Long:  `List your available API keys`,
	Run: func(cmd *cobra.Command, args []string) {
		keys, err := builder.ListAPIKeys(settings)
		if err != nil {
			printFailure(err.Error())
			return
		}
		values := make([]map[string]string, len(keys.Items))
		for index, key := range keys.Items {
			keydef := make(map[string]string)
			keydef["name"] = aws.StringValue(key.Name)
			keydef["description"] = aws.StringValue(key.Description)
			keydef["apikey"] = aws.StringValue(key.Id)
			values[index] = keydef
		}
		printSliceMaps(values)
	},
}

func init() {
	RootCmd.AddCommand(apikeyCmd)
}
