package godiff

import "io"

type DisplayPatchesOptions struct {
	WithLines bool
}

func WritePatches(w io.Writer, a []string, patches []Patch, opts *DisplayPatchesOptions) {

}
