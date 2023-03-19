package products

type ProductFilterDto struct {
	Name      string  `form:"name" json:"name"`
	FromPrice float32 `form:"fromPrice" json:"fromPrice"`
	ToPrice   float32 `form:"toPrice" json:"toPrice"`
	Branch    string  `form:"branch" json:"branch"`
}

type ProductSortDto struct {
	SortField     string `form:"sortField" json:"sortField"`
	SortDirection string `form:"sortDirection" json:"sortDirection"`
}

type ProductQueryDto struct {
	Filter ProductFilterDto `json:"filter"`
	Sort   ProductSortDto   `json:"sort"`
}

type ProductDto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImgURL      string  `json:"imgUrl"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Branch      string  `json:"branch"`
}

type IProductTransformer interface {
	Transform(product Product) ProductDto
	TransformMultiple(products []Product) []ProductDto
}

type ProductTransformer struct{}

func (t *ProductTransformer) Transform(product Product) ProductDto {
	return ProductDto(product)
}

func (t *ProductTransformer) TransformMultiple(products []Product) []ProductDto {
	data := make([]ProductDto, len(products))
	for index, product := range products {
		data[index] = t.Transform(product)
	}
	return data
}

func NewProductTransformer() IProductTransformer {
	return &ProductTransformer{}
}
