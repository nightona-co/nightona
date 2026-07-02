// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	hc_plugin "github.com/hashicorp/go-plugin"
	"github.com/nightona-co/nightona/apps/daemon/pkg/toolbox/computeruse"
	"github.com/nightona-co/nightona/apps/daemon/pkg/toolbox/computeruse/manager"
	cu "github.com/nightona-co/nightona/libs/computer-use/pkg/computeruse"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	hc_plugin.Serve(&hc_plugin.ServeConfig{
		HandshakeConfig: manager.ComputerUseHandshakeConfig,
		Plugins: map[string]hc_plugin.Plugin{
			"nightona-computer-use": &computeruse.ComputerUsePlugin{Impl: &cu.ComputerUse{}},
		},
		Logger: logger,
	})
}
