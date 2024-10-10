package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
	"net/url"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputURL := scanner.Text()

		if !strings.HasPrefix(inputURL, "http://") && !strings.HasPrefix(inputURL, "https://") {
			inputURL = "https://" + inputURL
		}

		parsedURL, err := url.Parse(inputURL)
		if err != nil {
			logrus.Error("failed to parse URL: ", err)
			panic(err)
		}

		host := parsedURL.Hostname()

		apexDomain, err := publicsuffix.EffectiveTLDPlusOne(host)
		if err != nil {
			logrus.Error("failed to extract apex domain: ", err)
			panic(err)
		}

		fmt.Println(apexDomain)
	}

	if err := scanner.Err(); err != nil {
		logrus.Error("error reading input: ", err)
		panic(err)
	}
}
