package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SubAreaSearchResponse struct {
	ErrorCode struct {
		ID          int    `json:"Id"`
		Description string `json:"Description"`
		ProductName string `json:"ProductName"`
		Version     string `json:"Version"`
	} `json:"ErrorCode"`
	SubArea []struct {
		Location       string  `json:"Location"`
		Longitude      string  `json:"Longitude"`
		Latitude       string  `json:"Latitude"`
		Height         float64 `json:"Height"`
		Width          float64 `json:"Width"`
		InternalResult string  `json:"InternalResult"`
		Viewport       struct {
			NorthEast struct {
				Latitude  string `json:"Latitude"`
				Longitude string `json:"Longitude"`
			} `json:"NorthEast"`
			SouthWest struct {
				Latitude  string `json:"Latitude"`
				Longitude string `json:"Longitude"`
			} `json:"SouthWest"`
		} `json:"Viewport"`
		GEOID    string        `json:"GEOId"`
		Polygons []interface{} `json:"Polygons"`
	} `json:"SubArea"`
}

type ListingResult struct {
	ErrorCode struct {
		ID          int    `json:"Id"`
		Description string `json:"Description"`
		Status      string `json:"Status"`
		ProductName string `json:"ProductName"`
		Version     string `json:"Version"`
	} `json:"ErrorCode"`
	Paging struct {
		RecordsPerPage int `json:"RecordsPerPage"`
		CurrentPage    int `json:"CurrentPage"`
		TotalRecords   int `json:"TotalRecords"`
		MaxRecords     int `json:"MaxRecords"`
		TotalPages     int `json:"TotalPages"`
		RecordsShowing int `json:"RecordsShowing"`
		Pins           int `json:"Pins"`
	} `json:"Paging"`
	Results []struct {
		ID            string `json:"Id"`
		MlsNumber     string `json:"MlsNumber"`
		PublicRemarks string `json:"PublicRemarks"`
		Building      struct {
			BathroomTotal string `json:"BathroomTotal"`
			Bedrooms      string `json:"Bedrooms"`
			SizeInterior  string `json:"SizeInterior"`
			StoriesTotal  string `json:"StoriesTotal"`
			Type          string `json:"Type"`
		} `json:"Building"`
		Individual []struct {
			IndividualID int    `json:"IndividualID"`
			Name         string `json:"Name"`
			Organization struct {
				OrganizationID int    `json:"OrganizationID"`
				Name           string `json:"Name"`
				Logo           string `json:"Logo"`
				Address        struct {
					AddressText       string      `json:"AddressText"`
					PermitShowAddress bool        `json:"PermitShowAddress"`
					DisseminationArea interface{} `json:"DisseminationArea"`
				} `json:"Address"`
				Phones []struct {
					PhoneType   string `json:"PhoneType"`
					PhoneNumber string `json:"PhoneNumber"`
					AreaCode    string `json:"AreaCode"`
					PhoneTypeID string `json:"PhoneTypeId"`
				} `json:"Phones"`
				Emails []struct {
					ContactID string `json:"ContactId"`
				} `json:"Emails"`
				Websites []struct {
					Website       string `json:"Website"`
					WebsiteTypeID string `json:"WebsiteTypeId"`
				} `json:"Websites"`
				OrganizationType      string `json:"OrganizationType"`
				Designation           string `json:"Designation"`
				HasEmail              bool   `json:"HasEmail"`
				PermitFreetextEmail   bool   `json:"PermitFreetextEmail"`
				PermitShowListingLink bool   `json:"PermitShowListingLink"`
				RelativeDetailsURL    string `json:"RelativeDetailsURL"`
				PhotoLastupdate       string `json:"PhotoLastupdate"`
			} `json:"Organization"`
			Phones []struct {
				PhoneType   string `json:"PhoneType"`
				PhoneNumber string `json:"PhoneNumber"`
				AreaCode    string `json:"AreaCode"`
				PhoneTypeID string `json:"PhoneTypeId"`
			} `json:"Phones"`
			Websites []struct {
				Website       string `json:"Website"`
				WebsiteTypeID string `json:"WebsiteTypeId"`
			} `json:"Websites"`
			Emails []struct {
				ContactID string `json:"ContactId"`
			} `json:"Emails"`
			Photo                    string `json:"Photo"`
			Position                 string `json:"Position"`
			PermitFreetextEmail      bool   `json:"PermitFreetextEmail"`
			FirstName                string `json:"FirstName"`
			LastName                 string `json:"LastName"`
			CorporationDisplayTypeID string `json:"CorporationDisplayTypeId"`
			PermitShowListingLink    bool   `json:"PermitShowListingLink"`
			RelativeDetailsURL       string `json:"RelativeDetailsURL"`
			AgentPhotoLastUpdated    string `json:"AgentPhotoLastUpdated"`
			PhotoHighRes             string `json:"PhotoHighRes"`
			RankMyAgentKey           string `json:"RankMyAgentKey"`
			RealSatisfiedKey         string `json:"RealSatisfiedKey"`
		} `json:"Individual"`
		Property struct {
			Price   string `json:"Price"`
			Type    string `json:"Type"`
			Address struct {
				AddressText       string      `json:"AddressText"`
				Longitude         string      `json:"Longitude"`
				Latitude          string      `json:"Latitude"`
				PermitShowAddress bool        `json:"PermitShowAddress"`
				DisseminationArea interface{} `json:"DisseminationArea"`
			} `json:"Address"`
			Photo []struct {
				SequenceID  string `json:"SequenceId"`
				HighResPath string `json:"HighResPath"`
				MedResPath  string `json:"MedResPath"`
				LowResPath  string `json:"LowResPath"`
				Description string `json:"Description"`
				LastUpdated string `json:"LastUpdated"`
			} `json:"Photo"`
			TypeID                string `json:"TypeId"`
			FarmType              string `json:"FarmType"`
			OwnershipType         string `json:"OwnershipType"`
			ZoningType            string `json:"ZoningType"`
			AmmenitiesNearBy      string `json:"AmmenitiesNearBy"`
			ConvertedPrice        string `json:"ConvertedPrice"`
			OwnershipTypeGroupIds []int  `json:"OwnershipTypeGroupIds"`
			PriceUnformattedValue string `json:"PriceUnformattedValue"`
		} `json:"Property"`
		Business struct {
		} `json:"Business"`
		Land struct {
		} `json:"Land"`
		AlternateURL struct {
			DetailsLink string `json:"DetailsLink"`
		} `json:"AlternateURL,omitempty"`
		PostalCode         string `json:"PostalCode"`
		RelativeDetailsURL string `json:"RelativeDetailsURL"`
		StatusID           string `json:"StatusId"`
		PhotoChangeDateUTC string `json:"PhotoChangeDateUTC"`
		Distance           string `json:"Distance"`
		RelativeURLEn      string `json:"RelativeURLEn"`
		RelativeURLFr      string `json:"RelativeURLFr"`
		Media              []struct {
			MediaCategoryID  string `json:"MediaCategoryId"`
			MediaCategoryURL string `json:"MediaCategoryURL"`
			Order            int    `json:"Order"`
		} `json:"Media"`
		PriceChangeDateUTC string `json:"PriceChangeDateUTC,omitempty"`
		HasNewImageUpdate  bool   `json:"HasNewImageUpdate,omitempty"`
	} `json:"Results"`
	Pins          []interface{} `json:"Pins"`
	GroupingLevel string        `json:"GroupingLevel"`
}

func getSubAreaSearchResponse(location string) SubAreaSearchResponse{

    req, _ := http.NewRequest("GET", fmt.Sprintf("https://api2.realtor.ca/Location.svc/SubAreaSearch?Area=%s&ApplicationId=1&CultureId=1&CurrentPage=1", location), nil)
    req.Header.Add("referer", "https://www.realtor.ca/");
	req.Header.Add("accept-language", "en-US,en;q=0.9");
	req.Header.Add("origin", "https://www.realtor.ca");
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8");
	req.Header.Add("accept", "*/*");
	req.Header.Add("authority", "api2.realtor.ca");
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
	fmt.Println(res.Status)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    var subAreaSearchResponse SubAreaSearchResponse
    err = json.Unmarshal(body, &subAreaSearchResponse)
    if err != nil {
        panic(err)
    }

    return subAreaSearchResponse
}

func getListings(geoId string) ListingResult{

    data := url.Values{}
    data.Set("GeoIds", geoId);
    data.Set("PropertyTypeGroupID", "1");
    data.Set("PropertySearchTypeId", "1");
    data.Set("TransactionTypeId", "1");
    data.Set("RecordsPerPage", "12");
    data.Set("ApplicationId", "1");
    data.Set("CultureId", "1");
    data.Set("CurrentPage", "1");

    req, _ := http.NewRequest("POST", "https://api2.realtor.ca/Listing.svc/PropertySearch_Post", strings.NewReader(data.Encode()))
	req.Header.Add("referer", "https://www.realtor.ca/");
	req.Header.Add("accept-language", "en-US,en;q=0.9");
	req.Header.Add("origin", "https://www.realtor.ca");
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8");
	req.Header.Add("accept", "*/*");
	req.Header.Add("authority", "api2.realtor.ca");
	

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    var listingResult ListingResult
    err = json.Unmarshal(body, &listingResult)
    if err != nil {
        panic(err)
    }

	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected Status", res.Status)
	} else {
		fmt.Println("Success.", res.Status)
	}

    return listingResult
}

func GetRealtorData(location string) ListingResult{
	locationData := getSubAreaSearchResponse(url.QueryEscape(location))
	return getListings(locationData.SubArea[0].GEOID)
}