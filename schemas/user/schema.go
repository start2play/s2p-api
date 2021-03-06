package user

import (
	"s2p-api/core/reflection"
)

var Schema = reflection.InternalSchema{
	Name: "user",
	Querys: []*reflection.RootField{
		FilterByField,
		Login,
		Recovery,
		EmailConfirmation,
	},
	Mutations: []*reflection.RootField{
		CreateField,
		UpdateField,
		DeleteField,
		ResetPassword,
	},
}
