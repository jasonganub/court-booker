/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// bookCmd represents the book command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Book court",
	Long:  `Book a court immediately:`,
	Run: func(cmd *cobra.Command, args []string) {
		err := validateRequest(cmd)
		if err != nil {
			fmt.Println(fmt.Sprintf("[error]: %s", err.Error()))
			return
		}

		fmt.Println("book called")
	},
}

func validateRequest(cmd *cobra.Command) error {
	date, _ := cmd.Flags().GetString("date")
	if date == "" {
		return errors.New("missing date")
	}

	firstName, _ := cmd.Flags().GetString("first_name")
	if firstName == "" {
		return errors.New("missing first_name")
	}

	lastName, _ := cmd.Flags().GetString("last_name")
	if lastName == "" {
		return errors.New("missing last_name")
	}

	timeSlots, _ := cmd.Flags().GetString("time_slots")
	if timeSlots == "" {
		return errors.New("missing time_slots")
	}

	ballboy, _ := cmd.Flags().GetBool("ballboy")
	fmt.Println(ballboy)

	return nil
}

func init() {
	bookCmd.PersistentFlags().String("date", "", "Date for booking")
	bookCmd.PersistentFlags().String("first_name", "", "First name")
	bookCmd.PersistentFlags().String("last_name", "", "Last name")
	bookCmd.PersistentFlags().String("time_slots", "", "Timeslots (multiple allowed with spaces)")
	bookCmd.PersistentFlags().String("ballboy", "", "Ballboy")
	rootCmd.AddCommand(bookCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
