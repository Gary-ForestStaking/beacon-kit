package middleware

import (
	"context"
	"time"

	"github.com/berachain/beacon-kit/mod/async/pkg/types"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/async"
)

// AwaitTimeout is the timeout for awaiting events.
const AwaitTimeout = 2 * time.Second

type VMMiddleware struct {
	// dispatcher is the central dispatcher to
	dispatcher types.EventDispatcher
	// logger is the logger for the middleware.
	logger log.Logger
	// subGenDataProcessed is the channel to hold GenesisDataProcessed events.
	subGenDataProcessed chan async.Event[validatorUpdates]
	// subBuiltBeaconBlock is the channel to hold BuiltBeaconBlock events.
	subBuiltBeaconBlock chan async.Event[BeaconBlockT]
	// subBuiltSidecars is the channel to hold BuiltSidecars events.
	subBuiltSidecars chan async.Event[BlobSidecarsT]
	// subBBVerified is the channel to hold BeaconBlockVerified events.
	subBBVerified chan async.Event[BeaconBlockT]
	// subSCVerified is the channel to hold SidecarsVerified events.
	subSCVerified chan async.Event[BlobSidecarsT]
	// subFinalValidatorUpdates is the channel to hold
	// FinalValidatorUpdatesProcessed events.
	subFinalValidatorUpdates chan async.Event[validatorUpdates]
}

func NewABCIMiddleware(
	dispatcher types.EventDispatcher,
	logger log.Logger,
) *VMMiddleware {
	return &VMMiddleware{
		dispatcher:               dispatcher,
		logger:                   logger,
		subGenDataProcessed:      make(chan async.Event[validatorUpdates]),
		subBuiltBeaconBlock:      make(chan async.Event[BeaconBlockT]),
		subBuiltSidecars:         make(chan async.Event[BlobSidecarsT]),
		subBBVerified:            make(chan async.Event[BeaconBlockT]),
		subSCVerified:            make(chan async.Event[BlobSidecarsT]),
		subFinalValidatorUpdates: make(chan async.Event[validatorUpdates]),
	}
}

// Should this be called upon VM.Initialize or VM.SetState(normalOp) ??
func (vm *VMMiddleware) Start(_ context.Context) error {
	var err error
	if err = vm.dispatcher.Subscribe(
		async.GenesisDataProcessed, vm.subGenDataProcessed,
	); err != nil {
		return err
	}
	if err = vm.dispatcher.Subscribe(
		async.BuiltBeaconBlock, vm.subBuiltBeaconBlock,
	); err != nil {
		return err
	}
	if err = vm.dispatcher.Subscribe(
		async.BuiltSidecars, vm.subBuiltSidecars,
	); err != nil {
		return err
	}
	if err = vm.dispatcher.Subscribe(
		async.BeaconBlockVerified, vm.subBBVerified,
	); err != nil {
		return err
	}
	if err = vm.dispatcher.Subscribe(
		async.SidecarsVerified, vm.subSCVerified,
	); err != nil {
		return err
	}
	if err = vm.dispatcher.Subscribe(
		async.FinalValidatorUpdatesProcessed, vm.subFinalValidatorUpdates,
	); err != nil {
		return err
	}
	return nil
}

// Name returns the name of the middleware.
func (vm *VMMiddleware) Name() string {
	return "abci-middleware"
}
