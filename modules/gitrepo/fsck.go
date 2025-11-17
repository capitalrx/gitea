// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package gitrepo

import (
	"context"
	"time"

	"github.com/capitalrx/gitea/modules/git/gitcmd"
)

// Fsck verifies the connectivity and validity of the objects in the database
func Fsck(ctx context.Context, repo Repository, timeout time.Duration, args gitcmd.TrustedCmdArgs) error {
	return gitcmd.NewCommand("fsck").AddArguments(args...).Run(ctx, &gitcmd.RunOpts{Timeout: timeout, Dir: repoPath(repo)})
}
