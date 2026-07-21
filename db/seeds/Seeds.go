package seeds

import (
	"simple_go_gin_gorm_postgres_be_template/db/seeds/seed"
	"simple_go_gin_gorm_postgres_be_template/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	savedImage *models.Images
	savedRole  *models.Roles
)

func All() []seed.Seed {
	return []seed.Seed{
		{
			Name: "CreateJohnSuperadmin",
			Run: func(db *gorm.DB) error {
				return seed.CreateSuperadmins(db, models.CreateSuperadmins{
					Username: "John",
					Password: "john123",
					Email:    "johndoe@email.com",
				})
			},
		},
		{
			Name: "Create a mock image",
			Run: func(db *gorm.DB) error {
				var err error
				savedImage, err = seed.CreateImages(db, models.CreateImages{
					ImageUrl:    "https://example.com/sidayu.jpg",
					Name:        "Pemandangan Sidayu",
					Description: "Foto indah",
				})
				return err
			},
		},
		{
			Name: "Create a role mock data",
			Run: func(db *gorm.DB) error {
				var err error
				savedRole, err = seed.CreateRoles(db, models.CreateRoles{
					Name:        "Kepala Kecamatan",
					Level:       1,
					Description: "Pemimpin tertinggi di kecamatan",
				})
				return err
			},
		},
		{
			Name: "CreateBlogsSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateBlogs(db, models.CreateBlogs{
					Title:       "Test Blog",
					Description: "This is test blog",
					Tags:        []string{"Tag1", "Tag2", "tag3"},
					Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ",
					Author:      "John Doe",
					BlogWrittenDatetime: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2026-07-20T09:41:48Z")
						return t
					}(),
					EstimatedMinutesRead: 60,
					ThumbnailID:          (*uuid.UUID)(&savedImage.ID),
					Location:             "Sidayu",
					ExternalLinks:        nil,
					Timeline: &models.CreateTimelines{
						Name:        "Timeline test blog",
						Description: "timeline untuk test blog",
						TimelineData: []models.CreateTimelinesElement{
							{
								Name: "Important event 1",
								TimelineDatetime: func() time.Time {
									t, _ := time.Parse(time.RFC3339, "2026-07-20T09:41:48Z")
									return t
								}(),
								Description:  "Keterangan event 1",
								ExternalLink: "https://google.com",
							},
							{
								Name: "Important event 2",
								TimelineDatetime: func() time.Time {
									t, _ := time.Parse(time.RFC3339, "2026-07-20T09:41:48Z")
									return t
								}(),
								Description:  "Keterangan event 2",
								ExternalLink: "https://instagram.com",
							},
						},
					},
				})
			},
		},
		{
			Name: "CreateCultureBlogsSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateCultureBlogs(db, models.CreateCultureBlogsReq{
					Title:       "Tradisi Lokal Sidayu",
					Description: "Membahas tradisi unik yang ada di Sidayu",
					Content:     "Konten mendalam mengenai sejarah dan budaya Sidayu...",
					Tags:        []string{"Budaya", "Sejarah", "Sidayu"},
					ThemeType:   "Tradisi",
					ThumbnailID: (*uuid.UUID)(&savedImage.ID),
					Location:    "Sidayu",
					Author:      "Admin Budaya",
					BlogWrittenDatetime: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2026-07-20T09:41:48Z")
						return t
					}(),
					EstimatedMinutesReadTime: 15,
					ExternalLinks:            nil,
					Timeline: &models.CreateTimelines{
						Name:        "Timeline test blog",
						Description: "timeline untuk test blog",
						TimelineData: []models.CreateTimelinesElement{
							{
								Name: "Important event on culture 1",
								TimelineDatetime: func() time.Time {
									t, _ := time.Parse(time.RFC3339, "2026-07-20T09:41:48Z")
									return t
								}(),
								Description:  "Keterangan event on culture 1",
								ExternalLink: "https://google.com",
							},
							{
								Name: "Important event on culture 2",
								TimelineDatetime: func() time.Time {
									t, _ := time.Parse(time.RFC3339, "2026-07-20T09:41:48Z")
									return t
								}(),
								Description:  "Keterangan event on culture 2",
								ExternalLink: "https://instagram.com",
							},
						},
					},
				})
			},
		},
		{
			Name: "CreateDemographiesSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateDemograhies(db, models.CreateDemographies{
					VillageName:            "Desa Sidayu",
					DemographyDataYear:     2026,
					MalePopulation:         5000,
					FemalePopulation:       5200,
					TotalPopulation:        10200,
					PopulationDensityUnit:  "Orang/km2",
					FamiliesNumber:         2500,
					NumberOfBirth:          120,
					NumberOfDeath:          45,
					WorkingPopulation:      6000,
					UnemployedPopulation:   500,
					HousekeepingPopulation: 2000,
					StudentPopulation:      1700,
					SourceName:             "BPS Kabupaten",
					ExternalLinkSource:     "https://bps.go.id",
				})
			},
		},
		{
			Name: "CreateGalleriesSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateGalleries(db, models.CreateGalleries{
					Name:        "Galeri Sidayu",
					Description: "Deskripsi galeri",
					ImageID:     savedImage.ID,
				})
			},
		},
		{
			Name: "CreateGeographiesSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateGeographies(db, models.CreateGeographies{
					VillageName:  "Desa Sidayu",
					Area:         150.5,
					AreaUnit:     "km2",
					RainfallRate: 2500.0,
					RainfallUnit: "mm/tahun",
					RainyDay:     120,
					ImageID:      nil,
					Source:       "Data Geografis Desa 2026",
				})
			},
		},
		{
			Name: "CreateIndustriesBlogSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateIndustriesBlog(db, models.CreateIndustriesBlogsReq{
					Title:                      "Industri Kreatif Sidayu",
					Content:                    "Analisis mendalam mengenai perkembangan industri di Sidayu...",
					Location:                   "Sidayu, Gresik",
					Rating:                     4.8,
					Revenue:                    500000000.0,
					ProducedProducts:           []string{"Kerajinan Tangan", "Alat Rumah Tangga"},
					ProductionRatesPiecePerDay: 50,
					ThumbnailID:                nil,
					YearFounded:                2020,
					EmployeesCount:             25,
					BusinessType:               "UMKM",
				})
			},
		},
		{
			Name: "CreateOfficialsSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateOfficial(db, models.CreateOfficials{
					Name:           "Bapak Sidayu",
					Description:    "Kepala kecamatan periode 2026-2031",
					RoleID:         savedRole.ID,
					ProfileImageID: (*uuid.UUID)(&savedImage.ID),
				})
			},
		},
		{
			Name: "CreateShopsAndUmkmsBlogSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateShopsAndUmkmsBlog(db, models.CreateShopsAndUmkmsBlogsReq{
					Title:                 "Toko Kelontong Berkah",
					Content:               "Toko kelontong terlengkap di pusat desa Sidayu.",
					Location:              "Jl. Utama Sidayu No. 12",
					Rating:                4.9,
					Revenue:               15000000.0,
					MarketedProducts:      []string{"Sembako", "Alat Tulis", "Minuman Ringan"},
					SalesRatesPiecePerDay: 120,
					ThumbnailID:           (*uuid.UUID)(&savedImage.ID),
				})
			},
		},
	}
}
