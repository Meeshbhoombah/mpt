package mpt

type node interface {
	fstring(int) string
}

type (
	fullNode struct {
		Children [17]node
	}

	shortNode struct {
		// contains prefix and shared nibbles for both extension/
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
}
