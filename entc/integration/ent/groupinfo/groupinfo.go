// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package groupinfo

const (
	// Label holds the string label denoting the groupinfo type in the database.
	Label = "group_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldMaxUsers holds the string denoting the max_users field in the database.
	FieldMaxUsers = "max_users"

	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"

	// Table holds the table name of the groupinfo in the database.
	Table = "group_infos"
	// GroupsTable is the table the holds the groups relation/edge.
	GroupsTable = "groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "group_info"
)

// Columns holds all SQL columns for groupinfo fields.
var Columns = []string{
	FieldID,
	FieldDesc,
	FieldMaxUsers,
}

var (
	// DefaultMaxUsers holds the default value on creation for the max_users field.
	DefaultMaxUsers int
)
