package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lejianwen/rustdesk-api/v2/global"
	"github.com/lejianwen/rustdesk-api/v2/http/controller/web"
	"net/http"
)

func WebInit(g *gin.Engine) {
	i := &web.Index{}
	g.GET("/", i.Index)

	if global.Config.App.WebClient == 1 {
		g.GET("/webclient-config/index.js", i.ConfigJs)
	}

	if global.Config.App.WebClient == 1 {
		g.StaticFS("/webclient", http.Dir(global.Config.Gin.ResourcesPath+"/web"))
		// webclient2 was removed upstream, but the admin UI still links to
		// /webclient2/#/<id>. Send those to the v1 web client; browsers keep the
		// #/<id> fragment across the redirect and the client prefills that ID.
		g.GET("/webclient2", redirectToWebClient)
		g.GET("/webclient2/*any", redirectToWebClient)
	}
	g.StaticFS("/_admin", http.Dir(global.Config.Gin.ResourcesPath+"/admin"))
}

func redirectToWebClient(c *gin.Context) {
	c.Redirect(http.StatusFound, "/webclient/")
}
