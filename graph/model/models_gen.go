// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Achievement struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AchievementInput struct {
	ID string `json:"id"`
}

type Chill struct {
	ID     string        `json:"id"`
	Traces []*TracePoint `json:"traces"`
	Photos []*Photo      `json:"photos"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CoordinateInput struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EndChillInput struct {
	ID           string              `json:"id"`
	Timestamp    string              `json:"timestamp"`
	Coordinate   *CoordinateInput    `json:"coordinate"`
	Achievements []*AchievementInput `json:"achievements"`
}

type Photo struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Timestamp string `json:"timestamp"`
}

type PhotoInput struct {
	URL       string `json:"url"`
	Timestamp string `json:"timestamp"`
}

type PhotosInput struct {
	ID     string        `json:"id"`
	Photos []*PhotoInput `json:"photos"`
}

type RegisterUserInput struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type StartChillInput struct {
	Timestamp  string           `json:"timestamp"`
	Coordinate *CoordinateInput `json:"coordinate"`
}

type TracePoint struct {
	ID         string      `json:"id"`
	Timestamp  string      `json:"timestamp"`
	Coordinate *Coordinate `json:"coordinate"`
}

type TracePointInput struct {
	Timestamp  string           `json:"timestamp"`
	Coordinate *CoordinateInput `json:"coordinate"`
}

type TracePointsInput struct {
	ID          string             `json:"id"`
	TracePoints []*TracePointInput `json:"tracePoints"`
}

type User struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Avatar       string         `json:"avatar"`
	Chills       []*Chill       `json:"chills"`
	Achievements []*Achievement `json:"achievements"`
}
