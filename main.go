package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"net/url"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputURL := scanner.Text()

		if inputURL == "" {
			continue
		}

		if !strings.HasPrefix(inputURL, "http://") && !strings.HasPrefix(inputURL, "https://") {
			inputURL = "https://" + inputURL
		}

		parsedURL, err := url.Parse(inputURL)
		if err != nil {
			logrus.Warnf("failed to parse URL %s: %v", inputURL, err)
		}

		host := parsedURL.Hostname()

		apexDomain, err := publicsuffix.EffectiveTLDPlusOne(host)
		if err != nil {
			logrus.Warnf("failed to get apex domain for %s: %v", host, err)
		}

		lowerApexDomain := strings.ToLower(apexDomain)
		fmt.Println(lowerApexDomain)
	}

	if err := scanner.Err(); err != nil {
		logrus.Error("error reading input: ", err)
		panic(err)
	}
}
