// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package card

import (
	"time"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeSpec holds the string denoting the spec edge name in mutations.
	EdgeSpec = "spec"

	// Table holds the table name of the card in the database.
	Table = "cards"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "cards"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_card"
	// SpecTable is the table the holds the spec relation/edge. The primary key declared below.
	SpecTable = "spec_card"
	// SpecInverseTable is the table name for the Spec entity.
	// It exists in this package in order to avoid circular dependency with the "spec" package.
	SpecInverseTable = "specs"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldNumber,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Card type.
var ForeignKeys = []string{
	"user_card",
}

var (
	// SpecPrimaryKey and SpecColumn2 are the table columns denoting the
	// primary key for the spec relation (M2M).
	SpecPrimaryKey = []string{"spec_id", "card_id"}
)

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// comment from another template.
