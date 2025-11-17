// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package indexer

import (
	code_indexer "github.com/capitalrx/gitea/modules/indexer/code"
	issue_indexer "github.com/capitalrx/gitea/modules/indexer/issues"
	stats_indexer "github.com/capitalrx/gitea/modules/indexer/stats"
	notify_service "github.com/capitalrx/gitea/services/notify"
)

// Init initialize the repo indexer
func Init() error {
	notify_service.RegisterNotifier(NewNotifier())

	issue_indexer.InitIssueIndexer(false)
	code_indexer.Init()
	return stats_indexer.Init()
}
