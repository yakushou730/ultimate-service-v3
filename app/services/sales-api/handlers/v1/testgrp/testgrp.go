package testgrp

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/yakushou730/ultimate-service-v3/foundation/web"

	"go.uber.org/zap"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		//return validate.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		panic("test panic")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
