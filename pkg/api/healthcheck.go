package api

// Healthcheck is a healt check response structure.
type Healthcheck struct {
	Version  string  `json:"version" example:"1"`
	Name     string  `json:"name" example:"Сервис получения информации по разделу"`
	RootPath string  `json:"root_path" example:"section-info"`
	DBAddr   *string `json:"DB_addr" example:"192.168.8.250"`
	DBName   *string `json:"DB_name" example:"dbqueue_korenovsk_actual"`
	DBTime   *string `json:"DB_time" example:"2020-11-23 15:47:00.900"`
}
