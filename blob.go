package main

import (
	"bytes"
	"io"

	"github.com/Azure/azure-sdk-for-go/storage"
)

type blob struct {
	b     *storage.Blob
	lease string
}

func newBlob() (*blob, error) {
	c, err := storage.NewClient(cfg.accountName, cfg.accountKey,
		"core.windows.net", storage.DefaultAPIVersion, cfg.useHTTPS)

	// c.HTTPClient.Transport = &debugRoundTripper{}

	bsc := c.GetBlobService()

	ctr := bsc.GetContainerReference(cfg.containerName)

	_, err = ctr.CreateIfNotExists(nil)
	if err != nil {
		return nil, err
	}

	b := ctr.GetBlobReference(cfg.blobName)
	b.Properties.ContentLength = cfg.size

	exists, err := b.Exists()
	if err != nil {
		return nil, err
	}

	if !exists {
		err = b.PutPageBlob(nil)
		if err != nil {
			return nil, err
		}
	}

	lease, err := b.AcquireLease(-1, "", nil)
	if err != nil {
		return nil, err
	}

	return &blob{b: b, lease: lease}, nil
}

func (b *blob) Close() error {
	return b.b.ReleaseLease(b.lease, nil)
}

func (b *blob) WriteAt(p []byte, off int64) (int, error) {
	err := b.b.WriteRange(storage.BlobRange{
		Start: uint64(off),
		End:   uint64(off) + uint64(len(p)) - 1,
	},
		bytes.NewBuffer(p),
		&storage.PutPageOptions{LeaseID: b.lease})

	return len(p), err
}

func (b *blob) ReadAt(p []byte, off int64) (int, error) {
	rc, err := b.b.GetRange(&storage.GetBlobRangeOptions{
		Range: &storage.BlobRange{
			Start: uint64(off),
			End:   uint64(off) + uint64(len(p)) - 1,
		},
	})
	if err != nil {
		return 0, err
	}
	defer rc.Close()

	return io.ReadFull(rc, p)
}
