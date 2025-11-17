// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package webhook

import (
	"testing"

	"github.com/capitalrx/gitea/models/unittest"
	"github.com/capitalrx/gitea/modules/hostmatcher"
	"github.com/capitalrx/gitea/modules/setting"

	_ "github.com/capitalrx/gitea/models"
	_ "github.com/capitalrx/gitea/models/actions"
)

func TestMain(m *testing.M) {
	// for tests, allow only loopback IPs
	setting.Webhook.AllowedHostList = hostmatcher.MatchBuiltinLoopback
	unittest.MainTest(m, &unittest.TestOptions{
		SetUp: func() error {
			setting.LoadQueueSettings()
			return Init()
		},
	})
}
