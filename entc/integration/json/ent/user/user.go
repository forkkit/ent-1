// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// FieldDirs holds the string denoting the dirs field in the database.
	FieldDirs = "dirs"
	// FieldInts holds the string denoting the ints field in the database.
	FieldInts = "ints"
	// FieldFloats holds the string denoting the floats field in the database.
	FieldFloats = "floats"
	// FieldStrings holds the string denoting the strings field in the database.
	FieldStrings = "strings"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldRaw,
	FieldDirs,
	FieldInts,
	FieldFloats,
	FieldStrings,
}
