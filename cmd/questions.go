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
	"github.com/spf13/cobra"
	"ftserver/proto"
	"context"
	"bufio"
	"os"
	"strings"
	"github.com/olekukonko/tablewriter"
)

// questionsCmd represents the questions command
var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "Shows list of all questions",
	Long:  `Shows the list of all available questions.`,
	Run: func(cmd *cobra.Command, args []string) {
		startQuiz()
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)
}

const (
	NETWORK = "network"
	FILE = "file"
	DB = "db"
)

type condition func(string) bool

func promptForInput(f condition, text string) string {
	input := ""

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
		return promptForInput(f, text)
	} else if (f != nil && f(string(sentence))) {
		fmt.Println(text)
		return promptForInput(f, text)
	}

	input = string(sentence)
	input = strings.TrimSpace(input)
	return input
}

func promptForPlayerName() string {
	fmt.Println("Enter player name: ")
	conditionFunc := func(text string) bool { return len(text) <= 2 }
	player := promptForInput(conditionFunc, "User must have at least 2 characters")

	return player
}

func prepareData(userResponses []string, results []string, answers []string) [][]string {
	// allocate 2d array
	data := make([][]string, len(results))
	for i := range data {
		data[i] = make([]string, 3)
	}

	for i := range data {
		data[i][0] = userResponses[i]
		data[i][1] = results[i]
		data[i][2] = answers[i]
	}

	return data
}

func drawTable(data [][]string, score string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Answered", "Status", "Correct"})
	table.SetFooter([]string{"", "Total", score})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

func startQuiz() {
	fmt.Println("\n########################################")
	fmt.Println("\n########## Welcome to FT QUIZ ##########\n")
	fmt.Println("########################################\n")

	client := proto.NewQuestionsServiceClient(conn)
	req := &proto.LoadQuestionsList{}

	player := promptForPlayerName()
	fmt.Println("\nWelcome " + player)
	fmt.Println("Lets play!!!")

	var userResponses []string
	if response, err := client.GetAllQuestions(context.Background(), req); err == nil {
		for i, elem := range response.Result {
			fmt.Printf("\n%d: %s\n", i, elem.Question)
			fmt.Println("___________________________________\n")

			for i, elem := range elem.Answers {
				fmt.Printf("%d: %s\n", i, elem)
			}

			input := promptForInput(nil, "")
			userResponses = append(userResponses, input)
			fmt.Println(input)
		}
	}

	reqUserResult := &proto.SendUserAnswers{Answers: userResponses, User: player}
	if responseUserResult, err := client.CheckUserAnswers(context.Background(), reqUserResult); err == nil {
		data := prepareData(userResponses, responseUserResult.Result, responseUserResult.Answers)
		drawTable(data, responseUserResult.Percentage)
	} else {
		fmt.Println("Unable to retrieve user results")
	}
}
