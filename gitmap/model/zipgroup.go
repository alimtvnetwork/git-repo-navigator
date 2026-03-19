// Package model defines the core data structures for gitmap.
package model

// ZipGroup represents a named collection of files/folders for archiving.
type ZipGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ArchiveName string `json:"archiveName"`
	CreatedAt   string `json:"createdAt"`
}

// ZipGroupItem links a file or folder path to a zip group.
type ZipGroupItem struct {
	GroupID  string `json:"groupId"`
	Path     string `json:"path"`
	IsFolder bool   `json:"isFolder"`
}
