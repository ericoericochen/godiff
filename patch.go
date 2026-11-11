package godiff

type Patch struct {
	Type     PatchType
	OldStart int
	OldEnd   int
	NewStart int
	NewEnd   int
	Value    []string
}

type PatchType string

const (
	PatchInsert PatchType = "insert"
	PatchDelete PatchType = "delete"
	PatchEqual  PatchType = "equal"
)
