// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package system_test

import (
	"testing"

	"github.com/capitalrx/gitea/models/unittest"

	_ "github.com/capitalrx/gitea/models" // register models
	_ "github.com/capitalrx/gitea/models/actions"
	_ "github.com/capitalrx/gitea/models/activities"
	_ "github.com/capitalrx/gitea/models/system" // register models of system
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
