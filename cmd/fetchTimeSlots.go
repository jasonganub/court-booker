/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"time"
)

type RequestBody struct {
	ExperienceId        int    `json:"experienceId"`
	DashboardSortFilter string `json:"dashboardSortFilter"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	FetchType           string `json:"fetchType"`
	TicketTypeId        int    `json:"ticketTypeId"`
}

type Response struct {
	Count int   `json:"count"`
	Rows  []Row `json:"rows"`
}

type Row struct {
	Id              int         `json:"id"`
	CustomerId      int         `json:"customer_id"`
	ExperienceId    int         `json:"experience_id"`
	AddonId         interface{} `json:"addon_id"`
	PriceId         int         `json:"price_id"`
	OrderPaymentId  int         `json:"order_payment_id"`
	Date            string      `json:"date"`
	StartTime       string      `json:"start_time"`
	EndTime         string      `json:"end_time"`
	Status          string      `json:"status"`
	AdditionalNotes string      `json:"additional_notes"`
	Price           string      `json:"price"`
	Quantity        int         `json:"quantity"`
	CheckedIn       bool        `json:"checked_in"`
	TicketPosition  string      `json:"ticket_position"`
	HasRescheduled  interface{} `json:"has_rescheduled"`
	UpdatedBy       string      `json:"updated_by"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	DeletedAt       interface{} `json:"deletedAt"`
	PriceBreakdown  Price       `json:"Price"`
}

type Price struct {
	Id                  int         `json:"id"`
	ExperienceId        int         `json:"experience_id"`
	TicketTitle         string      `json:"ticket_title"`
	TicketPrice         string      `json:"ticket_price"`
	MemberPrice         string      `json:"member_price"`
	TicketDescription   string      `json:"ticket_description"`
	TicketDuration      string      `json:"ticket_duration"`
	TicketQuota         int         `json:"ticket_quota"`
	TicketType          string      `json:"ticket_type"`
	TicketTypeDays      interface{} `json:"ticket_type_days"`
	OpeningHours        string      `json:"opening_hours"`
	ClosingHours        string      `json:"closing_hours"`
	HasBreakHours       bool        `json:"has_break_hours"`
	BreakStartHours     interface{} `json:"break_start_hours"`
	BreakEndHours       interface{} `json:"break_end_hours"`
	IsPrivate           bool        `json:"is_private"`
	IsActive            bool        `json:"is_active"`
	StartTicketPosition int         `json:"start_ticket_position"`
	TicketPax           int         `json:"ticket_pax"`
	CreatedAt           time.Time   `json:"createdAt"`
	UpdatedAt           time.Time   `json:"updatedAt"`
	DeletedAt           interface{} `json:"deletedAt"`
}

// fetchTimeSlotsCmd represents the fetchTimeSlots command
var fetchTimeSlotsCmd = &cobra.Command{
	Use:   "fetchTimeSlots",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		if date == "" {
			fmt.Println("missing date")
			return
		}
		fmt.Printf("date %s\n", date)

		client := http.Client{}
		requestBody := RequestBody{
			ExperienceId:        28,
			DashboardSortFilter: "all",
			StartDate:           date,
			EndDate:             date,
			FetchType:           "all",
			TicketTypeId:        251,
		}
		requestBodyJson, _ := json.Marshal(requestBody)

		req, err := http.NewRequest("POST", "https://ayola.co/api/fetch-public-orders", bytes.NewBuffer(requestBodyJson))
		if err != nil {
			fmt.Printf("[error] %s", err.Error())
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("[error] %s", err.Error())
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var result Response
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Printf("[error] %s", err.Error())
		}

		var timeSlots []string
		for _, row := range result.Rows {
			timeSlots = append(timeSlots, row.StartTime)
		}

		for _, timeSlot := range timeSlots {
			fmt.Println(timeSlot)
		}
		return
	},
}

func init() {
	fetchTimeSlotsCmd.PersistentFlags().String("date", "", "date for availability")
	rootCmd.AddCommand(fetchTimeSlotsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchTimeSlotsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchTimeSlotsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
