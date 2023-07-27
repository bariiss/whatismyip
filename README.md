# Go IP Geolocation CLI Tool (whatismyip)

This tool is a simple command line interface (CLI) application, written in Go, which fetches and displays your current IP and geolocation information.

## How it Works

The application sends a HTTP GET request to [ip-api.com](http://ip-api.com/json/), a free, non-login IP Geolocation API. The API returns a JSON response with the IP and geolocation details, which the application parses and displays in a human-readable format.

The displayed details include the following:

- IP
- Country
- Region
- City
- ISP (Internet Service Provider)
- Latitude
- Longitude
- Timezone

## How to Run

To run this application, you will need Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

Once you have Go installed, you can run the application using the following command:

```bash
go run main.go
