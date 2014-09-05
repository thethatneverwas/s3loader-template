package main

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"log"
	"os"
)

func main() {
	auth, err := aws.EnvAuth()
	chkerr(err)
	region := aws.Regions[os.Getenv("HACO_AWS_REGION")]
	connection := s3.New(auth, region)
	mybucket := connection.Bucket(os.Getenv("HACO_AWS_BUCKET"))

	file, err := os.Open(os.Args[1])
	chkerr(err)
	fi, err := file.Stat()
	chkerr(err)
	data := make([]byte, fi.Size())
	_, err = file.Read(data)
	chkerr(err)
	path := fi.Name()
	if 2 < len(os.Args) {
		path = os.Args[2]
	}
	err = mybucket.Put(path, data, "text/plain", s3.BucketOwnerFull)
	chkerr(err)
}

func chkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
