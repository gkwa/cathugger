package cathugger

import (
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/taylormonacelli/lemondrop"
)

func GetAWSConsoleUrl(region, service string) string {
	regions, err := lemondrop.GetRegionDetails()
	if err != nil {
		slog.Error("fetch regions", "error", err.Error())
	}

	isValidRegion := false
	for _, r := range regions {
		if r.RegionCode == region {
			slog.Debug("arg check", "found", true, "region", region)
			isValidRegion = true
			slog.Info("region", "region", r.RegionDesc)
			break
		}
	}

	if !isValidRegion {
		slog.Error("region arg bad", "region", region)
		return ""
	}

	services := []string{"ec2", "s3", "lambda"}

	isValidService := false
	for _, s := range services {
		if service == s {
			slog.Debug("arg check", "found", true, "service", service)
			isValidService = true
			break
		}
	}

	if !isValidService {
		slog.Error("service arg bad", "service", service)
		return ""
	}

	url := ""
	if service == "ec2" {
		url = fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#Instances:", region, region)
	}

	if service == "lambda" {
		url = fmt.Sprintf("https://%s.console.aws.amazon.com/lambda/home?region=%s#/functions", region, region)
	}

	if service == "s3" {
		url = fmt.Sprintf("https://s3.console.aws.amazon.com/s3/buckets?region=%s", region)
	}

	slog.Debug("region", "url", url)
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
