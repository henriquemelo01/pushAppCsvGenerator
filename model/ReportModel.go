package model

type ReportModel struct {
	AccelerationPerTime []Entry `json:"accelerationPerTime"`
	CreatedAt           int     `json:"CreatedAt"`
	Exercise            string  `json:"exercise"`
	ForcePerTime        []Entry `json:"forcePerTime"`
	Id                  string  `json:"id"`
	MeanForce           float64 `json:"meanForce"`
	MeanPower           float64 `json:"meanPower"`
	MeanVelocity        float64 `json:"meanVelocity"`
	OffsetMovements     []Entry `json:"offsetMovements"`
	PowerPerTime        []Entry `json:"powerPerTime"`
	TrainingMethodology string  `json:"trainingMethodology"`
	UserId              string  `json:"userId"`
	VelocityPerTime     []Entry `json:"velocityPerTime"`
	Weight              int     `json:"weight"`
}
