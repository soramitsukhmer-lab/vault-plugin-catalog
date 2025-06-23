package utils

import (
	"fmt"
	"io"
	"os"
)

// Using os.Rename() with mounted volume in Docker or similar environments
// will result in a invalid cross-device link error.
func RenameCrossDevice(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file from %s to %s: %w", src, dst, err)
	}

	if err := dstFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync destination file %s: %w", dst, err)
	}

	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("failed to remove source file %s after copy: %w", src, err)
	}

	return nil
}
