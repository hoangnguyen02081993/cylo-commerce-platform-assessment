package products_test

import (
	p "commerce-platform/modules/products"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Suite", func() {
	var productTransformer p.IProductTransformer

	BeforeEach(func() {
		productTransformer = &p.ProductTransformer{}
	})

	Describe("Transformer", func() {
		Context("All cases", func() {
			It("Should run as expected", func() {
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

				result := productTransformer.TransformMultiple(data)

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
	})
})
