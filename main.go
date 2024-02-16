package main

import (
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

type GeoInfo struct {
	IP       string  `json:"query"`
	Country  string  `json:"country"`
	Region   string  `json:"regionName"`
	City     string  `json:"city"`
	ISP      string  `json:"isp"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Timezone string  `json:"timezone"`
}

func main() {
	var ipURL string
	if len(os.Args) > 1 {
		ipURL = fmt.Sprintf("http://ip-api.com/json/%s", os.Args[1])
	} else {
		ipURL = "http://ip-api.com/json/"
	}

	resp, err := http.Get(ipURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error fetching IP:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading response:", err)
		os.Exit(1)
	}

	var geoInfo GeoInfo
	err = json.Unmarshal(body, &geoInfo)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error decoding response:", err)
		os.Exit(1)
	}

	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(w, "%s:\t%s\n", cyan("IP"), white(geoInfo.IP))
	_, _ = fmt.Fprintf(w, "%s:\t%s\n", cyan("Country"), yellow(geoInfo.Country))
	_, _ = fmt.Fprintf(w, "%s:\t%s\n", cyan("Region"), yellow(geoInfo.Region))
	_, _ = fmt.Fprintf(w, "%s:\t%s\n", cyan("City"), yellow(geoInfo.City))
	_, _ = fmt.Fprintf(w, "%s:\t%s\n", cyan("ISP"), yellow(geoInfo.ISP))
	_, _ = fmt.Fprintf(w, "%s:\t%f\n", cyan("Latitude"), geoInfo.Lat)
	_, _ = fmt.Fprintf(w, "%s:\t%f\n", cyan("Longitude"), geoInfo.Lon)
	_, _ = fmt.Fprintf(w, "%s:\t%s\n", cyan("Timezone"), white(geoInfo.Timezone))

	_ = w.Flush()
}
