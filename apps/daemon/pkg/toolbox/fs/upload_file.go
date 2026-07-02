// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package fs

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadFile godoc
//
//	@Summary		Upload a file
//	@Description	Upload a file to the specified path
//	@Tags			file-system
//	@Accept			multipart/form-data
//	@Param			path	query		string	true	"Destination path for the uploaded file"
//	@Param			file	formData	file	true	"File to upload"
//	@Success		200		{object}	gin.H
//	@Router			/files/upload [post]
//
//	@id				UploadFile
func UploadFile(c *gin.Context) {
	enableFullDuplex(c)
	defer drainBody(c)

	path := c.Query("path")
	if path == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("path is required"))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := saveUploadedFile(file, path); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}

// saveUploadedFile writes the multipart file to dst. Unlike gin's
// Context.SaveUploadedFile (which since gin v1.11 chmods the destination
// directory and fails when the daemon does not own it, e.g. /tmp), this
// only creates missing directories and never alters permissions of
// existing ones.
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if dir := filepath.Dir(dst); dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	_, copyErr := io.Copy(out, src)
	// Inspect Close() — on FUSE-backed filesystems (e.g. mount-s3 for
	// volume mounts) the actual remote write happens here, so a swallowed
	// close error means silent data loss.
	closeErr := out.Close()
	if copyErr != nil {
		return copyErr
	}
	return closeErr
}
