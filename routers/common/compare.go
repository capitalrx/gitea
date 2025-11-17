// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package common

import (
	repo_model "github.com/capitalrx/gitea/models/repo"
	user_model "github.com/capitalrx/gitea/models/user"
	"github.com/capitalrx/gitea/modules/git"
	pull_service "github.com/capitalrx/gitea/services/pull"
)

// CompareInfo represents the collected results from ParseCompareInfo
type CompareInfo struct {
	HeadUser         *user_model.User
	HeadRepo         *repo_model.Repository
	HeadGitRepo      *git.Repository
	CompareInfo      *pull_service.CompareInfo
	BaseBranch       string
	HeadBranch       string
	DirectComparison bool
}
