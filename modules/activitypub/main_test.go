// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub

import (
	"testing"

	"github.com/capitalrx/gitea/models/unittest"

	_ "github.com/capitalrx/gitea/models"
	_ "github.com/capitalrx/gitea/models/actions"
	_ "github.com/capitalrx/gitea/models/activities"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
