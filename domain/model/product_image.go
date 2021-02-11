/*
* @Time    : 2021-02-11 11:39
* @Author  : CoderCharm
* @File    : product_image.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/

package model

type ProductImage struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ImageName      string `json:"image_name"`
	ImageCode      string `gorm:"unique_index;not_null" json:"image_code"`
	ImageUrl       string `json:"image_url"`
	ImageProductID int64  `json:"image_product_id"`
}
