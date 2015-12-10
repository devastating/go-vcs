	gitCommands := []string{
			repo: makeGitRepositoryLibGit2(t, gitCommands...),
			repo: makeGitRepositoryCmd(t, gitCommands...),
			base: "testbase", head: "testhead",
			wantDiff: &vcs.Diff{
				Raw: "diff --git f f\nindex a29bdeb434d874c9b1d8969c40c42161b03fafdc..c0d0fb45c382919737f8d0c20aaf57cf89b74af8 100644\n--- f\n+++ f\n@@ -1 +1,2 @@\n line1\n+line2\n",
			},
		},
		"git go-git": {
			repo: makeGitRepositoryGoGit(t, gitCommands...),
	gitCommands := []string{
			repo: makeGitRepositoryLibGit2(t, gitCommands...),
			repo: makeGitRepositoryCmd(t, gitCommands...),
		"git go-git": {
			baseRepo: makeGitRepositoryGoGit(t, gitCmdsBase...),
			headRepo: makeGitRepositoryGoGit(t, gitCmdsHead...),
			base:     "testbase", head: "testhead",
			wantDiff: &vcs.Diff{
				Raw: "diff --git f f\nindex a29bdeb434d874c9b1d8969c40c42161b03fafdc..c0d0fb45c382919737f8d0c20aaf57cf89b74af8 100644\n--- f\n+++ f\n@@ -1 +1,2 @@\n line1\n+line2\n",
			},
		},