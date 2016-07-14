package osv

import (
	"github.com/emc-advanced-dev/unik/pkg/types"
)

type OSvCompilerBase struct {
	DockerImage string
	CreateImage func(params types.CompileImageParams) (*types.RawImage, error)
}

// CompileRawImage simply calls compiler's CreateImage function.
// This way we can define CrateImage function on declaration time.
func (compiler *OsvVirtualboxCompiler) CompileRawImage(params types.CompileImageParams) (_ *types.RawImage, err error) {
	return compiler.CreateImage(params)
}
