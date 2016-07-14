package osv

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/emc-advanced-dev/pkg/errors"
	"github.com/emc-advanced-dev/unik/pkg/types"
	unikutil "github.com/emc-advanced-dev/unik/pkg/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// javaProjectConfig defines available inputs
type nodeProjectConfig struct {
	MainFile string `yaml:"main_file"`
}

// CreateImageNodeJS creates OSv image for NodeJS project and returns its metadata.
func CreateImageNodeJS(params types.CompileImageParams) (*types.RawImage, error) {
	resultFile, err := compileRawImageNodeJS(params)
	if err != nil {
		return nil, errors.New("failed to compile raw OSv NodeJS image", err)
	}
	return &types.RawImage{
		LocalImagePath: resultFile,
		StageSpec: types.StageSpec{
			ImageFormat: types.ImageFormat_QCOW2,
		},
		RunSpec: types.RunSpec{
			DeviceMappings: []types.DeviceMapping{
				types.DeviceMapping{MountPoint: "/", DeviceName: "/dev/sda1"},
			},
			StorageDriver:         types.StorageDriver_SATA,
			DefaultInstanceMemory: OSV_VIRTUALBOX_MEMORY,
		},
	}, nil
}

// compileRawImage creates OSv image from project source code and returns filepath of it.
func compileRawImageNodeJS(params types.CompileImageParams) (string, error) {
	sourcesDir := params.SourcesDir

	var config nodeProjectConfig
	data, err := ioutil.ReadFile(filepath.Join(sourcesDir, "manifest.yaml"))
	if err != nil {
		return "", errors.New("failed to read manifest.yaml file", err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return "", errors.New("failed to parse yaml manifest.yaml file", err)
	}

	container := unikutil.NewContainer("compilers-osv-nodejs").WithVolume(sourcesDir+"/", "/project_directory").WithEnv("NODE_MAIN_FILE", config.MainFile)
	var args []string
	//args = append(args, "--run", fmt.Sprintf("/bin/node %s", config.MainFile))

	logrus.WithFields(logrus.Fields{
		"args": args,
	}).Debugf("running compilers-osv-nodejs container")

	if err := container.Run(args...); err != nil {
		return "", errors.New("failed running compilers-osv-nodejs on "+sourcesDir, err)
	}

	resultFile, err := ioutil.TempFile("", "osv-boot.vmdk.")
	if err != nil {
		return "", errors.New("failed to create tmpfile for result", err)
	}
	defer func() {
		if err != nil && !params.NoCleanup {
			os.Remove(resultFile.Name())
		}
	}()

	if err := os.Rename(filepath.Join(sourcesDir, "boot.qcow2"), resultFile.Name()); err != nil {
		return "", errors.New("failed to rename result file", err)
	}
	return resultFile.Name(), nil
}
