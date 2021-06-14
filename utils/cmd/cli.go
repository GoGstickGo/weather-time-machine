package cmd

import (
	"fmt"
	"log"
	"os"
	"time"
	"weather-api/rapidapis"

	"github.com/spf13/cobra"
)

func newCmdDailyRun() *cobra.Command {
	var (
		params = rapidapis.Params{Writer: os.Stdout}
		s      = []string{}
	)
	cmd := &cobra.Command{
		Use:   "daily",
		Short: "daily",
		Long:  `Get daily lowest/highest temperature for specific date & city`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := rapidapis.DarkSkyreturns(params)
			if err != nil {
				log.Fatalf("❌ Couldn't initliaze command line: %v", err)
			}

			cmd.SilenceUsage = true
			return nil
		},
	}
	cmd.Flags().StringVarP(&params.Year, "year", "y", "", "Add year between: 1940-"+fmt.Sprint(time.Now().Year()))
	cmd.MarkFlagRequired("year")
	cmd.Flags().StringVarP(&params.Month, "month", "m", "", "Please choose valid month between: 01 - 12")
	cmd.MarkFlagRequired("month")
	cmd.Flags().StringVarP(&params.Day, "day", "d", "", "Please choose valid day between: 01 - 31")
	cmd.MarkFlagRequired("day")
	cmd.Flags().StringVar(&params.Apikey, "apikey", "", "Please add valid Rapidapi key")
	cmd.MarkFlagRequired("apikey")
	cmd.Flags().StringArrayVarP(&params.City, "city", "c", s, "City with more than one name it must be in quotations")
	cmd.MarkFlagRequired("city")
	return cmd
}

func NewDefaultWTMCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		Use:   "wtm",
		Short: "Weather time machine porvides temperature for specific date and city",
		Long: `Weather time machine gets temperatures for specific date and city.
		You must have valid Rapidapi APIkey, please see:  https://docs.rapidapi.com/docs/keys.
		Cities with more than 1 name it must in quotations as --city "San Francisco".
		
		Examples:
		# Long version
		wtm daily --year 1972 --month 01 --day 12 --city "San Francisco" --apikey 23lk4jh234jkl23h5dsfh345

		#Shorthand version
		wtm daily -y 1972 -m 01 -d 12 -c Dublin --apikey 23lk4jh234jkl23h5dsfh345
		`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.AddCommand(
		newCmdDailyRun(),
	)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	return cmd
}
