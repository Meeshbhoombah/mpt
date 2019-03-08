package mpt

type node interface {
	fstring(int) string
}

type (
	fullNode struct {
		Children [17]node
	}

	// Leaf/extension nodes are determined by the prefixed nibble
	shortNode struct {
		Key   string
		Value node
	}

	hashNode string

	value string
)

func (n *fullNode) fstring(lvld int) string {
	prefix := ""
	for lvl := 0; lvl <= lvld; lvl++ {
		prefix = prefix + "\t"
	}

	fmt.Sprintf(`|-"%s"`)
}
