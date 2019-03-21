package workable

import (
	"time"
)

type Job struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	FullTitle      string    `json:"full_title"`
	Shortcode      string    `json:"shortcode"`
	Code           string    `json:"code"`
	State          string    `json:"state"`
	Department     string    `json:"department"`
	URL            string    `json:"url"`
	ApplicationURL string    `json:"application_url"`
	Shortlink      string    `json:"shortlink"`
	Location       Location  `json:"location"`
	CreatedAt      time.Time `json:"created_at"`
}

type Location struct {
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
	Region        string `json:"region"`
	RegionCode    string `json:"region_code"`
	City          string `json:"city"`
	ZipCode       string `json:"zip_code"`
	Telecommuting bool   `json:"telecommuting"`
}

type Jobs struct {
	Jobs   []Job  `json:"jobs"`
	Paging Paging `json:"paging"`
}

type Paging struct {
	Next string `json:"next"`
}
