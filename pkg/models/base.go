package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// WAJIB DICEK: Jika UUID belum ada nilainya (masih kosong), baru buat baru!
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}

// add table naming
func (CreateBlogs) TableName() string                 { return "blogs" }
func (CreateCultureBlogsReq) TableName() string       { return "culture_blogs" }
func (CreateDemographies) TableName() string          { return "demographies" }
func (CreateGalleries) TableName() string             { return "galleries" }
func (CreateGeographies) TableName() string           { return "geographies" }
func (CreateImages) TableName() string                { return "images" }
func (CreateIndustriesBlogsReq) TableName() string    { return "industries_blogs" }
func (CreateOfficials) TableName() string             { return "officials" }
func (CreateRoles) TableName() string                 { return "roles" }
func (CreateShopsAndUmkmsBlogsReq) TableName() string { return "shops_and_umkms_blogs" }
func (CreateSuperadmins) TableName() string           { return "superadmins" }
func (Superadmins) TableName() string                 { return "superadmins" }
func (CreateTimelines) TableName() string             { return "timelines" }

// add table insert mapping
func (b *CreateBlogs) ToModel() Blogs {
	var timelineModel *Timelines
	if b.Timeline != nil {
		var elements []TimelinesElement
		for _, el := range b.Timeline.TimelineData {
			elements = append(elements, TimelinesElement{
				Name:             el.Name,
				TimelineDatetime: el.TimelineDatetime,
				Description:      el.Description,
				ExternalLink:     el.ExternalLink,
			})
		}
		timelineModel = &Timelines{
			Name:         b.Timeline.Name,
			Description:  b.Timeline.Description,
			TimelineData: elements,
		}
	}

	return Blogs{
		Title:                b.Title,
		Description:          b.Description,
		Tags:                 b.Tags,
		Content:              b.Content,
		Author:               b.Author,
		BlogWrittenDatetime:  b.BlogWrittenDatetime,
		EstimatedMinutesRead: b.EstimatedMinutesRead,
		ThumbnailID:          b.ThumbnailID,
		Location:             b.Location,
		ExternalLinks:        b.ExternalLinks,
		Timeline:             timelineModel,
	}
}

func (c *CreateCultureBlogsReq) ToModel() CultureBlogs {
	var timelineModel *Timelines
	if c.Timeline != nil {
		timelineModel = &Timelines{
			Name:        c.Timeline.Name,
			Description: c.Timeline.Description,
		}
		var elements []TimelinesElement
		for _, el := range c.Timeline.TimelineData {
			elements = append(elements, TimelinesElement{
				Name:             el.Name,
				TimelineDatetime: el.TimelineDatetime,
				Description:      el.Description,
				ExternalLink:     el.ExternalLink,
			})
		}
		timelineModel.TimelineData = elements
	}

	return CultureBlogs{
		Title:                    c.Title,
		Description:              c.Description,
		Content:                  c.Content,
		Tags:                     pq.StringArray(c.Tags),
		ThemeType:                c.ThemeType,
		ThumbnailID:              c.ThumbnailID,
		Location:                 c.Location,
		Author:                   c.Author,
		BlogWrittenDatetime:      c.BlogWrittenDatetime,
		EstimatedMinutesReadTime: c.EstimatedMinutesReadTime,
		ExternalLinks:            pq.StringArray(c.ExternalLinks),
		Timeline:                 timelineModel,
	}
}

func (i *CreateIndustriesBlogsReq) ToModel() IndustriesBlogs {
	return IndustriesBlogs{
		Title:                      i.Title,
		Content:                    i.Content,
		Location:                   i.Location,
		Rating:                     i.Rating,
		Revenue:                    i.Revenue,
		ProducedProducts:           pq.StringArray(i.ProducedProducts),
		ProductionRatesPiecePerDay: i.ProductionRatesPiecePerDay,
		ThumbnailID:                i.ThumbnailID,
		YearFounded:                i.YearFounded,
		EmployeesCount:             i.EmployeesCount,
		BusinessType:               i.BusinessType,
	}
}

func (s *CreateShopsAndUmkmsBlogsReq) ToModel() ShopsAndUmkmsBlogs {
	return ShopsAndUmkmsBlogs{
		Title:                 s.Title,
		Content:               s.Content,
		Location:              s.Location,
		Rating:                s.Rating,
		Revenue:               s.Revenue,
		MarketedProducts:      pq.StringArray(s.MarketedProducts),
		SalesRatesPiecePerDay: s.SalesRatesPiecePerDay,
		ThumbnailID:           s.ThumbnailID,
	}
}
