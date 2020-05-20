package models

import "time"

// Node represents a remote kamji-node owner by a user
type Node struct {
	ID            string    `bson:"id"`
	TS            time.Time `bson:"ts"`
	LastHeartbeat time.Time `bson:"lastHeartbeat"`
	Version       string    `bson:"version"`
	HostData      HostData  `bson:"hostData"`
	Pods          []Pod     `bson:"pods"`
}

// HostData represents the metadata of a host machine where a kamaji-node is running
type HostData struct {
	OS            string `bson:"os"`
	DockerVersion string `bson:"dockerVersion"`
}
