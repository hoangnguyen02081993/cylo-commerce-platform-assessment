package main

import (
	c "commerce-platform/core/config"
	database "commerce-platform/core/database"
	a "commerce-platform/modules/activities"
	p "commerce-platform/modules/products"
)

func main() {
	manager := database.GetInstance(c.GetConfig())
	db, err := manager.GetDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&p.Product{})
	db.AutoMigrate(&a.Activity{})

	db.Exec("TRUNCATE TABLE products")
	db.Exec(`INSERT INTO products (name, img_url, description, price, branch, is_deleted, created_at, created_by, updated_at, updated_by) VALUES 
			('Sài Gòn sider', '/imgs/beer-sg/Sai_Gon_chai_1.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 32.0, 'Bia Sài Gòn', 0, NOW(), 'admin', NULL, NULL),
			('Sài Gòn Special (chai)', '/imgs/beer-sg/Sai_Gon_Chai.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 18.0, 'Bia Sài Gòn', 0, NOW(), 'admin', NULL, NULL),
			('Sài Gòn Lager', '/imgs/beer-sg/Sai_Gon_Lager.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 27.0, 'Bia Sài Gòn', 0, NOW(), 'admin', NULL, NULL),
			('Sài Gòn Special (lon)', '/imgs/beer-sg/Sai_Gon_Special.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 18.0, 'Bia Sài Gòn', 0, NOW(), 'admin', NULL, NULL),
			('Cơm Bento (combo 7)', '/imgs/bento/Com_bento_1.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 50.0, 'Cơm Bento', 0, NOW(), 'admin', NULL, NULL),
			('Cơm Bento (size đại)', '/imgs/bento/Com_bento_3.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 40.0, 'Cơm Bento', 0, NOW(), 'admin', NULL, NULL),
			('Cơm Bento (combo mini)', '/imgs/bento/Com_bento_4.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 52.0, 'Cơm Bento', 0, NOW(), 'admin', NULL, NULL),
			('Cơm Bento (cho bé)', '/imgs/bento/Com_bento_12.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 36.0, 'Cơm Bento', 0, NOW(), 'admin', NULL, NULL),
			('Bánh ngọt hình trái tim', '/imgs/cake/Banh_1.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 80.0, 'Bánh Ngọt', 0, NOW(), 'admin', NULL, NULL),
			('Sữa tươi (size lớn)', '/imgs/milk/Vinamilk_1.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 62.0, 'Sữa tươi', 0, NOW(), 'admin', NULL, NULL),
			('Sữa tươi (chai thuỷ tinh)', '/imgs/milk/Vinamilk_2.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 20.0, 'Sữa tươi', 0, NOW(), 'admin', NULL, NULL),
			('Sữa tươi (hộp nhỏ)', '/imgs/milk/Vinamilk_3.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 27.0, 'Sữa tươi', 0, NOW(), 'admin', NULL, NULL),
			('Sữa tươi (bé 10 tuổi)', '/imgs/milk/Vinamilk_4.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 17.0, 'Sữa tươi', 0, NOW(), 'admin', NULL, NULL),
			('Sữa tươi (combo 2 bịch)', '/imgs/milk/Vinamilk_5.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 49.0, 'Sữa tươi', 0, NOW(), 'admin', NULL, NULL),
			('Sữa tươi (yogurt vị dâu)', '/imgs/milk/Vinamilk_6.jpg', 'Thật khó để có thể tìm kiếm được description cho nó đẹp trang web. Nên bây giờ chỉ cố gắng viết description cho các sản phẩm y xì nhau thôi nha.', 12.0, 'Sữa tươi', 0, NOW(), 'admin', NULL, NULL);`)
}
