package gcp

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type UploadRes struct {
	Filename string `json:"file_name"`
	Data     []byte `json:"data"`
}

type GcpClient struct {
	Client *storage.Client
	Bucket string
}

func NewGcpAPi(a GcpClient) *GcpClient {
	return &GcpClient{
		Client: a.Client,
		Bucket: a.Bucket,
	}
}

func (u GcpClient) UploadFile(ctx context.Context, f multipart.File, filename string) error {

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := u.Client.Bucket(u.Bucket).Object(filename)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	// For an object that does not yet exist, set the DoesNotExist precondition.
	o = o.If(storage.Conditions{DoesNotExist: true})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	// 	return fmt.Errorf("object.Attrs: %v", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := o.NewWriter(ctx)
	if _, err := io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil

}

func (u GcpClient) GetFile(ctx context.Context, object string) (*storage.ObjectAttrs, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := u.Client.Bucket(u.Bucket).Object(object)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).Attrs: %v", object, err)
	}

	return attrs, nil
}

func (u GcpClient) DeleteFile(ctx context.Context, object string) error {

	o := u.Client.Bucket(u.Bucket).Object(object)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("object.Attrs: %v", err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", object, err)
	}

	return nil
}
