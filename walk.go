package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

type Reference struct {
	Title       string
	Authors     []string
	Editors     []string
	Description string
	Publisher   string
	Year        string
	Filename    string
	Path        string
	Mimetype    string
	AddedAt     time.Time
	Tags        []string
}

func (r *Reference) dump() {
	fmt.Printf("Title: %s\n", r.Title)
	fmt.Printf("Filename: %s\n", r.Filename)
	fmt.Printf("Path: %s\n", r.Path)
	fmt.Printf("Mimetype: %s\n", r.Mimetype)
	fmt.Printf("AddedAt: %s\n", r.AddedAt)
}

func titleFromFilename(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func creationTimeFromInfo(info os.FileInfo) time.Time {
	if sys, ok := info.Sys().(*syscall.Stat_t); ok {
		return time.Unix(int64(sys.Ctimespec.Sec), int64(sys.Ctimespec.Nsec))
	}
	return info.ModTime()
}

func createReference(path string, info os.FileInfo) Reference {
	filename := filepath.Base(path)
	return Reference{
		Title:    titleFromFilename(filename),
		Filename: filename,
		Path:     path,
		Mimetype: mime.TypeByExtension(filepath.Ext(filename)),
		AddedAt:  creationTimeFromInfo(info),
	}
}

func main() {
	startDirectory := "/Users/jamie/Downloads" // you can change this to any other path

	err := filepath.Walk(startDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ref := createReference(path, info)
		ref.dump()

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}
}
