package models

import "time"

// Pod reperesents a pod running on a kamaji-node
type Pod struct {
	ID            string    `bson:"id"`
	TS            time.Time `bson:"ts"`
	LastHeartbeat time.Time `bson:"lastHeartbeat"`
	MultiAddr     string    `bson:"multiAddr"`
	TemplateID    int       `bson:"tplId"`
	State         string    `bson:"state"`
}
