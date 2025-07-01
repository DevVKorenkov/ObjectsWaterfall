package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	bbl "objectswaterfall.com/BBL"
	"objectswaterfall.com/core/models"
	"objectswaterfall.com/stores"
)

func Start(ctx *gin.Context) {
	var workerSettings models.BackgroundWorkerSettings
	if err := ctx.ShouldBindBodyWithJSON(&workerSettings); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	store := stores.GetWorkerStore()
	duration := time.Now().Add(time.Minute * time.Duration(workerSettings.Timer))
	context, cancel := context.WithDeadline(context.Background(), duration)
	worker := bbl.NewSendWorker(workerSettings, cancel)
	workerId := store.Add(&worker)

	go worker.DoWork(context)

	ctx.JSON(http.StatusBadRequest, gin.H{"workerId": workerId})
}

func Stop(ctx *gin.Context) {

}

func Seed(ctx *gin.Context) {
	var seedProc bbl.SeedProcessor
	err := ctx.ShouldBindBodyWithJSON(&seedProc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errCh := make(chan error)

	go func() {
		errCh <- seedProc.ProcessJson(false, 0)
	}()

	if err = <-errCh; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
