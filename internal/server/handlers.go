package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Hello epta
func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, request: %+v", r)
	fmt.Println("response")
}

func (s *Server) getLocationURLByWasteType(w http.ResponseWriter, r *http.Request) {
	// _, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	// defer func() {
	// 	// log.Println("getLocationURLByWasteType: canceling context")
	// 	cancel()
	// }()
	// vars := mux.Vars(r)
	// typeId := vars["type_id"]
	// lat := r.FormValue("latitude")
	// lon := r.FormValue("longitude")
	// cityId := getNearPoint(s, typeId, lat, lon, s.GoogleApiKey)
	//
	// requestUrl := fmt.Sprintf("https://recyclemap.ru/index.php?option=com_greenmarkers&city=%d", cityId)
	// err := json.NewEncoder(w).Encode(&requestUrl)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

}

func (s *Server) getLocationPointByWasteType(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		typeId = params["type_id"]
		lat    = params["latitude"]
		lon    = params["longitude"]
	)

	log.Print(typeId, lat, lon)

	fmt.Fprint(w, "hi")

	return
}

//
// func findCity(s *Server, city string) int {
// 	cities := s.st.GetCities()
// 	for _, v := range cities {
// 		if v.Name == city {
// 			return v.Id
// 		}
//
// 	}
// 	return -1
// }
//
// func getNearPoint(s *Server, typeId string, lat string, lon string, gApi string) int {
// 	city := getCity(lat, lon, gApi)
// 	cityId := findCity(s, city)
// 	return cityId
// }
//
// type googleGeocodeResponse struct {
// 	Results []struct {
// 		AddressComponents []struct {
// 			LongName string   `json:"long_name"`
// 			Types    []string `json:"types"`
// 		} `json:"address_components"`
// 	} `json:"results"`
// }
//
// func getCity(lat string, lon string, gApi string) string {
// 	latF, _ := strconv.ParseFloat(lon, 64)
// 	lonF, _ := strconv.ParseFloat(lon, 64)
//
// 	fmt.Println(latF)
// 	fmt.Println(lonF)
//
// 	p := geo.NewPoint(latF, lonF)
// 	geo.SetGoogleAPIKey(gApi)
// 	fmt.Println("key:", gApi)
// 	geocoder := new(geo.GoogleGeocoder)
//
// 	geo.HandleWithSQL()
//
// 	data, err := geocoder.Request(fmt.Sprintf("latlng=%f,%f&key=%s", p.Lat(), p.Lng(), gApi))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var res googleGeocodeResponse
// 	if err := json.Unmarshal(data, &res); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(data))
// 	var city string
// 	fmt.Println(res.Results)
// 	if len(res.Results) > 0 {
// 		r := res.Results[0]
// 	outer:
// 		for _, comp := range r.AddressComponents {
// 			// See https://developers.google.com/maps/documentation/geocoding/#Types
// 			// for address types
// 			for _, compType := range comp.Types {
// 				if compType == "locality" {
// 					city = comp.LongName
// 					break outer
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(city)
// 	return city
// }
