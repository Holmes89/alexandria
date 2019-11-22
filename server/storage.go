package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
	"io"
	"time"
)

type DocumentSave interface {
	Save(ctx context.Context, fileName string, reader io.Reader) (path string, err error)
}

type DocumentGet interface {
	Get(ctx context.Context, path string) (string, error)
	List(ctx context.Context) <-chan string
}

type DocumentStorage interface {
	DocumentSave
	DocumentGet
}

type BucketStorage struct {
	Bucket *blob.Bucket
}

func NewBucketStorage(config BucketConfig) *BucketStorage {
	bucket, err := blob.OpenBucket(context.Background(), config.ConnectionString)
	if err != nil {
		logrus.WithError(err).Fatal("unable to connect to bucket")
	}
	return &BucketStorage{
		Bucket: bucket,
	}
}

func NewBucketDocumentStorage(storage *BucketStorage) DocumentStorage {
	return storage
}

func (s *BucketStorage) Save(ctx context.Context, fileName string, reader io.Reader) (path string, err error) {

	w, err := s.Bucket.NewWriter(ctx, fileName, nil)
	if err != nil {
		logrus.WithError(err).Error("unable to create upload writer")
		return "", errors.Wrap(err, "unable to creat upload writer")
	}

	defer w.Close()

	if _, err := io.Copy(w, reader); err != nil {
		logrus.WithError(err).Error("failed to upload file")
		return "", errors.Wrap(err, "failed to upload file")
	}

	return fileName, err //TODO allow for custom directory?
}


func (s *BucketStorage) Get(ctx context.Context, path string) (string, error) {
	opts := &blob.SignedURLOptions{
		Expiry: 15 * time.Hour,
		Method: "GET",
	}
	return s.Bucket.SignedURL(ctx, path, opts)
}

func (s *BucketStorage) List(ctx context.Context) <-chan string {
	opts := &blob.ListOptions{}
	iter := s.Bucket.List(opts)
	out := make(chan string)
	go func() {
		defer close(out)
		bgctx := context.Background()
		for{
			for {
				obj, err := iter.Next(bgctx)
				if err == io.EOF {
					logrus.Info("file search complete")
					return
				}
				if err != nil {
					logrus.WithError(err).Error("unable to fetch object")
					return
				}
				if obj.IsDir {
					logrus.WithField("name", obj.Key).Debug("directory found")
					continue
				}
				out <- obj.Key
			}
		}
	}()
	return out
}