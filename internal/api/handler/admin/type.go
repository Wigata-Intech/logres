package admin

import (
	"github.com/wigata-intech/logres/internal/api/service"
)

type AdminHandler struct {
	adminUserService    service.IAdminUserService
	adminFileService    service.IAdminFileService
	adminProductService service.IAdminProductService
	adminOrderService   service.IAdminOrderService
}

func New(
	adminUserService service.IAdminUserService,
	adminFileService service.IAdminFileService,
	adminProductService service.IAdminProductService,
	adminOrderService service.IAdminOrderService,
) *AdminHandler {
	return &AdminHandler{
		adminUserService:    adminUserService,
		adminFileService:    adminFileService,
		adminProductService: adminProductService,
		adminOrderService:   adminOrderService,
	}
}
