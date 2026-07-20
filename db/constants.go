package db

import "simple_go_gin_gorm_postgres_be_template/internal/models"

var ModelsRegistry = []any{
	&models.Blogs{},
	&models.CultureBlogs{},
	&models.Demographies{},
	&models.Galleries{},
	&models.Geographies{},
	&models.Images{},
	&models.IndustriesBlogs{},
	&models.Officials{},
	&models.Roles{},
	&models.ShopsAndUmkmsBlogs{},
	&models.Superadmins{},
	&models.Timelines{},
}
