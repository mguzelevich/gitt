package push

var testGitPushFixtures = []string{
	`Counting objects: 5, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (5/5), done.
Writing objects: 100% (5/5), 529 bytes | 0 bytes/s, done.
Total 5 (delta 4), reused 0 (delta 0)
To git@github.com:mguzelevich/gitt.git
   0fec14f..6cd668e  master -> master
`,
	`[master f55441b] repo restructurization
 9 files changed, 4 insertions(+), 4 deletions(-)
 rename tests_git_commit_fixtures.go => git/commit/commit_test_fixtures.go (88%)
 rename {pull => git/pull}/pull.go (100%)
 rename {pull => git/pull}/pull_test.go (100%)
 rename {pull => git/pull}/pull_test_fixtures.go (100%)
 rename tests_git_push_fixtures.go => git/push/push_test_fixtures.go (96%)
 rename {status => git/status}/status.go (100%)
 rename {status => git/status}/status_test.go (100%)
 rename {status => git/status}/status_test_fixtures.go (100%)
Counting objects: 8, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (7/7), done.
Writing objects: 100% (8/8), 1004 bytes | 0 bytes/s, done.
Total 8 (delta 2), reused 0 (delta 0)
To git@github.com:mguzelevich/gitt.git
   987083a..f55441b  master -> master
`,
}
