// Copyright (C) 2013-2017, The MetaCurrency Project (Eric Harris-Braun, Arthur Brock, et. al.)
// Use of this source code is governed by GPLv3 found in the LICENSE file
//----------------------------------------------------------------------------------------
// Scaffold implements loading DNA and other app information from a scaffold structure

package holochain

import (
	"io"
)

// LoadDNA decodes a DNA from scaffold file (via an io.reader)
func LoadDNAScaffold(reader io.Reader) (dnaP *DNA, err error) {
	var scaffold DNA
	err = Decode(reader, "yml", &scaffold)
	if err != nil {
		return
	}
	dnaP = &scaffold
	scaffold.PropertiesSchema = `{
	"title": "Properties Schema",
	"type": "object",
	"properties": {
		"description": {
			"type": "string"
		},
		"language": {
			"type": "string"
		}
	}
}
`
	return
}

var BasicTemplateScaffold string = `{
  # This is a holochain scaffold yaml definition. http://ceptr.org/projects/holochain

  # Scaffold Version
  # The hc-scaffold schema version of this file.
  "scaffoldVersion": "0.0.1",
  "generator": "hc-scaffold:0.0.2+d8b4ed4",

  # DNA File Version
  # Version indicator for changes to DNA
  "Version": 1,

  # DNA Unique ID
  # This ID differentiates your app from others. For example, to tell one Slack team from another with same code, so change it!
  "UUID": "00000000-0000-0000-0000-000000000000",

  # Application Name
  # What would you like to call your holochain app?
  "Name": "templateApp",

  # Properties
  # Properties that you want available across all Zomes.
  "Properties": {

    # Application Description
    # Briefly describe your holochain app.
    "description": "provides an application template",

    # Language
    # The base (human) language of this holochain app.
    "language": "en"
  },

  # Properties Schema File
  # Describes the entries in the Properties section of your dna file.
  "PropertiesSchemaFile": "properties_schema.json",

  # DHT Settings
  # Configure the properties of your Distributed Hash Table (e.g. hash algorithm, neighborhood size, etc.).
  "DHTConfig": {
    "HashType": "sha2-256"
  },

  # Zomes
  # List the Zomes your application will support.
  "Zomes": [
    {

      # Zome Name
      # The name of this code module.
      "Name": "testZome",

      # Zome Description
      # What is the purpose of this module?
      "Description": "provide an template zome",

      # Ribosome Type
      # What scripting language will you code in?
      "RibosomeType": "js",

      # Code File
      # Points to the main script file for this Zome.
      "CodeFile": "testZomezome.js",

      # Zome Entries
      # Data stored and tracked by your Zome.
      "Entries": [
        {
          "Name": "testEntry", # The name of this entry.
          "Required": true, # Is this entry required?
          "DataFormat": "json", # What type of data should this entry store?
          "Sharing": "public", # Should this entry be publicly accessible?
          "SchemaFile": "testEntry.json",
          "_": "cr"
        }
      ],

      # Zome Functions
      # Functions which can be called in your Zome's API.
      "Functions": [
        {
          "Name": "testEntryCreate", # The name of this function.
          "CallingType": "json", # Data format for parameters passed to this function.
          "Exposure": "public", # Level to which is this function exposed.
          "_": "c:testEntry"
        },
        {
          "Name": "testEntryRead", # The name of this function.
          "CallingType": "json", # Data format for parameters passed to this function.
          "Exposure": "public", # Level to which is this function exposed.
          "_": "r:testEntry"
        },
        {
          "Name": "doTestAction", # The name of this function.
          "CallingType": "json", # Data format for parameters passed to this function.
          "Exposure": "public", # Level to which is this function exposed.
        }
      ],

      # Zome Source Code
      # The logic that will control Zome behavior
      "Code": "'use strict';\n\n// -----------------------------------------------------------------\n//  This stub Zome code file was auto-generated\n// -----------------------------------------------------------------\n\n/**\n * Called only when your source chain is generated\n * @return {boolean} success\n */\nfunction genesis () {\n  // any genesis code here\n  return true;\n}\n\n// -----------------------------------------------------------------\n//  validation functions for every DHT entry change\n// -----------------------------------------------------------------\n\n/**\n * Called to validate any changes to the DHT\n * @param {string} entryName - the name of entry being modified\n * @param {*} entry - the entry data to be set\n * @param {?} header - ?\n * @param {?} pkg - ?\n * @param {?} sources - ?\n * @return {boolean} is valid?\n */\nfunction validateCommit (entryName, entry, header, pkg, sources) {\n  switch (entryName) {\n    case \"testEntry\":\n      // validation code here\n      return false;\n    default:\n      // invalid entry name!!\n      return false;\n  }\n}\n\n/**\n * Called to validate any changes to the DHT\n * @param {string} entryName - the name of entry being modified\n * @param {*}entry - the entry data to be set\n * @param {?} header - ?\n * @param {?} pkg - ?\n * @param {?} sources - ?\n * @return {boolean} is valid?\n */\nfunction validatePut (entryName, entry, header, pkg, sources) {\n  switch (entryName) {\n    case \"testEntry\":\n      // validation code here\n      return false;\ndefault:\n      // invalid entry name!!\n      return false;\n  }\n}\n\n/**\n * Called to validate any changes to the DHT\n * @param {string} entryName - the name of entry being modified\n * @param {*} entry- the entry data to be set\n * @param {?} header - ?\n * @param {*} replaces - the old entry data\n * @param {?} pkg - ?\n * @param {?} sources - ?\n * @return {boolean} is valid?\n */\nfunction validateMod (entryName, entry, header, replaces, pkg, sources) {\n  switch (entryName) {\n    case \"testEntry\":\n      // validation code here\n      return false;\n    default:\n      // invalid entry name!!\n      return false;\n  }\n}\n\n/**\n * Called to validate any changes to the DHT\n * @param {string} entryName - the name of entry being modified\n * @param {string} hash - the hash of the entry to remove\n * @param {?} pkg - ?\n * @param {?} sources - ?\n * @return {boolean} is valid?\n */\nfunction validateDel (entryName,hash, pkg, sources) {\n  switch (entryName) {\n    case \"testEntry\":\n      // validation code here\nreturn false;\n    default:\n      // invalid entry name!!\n      return false;\n  }\n}\n\n/**\n * Called to get the data needed to validate\n * @param {string} entryName - the name of entry to validate\n * @return {*} the data required for validation\n */\nfunction validatePutPkg (entryName) {\n  return null;\n}\n\n/**\n * Called to get the data needed to validate\n * @param {string} entryName - the name of entry to validate\n * @return {*} the data required for validation\n */\nfunction validateModPkg (entryName) {\n  return null;\n}\n\n/**\n * Called to get the data needed to validate\n * @param {string} entryName - the name of entry to validate\n * @return {*} the data required for validation\n */\nfunction validateDelPkg (entryName) {\n  return null;\n}"
    }
  ]
}
`
