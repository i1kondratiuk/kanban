package common

// Entity is the base type for domain models
type Entity interface {
	GetId() Id
}
