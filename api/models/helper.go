package models

import "time"

/*
HealthCheckResponse
Interface correlating to the response body for a health check request.
*/
type HealthCheckResponse struct {
	Date    time.Time
	Message string
}
