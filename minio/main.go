package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint        = "127.0.0.1:9000"
	accessKeyID     = "jim"
	secretAccessKey = "password"
	useSSL          = false
)

var (
	repStr = `<ReplicationConfiguration>
	<Rule>
		<ID>rule1</ID>
		<Status>Enabled</Status>
		<Priority>1</Priority>
		<DeleteMarkerReplication>
			<Status>Disabled</Status>
		</DeleteMarkerReplication>
			<Destination>
				<Bucket>arn:aws:s3:::test</Bucket>
			</Destination>
			<Filter>
				<And>
				<Prefix></Prefix>
				</And>
			</Filter>
	</Rule>
	</ReplicationConfiguration>`
)

type MinioClient struct {
	Client *minio.Client
}

type ClientImp interface {
	CreateBucket()
	GetBucket()
	UploadFile(string) error
}

// NewMinioClient : Initialize minio client object.
func NewMinioClient() (ClientImp, error) {
	c, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, err
	}

	return &MinioClient{
		Client: c,
	}, nil
}

func (c *MinioClient) GetBucket() {
	s3Client := c.Client
	s3Client.TraceOn(os.Stderr)

	// Get replication metrics for a bucket
	m, err := s3Client.GetBucketReplicationMetrics(context.Background(), "bucket-jim")
	if err != nil {
		log.Fatalf("get error %v\n", err)
	}
	fmt.Println("Replication metrics for my-bucketname:", m)
}

func (c *MinioClient) CreateBucket() {
	s3Client := c.Client
	s3Client.TraceOn(os.Stderr)
	bucketName := "bucket-jim"
	location := "TW"

	err := s3Client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := s3Client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			fmt.Println(fmt.Errorf("bucket %s alread exists at %s", bucketName, location))
		} else {
			fmt.Println("create bucket failed: ", err)
		}
	}

	fmt.Println("set bucket: bucket success")
}

func (c *MinioClient) UploadFile(f string) error {
	return nil
}

func main() {
	c, err := NewMinioClient()
	if err != nil {
		log.Fatal(err)
	}

	c.CreateBucket()
	c.GetBucket()
}
