package handler

import "github.com/ryan-albuquerque/gopportunities-api/internal/models"

type OpeningResponse struct {
	ID        uint   `json:"id"`
	Role      string `json:"role"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	Link      string `json:"link"`
	Salary    int64  `json:"salary"`
	Remote    bool   `json:"remote"`
	IsActived bool   `json:"is_actived"`
	CreatedAt string `json:"created_at"`
}

func (o *OpeningResponse) FromModel(opening models.Opening) {
	o.ID = opening.ID
	o.Role = opening.Role
	o.Company = opening.Company
	o.Location = opening.Location
	o.Link = opening.Link
	o.Salary = opening.Salary
	o.Remote = opening.Remote
	o.IsActived = opening.IsActived
	o.CreatedAt = opening.CreatedAt.Format("2006-01-02 15:04:05")
}

func ParseListOpeningResponse(openings []models.Opening) []OpeningResponse {
	var openingsResponse []OpeningResponse
	for _, opening := range openings {
		openingResponse := OpeningResponse{}
		openingResponse.FromModel(opening)
		openingsResponse = append(openingsResponse, openingResponse)
	}
	return openingsResponse
}
