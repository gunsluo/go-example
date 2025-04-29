package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	NWSAPIBase = "https://api.weather.gov"
	UserAgent  = "weather-app/1.0"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"weather",
		"1.6.0",
	)

	// Add tool
	getAlertsTool := mcp.NewTool("get_alerts",
		mcp.WithDescription("Get weather alerts for a US state."),
		mcp.WithString("state",
			mcp.Required(),
			mcp.Description("Two-letter US state code (e.g. CA, NY)"),
		),
	)

	getForecastTool := mcp.NewTool("get_forecast",
		mcp.WithDescription("Get weather forecast for a location."),
		mcp.WithNumber("latitude",
			mcp.Required(),
			mcp.Description("Latitude of the location"),
		),
		mcp.WithNumber("longitude",
			mcp.Required(),
			mcp.Description("Longitude of the location"),
		),
	)

	// Add tool handler
	s.AddTool(getAlertsTool, getAlertsHandler)
	s.AddTool(getForecastTool, getForecastHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func getAlertsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	state, ok := request.Params.Arguments["state"].(string)
	if !ok {
		return nil, errors.New("state must be a string")
	}

	alerts, err := getAlerts(state)
	if err != nil {
		return nil, err
	}

	var textAlerts []string
	for _, alert := range alerts {
		textAlerts = append(textAlerts,
			fmt.Sprintf("[%s]\n%s\n%s", alert.Properties.Event, alert.Properties.Headline, alert.Properties.Description),
		)
	}

	return mcp.NewToolResultText(joinWithSeparator(textAlerts, "\n---\n")), nil
}

func getForecastHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	latitude, ok := request.Params.Arguments["latitude"].(float64)
	if !ok {
		return nil, errors.New("latitude must be a number")
	}

	longitude, ok := request.Params.Arguments["longitude"].(float64)
	if !ok {
		return nil, errors.New("longitude must be a number")
	}

	periods, err := getForecast(latitude, longitude)
	if err != nil {
		return nil, err
	}

	var textPeriods []string
	for i, period := range periods {
		if i >= 5 {
			break
		}

		text := fmt.Sprintf(
			"%s:\nTemperature: %.0f°%s\nWind: %s %s\nForecast: %s",
			period.Name,
			period.Temperature,
			period.TemperatureUnit,
			period.WindSpeed,
			period.WindDirection,
			period.DetailedForecast,
		)

		textPeriods = append(textPeriods, text)
	}

	return mcp.NewToolResultText(joinWithSeparator(textPeriods, "\n---\n")), nil
}

// getAlerts retrieves weather alerts for a given state (e.g. "CA").
func getAlerts(state string) ([]AlertFeature, error) {
	url := fmt.Sprintf("%s/alerts/active/area/%s", NWSAPIBase, state)
	var data AlertResponse
	if err := doMakeRequest(url, &data); err != nil {
		return nil, errors.New("Unable to fetch alerts.")
	}

	if len(data.Features) == 0 {
		return nil, errors.New("No active alerts for this state.")
	}

	return data.Features, nil
}

// getForecast retrieves the forecast for a specific location by latitude and longitude.
func getForecast(latitude, longitude float64) ([]ForecastPeriod, error) {
	pointsURL := fmt.Sprintf("%s/points/%.4f,%.4f", NWSAPIBase, latitude, longitude)

	var data PointsData
	if err := doMakeRequest(pointsURL, &data); err != nil {
		return nil, errors.New("Unable to fetch forecast data for this location.")
	}

	var forecastData ForecastData
	if err := doMakeRequest(data.Properties.Forecast, &forecastData); err != nil {
		return nil, errors.New("Unable to fetch detailed forecast.")
	}

	return forecastData.Properties.Periods, nil
}

// doMakeRequest sends a GET request to the NWS API and returns parsed JSON.
func doMakeRequest(url string, data any) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}

func joinWithSeparator(items []string, sep string) string {
	result := ""
	for i, item := range items {
		if i > 0 {
			result += sep
		}
		result += item
	}
	return result
}

type AlertResponse struct {
	Features []AlertFeature `json:"features"`
}

type AlertFeature struct {
	Properties AlertProperties `json:"properties"`
}

type AlertProperties struct {
	Event       string `json:"event"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
}

type PointsData struct {
	Properties PointsProperties `json:"properties"`
}

type PointsProperties struct {
	Forecast string `json:"forecast"`
}

type ForecastData struct {
	Properties ForecastProperties `json:"properties"`
}

type ForecastProperties struct {
	Periods []ForecastPeriod `json:"periods"`
}

type ForecastPeriod struct {
	Name             string  `json:"name"`
	Temperature      float64 `json:"temperature"`
	TemperatureUnit  string  `json:"temperatureUnit"`
	WindSpeed        string  `json:"windSpeed"`
	WindDirection    string  `json:"windDirection"`
	DetailedForecast string  `json:"detailedForecast"`
}
