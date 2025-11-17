// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package doctor

import (
	"testing"

	actions_model "github.com/capitalrx/gitea/models/actions"
	"github.com/capitalrx/gitea/models/unittest"
	"github.com/capitalrx/gitea/modules/log"

	"github.com/stretchr/testify/assert"
)

func Test_fixUnfinishedRunStatus(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	fixUnfinishedRunStatus(t.Context(), log.GetLogger(log.DEFAULT), true)

	// check if the run is cancelled by id
	run := unittest.AssertExistsAndLoadBean(t, &actions_model.ActionRun{ID: 796})
	assert.Equal(t, actions_model.StatusCancelled, run.Status)
}
