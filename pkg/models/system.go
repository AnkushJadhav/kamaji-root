package models

// Bootupstate is represent the boot up state of the system
type Bootupstate int

const (
	// BootupStatePending is the 0 state for bootup process
	BootupStatePending = iota
	// BootupStateCompleted is the 2 state for bootup process
	BootupStateCompleted
)

// SystemConfig stores the config used by the system
type SystemConfig struct {
	ID          string      `bson:"id"`
	BootupState Bootupstate `bson:"bootupState"`
	RootToken   string      `bson:"rootToken"`
	JWTSecret   string      `bson:"jwtToken"`
}
