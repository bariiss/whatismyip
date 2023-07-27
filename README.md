# WhatIsMyIP: A Go CLI Tool for Fetching Your IP

WhatIsMyIP is a simple command-line interface (CLI) application written in Go that fetches and displays your current IP and associated details.

## Overview

The application sends a HTTP GET request to [ip-api.com](http://ip-api.com/json/), a free, non-login IP Geolocation API. The API responds with a JSON containing details about the IP, which the application parses and displays in a human-readable format.

The application presents the following details:

- IP
- Country
- Region
- City
- ISP (Internet Service Provider)
- Latitude
- Longitude
- Timezone

## Prerequisites

You need Go installed on your machine to run this application. You can download and install Go from the [official website](https://golang.org/dl/).

## Getting Started

1. Clone or download the repository to your local machine.
2. Navigate into the project directory.
3. Run the following command to fetch the necessary dependencies:

    ```bash
    go mod tidy
    ```

4. After fetching the dependencies, you can run the application with the following command:

    ```bash
    go run main.go
    ```

This command compiles and runs the application. You should see an output with your IP and related details printed in the terminal.

## Installation

Starting from Go 1.16, you can install this tool globally using the following command:

```bash
go install github.com/bariiss/whatismyip@latest
 ```

This command downloads the package from the GitHub repository, compiles it, and moves the resulting binary into the $GOPATH/bin directory. Ensure this directory is in your system's PATH. After installing the tool, you can use it anywhere in your terminal by running:

```bash
whatismyip
 ```

```bash
➜ whatismyip
IP:         123.45.67.89
Country:    Narnia
Region:     Far Far Away
City:       Emerald City
ISP:        Wizard of Oz Internet
Latitude:   12.345600
Longitude:  98.765400
Timezone:   Middle/Earth
```