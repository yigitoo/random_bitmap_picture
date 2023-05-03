package cmd

// uuid.New().String() -> bit to string
type UUIDString string
type Image string

type EncodingTypes struct {
	types []string
}

var EncodingType = []string{"png" /*"svg" not added because of i didnt write code fot it :D*/}

type ProfilePicture struct {
	UniqueID     UUIDString
	ImageDetails ImageData
	Picture      Image
}

type ImageData struct {
	ImageString  Image
	EncodingType string
	FileType     string
}
