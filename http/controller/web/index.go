package web

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lejianwen/rustdesk-api/v2/global"
)

type Index struct {
}

func (i *Index) Index(c *gin.Context) {
	c.Redirect(302, "/_admin/")
}

// jsString encodes s as a JavaScript string literal, so config values cannot break out of
// the generated script.
func jsString(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

// ConfigJs serves the web client's bootstrap config: only where the account API is. The ID
// server, relay and key are handed out after sign-in by /api/server-config-v2 (and to share
// links by /api/shared-peer), so anonymous visitors get nothing beyond the sign-in page.
func (i *Index) ConfigJs(c *gin.Context) {
	apiServer := jsString(global.Config.Rustdesk.ApiServer)
	tmp := fmt.Sprintf(`localStorage.setItem('api-server', %s);
`, apiServer)

	c.Header("Content-Type", "application/javascript")
	c.String(200, tmp)
}
