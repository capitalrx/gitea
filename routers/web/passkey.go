// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package web

import (
	"net/http"

	"github.com/capitalrx/gitea/modules/setting"
	"github.com/capitalrx/gitea/services/context"
)

type passkeyEndpointsType struct {
	Enroll string `json:"enroll"`
	Manage string `json:"manage"`
}

func passkeyEndpoints(ctx *context.Context) {
	url := setting.AppURL + "user/settings/security"
	ctx.JSON(http.StatusOK, passkeyEndpointsType{
		Enroll: url,
		Manage: url,
	})
}
