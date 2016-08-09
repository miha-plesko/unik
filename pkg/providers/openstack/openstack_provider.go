package openstack

import (
	"os"
	"path/filepath"

	"github.com/emc-advanced-dev/unik/pkg/config"
	"github.com/emc-advanced-dev/unik/pkg/state"
)

var debuggerTargetImageName string

type OpenstackProvider struct {
	config config.Openstack
	state  state.State
}

func OpenstackStateFile() string {
	return filepath.Join(config.Internal.UnikHome, "openstack/state.json")

}
func openstackImagesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "openstack/images/")
}

func openstackInstancesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "openstack/instances/")
}

func openstackVolumesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "openstack/volumes/")
}

func NewOpenstackProvider(config config.Openstack) (*OpenstackProvider, error) {

	os.MkdirAll(openstackImagesDirectory(), 0777)
	os.MkdirAll(openstackInstancesDirectory(), 0777)
	os.MkdirAll(openstackVolumesDirectory(), 0777)

	p := &OpenstackProvider{
		config: config,
		state:  state.NewBasicState(OpenstackStateFile()),
	}

	return p, nil
}

func (p *OpenstackProvider) WithState(state state.State) *OpenstackProvider {
	p.state = state
	return p
}

func getImagePath(imageName string) string {
	return filepath.Join(openstackImagesDirectory(), imageName, "boot.img")
}

func getKernelPath(imageName string) string {
	return filepath.Join(openstackImagesDirectory(), imageName, "program.bin")
}

func getCmdlinePath(imageName string) string {
	return filepath.Join(openstackImagesDirectory(), imageName, "cmdline")
}

func getVolumePath(volumeName string) string {
	return filepath.Join(openstackVolumesDirectory(), volumeName, "data.img")
}
