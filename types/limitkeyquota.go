package types

import (
	"context"
	"errors"
	"fmt"

	"github.com/kong/deck/crud"
	"github.com/kong/deck/state"
	"github.com/kong/deck/utils"
	"github.com/kong/go-kong/kong"
)

// limitKeyQuotaCRUD implements crud.Actions interface.
type limitKeyQuotaCRUD struct {
	client *kong.Client
}

func limitKeyQuotaFromStruct(arg crud.Event) *state.LimitKeyQuota {
	limitKeyQuota, ok := arg.Obj.(*state.LimitKeyQuota)
	if !ok {
		panic("unexpected type, expected *state.LimitKeyQuota")
	}

	return limitKeyQuota
}

// Create creates a Route in Kong.
// The arg should be of type crud.Event, containing the limitKeyQuota to be created,
// else the function will panic.
// It returns a the created *state.Route.
func (s *limitKeyQuotaCRUD) Create(ctx context.Context, arg ...crud.Arg) (crud.Arg, error) {
	event := crud.EventFromArg(arg[0])
	limitKeyQuota := limitKeyQuotaFromStruct(event)
	createdLimitKeyQuota, err := s.client.LimitKeyQuotas.Create(ctx, limitKeyQuota.Consumer.ID,
		&limitKeyQuota.LimitKeyQuota)
	if err != nil {
		return nil, err
	}
	return &state.LimitKeyQuota{LimitKeyQuota: *createdLimitKeyQuota}, nil
}

// Delete deletes a Route in Kong.
// The arg should be of type crud.Event, containing the limitKeyQuota to be deleted,
// else the function will panic.
// It returns a the deleted *state.Route.
func (s *limitKeyQuotaCRUD) Delete(ctx context.Context, arg ...crud.Arg) (crud.Arg, error) {
	event := crud.EventFromArg(arg[0])
	limitKeyQuota := limitKeyQuotaFromStruct(event)
	cid := ""
	if !utils.Empty(limitKeyQuota.Consumer.Username) {
		cid = *limitKeyQuota.Consumer.Username
	}
	if !utils.Empty(limitKeyQuota.Consumer.ID) {
		cid = *limitKeyQuota.Consumer.ID
	}
	err := s.client.LimitKeyQuotas.Delete(ctx, &cid, limitKeyQuota.ID)
	if err != nil {
		return nil, err
	}
	return limitKeyQuota, nil
}

// Update updates a Route in Kong.
// The arg should be of type crud.Event, containing the limitKeyQuota to be updated,
// else the function will panic.
// It returns a the updated *state.Route.
func (s *limitKeyQuotaCRUD) Update(ctx context.Context, arg ...crud.Arg) (crud.Arg, error) {
	event := crud.EventFromArg(arg[0])
	limitKeyQuota := limitKeyQuotaFromStruct(event)

	updatedLimitKeyQuota, err := s.client.LimitKeyQuotas.Create(ctx, limitKeyQuota.Consumer.ID,
		&limitKeyQuota.LimitKeyQuota)
	if err != nil {
		return nil, err
	}
	return &state.LimitKeyQuota{LimitKeyQuota: *updatedLimitKeyQuota}, nil
}

type limitKeyQuotaDiffer struct {
	kind crud.Kind

	currentState, targetState *state.KongState
}

func (d *limitKeyQuotaDiffer) Deletes(handler func(crud.Event) error) error {
	currentLimitKeyQuotas, err := d.currentState.LimitKeyQuotas.GetAll()
	if err != nil {
		return fmt.Errorf("error fetching limit-key-quotas from state: %w", err)
	}

	for _, limitKeyQuota := range currentLimitKeyQuotas {
		n, err := d.deleteLimitKeyQuota(limitKeyQuota)
		if err != nil {
			return err
		}
		if n != nil {
			err = handler(*n)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *limitKeyQuotaDiffer) deleteLimitKeyQuota(limitKeyQuota *state.LimitKeyQuota) (*crud.Event, error) {
	_, err := d.targetState.LimitKeyQuotas.Get(*limitKeyQuota.ID)
	if errors.Is(err, state.ErrNotFound) {
		return &crud.Event{
			Op:   crud.Delete,
			Kind: d.kind,
			Obj:  limitKeyQuota,
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("looking up limit-key-quota %q: %w", *limitKeyQuota.ID, err)
	}
	return nil, nil
}

func (d *limitKeyQuotaDiffer) CreateAndUpdates(handler func(crud.Event) error) error {
	targetLimitKeyQuotas, err := d.targetState.LimitKeyQuotas.GetAll()
	if err != nil {
		return fmt.Errorf("error fetching limit-key-quotas from state: %w", err)
	}

	for _, limitKeyQuota := range targetLimitKeyQuotas {
		n, err := d.createUpdateLimitKeyQuota(limitKeyQuota)
		if err != nil {
			return err
		}
		if n != nil {
			err = handler(*n)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *limitKeyQuotaDiffer) createUpdateLimitKeyQuota(limitKeyQuota *state.LimitKeyQuota) (*crud.Event, error) {
	limitKeyQuota = &state.LimitKeyQuota{LimitKeyQuota: *limitKeyQuota.DeepCopy()}
	currentLimitKeyQuota, err := d.currentState.LimitKeyQuotas.Get(*limitKeyQuota.ID)
	if errors.Is(err, state.ErrNotFound) {
		// limitKeyQuota not present, create it

		return &crud.Event{
			Op:   crud.Create,
			Kind: d.kind,
			Obj:  limitKeyQuota,
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error looking up limit-key-quota %q: %w",
			*limitKeyQuota.ID, err)
	}
	// found, check if update needed

	if !currentLimitKeyQuota.EqualWithOpts(limitKeyQuota, false, true, false) {
		return &crud.Event{
			Op:     crud.Update,
			Kind:   d.kind,
			Obj:    limitKeyQuota,
			OldObj: currentLimitKeyQuota,
		}, nil
	}
	return nil, nil
}
