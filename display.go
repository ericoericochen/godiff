package godiff

import (
	"fmt"
	"io"
	"strconv"
)

type DisplayPatchesOptions struct {
	WithLines bool
}

type displayLine struct {
	patchType PatchType
	oldLine   int
	newLine   int
	value     string
}

func WritePatches(w io.Writer, a []string, patches []Patch, opts *DisplayPatchesOptions) {
	displayLines := getDisplayLines(a, patches)
	maxLine := 0
	for _, displayLine := range displayLines {
		maxLine = max(maxLine, displayLine.oldLine)
		maxLine = max(maxLine, displayLine.newLine)
	}

	width := len(strconv.Itoa(maxLine))

	for idx, displayLine := range displayLines {
		line := ""
		if opts.WithLines {
			if displayLine.oldLine > 0 {
				line += fmt.Sprintf("%*d", width, displayLine.oldLine)
			} else {
				line += fmt.Sprintf("%*s", width, "")

			}
			line += " "
			if displayLine.newLine > 0 {
				line += fmt.Sprintf("%*d", width, displayLine.newLine)
			} else {
				line += fmt.Sprintf("%*s", width, "")

			}
			line += " "
		}

		switch displayLine.patchType {
		case PatchDelete:
			{
				line += fmt.Sprintf("\033[41m%s\033[0m\n",
					fmt.Sprintf("-%v", displayLine.value),
				)
			}
		case PatchInsert:
			{
				line += fmt.Sprintf("\033[42m%s\033[0m\n",
					fmt.Sprintf("+%v", displayLine.value),
				)
			}
		case PatchEqual:
			{
				line += " " + displayLine.value + "\n"
			}
		}

		if idx == len(displayLines)-1 {
			line = line[:len(line)-1]
		}

		w.Write([]byte(line))
	}
}

func getDisplayLines(a []string, patches []Patch) []displayLine {
	displayLines := []displayLine{}
	for _, patch := range patches {
		switch patch.Type {
		case PatchDelete:
			{
				for i := patch.OldStart; i < patch.OldEnd; i++ {
					line := a[i]
					displayLine := displayLine{patchType: patch.Type, oldLine: i + 1, newLine: -1, value: line}
					displayLines = append(displayLines, displayLine)
				}
			}
		case PatchInsert:
			{
				for i := patch.NewStart; i < patch.NewEnd; i++ {
					j := i - patch.NewStart
					line := patch.Value[j]
					displayLine := displayLine{patchType: patch.Type, oldLine: -1, newLine: i + 1, value: line}
					displayLines = append(displayLines, displayLine)
				}
			}
		case PatchEqual:
			{
				for i := patch.OldStart; i < patch.OldEnd; i++ {
					j := i - patch.OldStart + patch.NewStart
					k := i - patch.OldStart
					line := a[k]
					displayLine := displayLine{patchType: patch.Type, oldLine: i + 1, newLine: j + 1, value: line}
					displayLines = append(displayLines, displayLine)
				}
			}
		}

	}
	return displayLines
}
