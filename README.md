# Go IP Geolocation CLI Tool

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

Once you have Go installed, follow these steps:

1. Clone or download the repository to your local machine.

2. Navigate into the project directory.

3. Run the following command to fetch the necessary dependencies:

    ```bash
    go mod tidy
    ```

4. After fetching the dependencies, you can run the application with the following command:

    ```bash
    go run ./cmd/main.go
    ```

This command will compile and run the application, and you should see the output with your IP and geolocation details printed in the terminal.

## Dependencies

This application uses the following dependencies:

- [color](github.com/fatih/color): For printing colorful text in the terminal.
- [net/http](https://golang.org/pkg/net/http/): To make HTTP requests.
- [encoding/json](https://golang.org/pkg/encoding/json/): To parse the JSON response from the API.
- [io/ioutil](https://golang.org/pkg/io/ioutil/): To read the body of the API's HTTP response.
- [os](https://golang.org/pkg/os/): To interact with the operating system, mainly for printing error messages and exiting the application.
- [text/tabwriter](https://golang.org/pkg/text/tabwriter/): For aligning the output in a tabular format.