package godiff

type UnifiedDiffOptions struct {
	FromFile string
	ToFile   string
	N        int
}

func UnifiedDiff(old []string, new []string, opts *UnifiedDiffOptions) {

}
