package models

import "time"

type Appointment struct {
	ApptName string    `json:"apptName"`
	ApptType string    `json:"apptType"`
	UserID   int       `json:"userID"`
	ApptDate time.Time `json:"apptDate"`
	ApptTime string    `json:"apptTime"`
	ApptDesc string    `json:"apptDesc,omitempty"`
	ID       int       `json:"id,omitempty"`
}
