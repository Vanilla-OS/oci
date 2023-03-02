package oci

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDigest(t *testing.T) {
	image := "docker.io/nginx:latest"

	digest, err := Digest(image)
	if err != nil {
		t.Error(err)
	}
	if len(digest) == 0 {
		t.Error("zero length digest")
	}
}

func TestDigestPath(t *testing.T) {
	image := "docker.io/nginx:latest"

	digestPath, err := DigestPath(image)
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(digestPath, ":") {
		t.Error("digest path contains invalid characters")
	}
}

func TestSave(t *testing.T) {
	image := "docker.io/nginx:latest"
	tempDirectory := t.TempDir()
	err := Save(image, tempDirectory)
	if err != nil {
		t.Error(err)
	}
	digestPath, err := DigestPath(image)
	if err != nil {
		t.Error(err)
	}
	want := filepath.Join(tempDirectory, digestPath)
	_, err = os.Stat(want)
	if err != nil {
		t.Error(err)
	}

}
