package application

// Import the required packages for upload and admin.

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/admin/search"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

func cloudinaryConfig(c *gin.Context) string {
	file, header, err := c.Request.FormFile("image")

	if err != nil {
		panic(err)
	}

	filename := header.Filename

	fmt.Println("the file name is:", filename)
	fmt.Println("file details is", file)

	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	cld, _ := cloudinary.NewFromParams("dsprvlehn", "888255571214679", "l871yMNBHYHaY7CTTYZsZVyP5ik")

	var ctx = context.Background()

	var extension = filepath.Ext(filename)
	var name = filename[0 : len(filename)-len(extension)]

	// Upload the my_picture.jpg image and set the public_id to "my_image".

	// resp, err := cld.Upload.Upload(ctx, "my_picture.jpg", uploader.UploadParams{PublicID: "my_image"})

	// Get details about the image with public id "my_image" and log the secure URL.

	// Upload an image to your Cloudinary account from a specified URL.
	//
	// Alternatively you can provide a path to a local file on your filesystem,
	// base64 encoded string, io.Reader and more.
	//
	// For additional information see:
	// https://cloudinary.com/documentation/upload_images
	//
	// Upload can be greatly customized by specifying uploader.UploadParams,
	// in this case we set the Public ID of the uploaded asset to "logo".
	uploadResult, err := cld.Upload.Upload(
		ctx,
		"./public/"+filename,
		uploader.UploadParams{PublicID: name})

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	log.Println(uploadResult.SecureURL)
	// Prints something like:
	// https://res.cloudinary.com/<your cloud name>/image/upload/v1615875158/logo.png

	// uploadResult contains useful information about the asset, like Width, Height, Format, etc.
	// See uploader.UploadResult struct for more details.

	// Now we can use Admin API to see the details about the asset.
	// The request can be customised by providing AssetParams.
	asset, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "logo"})
	if err != nil {
		log.Fatalf("Failed to get asset details, %v\n", err)
	}

	// Print some basic information about the asset.
	log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)
	// Cloudinary also provides a very flexible Search API for filtering and retrieving
	// information on all the assets in your account with the help of query expressions
	// in a Lucene-like query language.

	searchQuery := search.Query{
		Expression: "resource_type:image AND uploaded_at>1d AND bytes<1m",
		SortBy:     []search.SortByField{{"created_at": search.Descending}},
		MaxResults: 30,
	}

	searchResult, err := cld.Admin.Search(ctx, searchQuery)

	if err != nil {
		log.Fatalf("Failed to search for assets, %v\n", err)
	}

	log.Printf("Assets found: %v\n", searchResult.TotalCount)

	for _, asset := range searchResult.Assets {
		log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)
	}

	return uploadResult.SecureURL
}
