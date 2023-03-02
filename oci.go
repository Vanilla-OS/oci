package oci

import (
	"fmt"
	"path/filepath"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"

	"github.com/kennygrant/sanitize"
)

// Digest returns the sha256 digest of an OCI image
// as a string
func Digest(image string) (string, error) {
	options := []crane.Option{}
	return crane.Digest(image, options...)
}

// DigestPath returns the sha256 digest of an OCI image
// as a string that has been sanitized for safety as
// a file name
func DigestPath(image string) (string, error) {
	digest, err := Digest(image)
	if err != nil {
		return "", err
	}
	return sanitize.Name(digest), nil
}

// Save creates a tar file in `basePath` with the
// sanitized `image` sha256 as the file name.
func Save(image, basePath string) error {
	options := []crane.Option{}

	imageMap := map[string]v1.Image{}
	o := crane.GetOptions(options...)

	ref, err := name.ParseReference(image, o.Name...)
	if err != nil {
		return fmt.Errorf("parsing reference %q: %w", image, err)
	}

	rmt, err := remote.Get(ref, o.Remote...)
	if err != nil {
		return err
	}

	img, err := rmt.Image()
	if err != nil {
		return err
	}

	imageMap[image] = img
	dp, err := DigestPath(image)
	if err != nil {
		return err
	}
	imagePath := filepath.Join(basePath, dp)
	return crane.MultiSave(imageMap, imagePath)
}
