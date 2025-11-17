// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"net/http"
	"strconv"

	user_model "github.com/capitalrx/gitea/models/user"
	"github.com/capitalrx/gitea/modules/json"
	"github.com/capitalrx/gitea/services/context"
)

func UpdatePreferences(ctx *context.Context) {
	type preferencesForm struct {
		CodeViewShowFileTree bool `json:"codeViewShowFileTree"`
	}
	form := &preferencesForm{}
	if err := json.NewDecoder(ctx.Req.Body).Decode(&form); err != nil {
		ctx.HTTPError(http.StatusBadRequest, "json decode failed")
		return
	}
	_ = user_model.SetUserSetting(ctx, ctx.Doer.ID, user_model.SettingsKeyCodeViewShowFileTree, strconv.FormatBool(form.CodeViewShowFileTree))
	ctx.JSONOK()
}
