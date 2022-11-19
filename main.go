package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"pushAppCsvGenerator/constant"
	"pushAppCsvGenerator/controller"
	"pushAppCsvGenerator/di"
)

func main() {

	serverHandler := chi.NewRouter()

	serverHandler.Use(middleware.DefaultLogger)

	serverHandler.Get(fmt.Sprintf("%s/ping", constant.ApiRoute), func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("Pong 🧙‍♂️"))
		return
	})

	reportCsvGeneratorController := di.CreateReportCsvGeneratorController()
	reportCsvGeneratorController.SetupRouter(serverHandler)

	reportsController := controller.CreateReportsController()
	reportsController.SetupRoute(serverHandler)

	var httpServerPort = os.Getenv("PORT")
	if httpServerPort == "" {
		httpServerPort = "8080"
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf("localhost:%v", httpServerPort),
		Handler: serverHandler,
	}

	if serverConnectionError := httpServer.ListenAndServe(); serverConnectionError != nil {
		log.Fatalln("Erro ao inicializar web service")
	}
}

//func getUserNameById(userId string, client *firestore.Client) (string, error) {
//
//	defer client.Close()
//
//	userDocument := client.Collection("users").Doc(userId)
//
//	userData, getUserError := userDocument.Get(context.Background())
//
//	if getUserError != nil {
//		return "", getUserError
//	}
//
//	var firebaseUser FirebaseUser
//
//	if parseUserDataError := userData.DataTo(&firebaseUser); parseUserDataError != nil {
//		return "", parseUserDataError
//	}
//
//	return strings.TrimSpace(firebaseUser.FirstName), nil
//
//}
//
//type FirebaseUser struct {
//	Email     string `json:"email"`
//	FirstName string `json:"firstName"`
//}
//
//type Point struct {
//	Timestamp float64 `json:"timestamp"`
//	Value     float64 `json:"value"`
//}
//
//type FirebaseReport struct {
//	AccelerationPerTime []Point `json:"accelerationPerTime"`
//	CreatedAt           int     `json:"CreatedAt"`
//	Exercise            string  `json:"exercise"`
//	ForcePerTime        []Point `json:"forcePerTime"`
//	Id                  string  `json:"id"`
//	MeanForce           float64 `json:"meanForce"`
//	MeanPower           float64 `json:"meanPower"`
//	MeanVelocity        float64 `json:"meanVelocity"`
//	OffsetMovements     []Point `json:"offsetMovements"`
//	PowerPerTime        []Point `json:"powerPerTime"`
//	TrainingMethodology string  `json:"trainingMethodology"`
//	UserId              string  `json:"userId"`
//	VelocityPerTime     []Point `json:"velocityPerTime"`
//	Weight              int     `json:"weight"`
//}
//
//func addPointElementsToRow(row *[][]string, variableName string, dataPoints []Point) {
//
//	emptyLine := []string{"", ""}
//	*row = append(*row, emptyLine)
//
//	variableNamePerTimeRow := []string{"Tempo", variableName}
//	*row = append(*row, variableNamePerTimeRow)
//
//	for _, point := range dataPoints {
//		dataPerTime := []string{fmt.Sprint(point.Timestamp), fmt.Sprint(point.Value)}
//		*row = append(*row, dataPerTime)
//	}
//}
//
//func addElementToRow(row *[][]string, header []string, dataPoints []any) {
//
//	emptyLine := []string{"", ""}
//	*row = append(*row, emptyLine)
//
//	*row = append(*row, header)
//
//	for _, point := range dataPoints {
//		dataPerTime := []string{fmt.Sprint(point), fmt.Sprint(point)}
//		*row = append(*row, dataPerTime)
//	}
//}
//
//func addEmptyLineToRow(row *[][]string) {
//	emptyLine := []string{"", ""}
//	*row = append(*row, emptyLine)
//}
//
//func addReportDataCsvRows(rows *[][]string, report FirebaseReport) {
//
//	addElementToRow(rows, []string{"User", "Report"}, []any{report.UserId, report.Id})
//
//	addPointElementsToRow(rows, "Velocidade", report.VelocityPerTime)
//
//	addPointElementsToRow(rows, "Forca", report.ForcePerTime)
//
//	addPointElementsToRow(rows, "Potencia", report.PowerPerTime)
//}
