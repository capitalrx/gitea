// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo_test

import (
	"testing"

	"github.com/capitalrx/gitea/models/unittest"

	_ "github.com/capitalrx/gitea/models" // register table model
	_ "github.com/capitalrx/gitea/models/actions"
	_ "github.com/capitalrx/gitea/models/activities"
	_ "github.com/capitalrx/gitea/models/perm/access" // register table model
	_ "github.com/capitalrx/gitea/models/repo"        // register table model
	_ "github.com/capitalrx/gitea/models/user"        // register table model
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
