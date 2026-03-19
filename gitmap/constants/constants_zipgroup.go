package constants

// Zip group command names.
const (
	CmdZipGroup        = "zip-group"
	CmdZipGroupShort   = "z"
	SubCmdZGCreate     = "create"
	SubCmdZGAdd        = "add"
	SubCmdZGRemove     = "remove"
	SubCmdZGList       = "list"
	SubCmdZGShow       = "show"
	SubCmdZGDelete     = "delete"
	SubCmdZGRename     = "rename"
)

// Zip group table names.
const (
	TableZipGroups     = "ZipGroups"
	TableZipGroupItems = "ZipGroupItems"
)

// SQL: create ZipGroups table.
const SQLCreateZipGroups = `CREATE TABLE IF NOT EXISTS ZipGroups (
	Id          TEXT PRIMARY KEY,
	Name        TEXT NOT NULL UNIQUE,
	ArchiveName TEXT DEFAULT '',
	CreatedAt   TEXT DEFAULT CURRENT_TIMESTAMP
)`

// SQL: create ZipGroupItems table.
const SQLCreateZipGroupItems = `CREATE TABLE IF NOT EXISTS ZipGroupItems (
	GroupId  TEXT NOT NULL REFERENCES ZipGroups(Id) ON DELETE CASCADE,
	Path     TEXT NOT NULL,
	IsFolder INTEGER DEFAULT 0,
	PRIMARY KEY (GroupId, Path)
)`

// SQL: zip group operations.
const (
	SQLInsertZipGroup = `INSERT INTO ZipGroups (Id, Name, ArchiveName) VALUES (?, ?, ?)`

	SQLSelectAllZipGroups = `SELECT Id, Name, ArchiveName, CreatedAt FROM ZipGroups ORDER BY Name`

	SQLSelectZipGroupByName = `SELECT Id, Name, ArchiveName, CreatedAt FROM ZipGroups WHERE Name = ?`

	SQLDeleteZipGroup = `DELETE FROM ZipGroups WHERE Name = ?`

	SQLUpdateZipGroupArchive = `UPDATE ZipGroups SET ArchiveName = ? WHERE Name = ?`
)

// SQL: zip group item operations.
const (
	SQLInsertZipGroupItem = `INSERT OR IGNORE INTO ZipGroupItems (GroupId, Path, IsFolder) VALUES (?, ?, ?)`

	SQLDeleteZipGroupItem = `DELETE FROM ZipGroupItems WHERE GroupId = ? AND Path = ?`

	SQLSelectZipGroupItems = `SELECT GroupId, Path, IsFolder FROM ZipGroupItems WHERE GroupId = ? ORDER BY Path`

	SQLCountZipGroupItems = `SELECT COUNT(*) FROM ZipGroupItems WHERE GroupId = ?`

	SQLSelectAllZipGroupsWithCount = `SELECT g.Id, g.Name, g.ArchiveName, g.CreatedAt,
		(SELECT COUNT(*) FROM ZipGroupItems i WHERE i.GroupId = g.Id) AS ItemCount
		FROM ZipGroups g ORDER BY g.Name`
)

// SQL: drop zip group tables.
const (
	SQLDropZipGroups     = "DROP TABLE IF EXISTS ZipGroups"
	SQLDropZipGroupItems = "DROP TABLE IF EXISTS ZipGroupItems"
)

// Zip group flag descriptions.
const (
	FlagDescZGArchive  = "Custom output archive filename"
	FlagDescZGZipGroup = "Include a persistent zip group as a release asset"
	FlagDescZGZipItem  = "Add ad-hoc file or folder to zip as a release asset"
	FlagDescZGBundle   = "Bundle all -Z items into a single named archive"
)

// Zip group messages.
const (
	MsgZGCreated      = "  ✓ Created zip group %q\n"
	MsgZGDeleted      = "  ✓ Deleted zip group %q\n"
	MsgZGItemAdded    = "  ✓ Added %s to %q\n"
	MsgZGItemRemoved  = "  ✓ Removed %s from %q\n"
	MsgZGArchiveSet   = "  ✓ Archive name set to %q for group %q\n"
	MsgZGListHeader   = "\n  Zip Groups (%d):\n\n"
	MsgZGListRow      = "  %-20s %3d item(s)  %s\n"
	MsgZGShowHeader   = "\n  %s (%d item(s)):\n\n"
	MsgZGShowFile     = "    📄 %s\n"
	MsgZGShowFolder   = "    📁 %s\n"
	MsgZGShowArchive  = "  Archive: %s\n"
	MsgZGCompressed   = "  ✓ Compressed %s → %s\n"
	MsgZGDryRunHeader = "  [dry-run] Would create %d zip archive(s):\n"
	MsgZGDryRunEntry  = "    → %s (%d items: %s)\n"
	MsgZGSkipEmpty    = "  ⚠ Skipping empty group %q\n"
	MsgZGSkipMissing  = "  ⚠ Skipping missing item: %s\n"
)

// Zip group error messages.
const (
	ErrZGNotFound    = "no zip group found: %s"
	ErrZGEmpty       = "zip group name cannot be empty"
	ErrZGCreate      = "failed to create zip group: %v"
	ErrZGQuery       = "failed to query zip groups: %v"
	ErrZGDelete      = "failed to delete zip group: %v"
	ErrZGAddItem     = "failed to add item to zip group: %v"
	ErrZGRemoveItem  = "failed to remove item from zip group: %v"
	ErrZGCompress    = "  ✗ Failed to create archive for %s: %v\n"
	ErrZGGroupNotDB  = "zip group %q not found in database"
)
