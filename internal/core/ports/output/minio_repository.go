package output

import "mime/multipart"

type MinIORepository interface {
	FileUploadSvc(file multipart.File, directoryPath, filename string, fileSize int64) (string, error)
	DeleteFileSvc(bucketName string) error
	UploadSingleFile(file *multipart.FileHeader, filepath string) (string, error)
	UploadMultipleFiles(files []*multipart.FileHeader, filepath string) ([]string, error)
	DeleteFile(objectName string) error
	GetFileURL(objectName string) string
}
