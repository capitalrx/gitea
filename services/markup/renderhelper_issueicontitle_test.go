// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package markup

import (
	"testing"

	"github.com/capitalrx/gitea/models/repo"
	"github.com/capitalrx/gitea/models/unittest"
	"github.com/capitalrx/gitea/modules/markup"
	"github.com/capitalrx/gitea/modules/templates"
	"github.com/capitalrx/gitea/modules/util"
	"github.com/capitalrx/gitea/services/contexttest"

	"github.com/stretchr/testify/assert"
)

func TestRenderHelperIssueIconTitle(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	ctx, _ := contexttest.MockContext(t, "/", contexttest.MockContextOption{Render: templates.HTMLRenderer()})
	ctx.Repo.Repository = unittest.AssertExistsAndLoadBean(t, &repo.Repository{ID: 1})
	htm, err := renderRepoIssueIconTitle(ctx, markup.RenderIssueIconTitleOptions{
		LinkHref:   "/link",
		IssueIndex: 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, `<a href="/link"><span>octicon-issue-opened(16/text green)</span> issue1 (#1)</a>`, string(htm))

	ctx, _ = contexttest.MockContext(t, "/", contexttest.MockContextOption{Render: templates.HTMLRenderer()})
	htm, err = renderRepoIssueIconTitle(ctx, markup.RenderIssueIconTitleOptions{
		OwnerName:  "user2",
		RepoName:   "repo1",
		LinkHref:   "/link",
		IssueIndex: 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, `<a href="/link"><span>octicon-issue-opened(16/text green)</span> issue1 (user2/repo1#1)</a>`, string(htm))

	ctx, _ = contexttest.MockContext(t, "/", contexttest.MockContextOption{Render: templates.HTMLRenderer()})
	_, err = renderRepoIssueIconTitle(ctx, markup.RenderIssueIconTitleOptions{
		OwnerName:  "user2",
		RepoName:   "repo2",
		LinkHref:   "/link",
		IssueIndex: 2,
	})
	assert.ErrorIs(t, err, util.ErrPermissionDenied)
}
