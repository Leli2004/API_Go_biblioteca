package worker

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/fine"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"log"
	"time"
)

type FineChecker struct {
	fineUC   fine.UseCase
	interval time.Duration
}

func NewFineChecker(uc fine.UseCase) *FineChecker {
	return &FineChecker{fineUC: uc, interval: time.Duration(helpers.TIME_CHECK_FINE_MINUTES) * time.Minute}
}

func (w *FineChecker) Run(ctx context.Context) {
	w.check(ctx)
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			w.check(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (w *FineChecker) check(ctx context.Context) {
	_, err, count := w.fineUC.ProcessOverdueLoans(ctx)
	if err != nil {
		log.Printf("fine checker error: %v", err)
		return
	}
	if count > 0 {
		log.Printf("fine checker created %d fine(s)", count)
	}
}
