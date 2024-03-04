package filestore

type IFilestoreRepository[T interface{}] interface {
	UploadFile(request *T) (*string, error)
	DownloadFile(request *T) ([]byte, error)
	GeneratePreSignedURl(request *T) (*string, error)
}
