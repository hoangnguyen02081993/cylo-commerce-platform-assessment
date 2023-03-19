package products_test

import (
	"regexp"

	p "commerce-platform/modules/products"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ = Describe("Product Suite", func() {
	var manager *MockDBConnectionManager
	var productService p.IProductService
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		db, m, err := sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		mock = m

		dialector := mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		})
		gdb, err := gorm.Open(dialector, &gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())

		manager = &MockDBConnectionManager{
			gdb: gdb,
		}
		productService = &p.ProductService{DBConnectionManager: manager}
		manager.GetDB()
	})

	Describe("Service", func() {
		Context("Filter product", func() {
			It("Should run as expected", func() {
				productQueryDto := p.ProductQueryDto{
					Filter: p.ProductFilterDto{
						Name:      "bia",
						FromPrice: 0,
						ToPrice:   20,
						Branch:    "Milk",
					},
					Sort: p.ProductSortDto{
						SortField:     "Name",
						SortDirection: "ASC",
					},
				}

				data := []p.Product{
					{
						ID:          1,
						Name:        "Name 1",
						Description: "Description 1",
						ImgURL:      "ImageUrl 1",
						Price:       14,
						Branch:      "Branch 1",
					},
					{
						ID:          2,
						Name:        "Name 2",
						Description: "Description 2",
						ImgURL:      "ImageUrl 2",
						Price:       30,
						Branch:      "Branch 2",
					},
				}

				expectedResult := sqlmock.NewRows([]string{"id", "name", "description", "ImgURL", "price", "branch"})
				for _, product := range data {
					expectedResult.AddRow(product.ID, product.Name, product.Description, product.ImgURL, product.Price, product.Branch)
				}

				sqlQuery := "SELECT * FROM `products` WHERE 1=1 and (name like ? or description like ?) and price >= ? and price <= ? and branch like ? ORDER BY Name ASC"
				mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).
					WithArgs(
						"%"+productQueryDto.Filter.Name+"%",
						"%"+productQueryDto.Filter.Name+"%",
						productQueryDto.Filter.FromPrice,
						productQueryDto.Filter.ToPrice,
						"%"+productQueryDto.Filter.Branch+"%",
					).
					WillReturnRows(expectedResult)

				result, err := productService.Filter(productQueryDto)
				Expect(err).To(BeNil())

				err = mock.ExpectationsWereMet()
				Expect(err).To(BeNil())

				Expect(len(result)).To(Equal(len(data)))
				for index, p := range result {
					Expect(p.ID).To(Equal(data[index].ID))
					Expect(p.Name).To(Equal(data[index].Name))
					Expect(p.Description).To(Equal(data[index].Description))
					Expect(p.Price).To(Equal(data[index].Price))
					Expect(p.Branch).To(Equal(data[index].Branch))
				}
			})

			It("Without filter and sort, should run as expected", func() {
				data := []p.Product{
					{
						ID:          1,
						Name:        "Name 1",
						Description: "Description 1",
						ImgURL:      "ImageUrl 1",
						Price:       14,
						Branch:      "Branch 1",
					},
					{
						ID:          2,
						Name:        "Name 2",
						Description: "Description 2",
						ImgURL:      "ImageUrl 2",
						Price:       30,
						Branch:      "Branch 2",
					},
				}

				expectedResult := sqlmock.NewRows([]string{"id", "name", "description", "ImgURL", "price", "branch"})
				for _, product := range data {
					expectedResult.AddRow(product.ID, product.Name, product.Description, product.ImgURL, product.Price, product.Branch)
				}

				sqlQuery := "SELECT * FROM `products` WHERE 1=1"
				mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).
					WillReturnRows(expectedResult)

				result, err := productService.Filter(p.ProductQueryDto{})
				Expect(err).To(BeNil())

				err = mock.ExpectationsWereMet()
				Expect(err).To(BeNil())

				Expect(len(result)).To(Equal(len(data)))
				for index, p := range result {
					Expect(p.ID).To(Equal(data[index].ID))
					Expect(p.Name).To(Equal(data[index].Name))
					Expect(p.Description).To(Equal(data[index].Description))
					Expect(p.Price).To(Equal(data[index].Price))
					Expect(p.Branch).To(Equal(data[index].Branch))
				}
			})
		})

		Context("Find a product", func() {
			It("Should run as expected", func() {
				id := 1

				data := p.Product{
					ID:          1,
					Name:        "Name 1",
					Description: "Description 1",
					ImgURL:      "ImageUrl 1",
					Price:       14,
					Branch:      "Branch 1",
				}

				expectedResult := sqlmock.NewRows([]string{"id", "name", "description", "ImgURL", "price", "branch"}).
					AddRow(data.ID, data.Name, data.Description, data.ImgURL, data.Price, data.Branch)

				sqlQuery := "SELECT * FROM `products` WHERE id = ?"
				mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).
					WithArgs(id).
					WillReturnRows(expectedResult)

				result, err := productService.FindByID(int(id))
				Expect(err).To(BeNil())

				err = mock.ExpectationsWereMet()
				Expect(err).To(BeNil())

				Expect(result.ID).To(Equal(data.ID))
				Expect(result.Name).To(Equal(data.Name))
				Expect(result.Description).To(Equal(data.Description))
				Expect(result.Price).To(Equal(data.Price))
				Expect(result.Branch).To(Equal(data.Branch))
			})
		})
	})
})
