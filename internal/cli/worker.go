package cli

import (
	"context"
	"sync"

	"github.com/recovery-flow/distributors-admin/internal/service"
	"github.com/recovery-flow/distributors-admin/internal/service/events"
)

func runServices(ctx context.Context, wg *sync.WaitGroup) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { events.Listener(ctx) })

	run(func() { service.Run(ctx) })
}
