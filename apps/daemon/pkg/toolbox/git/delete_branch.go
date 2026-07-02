// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package git

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nightona-co/nightona/apps/daemon/pkg/git"
	common_errors "github.com/nightona-co/nightona/libs/common-go/pkg/errors"
)

// DeleteBranch godoc
//
//	@Summary		Delete a branch
//	@Description	Delete a branch from the Git repository
//	@Tags			git
//	@Accept			json
//	@Produce		json
//	@Param			request	body	GitDeleteBranchRequest	true	"Delete branch request"
//	@Success		204
//	@Router			/git/branches [delete]
//
//	@id				DeleteBranch
func DeleteBranch(c *gin.Context) {
	var req GitDeleteBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(common_errors.NewInvalidBodyRequestError(fmt.Errorf("invalid request body: %w", err)))
		return
	}

	gitService := git.Service{
		WorkDir: req.Path,
	}

	if err := gitService.DeleteBranch(req.Name); err != nil {
		abortWithGitError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
