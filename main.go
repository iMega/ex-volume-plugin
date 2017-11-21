package main

import (
	"log"
	"sync"

	"github.com/docker/go-plugins-helpers/volume"
	"github.com/sirupsen/logrus"
)

type Driver struct {
	volumes    map[string]string
	m          *sync.Mutex
	mountPoint string
}

func (Driver) Create(r *volume.CreateRequest) error {
	logrus.Infof("Create volume: %s", r.Name)
	return nil
}

func (Driver) List() (*volume.ListResponse, error) {
	logrus.Info("list volume")
	return &volume.ListResponse{}, nil
}

func (Driver) Get(r *volume.GetRequest) (*volume.GetResponse, error) {
	logrus.Infof("Get volume: %s", r.Name)
	return &volume.GetResponse{}, nil
}

func (Driver) Remove(r *volume.RemoveRequest) error {
	logrus.Infof("Remove volume: %s", r.Name)
	return nil
}

func (Driver) Path(r *volume.PathRequest) (*volume.PathResponse, error) {
	logrus.Infof("Path volume: %s", r.Name)
	return &volume.PathResponse{}, nil
}

func (Driver) Mount(r *volume.MountRequest) (*volume.MountResponse, error) {
	logrus.Infof("Mount volume: %s", r.Name)
	return &volume.MountResponse{}, nil
}

func (Driver) Unmount(r *volume.UnmountRequest) error {
	logrus.Infof("Unmount volume: %s", r.Name)
	return nil
}

func (Driver) Capabilities() *volume.CapabilitiesResponse {
	logrus.Info("Capabilities volume")
	return &volume.CapabilitiesResponse{}
}

func NewDriver() Driver {
	return Driver{
		volumes:    make(map[string]string),
		m:          &sync.Mutex{},
		mountPoint: "/tmp/exampledriver",
	}
}

func main() {
	driver := NewDriver()
	handler := volume.NewHandler(driver)
	if err := handler.ServeUnix("root", 0); err != nil {
		log.Fatalf("Error %v", err)
	}
}
