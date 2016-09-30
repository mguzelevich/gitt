package pull

type gitOutputTest struct {
	input  string
	output string
}

var testGitPullFixtures = []gitOutputTest{
	gitOutputTest{`remote: Counting objects: 1, done.
remote: Total 1 (delta 0), reused 0 (delta 0)
Unpacking objects: 100% (1/1), done.
From github.com:mguzelevich/gitt
   8c5784e..ff4f990  master     -> origin/master
Updating 8c5784e..ff4f990
Fast-forward
 git.go                    |  30 +++++++---
 git_status_test.go        |  80 ++++++++++++++++++++++++++
 helpers.go                |  77 +++++++++++++------------
 main.go                   |   6 ++
 README.md                 | 282 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 test_git_pull_fixtures.go |  31 ++++++++++
 6 files changed, 461 insertions(+), 45 deletions(-)
 create mode 100644 README.md
 create mode 100644 test_git_pull_fixtures.go
 create mode 100644 git_status_test.go
`, `[github.com:mguzelevich/gitt (8c5784e..ff4f990)] 6*=461+/45-
	8c5784e..ff4f990 master -> origin/master
changes:
	+ README.md 282*=90+/0- (100644)
	+ git_status_test.go 80*=26+/0- (100644)
	+ test_git_pull_fixtures.go 31*=10+/0- (100644)
	* git.go 30*=7+/3-
	* helpers.go 77*=13+/12-
	* main.go 6*=2+/0-
`},
	gitOutputTest{`remote: Counting objects: 3, done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 3 (delta 1), reused 0 (delta 0)
Unpacking objects: 100% (3/3), done.
From github.com:mguzelevich/gitt
   6cd668e..20225f0  master     -> origin/master
Updating 6cd668e..20225f0
Fast-forward
 default.xml | 24 ++++++++++++++++++++++++
 1 file changed, 24 insertions(+)
 create mode 100644 default.xml
`, `[github.com:mguzelevich/gitt (6cd668e..20225f0)] 1*=24+/0-
	6cd668e..20225f0 master -> origin/master
changes:
	+ default.xml 24*=24+/0- (100644)
`},
	gitOutputTest{`Already up-to-date.
`, `up-to-date
`},
	gitOutputTest{`remote: Counting objects: 23, done.
remote: Compressing objects: 100% (21/21), done.
remote: Total 23 (delta 10), reused 2 (delta 0)
Unpacking objects: 100% (23/23), done.
From github.com:mguzelevich/gitt
   ff4f990..b601ef9  master     -> origin/master
 * [new branch]      BE-236     -> origin/BE-236
Updating ff4f990..b601ef9
Fast-forward
 a/a.go | 48 +++++++++++++++++++++++-------------------------
 b/b.go   |  7 +++++++
 c.go         |  2 ++
 d.go    | 15 +++++++--------
 4 files changed, 39 insertions(+), 33 deletions(-)
 create mode 100644 b/b.go
`, `[github.com:mguzelevich/gitt (ff4f990..b601ef9)] 4*=39+/33-
	ff4f990..b601ef9 master -> origin/master
	+ [new] BE-236 -> origin/BE-236
changes:
	+ b/b.go 7*=7+/0- (100644)
	* a/a.go 48*=23+/25-
	* c.go 2*=2+/0-
	* d.go 15*=7+/8-
`},
	gitOutputTest{`remote: Counting objects: 4, done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 4 (delta 1), reused 0 (delta 0)
Unpacking objects: 100% (4/4), done.
From github.com:mguzelevich/gitt
   c9c5372..03a39a5  master     -> origin/master
Updating c9c5372..03a39a5
Fast-forward
 .gitignore |  2 ++
 main.go    | 12 ++++++------
 2 files changed, 8 insertions(+), 6 deletions(-)
 create mode 100644 .gitignore
`, `[github.com:mguzelevich/gitt (c9c5372..03a39a5)] 2*=8+/6-
	c9c5372..03a39a5 master -> origin/master
changes:
	+ .gitignore 2*=2+/0- (100644)
	* main.go 12*=6+/6-
`},
	gitOutputTest{`remote: Counting objects: 9, done.
remote: Compressing objects: 100% (7/7), done.
remote: Total 9 (delta 5), reused 0 (delta 0)
Unpacking objects: 100% (9/9), done.
From github.com:mguzelevich/gitt
   ddeee21..e3bb5e0  master     -> origin/master
Updating ddeee21..e3bb5e0
Fast-forward
 .gitignore           |  2 ++
 etc/config.yaml      |  2 +-
 a.go | 19 +++++++++++--------
 b.go    | 23 +++++++++++++----------
 c.go | 22 +++++++++++-----------
 5 files changed, 38 insertions(+), 30 deletions(-)
 create mode 100644 .gitignore
`, `[github.com:mguzelevich/gitt (ddeee21..e3bb5e0)] 5*=38+/30-
	ddeee21..e3bb5e0 master -> origin/master
changes:
	+ .gitignore 2*=2+/0- (100644)
	* a.go 19*=11+/8-
	* b.go 23*=13+/10-
	* c.go 22*=11+/11-
	* etc/config.yaml 2*=1+/1-
`},
	gitOutputTest{`remote: Counting objects: 8, done.
remote: Compressing objects: 100% (7/7), done.
remote: Total 8 (delta 4), reused 0 (delta 0)
Unpacking objects: 100% (8/8), done.
From github.com:mguzelevich/gitt
   592f32d..f8c885f  master     -> origin/master
Updating 592f32d..f8c885f
Fast-forward
 .gitignore                                              |  2 ++
 main.go                                                 |  9 +--------
 a.go                                   | 36 ++++++++++++++++++++++++++++++++++++
 b.go => b_b.go                 | 16 ++++++++++------
 c.go                                    | 29 -----------------------------
 d.go => d_d.go | 18 +++++++++++-------
 6 files changed, 60 insertions(+), 50 deletions(-)
 create mode 100644 .gitignore
 create mode 100644 a.go
 rename b.go => b_b.go (56%)
 delete mode 100644 c.go
 rename d.go => d_d.go (57%)
`, `[github.com:mguzelevich/gitt (592f32d..f8c885f)] 6*=60+/50-
	592f32d..f8c885f master -> origin/master
changes:
	+ .gitignore 2*=2+/0- (100644)
	+ a.go 36*=36+/0- (100644)
	> b_b.go 16*=10+/6-
	> d_d.go 18*=11+/7-
	* main.go 9*=1+/8-
	- c.go 29*=0+/29- (100644)
`},
	// gitOutputTest{`From github.com:mguzelevich/gitt
	// 6057368..76ec6d6  master     -> origin/master
	// Auto-merging a/a.yaml
	// Merge made by the 'recursive' strategy.
	// a/a.yaml | 3 ++-
	// main.go  | 6 ++++--
	// 2 files changed, 6 insertions(+), 3 deletions(-)
	// `, ``},
	gitOutputTest{`remote: Counting objects: 22, done.
remote: Compressing objects: 100% (21/21), done.
remote: Total 22 (delta 12), reused 0 (delta 0)
Unpacking objects: 100% (22/22), done.
From github.com:mguzelevich/gitt
   a1b6ee1..ae2a6a9  master     -> origin/master
 * [new branch]      BE-244     -> origin/BE-244
Updating a1b6ee1..ae2a6a9
Fast-forward
 a/a.go          | 18 ++++--------------
 b/b.go    | 10 ++++------
 c/c.go      | 39 ++++++++++++++++++++++++++++++---------
 c/d.go |  2 +-
 4 files changed, 39 insertions(+), 30 deletions(-)
`, `[github.com:mguzelevich/gitt (a1b6ee1..ae2a6a9)] 4*=39+/30-
	a1b6ee1..ae2a6a9 master -> origin/master
	+ [new] BE-244 -> origin/BE-244
changes:
	* a/a.go 18*=4+/14-
	* b/b.go 10*=4+/6-
	* c/c.go 39*=30+/9-
	* c/d.go 2*=1+/1-
`},
	gitOutputTest{`remote: Counting objects: 11, done.
remote: Compressing objects: 100% (10/10), done.
remote: Total 11 (delta 8), reused 0 (delta 0)
Unpacking objects: 100% (11/11), done.
From github.com:mguzelevich/gitt
   b56d958..fc0935b  master     -> origin/master
Updating b56d958..fc0935b
Fast-forward
 .gitignore                    |  1 +
 main.go                       |  2 +-
 a.go     | 19 ++++++++++++-------
 b.go     | 22 +++++++++++++---------
 c.go         | 15 ++++++++++-----
 d.go            | 20 ++++++++++++--------
 e.go | 39 +++++++++++++--------------------------
 f.go         | 19 +++++++++++--------
 8 files changed, 73 insertions(+), 64 deletions(-)
`, `[github.com:mguzelevich/gitt (b56d958..fc0935b)] 8*=73+/64-
	b56d958..fc0935b master -> origin/master
changes:
	* .gitignore 1*=1+/0-
	* a.go 19*=12+/7-
	* b.go 22*=13+/9-
	* c.go 15*=10+/5-
	* d.go 20*=12+/8-
	* e.go 39*=13+/26-
	* f.go 19*=11+/8-
	* main.go 2*=1+/1-
`},
	gitOutputTest{`remote: Counting objects: 12, done.
remote: Compressing objects: 100% (9/9), done.
remote: Total 12 (delta 6), reused 0 (delta 0)
Unpacking objects: 100% (12/12), done.
From github.com:mguzelevich/gitt
   a6b99ee..64bdbcb  master     -> origin/master
Updating a6b99ee..64bdbcb
Fast-forward
 .gitignore       |  2 ++
 env/schema.sql   | 13 -------------
 etc/config.yaml  |  9 +++++----
 main.go          |  8 ++++++--
 a.go | 83 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 b.go | 44 ++++++++++++++++++++++++++++++++++++++++++++
 sql/001_init.sql | 16 ++++++++++++++++
 7 files changed, 156 insertions(+), 19 deletions(-)
 create mode 100644 .gitignore
 delete mode 100644 env/schema.sql
 create mode 100644 a.go
 create mode 100644 b.go
 create mode 100644 sql/001_init.sql
`, `[github.com:mguzelevich/gitt (a6b99ee..64bdbcb)] 7*=156+/19-
	a6b99ee..64bdbcb master -> origin/master
changes:
	+ .gitignore 2*=2+/0- (100644)
	+ a.go 83*=83+/0- (100644)
	+ b.go 44*=44+/0- (100644)
	+ sql/001_init.sql 16*=16+/0- (100644)
	* etc/config.yaml 9*=5+/4-
	* main.go 8*=6+/2-
	- env/schema.sql 13*=0+/13- (100644)
`},
	gitOutputTest{`remote: Counting objects: 22, done.
remote: Compressing objects: 100% (19/19), done.
remote: Total 22 (delta 8), reused 0 (delta 0)
Unpacking objects: 100% (22/22), done.
From github.com:mguzelevich/gitt
 * [new branch]      master     -> origin/master
`, `[github.com:mguzelevich/gitt ()] 0*=0+/0-
	+ [new] master -> origin/master
`},
	gitOutputTest{`remote: Counting objects: 12, done.
remote: Compressing objects: 100% (11/11), done.
remote: Total 12 (delta 7), reused 0 (delta 0)
Unpacking objects: 100% (12/12), done.
From github.com:mguzelevich/gitt
   15b0622..5c13243  master     -> origin/master
Updating 15b0622..5c13243
Fast-forward
 .gitignore        |  1 +
 a.go  | 54 ++++++++++++++++++++++++++++++++++++++++++++++++++++++
 b.go | 77 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 c.go    | 52 ++++++++++++++++++++++++++++++++++++++++++++++++++++
 d.go   | 53 -----------------------------------------------------
 e.go  | 73 -------------------------------------------------------------------------
 f.go     | 47 -----------------------------------------------
 7 files changed, 184 insertions(+), 173 deletions(-)
 create mode 100644 a.go
 create mode 100644 b.go
 create mode 100644 c.go
 delete mode 100644 d.go
 delete mode 100644 e.go
 delete mode 100644 f.go
`, `[github.com:mguzelevich/gitt (15b0622..5c13243)] 7*=184+/173-
	15b0622..5c13243 master -> origin/master
changes:
	+ a.go 54*=54+/0- (100644)
	+ b.go 77*=77+/0- (100644)
	+ c.go 52*=52+/0- (100644)
	* .gitignore 1*=1+/0-
	- d.go 53*=0+/53- (100644)
	- e.go 73*=0+/73- (100644)
	- f.go 47*=0+/47- (100644)
`},
}
