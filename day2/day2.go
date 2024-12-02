package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("reactorData.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open txt file: %v", err))
	}
	defer file.Close()

	totalSafeReports := CalculateNumSafeReports(file)

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to reset file poistion: %v", err))
	}

	totalSafeReportsWithDampener := CalculateNumSafeReportsWithDampener(file)

	println("Total safe reports:", totalSafeReports)
	println("Total safe reports w/ dampener: ", totalSafeReportsWithDampener)
}

func CalculateNumSafeReports(reportsData io.Reader) int {
	var numSafeReports int
	scanner := bufio.NewScanner(reportsData)

	for scanner.Scan() {
		report := scanner.Text()

		stringifiedReportData := strings.Fields(report)
		reportData, err := StringSliceToIntSlice(stringifiedReportData)
		if err != nil {
			fmt.Printf("Could not covert report to int vals, %s. %v", stringifiedReportData, err)
			continue
		}

		if IsSafe(reportData) {
			numSafeReports++
		}
	}

	return numSafeReports
}

func CalculateNumSafeReportsWithDampener(reportsData io.Reader) int {
	var numSafeReports int
	scanner := bufio.NewScanner(reportsData)

	for scanner.Scan() {
		report := scanner.Text()

		stringifiedReportData := strings.Fields(report)
		reportData, err := StringSliceToIntSlice(stringifiedReportData)
		if err != nil {
			fmt.Printf("Could not covert report to int vals, %s. %v", stringifiedReportData, err)
			continue
		}

		if IsSafe(reportData) {
			numSafeReports++
		} else {
			for i := range report {
				subReport := GetSubReport(i, reportData)
				if IsSafe(subReport) {
					numSafeReports++
					break
				}
			}
		}
	}

	return numSafeReports
}

func IsSafe(reportData []int) bool {
	if IsLevelsDecreasing(reportData) || IsLevelsIncreasing(reportData) {
		return true
	}
	return false
}

func IsLevelsIncreasing(reportLevels []int) bool {
	if len(reportLevels) == 0 {
		return false
	}

	for i := 1; i < len(reportLevels); i++ {
		if reportLevels[i] <= reportLevels[i-1] || reportLevels[i]-reportLevels[i-1] > 3 {
			return false
		}
	}

	return true
}

func IsLevelsDecreasing(reportLevels []int) bool {
	if len(reportLevels) == 0 {
		return false
	}

	for i := 1; i < len(reportLevels); i++ {
		if reportLevels[i] >= reportLevels[i-1] || reportLevels[i-1]-reportLevels[i] > 3 {
			return false
		}
	}

	return true
}

func StringSliceToIntSlice(strSlice []string) ([]int, error) {
	if len(strSlice) == 0 {
		return []int{}, nil
	}

	intSlice := make([]int, len(strSlice))

	for i, val := range strSlice {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("error converting val: %s to int. %v", val, err)
		}

		intSlice[i] = intVal
	}

	return intSlice, nil
}

func GetSubReport(indexToRemove int, report []int) []int {
	if indexToRemove < 0 || indexToRemove >= len(report) {
		return report
	}

	subReport := make([]int, 0)
	subReport = append(subReport, report[:indexToRemove]...)
	subReport = append(subReport, report[indexToRemove+1:]...)
	return subReport
}
