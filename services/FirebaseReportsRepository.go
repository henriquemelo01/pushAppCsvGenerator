package services

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"pushAppCsvGenerator/model"
	"time"
)

type FirebaseReportsRepository struct{}

func CreateFirebaseReportsRepository() *FirebaseReportsRepository {
	return &FirebaseReportsRepository{}
}

func (fr *FirebaseReportsRepository) connectToFirestoreClient(ctx context.Context) (*firestore.Client, error) {

	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, instantiateFirebaseAppError := firebase.NewApp(ctx, nil, opt)
	if instantiateFirebaseAppError != nil {
		return nil, fmt.Errorf("erro comunicação firebase - %s", instantiateFirebaseAppError.Error())
	}

	client, instantiateFirebaseClientError := app.Firestore(ctx)
	if instantiateFirebaseClientError != nil {
		log.Fatalf("Error ao obter o client do firestore: %s\n", instantiateFirebaseClientError.Error())
	}

	return client, nil
}

func (fr *FirebaseReportsRepository) GetReportById(id string) (model.ReportModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fbClient, connectToFirebaseClientError := fr.connectToFirestoreClient(ctx)
	if connectToFirebaseClientError != nil {
		return model.ReportModel{}, fmt.Errorf("erro ao conectar no client do Firebase: %s", connectToFirebaseClientError.Error())
	}

	defer fbClient.Close()

	reportDocument, getReportDocError := fbClient.Collection("reports").Doc(id).Get(context.Background())
	if getReportDocError != nil {
		return model.ReportModel{}, fmt.Errorf("report not found: %s", getReportDocError)
	}

	var report model.ReportModel

	parseReportDocToModelError := reportDocument.DataTo(&report)
	if parseReportDocToModelError != nil {
		return model.ReportModel{}, parseReportDocToModelError
	}

	return report, nil
}

func (fr *FirebaseReportsRepository) GetAll() ([]model.ReportModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fbClient, connectToFirebaseClientError := fr.connectToFirestoreClient(ctx)
	if connectToFirebaseClientError != nil {
		return nil, fmt.Errorf("erro ao conectar no client do Firebase: %s", connectToFirebaseClientError.Error())
	}

	defer fbClient.Close()

	reportsDocuments := fbClient.Collection("reports").Documents(context.Background())

	var firebaseReports []model.ReportModel

	for {
		doc, err := reportsDocuments.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var firebaseReport model.ReportModel
		if parseError := doc.DataTo(&firebaseReport); parseError != nil {
			return nil, fmt.Errorf("erro no parse: %s", parseError.Error())
		}

		firebaseReports = append(firebaseReports, firebaseReport)
	}

	return firebaseReports, nil
}
