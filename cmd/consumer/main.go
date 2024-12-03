package main

import (
	"context"
	"github.com/HafslundEcoVannkraft/samplesystem/pkg/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"sync"
	"time"
)

var tracer trace.Tracer

func main() {
	ctx := context.Background()
	tr, cleanup := telemetry.InitTracer("system", "app")
	defer cleanup()

	tracer = tr

	// Use the tracer instance throughout your application
	spanCtx, span := tracer.Start(ctx, "main")
	defer span.End()

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()
			_, span := tracer.Start(ctx, "worker", trace.WithAttributes(attribute.Int("worker", i)))
			defer span.End()

			<-time.After(time.Second * 1)

		}(spanCtx, i)
	}

	wg.Wait()
}
