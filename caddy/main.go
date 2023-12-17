// Copyright 2023 The OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"

	_ "github.com/corazawaf/coraza-caddy/v2"

	_ "github.com/corazawaf/coraza-coreruleset"

	_ "github.com/jcchavezs/mergefs"

	_ "github.com/jcchavezs/mergefs/io"
)

func main() {
	waf, err := coraza.NewWAF(
        coraza.NewWAFConfig().
            WithDirectives(`
                Include @owasp_crs/REQUEST-911-METHOD-ENFORCEMENT.conf
                Include myrules/*.conf
            `).
            WithRootFS(mergefs.Merge(coreruleset.FS, io.OSFS)),
    )

	caddycmd.Main()
}
