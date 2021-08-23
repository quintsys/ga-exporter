package handlers

import (
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/quintsys/ga-exporter/api/models"
	"github.com/quintsys/ga-exporter/api/restapi/operations/visitors"
)

// VisitorListHandlerImpl is the implementation of the visitor.VisitorListHandler
type VisitorListHandlerImpl struct{}

// NewVisitorListHandler creates a new VisitorListImpl
func NewVisitorListHandlerImpl() visitors.VisitorListHandler {
	return &VisitorListHandlerImpl{}
}

// Implement the Handle method of the visitor.VisitorListHandler interface
func (v *VisitorListHandlerImpl) Handle(params visitors.VisitorListParams) middleware.Responder {
	var list []*models.Visitor
	for i := 0; i < 10; i++ {
		visitor := &models.Visitor{
			ClientID:       swag.String("clientID" + strconv.Itoa(i)),
			AdContent:      new(string),
			AdGroup:        new(string),
			AdMatchedQuery: new(string),
			Campaign:       new(string),
			Keyword:        new(string),
			Medium:         new(string),
			Source:         new(string),
		}
		list = append(list, visitor)
	}
	return visitors.NewVisitorListOK().WithPayload(list)
}
