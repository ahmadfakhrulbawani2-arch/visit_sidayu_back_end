package seeds

import (
	"time"
	"visit-sidayu-backend/db/seeds/seed"
	"visit-sidayu-backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	SavedImage *models.Images
	SavedRole  *models.Roles
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
				SavedImage, err = seed.CreateImages(db, models.CreateImages{
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
				SavedRole, err = seed.CreateRoles(db, models.CreateRoles{
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
					ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
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
		// {
		// 	Name: "CreateExtraBlogsSample",
		// 	Run: func(db *gorm.DB) error {
		// 		blogs := []models.CreateBlogs{
		// 			{
		// 				Title:       "Pengenalan Pemrograman Go",
		// 				Description: "Dasar-dasar bahasa pemrograman Go untuk pemula",
		// 				Tags:        []string{"Golang", "Programming", "Backend"},
		// 				Content:     "Go adalah bahasa pemrograman sumber terbuka yang memudahkan pembuatan perangkat lunak yang sederhana, cepat, dan andal.",
		// 				Author:      "Ahmad Fauzi",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-01T08:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 15,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Surabaya",
		// 				ExternalLinks:        []string{"https://golang.org"},
		// 				Timeline: &models.CreateTimelines{
		// 					Name:        "Sejarah Singkat Go",
		// 					Description: "Perkembangan bahasa Go dari awal hingga stabil",
		// 					TimelineData: []models.CreateTimelinesElement{
		// 						{
		// 							Name: "Rilis Internal Google",
		// 							TimelineDatetime: func() time.Time {
		// 								t, _ := time.Parse(time.RFC3339, "2007-09-21T00:00:00Z")
		// 								return t
		// 							}(),
		// 							Description:  "Proyek mulai dikembangkan oleh Google.",
		// 							ExternalLink: "https://google.com",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Title:       "Arsitektur Microservices Modern",
		// 				Description: "Membangun sistem terdistribusi yang scalable",
		// 				Tags:        []string{"Microservices", "Architecture", "Backend"},
		// 				Content:     "Microservices adalah pendekatan arsitektur di mana aplikasi dibagi menjadi layanan-layanan independen kecil.",
		// 				Author:      "Siti Rahma",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-05T10:30:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 25,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Jakarta",
		// 				ExternalLinks:        nil,
		// 				Timeline: &models.CreateTimelines{
		// 					Name:        "Migrasi Monolit ke Microservices",
		// 					Description: "Tahapan transisi sistem",
		// 					TimelineData: []models.CreateTimelinesElement{
		// 						{
		// 							Name: "Analisis Domain",
		// 							TimelineDatetime: func() time.Time {
		// 								t, _ := time.Parse(time.RFC3339, "2026-06-01T09:00:00Z")
		// 								return t
		// 							}(),
		// 							Description:  "Memecah bounded context.",
		// 							ExternalLink: "https://example.com",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Title:       "Optimasi Query Database SQL",
		// 				Description: "Tips dan trik mempercepat eksekusi database",
		// 				Tags:        []string{"Database", "SQL", "Performance"},
		// 				Content:     "Indeks yang tepat dan struktur query yang efisien dapat memangkas waktu eksekusi secara drastis.",
		// 				Author:      "Budi Santoso",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-10T14:15:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 20,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Bandung",
		// 				ExternalLinks:        []string{"https://postgresql.org"},
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Memahami Docker & Containerization",
		// 				Description: "Panduan lengkap Docker untuk developer",
		// 				Tags:        []string{"DevOps", "Docker", "Container"},
		// 				Content:     "Docker memungkinkan pengembang mengemas aplikasi beserta seluruh dependensinya ke dalam container.",
		// 				Author:      "Dewi Lestari",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-12T11:20:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 30,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Yogyakarta",
		// 				ExternalLinks:        nil,
		// 				Timeline: &models.CreateTimelines{
		// 					Name:        "Adopsi Docker",
		// 					Description: "Timeline adopsi container di tim",
		// 					TimelineData: []models.CreateTimelinesElement{
		// 						{
		// 							Name: "Setup Dockerfile",
		// 							TimelineDatetime: func() time.Time {
		// 								t, _ := time.Parse(time.RFC3339, "2026-06-10T08:00:00Z")
		// 								return t
		// 							}(),
		// 							Description:  "Membuat image dasar.",
		// 							ExternalLink: "https://docker.com",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Title:       "Belajar Next.js App Router",
		// 				Description: "Eksplorasi fitur terbaru Next.js untuk frontend modern",
		// 				Tags:        []string{"React", "Next.js", "Frontend"},
		// 				Content:     "Next.js App Router membawa paradigma baru berbasis Server Components yang sangat efisien.",
		// 				Author:      "Fakhrul",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-15T16:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 40,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Surabaya",
		// 				ExternalLinks:        []string{"https://nextjs.org"},
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Pengantar Machine Learning dengan Python",
		// 				Description: "Mengenal dasar AI dan library populer seperti Scikit-Learn",
		// 				Tags:        []string{"AI", "Python", "Machine Learning"},
		// 				Content:     "Machine learning adalah cabang dari kecerdasan buatan yang fokus pada pembangunan sistem yang bisa belajar dari data.",
		// 				Author:      "Rian Hidayat",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-18T13:45:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 45,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Malang",
		// 				ExternalLinks:        nil,
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Manajemen State di React",
		// 				Description: "Context API vs Redux Toolkit vs Zustand",
		// 				Tags:        []string{"React", "Frontend", "State Management"},
		// 				Content:     "Memilih state management yang tepat sangat krusial untuk skala aplikasi frontend yang kompleks.",
		// 				Author:      "Siti Rahma",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-20T09:10:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 20,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Jakarta",
		// 				ExternalLinks:        nil,
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Keamanan Web: Mencegah SQL Injection",
		// 				Description: "Praktik terbaik mengamankan database aplikasi",
		// 				Tags:        []string{"Security", "Backend", "Web"},
		// 				Content:     "SQL Injection adalah celah keamanan klasik yang masih sering terjadi akibat kurangnya validasi input.",
		// 				Author:      "Ahmad Fauzi",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-22T15:30:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 25,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Surabaya",
		// 				ExternalLinks:        nil,
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Membuat RESTful API dengan Gin Framework",
		// 				Description: "Framework Go yang cepat dan ringan untuk web API",
		// 				Tags:        []string{"Golang", "Gin", "API"},
		// 				Content:     "Gin adalah web framework yang ditulis dengan Go yang menyediakan martini-like API dengan performa tinggi.",
		// 				Author:      "Fakhrul",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-25T11:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 35,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Surabaya",
		// 				ExternalLinks:        []string{"https://github.com/gin-gonic/gin"},
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Tips Desain UI/UX untuk Developer",
		// 				Description: "Memahami prinsip dasar estetika antarmuka pengguna",
		// 				Tags:        []string{"UI/UX", "Design", "Frontend"},
		// 				Content:     "Desain yang baik bukan hanya soal tampilan yang cantik, tetapi juga kemudahan navigasi bagi pengguna.",
		// 				Author:      "Dewi Lestari",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-28T08:30:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 15,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Yogyakarta",
		// 				ExternalLinks:        nil,
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Dasar-dasar Algoritma Competitive Programming",
		// 				Description: "Meningkatkan kemampuan logic dan struktur data",
		// 				Tags:        []string{"Algorithms", "C++", "Competitive Programming"},
		// 				Content:     "Kompetisi pemrograman melatih kecepatan berpikir dan pemilihan kompleksitas algoritma yang optimal.",
		// 				Author:      "Fakhrul",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-01T12:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 50,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Surabaya",
		// 				ExternalLinks:        []string{"https://codeforces.com"},
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Pengenalan Linux Command Line",
		// 				Description: "Navigasi terminal dan manajemen sistem operasi berbasis Linux",
		// 				Tags:        []string{"Linux", "DevOps", "OS"},
		// 				Content:     "Command line interface memberikan kontrol penuh dan efisiensi tinggi dalam mengelola server.",
		// 				Author:      "Budi Santoso",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-03T14:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 30,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Bandung",
		// 				ExternalLinks:        nil,
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Membangun CI/CD Pipeline dengan GitHub Actions",
		// 				Description: "Otomatisasi testing dan deployment aplikasi",
		// 				Tags:        []string{"CI/CD", "GitHub", "DevOps"},
		// 				Content:     "GitHub Actions mempermudah integrasi proses build dan rilis kode secara otomatis setiap kali push.",
		// 				Author:      "Rian Hidayat",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-06T10:15:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 35,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Malang",
		// 				ExternalLinks:        nil,
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Pentingnya Clean Code dalam Tim",
		// 				Description: "Menulis kode yang mudah dibaca dan dipelihara bersama",
		// 				Tags:        []string{"Clean Code", "Best Practices", "Software Engineering"},
		// 				Content:     "Kode ditulis sekali, tetapi dibaca berkali-kali oleh rekan satu tim maupun diri sendiri di masa depan.",
		// 				Author:      "Ahmad Fauzi",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-10T09:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 20,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Surabaya",
		// 				ExternalLinks:        nil,
		// 				Timeline:             nil,
		// 			},
		// 			{
		// 				Title:       "Eksplorasi Fitur Tailwind CSS v4",
		// 				Description: "Mengenal pembaruan performa dan konfigurasi di Tailwind versi terbaru",
		// 				Tags:        []string{"TailwindCSS", "CSS", "Frontend"},
		// 				Content:     "Tailwind CSS v4 hadir dengan kecepatan kompilasi yang jauh lebih tinggi dan pendekatan berbasis CSS file.",
		// 				Author:      "Fakhrul",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-15T16:20:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesRead: 25,
		// 				ThumbnailID:          (*uuid.UUID)(&SavedImage.ID),
		// 				Location:             "Surabaya",
		// 				ExternalLinks:        []string{"https://tailwindcss.com"},
		// 				Timeline:             nil,
		// 			},
		// 		}

		// 		for _, b := range blogs {
		// 			if err := seed.CreateBlogs(db, b); err != nil {
		// 				return err
		// 			}
		// 		}
		// 		return nil
		// 	},
		// },
		{
			Name: "CreateCultureBlogsSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateCultureBlogs(db, models.CreateCultureBlogsReq{
					Title:       "Tradisi Lokal Sidayu",
					Description: "Membahas tradisi unik yang ada di Sidayu",
					Content:     "Konten mendalam mengenai sejarah dan budaya Sidayu...",
					Tags:        []string{"Budaya", "Sejarah", "Sidayu"},
					ThemeType:   "Tradisi",
					ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
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
		// {
		// 	Name: "CreateExtraCultureBlogsSample",
		// 	Run: func(db *gorm.DB) error {
		// 		cultureBlogs := []models.CreateCultureBlogsReq{
		// 			{
		// 				Title:       "Sedekah Bumi di Pesisir Sidayu",
		// 				Description: "Ritual ungkapan rasa syukur atas hasil bumi dan laut",
		// 				Content:     "Sedekah bumi merupakan tradisi turun-temurun yang dilakukan oleh masyarakat petani dan nelayan di Sidayu sebagai bentuk rasa syukur kepada Tuhan Yang Maha Esa.",
		// 				Tags:        []string{"Budaya", "Tradisi", "Sidayu", "Lokal"},
		// 				ThemeType:   "Tradisi",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Budaya",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-02T09:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 12,
		// 				ExternalLinks:            nil,
		// 				Timeline: &models.CreateTimelines{
		// 					Name:        "Rangkaian Acara Sedekah Bumi",
		// 					Description: "Tahapan pelaksanaan tradisi",
		// 					TimelineData: []models.CreateTimelinesElement{
		// 						{
		// 							Name: "1. Bersih Desa",
		// 							TimelineDatetime: func() time.Time {
		// 								t, _ := time.Parse(time.RFC3339, "2026-06-01T07:00:00Z")
		// 								return t
		// 							}(),
		// 							Description:  "Warga kerja bakti membersihkan lingkungan sekitar.",
		// 							ExternalLink: "https://example.com",
		// 						},
		// 						{
		// 							Name: "2. Doa Bersama & Kenduri",
		// 							TimelineDatetime: func() time.Time {
		// 								t, _ := time.Parse(time.RFC3339, "2026-06-02T15:00:00Z")
		// 								return t
		// 							}(),
		// 							Description:  "Puncak acara makan bersama dan doa keselamatan.",
		// 							ExternalLink: "https://example.com",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Title:       "Jejak Sejarah Kanjeng Sepuh Sidayu",
		// 				Description: "Mengenang jasa Bupati pertama Sidayu dalam penyebaran Islam",
		// 				Content:     "Kanjeng Sepuh (K.R.T. SetrowIDjojo II) adalah tokoh penting yang memimpin Sidayu dengan penuh kebijaksanaan serta memperkuat nilai-nilai keislaman.",
		// 				Tags:        []string{"Sejarah", "Tokoh", "Sidayu", "Religi"},
		// 				ThemeType:   "Sejarah",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Sejarah",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-05T10:30:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 18,
		// 				ExternalLinks:            []string{"https://id.wikipedia.org"},
		// 				Timeline: &models.CreateTimelines{
		// 					Name:        "Masa Kepemimpinan Kanjeng Sepuh",
		// 					Description: "Tonggak sejarah penting",
		// 					TimelineData: []models.CreateTimelinesElement{
		// 						{
		// 							Name: "Pengangkatan Menjadi Bupati",
		// 							TimelineDatetime: func() time.Time {
		// 								t, _ := time.Parse(time.RFC3339, "1816-01-01T00:00:00Z")
		// 								return t
		// 							}(),
		// 							Description:  "Memulai masa bakti memajukan wilayah Sidayu.",
		// 							ExternalLink: "https://example.com",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Title:       "Arsitektur Kuno Masjid Besar Sidayu",
		// 				Description: "Keunikan bentuk bangunan peninggalan masa lampau",
		// 				Content:     "Masjid Besar Sidayu menyimpan arsitektur klasik perpaduan antara gaya Jawa kuno dan sentuhan kolonial, menjadi saksi bisu perkembangan dakwah Islam.",
		// 				Tags:        []string{"Arsitektur", "Religi", "Sejarah", "Sidayu"},
		// 				ThemeType:   "Arsitektur",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Budaya",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-10T14:15:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 10,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Kuliner Khas Pesisiran: Pindang Bandeng Sidayu",
		// 				Description: "Cita rasa olahan bandeng legendaris khas Gresik utara",
		// 				Content:     "Kecamatan Sidayu terkenal dengan produksi tambak bandeng berkualitas. Salah satu olahan khasnya adalah pindang bandeng dengan bumbu rempah khas tradisional.",
		// 				Tags:        []string{"Kuliner", "Tradisi", "Makanan", "Sidayu"},
		// 				ThemeType:   "Kuliner",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Kuliner",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-12T11:20:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 8,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Kesenian Lokal: Patrol Sidayu",
		// 				Description: "Tradisi membangunkan sahur dengan alat musik bambu keliling kampung",
		// 				Content:     "Menjelang bulan Ramadhan hingga hari H, kelompok pemuda di Sidayu kerap menghidupkan tradisi musik patrol keliling lorong desa untuk membangunkan warga.",
		// 				Tags:        []string{"Kesenian", "Musik", "Tradisi", "Ramadhan"},
		// 				ThemeType:   "Kesenian",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Budaya",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-15T16:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 14,
		// 				ExternalLinks:            nil,
		// 				Timeline: &models.CreateTimelines{
		// 					Name:        "Jadwal Keliling Patrol",
		// 					Description: "Rute malam ganjil Ramadhan",
		// 					TimelineData: []models.CreateTimelinesElement{
		// 						{
		// 							Name: "Start Alun-Alun Sidayu",
		// 							TimelineDatetime: func() time.Time {
		// 								t, _ := time.Parse(time.RFC3339, "2026-06-15T01:00:00Z")
		// 								return t
		// 							}(),
		// 							Description:  "Berkumpul dan mulai rute keliling kampung.",
		// 							ExternalLink: "https://example.com",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Title:       "Mitos dan Legenda Sumur Kuno di Sidayu",
		// 				Description: "Cerita rakyat yang diwariskan secara turun-temurun",
		// 				Content:     "Berbagai sumur tua peninggalan zaman wali atau bangsawan lokal di Sidayu menyimpan cerita unik yang melekat di ingatan masyarakat setempat.",
		// 				Tags:        []string{"Legenda", "Cerita Rakyat", "Sidayu"},
		// 				ThemeType:   "Folklore",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Sejarah",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-18T13:45:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 20,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Peran Pesantren Tua di Wilayah Sidayu",
		// 				Description: "Pusat pendidikan Islam klasik yang melahirkan banyak ulama",
		// 				Content:     "Sidayu sejak dulu dikenal sebagai salah satu basis pendidikan Islam tradisional di kawasan pantura Gresik.",
		// 				Tags:        []string{"Pendidikan", "Religi", "Sejarah", "Pesantren"},
		// 				ThemeType:   "Pendidikan",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Religi",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-20T09:10:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 16,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Potensi Kerajinan Tangan Lokal Sidayu",
		// 				Description: "Kreativitas warga dalam mengolah bahan baku lokal menjadi produk bernilai",
		// 				Content:     "UMKM lokal di Sidayu terus berkembang melalui berbagai kerajinan tangan kreatif khas pedesaan.",
		// 				Tags:        []string{"Ekonomi", "Kerajinan", "UMKM", "Sidayu"},
		// 				ThemeType:   "Ekonomi Kreatif",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Ekonomi",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-22T15:30:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 11,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Menelusuri Pasar Tradisional Legendaris Sidayu",
		// 				Description: "Pusat denyut nadi perekonomian rakyat tempo dulu hingga kini",
		// 				Content:     "Pasar tradisional di Sidayu tidak hanya menjadi tempat bertransaksi jual beli, tetapi juga ruang interaksi sosial budaya masyarakat.",
		// 				Tags:        []string{"Pasar", "Tradisi", "Ekonomi", "Sidayu"},
		// 				ThemeType:   "Sosial",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Budaya",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-25T11:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 13,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Kearifan Lokal Pengelolaan Tambak di Sidayu",
		// 				Description: "Teknik turun-temurun budidaya ikan bandeng dan udang",
		// 				Content:     "Masyarakat pesisir Sidayu memiliki pengetahuan ekologis lokal dalam menjaga sirkulasi air tambak agar hasil panen tetap optimal.",
		// 				Tags:        []string{"Lingkungan", "Tradisi", "Perikanan", "Sidayu"},
		// 				ThemeType:   "Lingkungan",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Lingkungan",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-06-28T08:30:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 15,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Bahasa dan Dialek Khas Pesisiran Sidayu",
		// 				Description: "Keunikan logat dan kosakata keseharian warga setempat",
		// 				Content:     "Meskipun menggunakan bahasa Jawa, logat dan beberapa istilah khusus di wilayah Sidayu memiliki ciri khas tersendiri dibanding daerah lain di Jawa Timur.",
		// 				Tags:        []string{"Bahasa", "Budaya", "Lokal", "Sidayu"},
		// 				ThemeType:   "Linguistik",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Budaya",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-01T12:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 9,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Situs Makam Kuno dan Wisata Religi Sidayu",
		// 				Description: "Destinasi ziarah yang sering dikunjungi peziarah luar kota",
		// 				Content:     "Keberadaan makam para tokoh penyebar agama di Sidayu menjadikannya salah satu titik penting wisata religi di Gresik utara.",
		// 				Tags:        []string{"Religi", "Ziarah", "Sejarah", "Sidayu"},
		// 				ThemeType:   "Religi",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Religi",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-03T14:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 12,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Perayaan Hari Besar Islam di Kampung-kampung Sidayu",
		// 				Description: "Meriahnya tradisi memperingati Maulid Nabi dan Isra Mi'raj",
		// 				Content:     "Warga Sidayu selalu antusias merayakan hari besar keagamaan dengan berbagai tradisi unik, seperti pembagian berkat dan pawai obor.",
		// 				Tags:        []string{"Tradisi", "Religi", "Islam", "Sidayu"},
		// 				ThemeType:   "Tradisi",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Budaya",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-06T10:15:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 17,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Potensi Wisata Mangrove Pesisir Sidayu",
		// 				Description: "Keindahan alam hijau dan pelestarian ekosistem pantai",
		// 				Content:     "Kawasan konservasi mangrove di pesisir Sidayu kini mulai dikembangkan menjadi destinasi ekowisata edukatif yang menarik untuk dikunjungi.",
		// 				Tags:        []string{"Wisata", "Alam", "Ekowisata", "Sidayu"},
		// 				ThemeType:   "Wisata",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Wisata",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-10T09:00:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 10,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 			{
		// 				Title:       "Harmoni Kehidupan Masyarakat Multikultural Sidayu",
		// 				Description: "Menjaga kerukunan antarwarga dalam balutan tradisi lokal",
		// 				Content:     "Sikap saling menghormati dan gotong royong yang kuat menjadi fondasi utama kokohnya kebersamaan antarwarga di Sidayu dari masa ke masa.",
		// 				Tags:        []string{"Sosial", "Harmoni", "Budaya", "Sidayu"},
		// 				ThemeType:   "Sosial",
		// 				ThumbnailID: (*uuid.UUID)(&SavedImage.ID),
		// 				Location:    "Sidayu",
		// 				Author:      "Admin Budaya",
		// 				BlogWrittenDatetime: func() time.Time {
		// 					t, _ := time.Parse(time.RFC3339, "2026-07-15T16:20:00Z")
		// 					return t
		// 				}(),
		// 				EstimatedMinutesReadTime: 11,
		// 				ExternalLinks:            nil,
		// 				Timeline:                 nil,
		// 			},
		// 		}

		// 		for _, cb := range cultureBlogs {
		// 			if err := seed.CreateCultureBlogs(db, cb); err != nil {
		// 				return err
		// 			}
		// 		}
		// 		return nil
		// 	},
		// },
		{
			Name: "CreateGeographiesSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateGeographies(db, models.CreateGeographies{
					VillageName: "Desa Sidayu",
					Area:        150.5,
					AreaUnit:    "km2",
					// RainfallRate: 2500.0,
					RainfallUnit: "mm/tahun",
					RainyDay:     120,
					ImageID:      nil,
					Source:       "Data Geografis Desa 2026",
				})
			},
		},
		// {
		// 	Name: "CreateExtraGeographiesSample",
		// 	Run: func(db *gorm.DB) error {
		// 		geographies := []models.CreateGeographies{
		// 			{
		// 				VillageName:  "Desa Golokan",
		// 				Area:         85.2,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2400.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     115,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Kertosono",
		// 				Area:         110.0,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2450.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     118,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Purwodadi",
		// 				Area:         75.4,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2350.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     110,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Randuboto",
		// 				Area:         95.8,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2420.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     116,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Ngaglik",
		// 				Area:         60.1,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2300.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     105,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Lasem",
		// 				Area:         105.3,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2480.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     119,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Mentaras",
		// 				Area:         88.9,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2390.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     112,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Wadeng",
		// 				Area:         130.2,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2520.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     122,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Mojo",
		// 				Area:         70.6,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2360.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     108,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Sidomulyo",
		// 				Area:         98.4,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2430.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     117,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Kauman",
		// 				Area:         55.0,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2280.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     102,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Cerme",
		// 				Area:         115.7,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2460.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     118,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Petung",
		// 				Area:         82.3,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2370.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     111,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 			{
		// 				VillageName:  "Desa Bunder",
		// 				Area:         122.6,
		// 				AreaUnit:     "km2",
		// 				RainfallRate: 2500.0,
		// 				RainfallUnit: "mm/tahun",
		// 				RainyDay:     120,
		// 				ImageID:      nil,
		// 				Source:       "Data Geografis Desa 2026",
		// 			},
		// 		}

		// 		for _, g := range geographies {
		// 			if err := seed.CreateGeographies(db, g); err != nil {
		// 				return err
		// 			}
		// 		}
		// 		return nil
		// 	},
		// },
		{
			Name: "CreateDemographiesSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateDemograhies(db, models.CreateDemographies{
					VillageName:            "Desa Sidayu",
					DemographyDataYear:     2026,
					MalePopulation:         5000,
					FemalePopulation:       5200,
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
		// {
		// 	Name: "CreateExtraDemographiesSample",
		// 	Run: func(db *gorm.DB) error {
		// 		demographics := []models.CreateDemographies{
		// 			{
		// 				VillageName:            "Desa Golokan",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         3100,
		// 				FemalePopulation:       3050,
		// 				FamiliesNumber:         1500,
		// 				NumberOfBirth:          75,
		// 				NumberOfDeath:          28,
		// 				WorkingPopulation:      3600,
		// 				UnemployedPopulation:   300,
		// 				HousekeepingPopulation: 1200,
		// 				StudentPopulation:      1050,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Kertosono",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         4200,
		// 				FemalePopulation:       4300,
		// 				FamiliesNumber:         2100,
		// 				NumberOfBirth:          95,
		// 				NumberOfDeath:          35,
		// 				WorkingPopulation:      4900,
		// 				UnemployedPopulation:   400,
		// 				HousekeepingPopulation: 1700,
		// 				StudentPopulation:      1500,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Purwodadi",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         2800,
		// 				FemalePopulation:       2750,
		// 				FamiliesNumber:         1350,
		// 				NumberOfBirth:          60,
		// 				NumberOfDeath:          22,
		// 				WorkingPopulation:      3200,
		// 				UnemployedPopulation:   250,
		// 				HousekeepingPopulation: 1100,
		// 				StudentPopulation:      1000,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Randuboto",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         3500,
		// 				FemalePopulation:       3600,
		// 				FamiliesNumber:         1750,
		// 				NumberOfBirth:          85,
		// 				NumberOfDeath:          30,
		// 				WorkingPopulation:      4100,
		// 				UnemployedPopulation:   350,
		// 				HousekeepingPopulation: 1400,
		// 				StudentPopulation:      1250,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Ngaglik",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         2400,
		// 				FemalePopulation:       2500,
		// 				FamiliesNumber:         1200,
		// 				NumberOfBirth:          50,
		// 				NumberOfDeath:          18,
		// 				WorkingPopulation:      2800,
		// 				UnemployedPopulation:   200,
		// 				HousekeepingPopulation: 1000,
		// 				StudentPopulation:      900,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Lasem",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         3900,
		// 				FemalePopulation:       3850,
		// 				FamiliesNumber:         1900,
		// 				NumberOfBirth:          90,
		// 				NumberOfDeath:          32,
		// 				WorkingPopulation:      4500,
		// 				UnemployedPopulation:   380,
		// 				HousekeepingPopulation: 1550,
		// 				StudentPopulation:      1320,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Mentaras",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         3300,
		// 				FemalePopulation:       3400,
		// 				FamiliesNumber:         1650,
		// 				NumberOfBirth:          78,
		// 				NumberOfDeath:          27,
		// 				WorkingPopulation:      3900,
		// 				UnemployedPopulation:   310,
		// 				HousekeepingPopulation: 1300,
		// 				StudentPopulation:      1190,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Wadeng",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         4600,
		// 				FemalePopulation:       4700,
		// 				FamiliesNumber:         2300,
		// 				NumberOfBirth:          110,
		// 				NumberOfDeath:          40,
		// 				WorkingPopulation:      5400,
		// 				UnemployedPopulation:   450,
		// 				HousekeepingPopulation: 1850,
		// 				StudentPopulation:      1600,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Mojo",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         2900,
		// 				FemalePopulation:       2950,
		// 				FamiliesNumber:         1450,
		// 				NumberOfBirth:          65,
		// 				NumberOfDeath:          24,
		// 				WorkingPopulation:      3400,
		// 				UnemployedPopulation:   270,
		// 				HousekeepingPopulation: 1150,
		// 				StudentPopulation:      1030,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Sidomulyo",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         3700,
		// 				FemalePopulation:       3650,
		// 				FamiliesNumber:         1800,
		// 				NumberOfBirth:          82,
		// 				NumberOfDeath:          29,
		// 				WorkingPopulation:      4250,
		// 				UnemployedPopulation:   340,
		// 				HousekeepingPopulation: 1450,
		// 				StudentPopulation:      1310,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Kauman",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         2200,
		// 				FemalePopulation:       2300,
		// 				FamiliesNumber:         1100,
		// 				NumberOfBirth:          45,
		// 				NumberOfDeath:          15,
		// 				WorkingPopulation:      2600,
		// 				UnemployedPopulation:   180,
		// 				HousekeepingPopulation: 900,
		// 				StudentPopulation:      820,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Cerme",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         4100,
		// 				FemalePopulation:       4150,
		// 				FamiliesNumber:         2050,
		// 				NumberOfBirth:          92,
		// 				NumberOfDeath:          33,
		// 				WorkingPopulation:      4750,
		// 				UnemployedPopulation:   390,
		// 				HousekeepingPopulation: 1650,
		// 				StudentPopulation:      1460,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Petung",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         3000,
		// 				FemalePopulation:       3100,
		// 				FamiliesNumber:         1500,
		// 				NumberOfBirth:          70,
		// 				NumberOfDeath:          25,
		// 				WorkingPopulation:      3550,
		// 				UnemployedPopulation:   290,
		// 				HousekeepingPopulation: 1200,
		// 				StudentPopulation:      1060,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 			{
		// 				VillageName:            "Desa Bunder",
		// 				DemographyDataYear:     2026,
		// 				MalePopulation:         4400,
		// 				FemalePopulation:       4450,
		// 				FamiliesNumber:         2200,
		// 				NumberOfBirth:          100,
		// 				NumberOfDeath:          38,
		// 				WorkingPopulation:      5100,
		// 				UnemployedPopulation:   420,
		// 				HousekeepingPopulation: 1750,
		// 				StudentPopulation:      1580,
		// 				SourceName:             "BPS Kabupaten",
		// 				ExternalLinkSource:     "https://bps.go.id",
		// 			},
		// 		}

		// 		for _, d := range demographics {
		// 			if err := seed.CreateDemograhies(db, d); err != nil {
		// 				return err
		// 			}
		// 		}
		// 		return nil
		// 	},
		// },
		{
			Name: "CreateGalleriesSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateGalleries(db, models.CreateGalleries{
					Name:        "Galeri Sidayu",
					Description: "Deskripsi galeri",
					ImageID:     SavedImage.ID,
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
		// {
		// 	Name: "CreateExtraIndustriesBlogSample",
		// 	Run: func(db *gorm.DB) error {
		// 		industries := []models.CreateIndustriesBlogsReq{
		// 			{
		// 				Title:                      "Sentra Pengolahan Bandeng Asap Sidayu",
		// 				Content:                    "Mengolah hasil tambak segar menjadi produk olahan bandeng asap khas pesisir yang tahan lama.",
		// 				Location:                   "Golokan, Sidayu",
		// 				Rating:                     4.9,
		// 				Revenue:                    750000000.0,
		// 				ProducedProducts:           []string{"Bandeng Asap", "Otak-Otak Bandeng", "Nugget Bandeng"},
		// 				ProductionRatesPiecePerDay: 120,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2018,
		// 				EmployeesCount:             18,
		// 				BusinessType:               "UMKM",
		// 			},
		// 			{
		// 				Title:                      "Kerajinan Bambu & Mebel Tradisional",
		// 				Content:                    "Produksi perabot rumah tangga berbahan dasar bambu dan kayu pilihan dengan sentuhan estetik.",
		// 				Location:                   "Kertosono, Sidayu",
		// 				Rating:                     4.7,
		// 				Revenue:                    350000000.0,
		// 				ProducedProducts:           []string{"Kursi Bambu", "Bilik Rumah", "Dekorasi Dinding"},
		// 				ProductionRatesPiecePerDay: 15,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2015,
		// 				EmployeesCount:             12,
		// 				BusinessType:               "Perseorangan",
		// 			},
		// 			{
		// 				Title:                      "Konveksi Hijab & Busana Muslimah",
		// 				Content:                    "Produsen perlengkapan muslim yang memasarkan produknya hingga ke luar kota Gresik.",
		// 				Location:                   "Purwodadi, Sidayu",
		// 				Rating:                     4.6,
		// 				Revenue:                    600000000.0,
		// 				ProducedProducts:           []string{"Jilbab Instan", "Gamis", "Mukena"},
		// 				ProductionRatesPiecePerDay: 80,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2021,
		// 				EmployeesCount:             30,
		// 				BusinessType:               "CV",
		// 			},
		// 			{
		// 				Title:                      "Pabrik Kerupuk Ikan Laut Pesisir",
		// 				Content:                    "Memanfaatkan hasil tangkapan nelayan lokal untuk diolah menjadi kerupuk renyah gurih.",
		// 				Location:                   "Randuboto, Sidayu",
		// 				Rating:                     4.8,
		// 				Revenue:                    450000000.0,
		// 				ProducedProducts:           []string{"Kerupuk Udang", "Kerupuk Ikan Tenggiri", "Kerupuk Bawang"},
		// 				ProductionRatesPiecePerDay: 200,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2017,
		// 				EmployeesCount:             15,
		// 				BusinessType:               "UMKM",
		// 			},
		// 			{
		// 				Title:                      "Sentra Pembuatan Songkok Bordir",
		// 				Content:                    "Kerajinan kopiah atau songkok dengan hiasan bordir motif islami yang rapi dan elegan.",
		// 				Location:                   "Kauman, Sidayu",
		// 				Rating:                     4.9,
		// 				Revenue:                    850000000.0,
		// 				ProducedProducts:           []string{"Songkok Hitam", "Songkok Bordir Emas", "Kopiah Anak"},
		// 				ProductionRatesPiecePerDay: 90,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2012,
		// 				EmployeesCount:             40,
		// 				BusinessType:               "CV",
		// 			},
		// 			{
		// 				Title:                      "Industri Makanan Ringan Enting-Enting",
		// 				Content:                    "Penganan tradisional berbahan dasar gula dan kacang yang manis dan renyah.",
		// 				Location:                   "Ngaglik, Sidayu",
		// 				Rating:                     4.5,
		// 				Revenue:                    200000000.0,
		// 				ProducedProducts:           []string{"Enting-Enting Kacang", "Wajik Tradisional", "Rengginang"},
		// 				ProductionRatesPiecePerDay: 150,
		// 				ThumbnailID:                nil,
		// 				YearFounded:                2019,
		// 				EmployeesCount:             8,
		// 				BusinessType:               "UMKM",
		// 			},
		// 			{
		// 				Title:                      "Bengkel Las & Bubut Logam Mandiri",
		// 				Content:                    "Melayani pembuatan pagar besi, teralis, serta perbaikan komponen mesin pertanian dan tambak.",
		// 				Location:                   "Wadeng, Sidayu",
		// 				Rating:                     4.7,
		// 				Revenue:                    550000000.0,
		// 				ProducedProducts:           []string{"Pagar Besi", "Kincir Air Tambak", "Tralis Jendela"},
		// 				ProductionRatesPiecePerDay: 10,
		// 				ThumbnailID:                nil,
		// 				YearFounded:                2016,
		// 				EmployeesCount:             10,
		// 				BusinessType:               "Perseorangan",
		// 			},
		// 			{
		// 				Title:                      "Pengrajin Tas dan Dompet Kulit Sintetis",
		// 				Content:                    "Produksi aneka tas lokal dengan desain modern yang digemari kalangan muda.",
		// 				Location:                   "Lasem, Sidayu",
		// 				Rating:                     4.8,
		// 				Revenue:                    900000000.0,
		// 				ProducedProducts:           []string{"Tas Selempang", "Dompet Kulit", "Pouch Custom"},
		// 				ProductionRatesPiecePerDay: 60,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2022,
		// 				EmployeesCount:             35,
		// 				BusinessType:               "PT",
		// 			},
		// 			{
		// 				Title:                      "Sentra Pembuatan Bata Merah Berkualitas",
		// 				Content:                    "Penyedia material konstruksi bangunan yang kokoh dari tanah liat pilihan asli Sidayu.",
		// 				Location:                   "Mentaras, Sidayu",
		// 				Rating:                     4.4,
		// 				Revenue:                    300000000.0,
		// 				ProducedProducts:           []string{"Bata Merah Press", "Genteng Tradisional"},
		// 				ProductionRatesPiecePerDay: 500,
		// 				ThumbnailID:                nil,
		// 				YearFounded:                2010,
		// 				EmployeesCount:             14,
		// 				BusinessType:               "UMKM",
		// 			},
		// 			{
		// 				Title:                      "Industri Pengolahan Terasi Udang Asli",
		// 				Content:                    "Menghasilkan terasi berkualitas tinggi dengan aroma khas menggunakan udang rebon segar.",
		// 				Location:                   "Mojo, Sidayu",
		// 				Rating:                     4.9,
		// 				Revenue:                    400000000.0,
		// 				ProducedProducts:           []string{"Terasi Udang Super", "Petis Udang"},
		// 				ProductionRatesPiecePerDay: 100,
		// 				ThumbnailID:                nil,
		// 				YearFounded:                2014,
		// 				EmployeesCount:             9,
		// 				BusinessType:               "UMKM",
		// 			},
		// 			{
		// 				Title:                      "Konveksi Kaos Sablon & Merchandise",
		// 				Content:                    "Pusat pembuatan kaos event, seragam komunitas, dan atribut sablon digital maupun manual.",
		// 				Location:                   "Sidomulyo, Sidayu",
		// 				Rating:                     4.7,
		// 				Revenue:                    480000000.0,
		// 				ProducedProducts:           []string{"Kaos Sablon", "Topi", "Lanyard"},
		// 				ProductionRatesPiecePerDay: 70,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2021,
		// 				EmployeesCount:             11,
		// 				BusinessType:               "UMKM",
		// 			},
		// 			{
		// 				Title:                      "Budidaya & Pengolahan Madu Murni Lokal",
		// 				Content:                    "Peternakan lebah madu lokal di sekitar kawasan hijau Sidayu yang higienis.",
		// 				Location:                   "Cerme, Sidayu",
		// 				Rating:                     4.9,
		// 				Revenue:                    320000000.0,
		// 				ProducedProducts:           []string{"Madu Randu", "Madu Klanceng", "Propolis"},
		// 				ProductionRatesPiecePerDay: 30,
		// 				ThumbnailID:                nil,
		// 				YearFounded:                2023,
		// 				EmployeesCount:             6,
		// 				BusinessType:               "Perseorangan",
		// 			},
		// 			{
		// 				Title:                      "Sentra Pembuatan Sepatu & Sandal Lokal",
		// 				Content:                    "Produksi alas kaki kasual yang nyaman dan tahan lama untuk pasar domestik.",
		// 				Location:                   "Petung, Sidayu",
		// 				Rating:                     4.6,
		// 				Revenue:                    650000000.0,
		// 				ProducedProducts:           []string{"Sandal Kulit", "Sepatu Kasual", "Sandal Hotel"},
		// 				ProductionRatesPiecePerDay: 85,
		// 				ThumbnailID:                (*uuid.UUID)(&SavedImage.ID),
		// 				YearFounded:                2018,
		// 				EmployeesCount:             22,
		// 				BusinessType:               "CV",
		// 			},
		// 			{
		// 				Title:                      "Industri Kerajinan Anyaman Purun & Pandan",
		// 				Content:                    "Memanfaatkan tumbuhan rawa dan daun pandan menjadi tikar serta tas anyaman artistik.",
		// 				Location:                   "Bunder, Sidayu",
		// 				Rating:                     4.8,
		// 				Revenue:                    280000000.0,
		// 				ProducedProducts:           []string{"Tikar Anyaman", "Tas Belanja Purun", "Kotak Tisu"},
		// 				ProductionRatesPiecePerDay: 40,
		// 				ThumbnailID:                nil,
		// 				YearFounded:                2016,
		// 				EmployeesCount:             16,
		// 				BusinessType:               "UMKM",
		// 			},
		// 		}

		// 		for _, ind := range industries {
		// 			if err := seed.CreateIndustriesBlog(db, ind); err != nil {
		// 				return err
		// 			}
		// 		}
		// 		return nil
		// 	},
		// },
		{
			Name: "CreateOfficialsSample",
			Run: func(db *gorm.DB) error {
				return seed.CreateOfficial(db, models.CreateOfficials{
					Name:           "Bapak Sidayu",
					Description:    "Kepala kecamatan periode 2026-2031",
					RoleID:         SavedRole.ID,
					ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
				})
			},
		},
		// {
		// 	Name: "CreateExtraOfficialsSample",
		// 	Run: func(db *gorm.DB) error {
		// 		officials := []models.CreateOfficials{
		// 			{
		// 				Name:           "Bapak H. Ahmad Fauzi, S.Sos.",
		// 				Description:    "Camat Sidayu periode 2026-2031",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Ibu Dra. Hj. Siti Aminah",
		// 				Description:    "Sekretaris Kecamatan Sidayu",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Drs. Budi Santoso",
		// 				Description:    "Kepala Seksi Pemerintahan dan Ketentraman",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak H. Moh. Rifa'i, S.E.",
		// 				Description:    "Kepala Seksi Pembangunan dan Pemberdayaan Masyarakat",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Ibu Dewi Lestari, S.Pd.",
		// 				Description:    "Kepala Seksi Kesejahteraan Sosial",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Zainul Arifin, S.H.",
		// 				Description:    "Kepala Seksi Ketentraman dan Ketertiban Umum",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Suprayitno",
		// 				Description:    "Kepala Desa Sidayu",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Mukhtarom",
		// 				Description:    "Kepala Desa Golokan",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Abdul Malik",
		// 				Description:    "Kepala Desa Kertosono",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak H. Slamet Riyadi",
		// 				Description:    "Kepala Desa Purwodadi",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Nur Hidayat",
		// 				Description:    "Kepala Desa Randuboto",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak M. Cholil",
		// 				Description:    "Kepala Desa Ngaglik",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Drs. H. Makmun",
		// 				Description:    "Kepala Desa Lasem",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak Anang Syaifuddin",
		// 				Description:    "Kepala Desa Mentaras",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Name:           "Bapak H. Syaiful Anwar",
		// 				Description:    "Kepala Desa Wadeng",
		// 				RoleID:         SavedRole.ID,
		// 				ProfileImageID: (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 		}

		// 		for _, o := range officials {
		// 			if err := seed.CreateOfficial(db, o); err != nil {
		// 				return err
		// 			}
		// 		}
		// 		return nil
		// 	},
		// },
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
					ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
				})
			},
		},
		// {
		// 	Name: "CreateExtraShopsAndUmkmsBlogSample",
		// 	Run: func(db *gorm.DB) error {
		// 		shops := []models.CreateShopsAndUmkmsBlogsReq{
		// 			{
		// 				Title:                 "Warung Kopi & Makanan Ringan Mbok Min",
		// 				Content:               "Tempat nongkrong favorit warga lokal sambil menikmati kopi tubruk dan jajanan pasar.",
		// 				Location:              "Golokan, Sidayu",
		// 				Rating:                4.8,
		// 				Revenue:               8500000.0,
		// 				MarketedProducts:      []string{"Kopi Hitam", "Gorengan", "Kue Tradisional"},
		// 				SalesRatesPiecePerDay: 150,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Title:                 "Kios Buah Segar Nusantara",
		// 				Content:               "Menjual berbagai buah-buahan segar pilihan lokal maupun impor kualitas terbaik.",
		// 				Location:              "Kertosono, Sidayu",
		// 				Rating:                4.7,
		// 				Revenue:               12000000.0,
		// 				MarketedProducts:      []string{"Jeruk Lokal", "Pisang", "Semangka"},
		// 				SalesRatesPiecePerDay: 90,
		// 				ThumbnailID:           nil,
		// 			},
		// 			{
		// 				Title:                 "Depot Air Minum & Gas Amanah",
		// 				Content:               "Layanan isi ulang air galon dan penyediaan gas elpiji cepat antar ke rumah warga.",
		// 				Location:              "Purwodadi, Sidayu",
		// 				Rating:                4.9,
		// 				Revenue:               10000000.0,
		// 				MarketedProducts:      []string{"Air Galon", "Gas 3kg", "Gas 12kg"},
		// 				SalesRatesPiecePerDay: 75,
		// 				ThumbnailID:           nil,
		// 			},
		// 			{
		// 				Title:                 "Toko Kelontong & Pakan Ternak Sumber Rejeki",
		// 				Content:               "Menyediakan kebutuhan pokok harian serta lengkap dengan pakan ikan dan unggas.",
		// 				Location:              "Randuboto, Sidayu",
		// 				Rating:                4.6,
		// 				Revenue:               18000000.0,
		// 				MarketedProducts:      []string{"Beras", "Pakan Ikan", "Pupuk Tanaman"},
		// 				SalesRatesPiecePerDay: 110,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Title:                 "Pusat Oleh-Oleh & Jajanan Pasar Sidayu",
		// 				Content:               "Surga kuliner bagi pelintas yang mencari camilan khas tradisional asli buatan warga lokal.",
		// 				Location:              "Kauman, Sidayu",
		// 				Rating:                4.9,
		// 				Revenue:               22000000.0,
		// 				MarketedProducts:      []string{"Kerupuk Udang", "Emping", "Kue Kering"},
		// 				SalesRatesPiecePerDay: 200,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Title:                 "Warung Sembako & Ritel Mandiri",
		// 				Content:               "Menyediakan kebutuhan dapur sehari-hari dengan harga bersahabat dan pelayanan ramah.",
		// 				Location:              "Ngaglik, Sidayu",
		// 				Rating:                4.5,
		// 				Revenue:               9500000.0,
		// 				MarketedProducts:      []string{"Minyak Goreng", "Gula Pasir", "Telur Ayam"},
		// 				SalesRatesPiecePerDay: 85,
		// 				ThumbnailID:           nil,
		// 			},
		// 			{
		// 				Title:                 "Toko Bangunan & Alat Listrik Barokah",
		// 				Content:               "Pusat material bangunan kecil dan perlengkapan instalasi listrik untuk renovasi rumah.",
		// 				Location:              "Wadeng, Sidayu",
		// 				Rating:                4.8,
		// 				Revenue:               25000000.0,
		// 				MarketedProducts:      []string{"Semen", "Kabel Listrik", "Pipa PVC"},
		// 				SalesRatesPiecePerDay: 40,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Title:                 "Kios Bibit Tanaman & Perlengkapan Tani",
		// 				Content:               "Menjual berbagai jenis bibit unggul tanaman holtikultura dan obat-obatan pertanian.",
		// 				Location:              "Lasem, Sidayu",
		// 				Rating:                4.7,
		// 				Revenue:               11000000.0,
		// 				MarketedProducts:      []string{"Bibit Cabai", "Pupuk Organik", "Pestisida"},
		// 				SalesRatesPiecePerDay: 60,
		// 				ThumbnailID:           nil,
		// 			},
		// 			{
		// 				Title:                 "Minimarket Lokal Sidayu Mart",
		// 				Content:               "Swalayan modern berukuran mini yang menyediakan produk higienis dan ber-AC.",
		// 				Location:              "Mentaras, Sidayu",
		// 				Rating:                4.8,
		// 				Revenue:               30000000.0,
		// 				MarketedProducts:      []string{"Makanan Ringan", "Produk Susu", "Keperluan Mandi"},
		// 				SalesRatesPiecePerDay: 250,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Title:                 "Warung Pecel Lele & Ayam Goreng Cak Mat",
		// 				Content:               "Kuliner malam hari yang ramai dikunjungi pembeli karena sambalnya yang khas.",
		// 				Location:              "Mojo, Sidayu",
		// 				Rating:                4.9,
		// 				Revenue:               14000000.0,
		// 				MarketedProducts:      []string{"Pecel Lele", "Ayam Goreng", "Es Teh Manis"},
		// 				SalesRatesPiecePerDay: 130,
		// 				ThumbnailID:           nil,
		// 			},
		// 			{
		// 				Title:                 "Toko Buku & Perlengkapan Sekolah Pelajar",
		// 				Content:               "Menyediakan buku tulis, seragam, dan peralatan sekolah lengkap untuk anak-anak.",
		// 				Location:              "Sidomulyo, Sidayu",
		// 				Rating:                4.6,
		// 				Revenue:               13500000.0,
		// 				MarketedProducts:      []string{"Buku Tulis", "Pulpen", "Seragam Sekolah"},
		// 				SalesRatesPiecePerDay: 95,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Title:                 "Kios Herbal & Jamu Tradisional",
		// 				Content:               "Menjual racikan jamu tradisional dan bahan rempah-rempah alami untuk kesehatan.",
		// 				Location:              "Cerme, Sidayu",
		// 				Rating:                4.7,
		// 				Revenue:               7000000.0,
		// 				MarketedProducts:      []string{"Jamu Beras Kencur", "Kunyit Asam", "Rempah Kering"},
		// 				SalesRatesPiecePerDay: 50,
		// 				ThumbnailID:           nil,
		// 			},
		// 			{
		// 				Title:                 "Depot Roti & Kue Basah Melati",
		// 				Content:               "Produsen sekaligus penjual aneka roti lembut dan jajanan pasar segar setiap pagi.",
		// 				Location:              "Petung, Sidayu",
		// 				Rating:                4.8,
		// 				Revenue:               16000000.0,
		// 				MarketedProducts:      []string{"Roti Manis", "Kue Lemper", "Onde-Onde"},
		// 				SalesRatesPiecePerDay: 180,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 			{
		// 				Title:                 "Warung Sate Ayam & Kambing Madura",
		// 				Content:               "Sate dengan bumbu kacang gurih khas Madura yang menggugah selera di pinggir jalan raya.",
		// 				Location:              "Bunder, Sidayu",
		// 				Rating:                4.9,
		// 				Revenue:               20000000.0,
		// 				MarketedProducts:      []string{"Sate Ayam", "Sate Kambing", "Lontong"},
		// 				SalesRatesPiecePerDay: 140,
		// 				ThumbnailID:           (*uuid.UUID)(&SavedImage.ID),
		// 			},
		// 		}

		// 		for _, s := range shops {
		// 			if err := seed.CreateShopsAndUmkmsBlog(db, s); err != nil {
		// 				return err
		// 			}
		// 		}
		// 		return nil
		// 	},
		// },
	}
}
