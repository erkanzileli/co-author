package git

import _ "embed"

var (
	//go:embed prepare-commit-msg.template
	PrepareCommitMsgHookTemplate string
)
