package vo

import "purchase/src/dao"

type SupplierVo struct {
	Name string `form:"name"`
	Code string `form:"code"`
}

func (supplierVo *SupplierVo) VoToModel() *dao.Supplier {
	return &dao.Supplier{Code: supplierVo.Code, Name: supplierVo.Name}
}
