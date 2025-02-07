package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-client-electron/service/autoclean"
	"github.com/pritunl/pritunl-client-electron/service/profile"
)

func stopPost(c *gin.Context) {
	prfls := profile.GetProfiles()
	for _, prfl := range prfls {
		prfl.StopBackground()
	}

	for _, prfl := range prfls {
		prfl.Wait()
	}

	autoclean.CheckAndCleanWatch()

	c.JSON(200, nil)
}
