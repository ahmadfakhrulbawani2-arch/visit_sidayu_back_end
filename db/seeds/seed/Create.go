package seed

import (
	"simple_go_gin_gorm_postgres_be_template/internal/helpers"
	"simple_go_gin_gorm_postgres_be_template/internal/models"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func CreateBlogs(db *gorm.DB, b models.CreateBlogs) error {
	blog := b.ToModel()
	tml := blog.Timeline
	blog.Timeline = nil

	return db.Transaction(func(tx *gorm.DB) error {
		// Gunakan .Create(), JANGAN gunakan .Save() atau .FirstOrCreate()
		if err := tx.Create(&blog).Error; err != nil {
			return err
		}

		if tml == nil {
			return nil
		}

		tml.ID = uuid.Nil
		tml.BlogID = &blog.ID
		if err := tx.Create(tml).Error; err != nil {
			return err
		}

		if len(tml.TimelineData) > 0 {
			for i := range tml.TimelineData {
				tml.TimelineData[i].TimelinesID = tml.ID
			}
			if err := tx.Create(&tml.TimelineData).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func CreateCultureBlogs(db *gorm.DB, b models.CreateCultureBlogsReq) error {
	culture := b.ToModel()
	return db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&culture).Error
	})
}

func CreateDemograhies(db *gorm.DB, d models.CreateDemographies) error {
	var demo models.Demographies
	copier.Copy(&demo, &d)
	return db.Create(&demo).Error
}

func CreateGalleries(db *gorm.DB, g models.CreateGalleries) error {
	var gallery models.Galleries
	copier.Copy(&gallery, &g)
	return db.Create(&gallery).Error
}

func CreateGeographies(db *gorm.DB, g models.CreateGeographies) error {
	var geo models.Geographies
	copier.Copy(&geo, &g)
	return db.Create(&geo).Error
}

func CreateImages(db *gorm.DB, req models.CreateImages) (*models.Images, error) {
	img := models.Images{}
	copier.Copy(&img, &req)
	err := db.Create(&img).Error
	return &img, err
}

func CreateIndustriesBlog(db *gorm.DB, i models.CreateIndustriesBlogsReq) error {
	ind := i.ToModel()

	return db.Create(&ind).Error
}

func CreateOfficial(db *gorm.DB, o models.CreateOfficials) error {
	var off models.Officials
	copier.Copy(&off, &o)
	return db.Create(&off).Error
}

func CreateShopsAndUmkmsBlog(db *gorm.DB, s models.CreateShopsAndUmkmsBlogsReq) error {
	shop := s.ToModel()

	return db.Create(&shop).Error
}

func CreateSuperadmins(db *gorm.DB, s models.CreateSuperadmins) error {
	var sa models.Superadmins
	copier.Copy(&sa, &s)
	hashedPw, err := helpers.HashPassword(s.Password)
	if err != nil {
		return err
	}
	sa.Password = hashedPw

	return db.Create(&sa).Error
}

func CreateTimelines(db *gorm.DB, t models.CreateTimelines) error {
	var time models.Timelines
	copier.Copy(&time, &t)
	return db.Create(&time).Error
}

func CreateRoles(db *gorm.DB, r models.CreateRoles) (*models.Roles, error) {
	role := models.Roles{}
	copier.Copy(&role, &r)
	err := db.Create(&role).Error
	return &role, err
}
