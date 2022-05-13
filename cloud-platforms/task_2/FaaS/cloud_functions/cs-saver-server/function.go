package cs_saver_server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("CSSaverServer", csSaverServer)
}

func csSaverServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Hello World!\n")
	case http.MethodPut, http.MethodPost:
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		contentString := string(bodyBytes)

		ctx := context.Background()

		// Sets your Google Cloud Platform project ID.
		projectID := os.Getenv("PROJECT_ID")

		// Creates a client.
		client, err := storage.NewClient(ctx)
		if err != nil {
			msg := fmt.Sprintf("Failed to create storage client: %v\n", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		defer client.Close()

		// Sets the name for the new bucket.
		bucketName := "cl-bucket-2205"

		// Check if Storage Bucket exists.
		// If not, create it.
		bucket := client.Bucket(bucketName)
		_, err = bucket.Attrs(ctx)
		if err == storage.ErrBucketNotExist {
			fmt.Fprintln(w, "The bucket does not exist. Creating new bucket..")
			ctx, cancel := context.WithTimeout(ctx, time.Second*20)
			defer cancel()
			if err := bucket.Create(ctx, projectID, nil); err != nil {
				msg := fmt.Sprintf("Failed to create bucket: %v\n", err)
				http.Error(w, msg, http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Bucket %v created.\n", bucketName)
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "The bucket %s exists!\n", bucketName)

		// Write an object to the bucket.
		objName := generateObjName()
		obj := bucket.Object(objName)
		// Write request content to obj.
		// writer implements io.Writer.
		writer := obj.NewWriter(ctx)
		// This will either create the object or overwrite whatever is there already.
		_, err = fmt.Fprint(writer, contentString)
		if err != nil {
			msg := fmt.Sprintf("Failed to create an object: %v\n", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		// Close, just like writing a file.
		if err := writer.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Object %s was successfully created!\n", objName)
		log.Printf("Object %s was successfully created!\n", objName)
	}
}

func generateObjName() string {
	objectName := fmt.Sprintf("obj-%d.json", time.Now().Unix())
	return objectName
}
