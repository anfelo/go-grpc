//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/anfelo/go-grpc/internal/rocket Store

package rocket

import "context"

// Rocket - defines the model of the rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(r Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service, responsible for
// updating the rocket inventory
type Service struct {
	Store Store
}

// New - returns a new instance of out rocket service
func New(s Store) Service {
	return Service{
		Store: s,
	}
}

// GetRocketByID - retrieves a rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	r, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return r, nil
}

// InsertRocket - insert a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, r Rocket) (Rocket, error) {
	r, err := s.Store.InsertRocket(r)
	if err != nil {
		return Rocket{}, err
	}
	return r, nil
}

// DeleteRocket - deletes a rocket from our store
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
