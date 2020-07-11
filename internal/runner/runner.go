package runner

import (
	"context"
	"fil-pusher/internal/collector/module"
	"fil-pusher/internal/log"
	"fmt"
)

type Options struct {
	Context string
}

type Runner struct {
	moduleManager *module.Manager
}

func NewRunner(ctx context.Context, logger log.Logger, options Options) (*Runner, error) {
	ctx = log.WithLoggerContext(ctx, logger)

	r := Runner{}

	if options.Context != "" {
		logger.With("initial-context", options.Context).Infof("Settiing initial context from user flags")
	}

	moduleManager, err := initModuleManager(logger)
	if err != nil {
		return nil, fmt.Errorf("init module manager: %w", err)
	}
	r.moduleManager = moduleManager

	moduleList, err := initModules(ctx)

	return &Runner{}, nil
}

func (r *Runner) Start(ctx context.Context, startupCh, shutdownCh chan bool) {
	logger := log.From(ctx)

	if startupCh != nil {
		startupCh <- true
	}

	go func() {
		fmt.Println("test")
	}()

	<-ctx.Done()

	shutdownCtx := log.WithLoggerContext(context.Background(), logger)
	r.Stop(shutdownCtx)
	shutdownCh <- true
}

func (r *Runner) Stop(ctx context.Context) {

}

func initModuleManager(logger log.Logger) (*module.Manager, error) {
	moduleManager, err := module.NewManager(logger)
	if err != nil {
		return nil, fmt.Errorf("create module manager: %w", err)
	}

	return moduleManager, nil
}

func initModules(ctx context.Context) ([]module.Module, error) {
	var list []module.Module

	return list, nil
}
