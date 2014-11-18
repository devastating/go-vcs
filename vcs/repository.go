package vcs

import (
	"errors"
	"time"

	"golang.org/x/tools/godoc/vfs"
)

type Repository interface {
	ResolveRevision(spec string) (CommitID, error)
	ResolveTag(name string) (CommitID, error)
	ResolveBranch(name string) (CommitID, error)

	Branches() ([]*Branch, error)
	Tags() ([]*Tag, error)

	GetCommit(CommitID) (*Commit, error)

	// Commits returns all commits matching the options, as well as
	// the total number of commits (the count of which is not subject
	// to the N/Skip options).
	Commits(CommitsOptions) (commits []*Commit, total uint, err error)

	FileSystem(at CommitID) (vfs.FileSystem, error)
}

// A Blamer is a repository that can blame portions of a file.
type Blamer interface {
	BlameFile(path string, opt *BlameOptions) ([]*Hunk, error)
}

// BlameOptions configures a blame.
type BlameOptions struct {
	NewestCommit CommitID `json:",omitempty" url:",omitempty"`
	OldestCommit CommitID `json:",omitempty" url:",omitempty"` // or "" for the root commit

	StartLine int `json:",omitempty" url:",omitempty"` // 1-indexed start byte (or 0 for beginning of file)
	EndLine   int `json:",omitempty" url:",omitempty"` // 1-indexed end byte (or 0 for end of file)
}

// A Hunk is a contiguous portion of a file associated with a commit.
type Hunk struct {
	StartLine int // 1-indexed start line number
	EndLine   int // 1-indexed end line number
	StartByte int // 0-indexed start byte position (inclusive)
	EndByte   int // 0-indexed end byte position (exclusive)
	CommitID
	Author Signature
}

// A Differ is a repository that can compute diffs between two
// commits.
type Differ interface {
	// Diff shows changes between two commits.
	Diff(base, head CommitID, opt *DiffOptions) (*Diff, error)
}

// A CrossRepoDiffer is a repository that can compute diffs with
// respect to a commit in a different repository.
type CrossRepoDiffer interface {
	// CrossRepoDiff shows changes between two commits in different
	// repositories.
	CrossRepoDiff(base CommitID, headRepo Repository, head CommitID, opt *DiffOptions) (*Diff, error)
}

var (
	ErrBranchNotFound   = errors.New("branch not found")
	ErrCommitNotFound   = errors.New("commit not found")
	ErrRevisionNotFound = errors.New("revision not found")
	ErrTagNotFound      = errors.New("tag not found")
)

type CommitID string

type Commit struct {
	ID        CommitID
	Author    Signature
	Committer *Signature `json:",omitempty"`
	Message   string
	Parents   []CommitID `json:",omitempty"`
}

type Signature struct {
	Name  string
	Email string
	Date  time.Time
}

// CommitsOptions specifies limits on the list of commits returned by
// (Repository).Commits.
type CommitsOptions struct {
	Head CommitID // include all commits reachable from this commit (required)

	N    uint // limit the number of returned commits to this many (0 means no limit)
	Skip uint // skip this many commits at the beginning
}

// DiffOptions configures a diff.
type DiffOptions struct {
	Paths []string // constrain diff to these pathspecs
}

// A Diff represents changes between two commits.
type Diff struct {
	Raw string // the raw diff output
}

// A Branch is a VCS branch.
type Branch struct {
	Name string
	Head CommitID
}

type Branches []*Branch

func (p Branches) Len() int           { return len(p) }
func (p Branches) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p Branches) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// A Tag is a VCS tag.
type Tag struct {
	Name     string
	CommitID CommitID

	// TODO(sqs): A git tag can point to other tags, or really any
	// other object. How should we handle this case? For now, we're
	// just assuming they're all commit IDs.
}

type Tags []*Tag

func (p Tags) Len() int           { return len(p) }
func (p Tags) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p Tags) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
