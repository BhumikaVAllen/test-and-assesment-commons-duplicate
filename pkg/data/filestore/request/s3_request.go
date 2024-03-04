package request

type S3Request struct {
	Bucket   string
	Key      string
	ByteData []byte
}
