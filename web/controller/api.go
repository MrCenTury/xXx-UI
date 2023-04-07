package controller

import "github.com/gin-gonic/gin"

type APIController struct {
	BaseController
	inboundController *InboundController
}

func NewAPIController(g *gin.RouterGroup) *APIController {
	a := &APIController{}
	a.initRouter(g)
	return a
}

func (a *APIController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/xui/API/inbounds")
	g.Use(a.checkLogin)

	g.GET("/", a.inbounds)
	g.GET("/get/:id", a.inbound)
	g.POST("/add", a.addInbound)
	g.POST("/del/:id", a.delInbound)
	g.POST("/update/:id", a.updateInbound)
	g.POST("/addClient/", a.addInboundClient)
	g.POST("/delClient/:email", a.delInboundClient)
	g.POST("/updateClient/:index", a.updateInboundClient)
	g.POST("/:id/resetClientTraffic/:email", a.resetClientTraffic)
	g.POST("/resetAllTraffics", a.resetAllTraffics)
	g.POST("/resetAllClientTraffics/:id", a.resetAllClientTraffics)

	a.inboundController = NewInboundController(g)
}

func (a *APIController) inbounds(c *gin.Context) {
	a.inboundController.getInbounds(c)
}
func (a *APIController) inbound(c *gin.Context) {
	a.inboundController.getInbound(c)
}
func (a *APIController) addInbound(c *gin.Context) {
	a.inboundController.addInbound(c)
}
func (a *APIController) delInbound(c *gin.Context) {
	a.inboundController.delInbound(c)
}
func (a *APIController) updateInbound(c *gin.Context) {
	a.inboundController.updateInbound(c)
}
func (a *APIController) addInboundClient(c *gin.Context) {
	a.inboundController.addInboundClient(c)
}
func (a *APIController) delInboundClient(c *gin.Context) {
	a.inboundController.delInboundClient(c)
}
func (a *APIController) updateInboundClient(c *gin.Context) {
	a.inboundController.updateInboundClient(c)
}
func (a *APIController) resetClientTraffic(c *gin.Context) {
	a.inboundController.resetClientTraffic(c)
}
func (a *APIController) resetAllTraffics(c *gin.Context) {
	a.inboundController.resetAllTraffics(c)
}
func (a *APIController) resetAllClientTraffics(c *gin.Context) {
	a.inboundController.resetAllClientTraffics(c)
}
