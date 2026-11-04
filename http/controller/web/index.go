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

// ConfigJs serves the web client's bootstrap config. Besides the API server it presets the
// self-hosted ID server and public key (unless the visitor already chose their own), so the
// web client talks to this deployment instead of the public RustDesk servers.
func (i *Index) ConfigJs(c *gin.Context) {
	apiServer := jsString(global.Config.Rustdesk.ApiServer)
	idServer := jsString(global.Config.Rustdesk.IdServer)
	key := jsString(global.Config.Rustdesk.Key)
	magicQueryonline := global.Config.Rustdesk.WebclientMagicQueryonline
	tmp := fmt.Sprintf(`localStorage.setItem('api-server', %s);
const ws2_prefix = 'wc-';
localStorage.setItem(ws2_prefix+'api-server', %s);
if (%s && !localStorage.getItem('custom-rendezvous-server')) {
  localStorage.setItem('custom-rendezvous-server', %s);
}
if (%s && !localStorage.getItem('key')) {
  localStorage.setItem('key', %s);
}

window.webclient_magic_queryonline = %d;
window.ws_host = %s;
`, apiServer, apiServer, idServer, idServer, key, key, magicQueryonline, jsString(global.Config.Rustdesk.WsHost))

	c.Header("Content-Type", "application/javascript")
	c.String(200, tmp)
}
