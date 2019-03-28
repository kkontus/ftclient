// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"context"
	"ftserver/proto"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// rankingCmd represents the ranking command
var rankingCmd = &cobra.Command{
	Use:   "ranking",
	Short: "Shows user ranking",
	Long:  `Shows user ranking if rank is set and if searched user is set.`,
	Run: func(cmd *cobra.Command, args []string) {
		showRanking()
	},
}

func init() {
	rootCmd.AddCommand(rankingCmd)
}

func showRanking() {
	fmt.Println("ranking called")

	client := proto.NewQuestionsServiceClient(conn)

	player := promptForPlayerName()
	req := &proto.LoadUserRanking{User: player}

	if response, err := client.CheckUserRanking(context.Background(), req); err == nil {
		data := [][]string{
			[]string{player, response.Score, response.ScoreOverall},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"User", "Score", "Rating"})

		for _, v := range data {
			table.Append(v)
		}

		table.Render()
	} else {
		fmt.Println("Unable to retrieve user ranking")
	}
}
