package cmd

import (
	"errors"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"

	"golang.org/x/exp/slices"
)

func NewProfile(file_type string) string {
	return CreateProfilePicture(file_type)
}

func CreateProfilePicture(file_type string) string {
	new_uuid := CreateID()
	new_fileName, err := CreateNewImage(string(new_uuid), file_type)
	if err != nil {
		panic(err)
	}
	return new_fileName
}

func CreateID() UUIDString {
	return UUIDString(uuid.NewString())
}

func CreateNewImage(uuid_string string, file_type string) (string, error) {
	if slices.Contains(EncodingType, file_type) {
		file_name := CreateImage(uuid_string)
		return file_name, nil
	}
	return "Cannot create file.", errors.New("Error file type is not supported. Must be [jpg, jpeg, png]")
}

/*
C ImageData {
 	ImageString.:  "nil",
 	EncodingType: "nil",
 	FileType:     "nil",
  }
*/
