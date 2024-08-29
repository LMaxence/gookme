package filters

import gitclient "github.com/LMaxence/gookme/packages/git-client"

type ChangesetResolvingStrategy interface {
	Resolve() ([]string, error)
}

// StagedChangesResolvingStrategy is a changeset resolving strategy that
// resolves the changeset by looking at the staged files in the repository.
type StagedChangesResolvingStrategy struct {
	repositoryPath string
}

func NewStagedChangesResolvingStrategy(repositoryPath string) *StagedChangesResolvingStrategy {
	return &StagedChangesResolvingStrategy{
		repositoryPath: repositoryPath,
	}
}

func (s *StagedChangesResolvingStrategy) Resolve() ([]string, error) {
	return gitclient.GetStagedFiles(&s.repositoryPath)
}

// FromToChangesResolvingStrategy is a changeset resolving strategy that
// resolves the changeset by looking at the changes between two refs.
type FromToChangesResolvingStrategy struct {
	repositoryPath string
	from           string
	to             string
}

func NewFromToChangesResolvingStrategy(repositoryPath string, from string, to string) *FromToChangesResolvingStrategy {
	return &FromToChangesResolvingStrategy{
		repositoryPath: repositoryPath,
		from:           from,
		to:             to,
	}
}

func (s *FromToChangesResolvingStrategy) Resolve() ([]string, error) {
	return gitclient.GetChangedFilesBetweenRefs(&s.repositoryPath, s.from, s.to)
}

// PrePushChangesResolvingStrategy is a changeset resolving strategy that
// resolves the changeset by looking at the files about to be pushed in the repository.
type PrePushChangesResolvingStrategy struct {
	repositoryPath string
}

func NewPrePushChangesResolvingStrategy(repositoryPath string) *PrePushChangesResolvingStrategy {
	return &PrePushChangesResolvingStrategy{
		repositoryPath: repositoryPath,
	}
}

func (s *PrePushChangesResolvingStrategy) Resolve() ([]string, error) {
	return gitclient.GetFilesToBePushed(&s.repositoryPath)
}

// NCommitsBeforeHeadChangesResolvingStrategy is a changeset resolving strategy that
// resolves the changeset by looking at the n commits before HEAD.

type NCommitsBeforeHeadChangesResolvingStrategy struct {
	repositoryPath string
	n              int
}

func NewNCommitsBeforeHeadChangesResolvingStrategy(repositoryPath string, n int) *NCommitsBeforeHeadChangesResolvingStrategy {
	return &NCommitsBeforeHeadChangesResolvingStrategy{
		repositoryPath: repositoryPath,
		n:              n,
	}
}

func (s *NCommitsBeforeHeadChangesResolvingStrategy) Resolve() ([]string, error) {
	return gitclient.GetFilesChangedNCommitsBefore(&s.repositoryPath, s.n)
}
