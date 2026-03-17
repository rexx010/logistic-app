package upload

import (
	"context"
	"fmt"
	"logisticApp/config"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var AllowedImageTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
}

const MaxFileSizeBytes = 5 * 1024 * 1024

type UploadResult struct {
	URL      string // the permanent HTTPS URL to the image
	PublicID string // Cloudinary's internal ID (used to delete/transform later)
}

func newCloudinaryClient() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(
		config.AppConfig.CloudinaryCloudName,
		config.AppConfig.CloudinaryAPIKey,
		config.AppConfig.CloudinaryAPISecret,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialized cloudinary client: %w", err)
	}
	return cld, nil
}

func UploadProfilePicture(file multipart.File, header *multipart.FileHeader, userID string) (*UploadResult, error) {
	//Validate file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !AllowedImageTypes[ext] {
		return nil, fmt.Errorf("file type %q is not allowed — use jpg, jpeg, png or webp", ext)
	}
	//Validate file size
	if header.Size > MaxFileSizeBytes {
		return nil, fmt.Errorf("file size %.2fMB exceeds the 5MB limit", float64(header.Size)/1024/1024)
	}
	//Upload to Cloudinary
	cld, err := newCloudinaryClient()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	publicID := fmt.Sprintf("logisticapp/profiles/%s", userID)

	response, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       publicID,
		Folder:         "logisticapp/profiles",
		Overwrite:      boolPtr(true),
		UniqueFilename: boolPtr(false),
		// Transformation applied at upload time:
		// crop to a square (fill) centered on the face
		// resize to 400x400px — enough for a profile picture
		// auto-format: Cloudinary picks the best format (webp for Chrome, etc.)
		// auto-quality: reduces file size without visible quality loss
		Transformation: "c_fill,g_face,h_400,w_400/f_auto/q_auto",
	})
	if err != nil {
		return nil, fmt.Errorf("cloudinary upload failed: %w", err)
	}
	return &UploadResult{
		URL:      response.SecureURL,
		PublicID: response.PublicID,
	}, nil
}

func UploadProofImage(file multipart.File, header *multipart.FileHeader, deliveryID string) (*UploadResult, error) {
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !AllowedImageTypes[ext] {
		return nil, fmt.Errorf("file type %q is not allowed — use jpg, jpeg, png or webp", ext)
	}
	if header.Size > MaxFileSizeBytes {
		return nil, fmt.Errorf("file size %.2fMB exceeds the 5MB limit", float64(header.Size)/1024/1024)
	}
	cld, err := newCloudinaryClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	// Each delivery has its own proof image.
	publicID := fmt.Sprintf("logisticapp/proofs/%s", deliveryID)

	response, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       publicID,
		Folder:         "logisticapp/proofs",
		Overwrite:      boolPtr(true),
		UniqueFilename: boolPtr(false),
		//uploading the full image
		Transformation: "f_auto/q_auto",
	})
	if err != nil {
		return nil, fmt.Errorf("cloudinary upload failed: %w", err)
	}
	return &UploadResult{
		URL:      response.SecureURL,
		PublicID: response.PublicID,
	}, nil
}

func DeleteImage(publicID string) error {
	cld, err := newCloudinaryClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}

func boolPtr(b bool) *bool {
	return &b
}
