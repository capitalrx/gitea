// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package conan

import (
	"net/http"

	user_model "github.com/capitalrx/gitea/models/user"
	"github.com/capitalrx/gitea/modules/log"
	"github.com/capitalrx/gitea/services/auth"
	"github.com/capitalrx/gitea/services/packages"
)

var _ auth.Method = &Auth{}

type Auth struct{}

func (a *Auth) Name() string {
	return "conan"
}

// Verify extracts the user from the Bearer token
func (a *Auth) Verify(req *http.Request, w http.ResponseWriter, store auth.DataStore, sess auth.SessionStore) (*user_model.User, error) {
	packageMeta, err := packages.ParseAuthorizationRequest(req)
	if err != nil {
		log.Trace("ParseAuthorizationToken: %v", err)
		return nil, err
	}

	if packageMeta == nil || packageMeta.UserID == 0 {
		return nil, nil
	}

	u, err := user_model.GetUserByID(req.Context(), packageMeta.UserID)
	if err != nil {
		return nil, err
	}
	if packageMeta.Scope != "" {
		store.GetData()["IsApiToken"] = true
		store.GetData()["ApiTokenScope"] = packageMeta.Scope
	}

	return u, nil
}
