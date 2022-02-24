package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/*
*	Creates output structs for Uniinfo
 */
func GetUniInfo(URL string) []UniInfo {
	var Universities []UniInfo
	var InUni []InUniversity = GetUnis(URL)

	//finds the country the university resides in and merges necessary info
	for _, uni := range InUni {
		countrySearch := countryURL + "alpha?codes=" + uni.Isocode
		Universities = append(Universities, mergeStructs(uni, GetCountry(countrySearch)[0]))
	}

	return Universities
}

/*
*	Creates output structs for Neighbourunis
 */
func GetNeighbourUnis(URL string, uniExtra string, limit int) []UniInfo {
	var Universities []UniInfo
	var InUni []InUniversity

	//finds the country that is searched for
	var InCoun []InCountry = GetCountry(URL)

	//checks if searched country has neighbours and finds them as well
	if len(InCoun[0].Borders) > 0 {
		neighbourSearch := countryURL + "alpha?codes="
		for _, code := range InCoun[0].Borders {
			neighbourSearch += code + ","
		}
		InCoun = append(InCoun, GetCountry(neighbourSearch)...)
	}

	//goes through all the countries to find the universities in these countries
	for _, country := range InCoun {

		//finds the universities according to the country and searched uni-name
		uniSearch := uniURL + "search?country=" + country.Name.CommonName
		if uniExtra != "" {
			uniSearch += "&name=" + uniExtra
		}

		//saves the universities that matches the search
		InUni = GetUnis(uniSearch)

		//checks if there is a limit -> if not the limit is set to be amount of universities
		if limit == 0 {
			limit = len(InUni)
		}

		//merges the country and the university
		for i := 0; i < limit && i < len(InUni); i++ {
			if i < len(InUni) && InUni[i].Country == country.Name.CommonName {
				Universities = append(Universities, mergeStructs(InUni[i], country))
			}
		}
	}

	return Universities
}

/*
*	Gets the Universities from the universities api
 */
func GetUnis(url string) []InUniversity {
	write, _ := http.Get(url)
	body, _ := ioutil.ReadAll(write.Body)

	var Unis []InUniversity

	//unmrshals the content from university api
	mashErr := json.Unmarshal(body, &Unis)
	if mashErr != nil {
		log.Fatal("Error when marshalling:", mashErr)
	}

	return Unis
}

/*
*	Gets the Countries from the countries api
 */
func GetCountry(url string) []InCountry {
	write, _ := http.Get(url)
	body, _ := ioutil.ReadAll(write.Body)

	var Countries []InCountry

	//unmrshals the content from country api
	mashErr := json.Unmarshal(body, &Countries)
	if mashErr != nil {
		log.Fatal("Error when marshalling:", mashErr)
	}

	return Countries
}

/*
*	Merges the university struct and country struct into the output struct
 */
func mergeStructs(uni InUniversity, country InCountry) UniInfo {
	var University = UniInfo{
		Name:      uni.Name,
		Country:   uni.Country,
		Isocode:   uni.Isocode,
		Webpages:  uni.Webpage,
		Languages: country.Languages,
		Map:       country.Map,
	}
	return University
}
