package godiff

import (
	"fmt"
	"io"
)

type DisplayPatchesOptions struct {
	WithLines bool
}

func WritePatches(w io.Writer, a []string, patches []Patch, opts *DisplayPatchesOptions) {
	for _, patch := range patches {
		switch patch.Type {
		case PatchDelete:
			{
				lines := a[patch.OldStart:patch.OldEnd]
				for _, line := range lines {
					content := fmt.Sprintf("\033[41m%s\033[0m\n",
						fmt.Sprintf("-%v", line),
					)
					w.Write([]byte(content))
				}
			}
		case PatchInsert:
			{
				lines := patch.Value
				for _, line := range lines {
					content := fmt.Sprintf("\033[42m%s\033[0m\n",
						fmt.Sprintf("+%v", line),
					)
					w.Write([]byte(content))
				}
			}
		case PatchEqual:
			{
				lines := a[patch.OldStart:patch.OldEnd]
				for _, line := range lines {
					content := fmt.Sprintf(" %v\n", line)
					w.Write([]byte(content))
				}
			}
		}
	}
}
