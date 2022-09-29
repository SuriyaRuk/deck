package types

import (
	"context"
	"fmt"

	"github.com/kong/deck/crud"
	"github.com/kong/deck/state"
	"github.com/kong/go-kong/kong"
	"github.com/samber/lo"
)

// consumerGroupCRUD implements crud.Actions interface.
type consumerGroupCRUD struct {
	client *kong.Client
}

func consumerGroupFromStruct(arg crud.Event) *state.ConsumerGroupObject {
	consumerGroup, ok := arg.Obj.(*state.ConsumerGroupObject)
	if !ok {
		panic("unexpected type, expected *state.ConsumerGroupObject")
	}
	return consumerGroup
}

// Create creates a consumerGroup in Kong.
// The arg should be of type crud.Event, containing the consumerGroup to be created,
// else the function will panic.
// It returns the created *state.consumerGroup.
func (s *consumerGroupCRUD) Create(ctx context.Context, arg ...crud.Arg) (crud.Arg, error) {
	event := crud.EventFromArg(arg[0])
	fmt.Println("HEHEHEHEH")
	consumerGroup := consumerGroupFromStruct(event)
	createdConsumerGroup, err := s.client.ConsumerGroups.Create(ctx, consumerGroup.ConsumerGroup)
	if err != nil {
		return nil, err
	}
	existingConsumers := []string{}
	cg, err := s.client.ConsumerGroups.Get(ctx, createdConsumerGroup.ID)
	if err != nil {
		return nil, err
	}
	for _, consumer := range cg.Consumers {
		existingConsumers = append(existingConsumers, *consumer.Username)
	}
	for _, consumer := range consumerGroup.Consumers {
		if !lo.Contains(existingConsumers, *consumer.Username) {
			_, err := s.client.ConsumerGroupConsumers.Create(
				ctx, createdConsumerGroup.ID, consumer.Username)
			if err != nil {
				return nil, err
			}
		}
	}
	return &state.ConsumerGroupObject{
		ConsumerGroupObject: kong.ConsumerGroupObject{
			ConsumerGroup: createdConsumerGroup,
			Consumers:     consumerGroup.Consumers,
		},
	}, nil
}

// Delete deletes a consumerGroup in Kong.
// The arg should be of type crud.Event, containing the consumerGroup to be deleted,
// else the function will panic.
// It returns a the deleted *state.consumerGroup.
func (s *consumerGroupCRUD) Delete(ctx context.Context, arg ...crud.Arg) (crud.Arg, error) {
	event := crud.EventFromArg(arg[0])
	consumerGroup := consumerGroupFromStruct(event)
	err := s.client.ConsumerGroups.Delete(ctx, consumerGroup.ConsumerGroup.ID)
	if err != nil {
		return nil, err
	}
	return consumerGroup, nil
}

// Update updates a consumerGroup in Kong.
// The arg should be of type crud.Event, containing the consumerGroup to be updated,
// else the function will panic.
// It returns a the updated *state.consumerGroup.
func (s *consumerGroupCRUD) Update(ctx context.Context, arg ...crud.Arg) (crud.Arg, error) {
	event := crud.EventFromArg(arg[0])
	consumerGroup := consumerGroupFromStruct(event)
	fmt.Println("HEHEHEHEH1")
	updatedconsumerGroup, err := s.client.ConsumerGroups.Create(ctx, consumerGroup.ConsumerGroup)
	if err != nil {
		return nil, err
	}
	existingConsumers := []string{}
	proposedConsumers := []string{}
	cg, err := s.client.ConsumerGroups.Get(ctx, updatedconsumerGroup.ID)
	if err != nil {
		return nil, err
	}
	for _, consumer := range cg.Consumers {
		existingConsumers = append(existingConsumers, *consumer.Username)
	}
	for _, consumer := range consumerGroup.Consumers {
		proposedConsumers = append(proposedConsumers, *consumer.Username)
	}
	for _, consumer := range consumerGroup.Consumers {
		if !lo.Contains(existingConsumers, *consumer.Username) {
			_, err := s.client.ConsumerGroupConsumers.Create(
				ctx, updatedconsumerGroup.ID, consumer.Username)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, consumer := range existingConsumers {
		if !lo.Contains(proposedConsumers, consumer) {
			err := s.client.ConsumerGroupConsumers.Delete(
				ctx, updatedconsumerGroup.ID, &consumer)
			if err != nil {
				return nil, err
			}
		}
	}
	return &state.ConsumerGroupObject{
		ConsumerGroupObject: kong.ConsumerGroupObject{
			ConsumerGroup: updatedconsumerGroup,
			Consumers:     consumerGroup.Consumers,
		},
	}, nil
}

type consumerGroupDiffer struct {
	kind crud.Kind

	currentState, targetState *state.KongState
}

func (d *consumerGroupDiffer) Deletes(handler func(crud.Event) error) error {
	currentconsumerGroups, err := d.currentState.ConsumerGroups.GetAll()
	if err != nil {
		return fmt.Errorf("error fetching consumerGroups from state: %w", err)
	}

	for _, consumerGroup := range currentconsumerGroups {
		n, err := d.deleteconsumerGroup(consumerGroup)
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

func (d *consumerGroupDiffer) deleteconsumerGroup(consumerGroup *state.ConsumerGroupObject) (*crud.Event, error) {
	_, err := d.targetState.ConsumerGroups.Get(*consumerGroup.ConsumerGroup.ID)
	if err == state.ErrNotFound {
		return &crud.Event{
			Op:   crud.Delete,
			Kind: "consumer-group",
			Obj:  consumerGroup,
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("looking up consumerGroup %q: %w",
			*consumerGroup.ConsumerGroup.Name, err)
	}
	return nil, nil
}

func (d *consumerGroupDiffer) CreateAndUpdates(handler func(crud.Event) error) error {
	targetconsumerGroups, err := d.targetState.ConsumerGroups.GetAll()
	if err != nil {
		return fmt.Errorf("error fetching consumerGroups from state: %w", err)
	}

	for _, consumerGroup := range targetconsumerGroups {
		n, err := d.createUpdateconsumerGroup(consumerGroup)
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

func (d *consumerGroupDiffer) createUpdateconsumerGroup(consumerGroup *state.ConsumerGroupObject) (*crud.Event,
	error,
) {
	consumerGroupCopy := &state.ConsumerGroupObject{
		ConsumerGroupObject: *consumerGroup.DeepCopy(),
	}
	currentconsumerGroup, err := d.currentState.ConsumerGroups.Get(*consumerGroup.ConsumerGroup.Name)

	if err == state.ErrNotFound {
		return &crud.Event{
			Op:   crud.Create,
			Kind: "consumer-group",
			Obj:  consumerGroupCopy,
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error looking up consumerGroup %v: %w",
			*consumerGroup.ConsumerGroup.Name, err)
	}

	// found, check if update needed
	if !currentconsumerGroup.EqualWithOpts(consumerGroupCopy, true, true) {
		return &crud.Event{
			Op:     crud.Update,
			Kind:   "consumer-group",
			Obj:    consumerGroupCopy,
			OldObj: currentconsumerGroup,
		}, nil
	}
	return nil, nil
}