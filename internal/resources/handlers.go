package resources

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/ola-maps/internal"
	"github.com/ola-maps/internal/app"
)

// Get directions
func (o *OLAMap) GetDirections(origin, destination string) (interface{}, error) {
	if origin == "" || destination == "" {
		return nil, errors.New("Missing required query parameters: 'origin' and/or 'destination'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	url := fmt.Sprintf(app.DirectionsURL, origin, destination)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make external request
	err := internal.MakeExternalRequest("POST", url, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// PlaceAutoComplete
func (o *OLAMap) PlaceAutoComplete(input string) (interface{}, error) {
	if input == "" {
		return nil, errors.New("Missing required query parameters: 'input'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf(app.PlaceAutoCompleteURL, input)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", url, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GeoCode
func (o *OLAMap) GeoCode(address, bounds, language string) (interface{}, error) {
	if address == "" {
		return nil, errors.New("Missing required query parameters: 'address'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf(app.GeoCodeURL, address, bounds, language)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", url, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// ReverseGeocode
func (o *OLAMap) ReverseGeocode(latlng string) (interface{}, error) {
	if latlng == "" {
		return nil, errors.New("Missing required query parameters: 'latlng'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the API request
	urlWithParams := fmt.Sprintf(app.ReverseGeocodeURL, url.QueryEscape(latlng))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", urlWithParams, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetDistanceMatrix
func (o *OLAMap) GetDistanceMatrix(origins, destinations string) (interface{}, error) {
	if origins == "" || destinations == "" {
		return nil, errors.New("Missing required query parameters: 'origin' and/or 'destination'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf(app.DistanceMatrixURL,
		url.QueryEscape(origins), url.QueryEscape(destinations))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", url, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// ArrayOfData
func (o *OLAMap) ArrayOfData(datasetName string) (interface{}, error) {
	if datasetName == "" {
		return nil, errors.New("Missing required query parameters: 'datasetname'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	apiURL := fmt.Sprintf(app.ArrayOfDataURL, url.QueryEscape(datasetName))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetStyleDetails
func (o *OLAMap) GetStyleDetails(styleName string) (interface{}, error) {
	if styleName == "" {
		return nil, errors.New("Missing required query parameters: 'stylename'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	apiURL := fmt.Sprintf(app.StyleDetailsURL, url.QueryEscape(styleName))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetMapStyle
func (o *OLAMap) GetMapStyle() (interface{}, error) {
	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	apiURL := app.MapStyleURL

	// Define a variable to hold the API response
	var apiResponse []map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetPlaceDetail
func (o *OLAMap) GetPlaceDetail(placeID string) (interface{}, error) {
	if placeID == "" {
		return nil, errors.New("Missing required query parameters: 'placeid'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	apiURL := fmt.Sprintf(app.PlaceDetailURL, placeID) // Replace with your actual endpoint

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetNearBySearch
func (o *OLAMap) GetNearBySearch(nearBySearch internal.NearBySearch) (interface{}, error) {
	if nearBySearch.Layers == "" || nearBySearch.Location == "" {
		return nil, errors.New("Missing required query parameters: 'layers' and/or 'location'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	apiURL := fmt.Sprintf(app.NearBySearchURL, nearBySearch.Layers, nearBySearch.Location, nearBySearch.Types, nearBySearch.Radius, nearBySearch.Strictbounds, nearBySearch.WithCentroid, nearBySearch.Limit)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetTextSearch
func (o *OLAMap) GetTextSearch(textSearch internal.TextSearch) (interface{}, error) {
	// Extract query parameters
	if textSearch.Input == "" {
		return nil, errors.New("Missing required query parameters: 'input'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(app.TextSearchURL, url.QueryEscape(textSearch.Input), textSearch.Location, textSearch.Radius, textSearch.Types, textSearch.Size)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetSnapToRoad
func (o *OLAMap) GetSnapToRoad(points, enhancePath string) (interface{}, error) {
	if points == "" {
		return nil, errors.New("Missing required query parameters: 'points'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Build URL
	apiURL := fmt.Sprintf(app.SnapToRoadURL, url.Values{
		"points":      {points},
		"enhancePath": {enhancePath},
	}.Encode())

	// Make the external request
	var apiResponse map[string]interface{}
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetNearestRoads
func (o *OLAMap) GetNearestRoads(points string, radius string) (interface{}, error) {
	if points == "" {
		return nil, errors.New("Missing required query parameters: 'points' and/or 'radius'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Build the API URL
	apiURL := fmt.Sprintf(app.NearestRoadsURL, points, radius)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetStaticMapImageCenter
func (o *OLAMap) GetStaticMapImageCenter(mapImageCenter internal.MapImageCenter) (interface{}, error) {
	// Validate required parameters
	if mapImageCenter.Stylename == "" || mapImageCenter.Longitude == "" || mapImageCenter.Latitude == "" || mapImageCenter.Zoomlevel == "" || mapImageCenter.Imagewidth == "" || mapImageCenter.Imageheight == "" || mapImageCenter.Imageformat == "" {
		return nil, errors.New("Missing required query parameters: 'stylename' or 'longitude' or 'latitude' or 'zoomlevel' or 'width' or 'height' or 'format'")
	}

	longitude, err := strconv.ParseFloat(mapImageCenter.Longitude, 64)
	if err != nil {
		return nil, errors.New("Invalid longitude value")
	}

	latitude, err := strconv.ParseFloat(mapImageCenter.Latitude, 64)
	if err != nil {
		return nil, errors.New("Invalid latitude value")
	}

	zoomLevel, err := strconv.Atoi(mapImageCenter.Zoomlevel)
	if err != nil {
		return nil, errors.New("Invalid zoom level value")
	}

	imageWidth, err := strconv.Atoi(mapImageCenter.Imagewidth)
	if err != nil {
		return nil, errors.New("Invalid image width value")
	}

	imageHeight, err := strconv.Atoi(mapImageCenter.Imageheight)
	if err != nil {
		return nil, errors.New("Invalid image height value")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(app.StaticMapImageCenterURL,
		url.QueryEscape(mapImageCenter.Stylename), longitude, latitude, zoomLevel, imageWidth, imageHeight, mapImageCenter.Imageformat)

	// Construct query parameters
	queryParams := url.Values{}
	if len(mapImageCenter.Markers) > 0 {
		queryParams.Add("marker", strings.Join(mapImageCenter.Markers, ","))
	}

	if mapImageCenter.Path != "" {
		queryParams.Add("path", mapImageCenter.Path)
	}

	if len(queryParams) > 0 {
		apiURL += "?" + queryParams.Encode()
	}

	// Create and send the external request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	req.Header.Add("X-Request-Id", o.RequestId)

	if oauthToken != "" {
		req.Header.Add("Authorization", oauthToken)
	}

	xCorrelationID := uuid.New().String()
	req.Header.Add("X-Correlation-Id", xCorrelationID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to make external request")
	}
	defer resp.Body.Close()

	return resp, nil
}

// GetStaticMapImageBounded
func (o *OLAMap) GetStaticMapImageBounded(mapImageBounded internal.MapImageBounded) (interface{}, error) {
	if mapImageBounded.Stylename == "" || mapImageBounded.Minxstr == "" || mapImageBounded.Minystr == "" || mapImageBounded.Maxxstr == "" || mapImageBounded.Maxystr == "" || mapImageBounded.Imagewidth == "" || mapImageBounded.Imageheight == "" || mapImageBounded.Imageformat == "" {
		return nil, errors.New("Missing required query parameters: 'styleName' or 'minXStr' or 'minYStr' or 'maxXStr' or 'maxYStr' or 'imageWidthStr' or 'imageHeightStr' or 'imageFormat'")
	}

	minX, err := strconv.ParseFloat(mapImageBounded.Minxstr, 64)
	if err != nil {
		return nil, errors.New("Invalid min_x value")
	}

	minY, err := strconv.ParseFloat(mapImageBounded.Minystr, 64)
	if err != nil {
		return nil, errors.New("Invalid min_y value")
	}

	maxX, err := strconv.ParseFloat(mapImageBounded.Maxxstr, 64)
	if err != nil {
		return nil, errors.New("Invalid max_x value")
	}

	maxY, err := strconv.ParseFloat(mapImageBounded.Maxystr, 64)
	if err != nil {
		return nil, errors.New("Invalid max_y value")
	}

	imageWidth, err := strconv.Atoi(mapImageBounded.Imagewidth)
	if err != nil {
		return nil, errors.New("Invalid image width value")
	}

	imageHeight, err := strconv.Atoi(mapImageBounded.Imageheight)
	if err != nil {
		return nil, errors.New("Invalid image height value")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(app.StaticMapImageBoundedURL,
		url.QueryEscape(mapImageBounded.Stylename), minX, minY, maxX, maxY, imageWidth, imageHeight, mapImageBounded.Imageformat)

	// Construct query parameters
	queryParams := url.Values{}
	if len(mapImageBounded.Markers) > 0 {
		queryParams.Add("marker", strings.Join(mapImageBounded.Markers, ","))
	}

	if mapImageBounded.Path != "" {
		queryParams.Add("path", mapImageBounded.Path)
	}

	if len(queryParams) > 0 {
		apiURL += "?" + queryParams.Encode()
	}

	// Make the external request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	req.Header.Add("X-Request-Id", o.RequestId)
	req.Header.Add("X-Correlation-Id", o.RequestId)

	if oauthToken != "" {
		req.Header.Add("Authorization", oauthToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to make external request")
	}
	defer resp.Body.Close()

	return resp, err
}

// StaticMapImage
func (o *OLAMap) StaticMapImage(mapImage internal.MapImage) (interface{}, error) {
	if mapImage.Stylename == "" || mapImage.Imagewidth == "" || mapImage.Imageheight == "" || mapImage.Imageformat == "" || mapImage.Path == "" {
		return nil, errors.New("Missing required query parameters: 'stylename' or 'imagewidth' or 'imageheight' or 'imageformat' or path")
	}

	imageWidth, err := strconv.Atoi(mapImage.Imagewidth)
	if err != nil {
		return nil, errors.New("Invalid image width value")
	}

	imageHeight, err := strconv.Atoi(mapImage.Imageheight)
	if err != nil {
		return nil, errors.New("Invalid image height value")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(app.StaticMapImageURL, url.QueryEscape(mapImage.Stylename), imageWidth, imageHeight, mapImage.Imageformat)

	// Construct query parameters
	queryParams := url.Values{}
	if len(mapImage.Markers) > 0 {
		queryParams.Add("marker", strings.Join(mapImage.Markers, ","))
	}

	if mapImage.Path != "" {
		queryParams.Add("path", mapImage.Path)
	}

	if len(queryParams) > 0 {
		apiURL += "?" + queryParams.Encode()
	}
	// Create a new request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, errors.New("Failed to create request")
	}

	if oauthToken != "" {
		req.Header.Add("Authorization", oauthToken)
	}

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to send request")
	}
	defer resp.Body.Close()

	return resp, nil
}
