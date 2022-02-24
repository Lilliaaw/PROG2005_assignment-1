package handler

/**
*	Struct for the input info about country from countries api
 */
type InCountry struct {
	Name        CountryName       `json:"name"`
	Languages   map[string]string `json:"languages"`
	Map         MapTypes          `json:"maps"`
	Borders     []string          `json:"borders"`
	CountryCode string            `json:"cca3"`
}

/**
*	Sub-struct for InCountry to get the common name for a country
 */
type CountryName struct {
	CommonName string `json:"common"`
}

/**
*	Sub-struct for InCountry to get the openStreetMap for a country
 */
type MapTypes struct {
	OpenStreetMap string `json:"openStreetMaps"`
}

/**
*	Struct for the input info about university from the universities api
 */
type InUniversity struct {
	Name    string   `json:"name"`
	Country string   `json:"country,omitempty"`
	Isocode string   `json:"alpha_two_code"`
	Webpage []string `json:"web_pages"`
}

/**
*	Struct for the output that combines the InCountry and InUniversity information
 */
type UniInfo struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       MapTypes          `json:"map"`
}

/**
*	Struct for the diagnostics output
 */
type Diagnostic struct {
	Statusuniapi     int    `json:"universitiesapi"`
	Statuscountryapi int    `json:"countriesapi"`
	Version          string `json:"version"`
	Duration         string `json:"uptime"`
}
