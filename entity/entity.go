package entity

import "github.com/DavinciUI/backend/code"

type Entity interface {
	GetCode() code.Code
}
