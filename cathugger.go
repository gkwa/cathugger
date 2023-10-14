package cathugger

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/taylormonacelli/bluecare"
	"github.com/taylormonacelli/eachgoose"
	"github.com/taylormonacelli/lemondrop"
)

func Execute(resources []eachgoose.Resource) {
	var urls []string

	for _, res := range resources {
		slog.Debug("args", "service", res.Service, "region", res.Regions)
		for _, region := range res.Regions {
			url, err := bluecare.GetServiceURLInRegion(res.Service, region)
			if err != nil {
				slog.Error("get url fail", "region", region, "service", res.Service, "error", err.Error())
			}
			urls = append(urls, url)
		}
	}

	for _, url := range urls {
		slog.Debug("will open url", "url", url)
		RunCmdOpenUrl(url)
	}
}

func isRegionValid(region string) bool {
	regions, err := lemondrop.GetRegionDetails()
	if err != nil {
		slog.Error("fetch regions", "error", err.Error())
	}

	_, found := regions[region]
	return found
}

func GetAWSConsoleUrl(region, service string) string {
	if !isRegionValid(region) {
		slog.Error("region arg bad", "region", region)
		return ""
	}

	serviceMap, err := bluecare.GetServiceURLMap()
	if err != nil {
		slog.Error("matching service", "error", err.Error())
		return ""
	}

	url, ok := serviceMap[service]
	if !ok {
		return ""
	}

	url = strings.Replace(url, "us-west-1", region, -1)

	slog.Debug("url", "url", url)
	return url
}

func RunCmdOpenUrl(url string) {
	cmd := exec.Command("open", url)

	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
