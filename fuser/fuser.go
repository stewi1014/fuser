package fuser

import (
	"fuser/openapi"
	"log"
	"net/http"
	"net/url"
	"os"

	"bazil.org/fuse/fs"
)

func ParseSource(in *string) fs.FS {
	address, err := url.Parse(*in)
	if err != nil {
		log.Fatal(err)
	}

	switch address.Scheme {
	case "http", "https":
		resp, err := http.Get(address.String())
		if err != nil {
			log.Fatal(err)
		}

		// TODO; Handle more cases
		// For now, assume it is an openapi.json file

		defer resp.Body.Close()
		return openapi.NewFilesystem(resp.Body)

	case "file":
		file, err := os.Open(address.Path)
		if err != nil {
			log.Fatal(err)
		}

		// TODO; Handle more cases
		// For now, assume it is an openapi.json file

		defer file.Close()
		return openapi.NewFilesystem(file)

	default:
		log.Fatalf("Unknown scheme %s", address.Scheme)
		return nil // not reached
	}
}

func ParseDestination(out *string) string {
	return *out
}
