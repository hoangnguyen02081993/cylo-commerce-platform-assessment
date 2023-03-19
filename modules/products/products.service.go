package products

import (
	c "commerce-platform/core/config"
	database "commerce-platform/core/database"
	model "commerce-platform/core/models"
)

type Product struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string  `gorm:"column:name;not null"`
	ImgURL      string  `gorm:"column:img_url;not null"`
	Description string  `gorm:"column:description"`
	Price       float32 `gorm:"column:price;not null"`
	Branch      string  `gorm:"column:branch;not null"`
	model.BaseModel
}

type IProductService interface {
	Filter(query ProductQueryDto) ([]Product, error)
	FindByID(id int) (Product, error)
}

type ProductService struct {
	DBConnectionManager database.IDBConnectionManager
}

func (p *ProductService) Filter(query ProductQueryDto) ([]Product, error) {
	db, err := p.DBConnectionManager.GetDB()
	if err != nil {
		return []Product{}, err
	}

	data := []Product{}

	filter := query.Filter
	queryStr := "1=1"
	conditionValues := []interface{}{}
	if filter.Name != "" {
		queryStr += " and (name like ? or description like ?)"
		conditionValues = append(append(conditionValues, "%"+filter.Name+"%"), "%"+filter.Name+"%")
	}
	if filter.FromPrice >= 0 {
		queryStr += " and price >= ?"
		conditionValues = append(conditionValues, filter.FromPrice)
	}
	if filter.ToPrice > 0 {
		queryStr += " and price <= ?"
		conditionValues = append(conditionValues, filter.ToPrice)
	}
	if filter.Branch != "" {
		queryStr += " and branch like ?"
		conditionValues = append(conditionValues, "%"+filter.Branch+"%")
	}

	sort := query.Sort
	sortStr := ""
	if sort.SortField != "" && sort.SortDirection != "" {
		sortStr += sort.SortField + " " + sort.SortDirection
	}

	db.Where(queryStr, conditionValues...).Order(sortStr).Find(&data)
	return data, nil
}

func (p *ProductService) FindByID(id int) (Product, error) {
	db, err := p.DBConnectionManager.GetDB()
	if err != nil {
		return Product{}, err
	}

	data := Product{}
	db.Where("id = ?", id).First(&data)
	return data, nil
}

func NewService() IProductService {
	config := c.GetConfig()
	manager := database.GetInstance(config)
	return &ProductService{DBConnectionManager: manager}
}
