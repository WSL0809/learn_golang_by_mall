package dao

import "init_golang/model"

func migration() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Category{}, 
		&model.Product{}, 
		&model.ProductImg{}, 
		&model.Carousel{}, 
		&model.Notice{}, 
		&model.Address{}, 
		&model.Admin{}, 
		&model.User{}, 
		&model.Favourite{}, 
		&model.Cart{}, 
		&model.Order{},
	)
	return
}
