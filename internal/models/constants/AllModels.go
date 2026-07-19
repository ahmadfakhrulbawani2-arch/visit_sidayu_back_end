package constants

import "simple_go_gin_gorm_postgres_be_template/internal/models"

var ModelsRegistry = []any{
	&models.Blogs{},
	&models.CultureBlog{},
	&models.Demographies{},
	&models.Galleries{},
	&models.Geographies{},
	&models.Images{},
	&models.IndustriesBlog{},
	&models.Officials{},
	&models.Roles{},
	&models.ShopsAndUmkmsBlog{},
	&models.Superadmins{},
	&models.Timelines{},
}
