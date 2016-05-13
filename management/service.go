package management

import (
	"github.com/micromdm/dep"
	"github.com/micromdm/micromdm/device"
	"github.com/micromdm/micromdm/profile"
	"github.com/pkg/errors"
)

// Service is the interface that provides methods for managing devices
type Service interface {
	AddProfile(prf *profile.Profile) (*profile.Profile, error)
	Profiles() ([]profile.Profile, error)
	Profile(identifier string) (*profile.Profile, error)
	FetchDEPDevices() error
}

type service struct {
	depClient dep.Client
	devices   device.Datastore
	profiles  profile.Datastore
}

func (svc service) AddProfile(prf *profile.Profile) (*profile.Profile, error) {
	return svc.profiles.Add(prf)
}

func (svc service) Profiles() ([]profile.Profile, error) {
	return svc.profiles.Profiles()
}

func (svc service) Profile(identifier string) (*profile.Profile, error) {
	panic("not implemented")
}

func (svc service) FetchDEPDevices() error {
	fetched, err := svc.depClient.FetchDevices(dep.Limit(100))
	if err != nil {
		return errors.Wrap(err, "management: dep fetch")
	}
	for _, d := range fetched.Devices {
		dev := device.NewFromDEP(d)
		_, err := svc.devices.New("fetch", dev)
		if err != nil {
			return errors.Wrap(err, "management: dep fetch")
		}
	}
	return nil
}

// NewService creates a management service
func NewService(ds device.Datastore, ps profile.Datastore, dc dep.Client) Service {
	return &service{
		devices:   ds,
		depClient: dc,
		profiles:  ps,
	}
}
