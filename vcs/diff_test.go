	"sync"
	"sourcegraph.com/sourcegraph/go-vcs/vcs"
	t.Parallel()
			repo: makeGitRepositoryCmd(t, cmds...),
			repo: makeHgRepositoryCmd(t, hgCommands...),

		if _, err := test.repo.Diff(nonexistentCommitID, headCommitID, test.opt); err != vcs.ErrCommitNotFound {
			t.Errorf("%s: Diff with bad base commit ID: want ErrCommitNotFound, got %v", label, err)
			continue
		}

		if _, err := test.repo.Diff(baseCommitID, nonexistentCommitID, test.opt); err != vcs.ErrCommitNotFound {
			t.Errorf("%s: Diff with bad head commit ID: want ErrCommitNotFound, got %v", label, err)
			continue
		}
	t.Parallel()
	gitCmdsBase := []string{
		"echo line1 > f",
		"git add f",
		"GIT_COMMITTER_NAME=a GIT_COMMITTER_EMAIL=a@a.com GIT_COMMITTER_DATE=2006-01-02T15:04:05Z git commit -m foo --author='a <a@a.com>' --date 2006-01-02T15:04:05Z",
		"git tag testbase",
	}
	gitCmdsHead := []string{
			baseRepo: makeGitRepositoryCmd(t, gitCmdsBase...),
			headRepo: makeGitRepositoryCmd(t, gitCmdsHead...),
			base:     "testbase", head: "testhead",
			wantDiff: &vcs.Diff{
				Raw: "diff --git a/f b/f\nindex a29bdeb434d874c9b1d8969c40c42161b03fafdc..c0d0fb45c382919737f8d0c20aaf57cf89b74af8 100644\n--- a/f\n+++ b/f\n@@ -1 +1,2 @@\n line1\n+line2\n",
			},
		},
		"git libgit2": {
			baseRepo: makeGitRepositoryLibGit2(t, gitCmdsBase...),
			headRepo: makeGitRepositoryLibGit2(t, gitCmdsHead...),
	// TODO(sqs): implement diff for hg native
		// Try calling CrossRepoDiff a lot. The git impls do some
		// global state stuff (creating a new remote, fetching into
		// the base). See if this panics or segfaults (is libgit2
		// concurrent with respect to all of these operations?).
		const n = 100
		var wg sync.WaitGroup
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := test.baseRepo.CrossRepoDiff(baseCommitID, test.headRepo, headCommitID, test.opt)
				if err != nil {
					t.Errorf("%s: in concurrency test for CrossRepoDiff(%s, %v, %s, %v): %s", label, baseCommitID, test.headRepo, headCommitID, test.opt, err)
				}
			}()
		}
		wg.Wait()


		if _, err := test.baseRepo.CrossRepoDiff(nonexistentCommitID, test.headRepo, headCommitID, test.opt); err != vcs.ErrCommitNotFound {
			t.Errorf("%s: CrossRepoDiff with bad base commit ID: want ErrCommitNotFound, got %v", label, err)
			continue
		}

		if _, err := test.baseRepo.CrossRepoDiff(baseCommitID, test.headRepo, nonexistentCommitID, test.opt); err != vcs.ErrCommitNotFound {
			if label == "git cmd" {
				t.Log("skipping failure on git cmd because unimplemented")
				continue
			}
			t.Errorf("%s: CrossRepoDiff with bad head commit ID: want ErrCommitNotFound, got %v", label, err)
			continue
		}