package models

import "context"

type Work func(ctx context.Context) int

type BackgroundWorkerSettings struct {
	TableName          string           `json:"tableName"`
	Timer              float64          `json:"timer"`
	RequestDelay       int              `json:"requestDellay"`
	Random             bool             `json:"random"`
	WritesNumberToSend int              `json:"writesNumberToSend"`
	TotalToSend        int64            `json:"totalToSend"`
	StopWhenTableEnds  bool             `json:"stopWhenTableEnds"`
	ConsumerSettings   ConsumerSettings `json:"consumerSettings"`
}

type ConsumerSettings struct {
	Host      string `json:"host"`
	AuthModel string `json:"authModel"`
}
