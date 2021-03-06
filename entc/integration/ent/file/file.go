// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package file

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID    = "id"    // FieldSize holds the string denoting the size vertex property in the database.
	FieldSize  = "fsize" // FieldName holds the string denoting the name vertex property in the database.
	FieldName  = "name"  // FieldUser holds the string denoting the user vertex property in the database.
	FieldUser  = "user"  // FieldGroup holds the string denoting the group vertex property in the database.
	FieldGroup = "group"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"

	// Table holds the table name of the file in the database.
	Table = "files"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "files"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_files"
	// TypeTable is the table the holds the type relation/edge.
	TypeTable = "files"
	// TypeInverseTable is the table name for the FileType entity.
	// It exists in this package in order to avoid circular dependency with the "filetype" package.
	TypeInverseTable = "file_types"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "file_type_files"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldSize,
	FieldName,
	FieldUser,
	FieldGroup,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the File type.
var ForeignKeys = []string{
	"file_type_files",
	"group_files",
	"user_files",
}

var (
	// DefaultSize holds the default value on creation for the size field.
	DefaultSize int
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int) error
)
