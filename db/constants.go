package db

import "visit-sidayu-backend/internal/models"

var ModelsRegistry = []any{
	&models.Roles{},
	&models.Officials{},

	&models.Images{},

	&models.Blogs{},
	&models.CultureBlogs{},
	&models.IndustriesBlogs{},
	&models.ShopsAndUmkmsBlogs{},

	&models.Timelines{},
	&models.TimelinesElement{},

	&models.Galleries{},
	&models.Geographies{},
	&models.Demographies{},
	&models.Superadmins{},
}
