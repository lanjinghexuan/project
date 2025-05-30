import service from '@/utils/request'
// @Tags Product
// @Summary 创建product表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "创建product表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /product/createProduct [post]
export const createProduct = (data) => {
  return service({
    url: '/product/createProduct',
    method: 'post',
    data
  })
}

// @Tags Product
// @Summary 删除product表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "删除product表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteProduct [delete]
export const deleteProduct = (params) => {
  return service({
    url: '/product/deleteProduct',
    method: 'delete',
    params
  })
}

// @Tags Product
// @Summary 批量删除product表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除product表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteProduct [delete]
export const deleteProductByIds = (params) => {
  return service({
    url: '/product/deleteProductByIds',
    method: 'delete',
    params
  })
}

// @Tags Product
// @Summary 更新product表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "更新product表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /product/updateProduct [put]
export const updateProduct = (data) => {
  return service({
    url: '/product/updateProduct',
    method: 'put',
    data
  })
}

// @Tags Product
// @Summary 用id查询product表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Product true "用id查询product表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /product/findProduct [get]
export const findProduct = (params) => {
  return service({
    url: '/product/findProduct',
    method: 'get',
    params
  })
}

// @Tags Product
// @Summary 分页获取product表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取product表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/getProductList [get]
export const getProductList = (params) => {
  return service({
    url: '/product/getProductList',
    method: 'get',
    params
  })
}

// @Tags Product
// @Summary 不需要鉴权的product表接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.ProductSearch true "分页获取product表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /product/getProductPublic [get]
export const getProductPublic = () => {
  return service({
    url: '/product/getProductPublic',
    method: 'get',
  })
}
