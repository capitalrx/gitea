// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package convert

import (
	"testing"

	"github.com/capitalrx/gitea/models/unittest"

	_ "github.com/capitalrx/gitea/models/actions"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)
}
