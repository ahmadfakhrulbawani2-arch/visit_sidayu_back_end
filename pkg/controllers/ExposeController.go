package controllers

import (
	"net/http"
	"visit-sidayu-backend/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type EndpointDoc struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	Protected   bool     `json:"protected"`
	QueryParams []string `json:"query_params,omitempty"`
	Description string   `json:"description,omitempty"`
}

type ResourceDoc struct {
	Name        string        `json:"name"`
	ContentType string        `json:"content_type"`
	Endpoints   []EndpointDoc `json:"endpoints"`
}

type ApiDocumentation struct {
	BaseURL  string        `json:"base_url"`
	Public   []EndpointDoc `json:"public"`
	Resources []ResourceDoc `json:"resources"`
}

func ExposeHandler(ctx *gin.Context) {

	doc := ApiDocumentation{
		BaseURL: "/api/v1",

		Public: []EndpointDoc{
			{
				Method:      "GET",
				Path:        "/ping",
				Description: "Health check",
			},
			{
				Method:      "POST",
				Path:        "/superadmins/auth/login",
				Description: "Login as superadmin",
			},
			{
				Method:      "POST",
				Path:        "/superadmins/auth/register",
				Description: "Register superadmin",
			},
		},

		Resources: []ResourceDoc{

			{
				Name:        "Superadmins",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/superadmins/", Protected: true, Description: "Get all superadmins"},
					{Method: "GET", Path: "/superadmins/me", Protected: true, Description: "Current logged in superadmin"},
					{Method: "GET", Path: "/superadmins/id/:id", Protected: true, Description: "Get superadmin by ID"},
					{Method: "PATCH", Path: "/superadmins/id/:id", Protected: true, Description: "Update superadmin"},
					{Method: "DELETE", Path: "/superadmins/id/:id", Protected: true, Description: "Delete superadmin"},
				},
			},

			{
				Name:        "Images",
				ContentType: "multipart/form-data",
				Endpoints: []EndpointDoc{
					{
						Method:      "GET",
						Path:        "/images",
						QueryParams: []string{"search", "page", "limit"},
						Description: "Get all images",
					},
					{
						Method:      "GET",
						Path:        "/images/id/:id",
						Description: "Get image by ID",
					},
					{
						Method:      "POST",
						Path:        "/images",
						Protected:   true,
						Description: "Upload image",
					},
					{
						Method:      "PUT",
						Path:        "/images/id/:id",
						Protected:   true,
						Description: "Update image",
					},
					{
						Method:      "DELETE",
						Path:        "/images/id/:id",
						Protected:   true,
						Description: "Delete image",
					},
				},
			},

			{
				Name:        "Blogs",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{
						Method:      "GET",
						Path:        "/blogs",
						QueryParams: []string{"search", "page", "limit"},
						Description: "Get all blogs",
					},
					{Method: "GET", Path: "/blogs/id/:id", Description: "Get blog by ID"},
					{Method: "GET", Path: "/blogs/slug/:slug", Description: "Get blog by slug"},
					{Method: "POST", Path: "/blogs", Protected: true, Description: "Create blog"},
					{Method: "PATCH", Path: "/blogs/id/:id", Protected: true, Description: "Update blog"},
					{Method: "DELETE", Path: "/blogs/id/:id", Protected: true, Description: "Delete blog"},
				},
			},

			{
				Name:        "Culture Blogs",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/culture-blogs", QueryParams: []string{"search", "page", "limit"}, Description: "Get all culture blogs"},
					{Method: "GET", Path: "/culture-blogs/id/:id", Description: "Get culture blog by ID"},
					{Method: "GET", Path: "/culture-blogs/slug/:slug", Description: "Get culture blog by slug"},
					{Method: "POST", Path: "/culture-blogs", Protected: true, Description: "Create culture blog"},
					{Method: "PATCH", Path: "/culture-blogs/id/:id", Protected: true, Description: "Update culture blog"},
					{Method: "DELETE", Path: "/culture-blogs/id/:id", Protected: true, Description: "Delete culture blog"},
				},
			},

			{
				Name:        "Demographies",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/demographies", QueryParams: []string{"search", "page", "limit"}, Description: "Get all demographies"},
					{Method: "GET", Path: "/demographies/id/:id", Description: "Get demography by ID"},
					{Method: "GET", Path: "/demographies/district", Description: "Get districts"},
					{Method: "POST", Path: "/demographies", Protected: true},
					{Method: "PATCH", Path: "/demographies/id/:id", Protected: true},
					{Method: "DELETE", Path: "/demographies/id/:id", Protected: true},
				},
			},

			{
				Name:        "Galleries",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/galleries", QueryParams: []string{"search", "page", "limit"}},
					{Method: "GET", Path: "/galleries/id/:id"},
					{Method: "GET", Path: "/galleries/slug/:slug"},
					{Method: "POST", Path: "/galleries", Protected: true},
					{Method: "PATCH", Path: "/galleries/id/:id", Protected: true},
					{Method: "DELETE", Path: "/galleries/id/:id", Protected: true},
				},
			},

			{
				Name:        "Geographies",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/geographies", QueryParams: []string{"search", "page", "limit"}},
					{Method: "GET", Path: "/geographies/id/:id"},
					{Method: "GET", Path: "/geographies/district"},
					{Method: "POST", Path: "/geographies", Protected: true},
					{Method: "PATCH", Path: "/geographies/id/:id", Protected: true},
					{Method: "DELETE", Path: "/geographies/id/:id", Protected: true},
				},
			},

			{
				Name:        "Industries Blogs",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/industries-blogs", QueryParams: []string{"search", "page", "limit"}},
					{Method: "GET", Path: "/industries-blogs/id/:id"},
					{Method: "GET", Path: "/industries-blogs/slug/:slug"},
					{Method: "POST", Path: "/industries-blogs", Protected: true},
					{Method: "PATCH", Path: "/industries-blogs/id/:id", Protected: true},
					{Method: "DELETE", Path: "/industries-blogs/id/:id", Protected: true},
				},
			},

			{
				Name:        "Officials",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/officials", QueryParams: []string{"search", "page", "limit"}},
					{Method: "GET", Path: "/officials/id/:id"},
					{Method: "POST", Path: "/officials", Protected: true},
					{Method: "PATCH", Path: "/officials/id/:id", Protected: true},
					{Method: "DELETE", Path: "/officials/id/:id", Protected: true},
				},
			},

			{
				Name:        "Shops & UMKMs",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/shops-umkms", QueryParams: []string{"search", "page", "limit"}},
					{Method: "GET", Path: "/shops-umkms/id/:id"},
					{Method: "GET", Path: "/shops-umkms/slug/:slug"},
					{Method: "POST", Path: "/shops-umkms", Protected: true},
					{Method: "PATCH", Path: "/shops-umkms/id/:id", Protected: true},
					{Method: "DELETE", Path: "/shops-umkms/id/:id", Protected: true},
				},
			},

			{
				Name:        "Timelines",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/timelines", QueryParams: []string{"search", "page", "limit"}},
					{Method: "GET", Path: "/timelines/id/:id"},
					{Method: "PATCH", Path: "/timelines/id/:id", Protected: true},
					{Method: "DELETE", Path: "/timelines/id/:id", Protected: true},
				},
			},

			{
				Name:        "Roles",
				ContentType: "application/json",
				Endpoints: []EndpointDoc{
					{Method: "GET", Path: "/roles", QueryParams: []string{"search", "page", "limit"}},
					{Method: "GET", Path: "/roles/id/:id"},
					{Method: "POST", Path: "/roles", Protected: true},
					{Method: "PATCH", Path: "/roles/id/:id", Protected: true},
					{Method: "DELETE", Path: "/roles/id/:id", Protected: true},
				},
			},
		},
	}

	helpers.RespSuccess(
		ctx,
		http.StatusOK,
		"Data are the list of all available endpoints",
		doc,
		"",
		nil,
	)
}
