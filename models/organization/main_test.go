// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package organization_test

import (
	"testing"

	"github.com/capitalrx/gitea/models/unittest"

	_ "github.com/capitalrx/gitea/models"
	_ "github.com/capitalrx/gitea/models/actions"
	_ "github.com/capitalrx/gitea/models/activities"
	_ "github.com/capitalrx/gitea/models/organization"
	_ "github.com/capitalrx/gitea/models/repo"
	_ "github.com/capitalrx/gitea/models/user"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
