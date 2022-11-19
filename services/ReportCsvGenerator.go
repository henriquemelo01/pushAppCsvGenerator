package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"pushAppCsvGenerator/model"
)

type ReportCsvGeneratorService struct{}

func CreateReportCsvGeneratorService() *ReportCsvGeneratorService {
	return &ReportCsvGeneratorService{}
}

func (reportCsvGeneratorService *ReportCsvGeneratorService) Export(report model.ReportModel) *model.ReportCsvFile {

	var rows [][]string

	addReportDataCsvRows(&rows, report)

	b := &bytes.Buffer{}          // creates IO Writer
	csvWriter := csv.NewWriter(b) // creates a csv writer that uses the io buffer.

	for _, row := range rows {
		_ = csvWriter.Write(row)
	}

	csvWriter.Flush()

	fileName := fmt.Sprintf("report_%s.csv", report.Id)

	return &model.ReportCsvFile{FileName: fileName, FileBuffer: b}
}

func addPointElementsToRow(row *[][]string, variableName string, dataPoints []model.Entry) {

	emptyLine := []string{"", ""}
	*row = append(*row, emptyLine)

	variableNamePerTimeRow := []string{"Tempo", variableName}
	*row = append(*row, variableNamePerTimeRow)

	for _, point := range dataPoints {
		dataPerTime := []string{fmt.Sprint(point.Timestamp), fmt.Sprint(point.Value)}
		*row = append(*row, dataPerTime)
	}
}

func addElementToRow(row *[][]string, header []string, dataPoints []any) {

	emptyLine := []string{"", ""}
	*row = append(*row, emptyLine)

	*row = append(*row, header)

	for _, point := range dataPoints {
		dataPerTime := []string{fmt.Sprint(point), fmt.Sprint(point)}
		*row = append(*row, dataPerTime)
	}
}

func addEmptyLineToRow(row *[][]string) {
	emptyLine := []string{"", ""}
	*row = append(*row, emptyLine)
}

func addReportDataCsvRows(rows *[][]string, report model.ReportModel) {

	addElementToRow(rows, []string{"User", "Report"}, []any{report.UserId, report.Id})

	addPointElementsToRow(rows, "Velocidade", report.VelocityPerTime)

	addPointElementsToRow(rows, "Forca", report.ForcePerTime)

	addPointElementsToRow(rows, "Potencia", report.PowerPerTime)
}
