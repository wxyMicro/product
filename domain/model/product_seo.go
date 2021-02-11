/*
* @Time    : 2021-02-11 11:40
* @Author  : CoderCharm
* @File    : product_seo.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package model

type ProductSeo struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	SeoTitle       string `json:"seo_title"`
	SeoKeywords    string `json:"seo_keywords"`
	SeoDescription string `json:"seo_description"`
	SeoCode        string `json:"seo_code"`
	SeoProductID   int64  `json:"seo_product_id"`
}
