package db

import "visit-sidayu-backend/internal/models"

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
	&models.TimelinesElement{},
}
