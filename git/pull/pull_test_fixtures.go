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
	gitOutputTest{`From github.com:mguzelevich/gitt
   6057368..76ec6d6  master     -> origin/master
Auto-merging a/a.yaml
Merge made by the 'recursive' strategy.
 a/a.yaml | 3 ++-
 main.go  | 6 ++++--
 2 files changed, 6 insertions(+), 3 deletions(-)
`, ``},
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
	gitOutputTest{`From https://github.com/mapsme/omim
   6299e3b..1444951  master     -> origin/master
 + 7df9017...35887d4 drape-2.0  -> origin/drape-2.0  (forced update)
 * [new branch]      drape-2.0-map-style-editor -> origin/drape-2.0-map-style-editor
 + 2695f1d...fc1693b map-style-editor-new -> origin/map-style-editor-new  (forced update)
   625a4b6..1d93371  new-search -> origin/new-search
Fetching submodule tools/kothic
From git://github.com/mapsme/kothic
   60c342c..b8d9357  master     -> origin/master
Fetching submodule tools/osmctools
From git://github.com/mapsme/osmctools
   6da67ed..96953b5  master     -> origin/master
Updating 6299e3b..1444951
Fast-forward
 .gitignore                                         |     1 +
 .../server/data_structures/datafacade_base.hpp     |     2 +-
 .../server/data_structures/internal_datafacade.hpp |     2 +-
 .../server/data_structures/shared_datafacade.hpp   |     2 +-
 .../routing_test_tools.hpp                         |     0
 .../clear/style-clear/symbols/cemetery-l.svg       |    14 +
 .../MWMMapViewControlsCommon.h                     |     0
 .../MapViewControls/MWMMapViewControlsCommon.h     |    11 +
 docs/MWM.md                                        |    14 +
 CONTRIBUTORS                                       |     2 +-
 README.md                                          |     2 +-
 android/UnitTests/build.gradle                     |     2 +-
 configure.sh                                       |     7 +-
 data/categories.txt                                |     2 +-
 data/drules_proto.bin                              |   Bin 201238 -> 205817 bytes
 .../ic_follow_mode_1_light.png                     |   Bin 0 -> 477 bytes
 .../follow_to_followandrotate_1.png                |   Bin 1569 -> 0 bytes
 1320 files changed, 33610 insertions(+), 28746 deletions(-)
 create mode 100644 data/styles/clear/style-clear/symbols/cemetery-l.svg
 delete mode 100644 docs/MWM.md
 rename iphone/Maps/Classes/CustomViews/MapViewControls/{SideMenu => APIBar}/MWMMapViewControlsCommon.h (100%)
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_1.imageset/ic_follow_mode_1_light.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_1.imageset/follow_to_followandrotate_1.png
 rename {integration_tests => routing/routing_integration_tests}/routing_test_tools.hpp (100%)
`, ``},

	gitOutputTest{`From https://github.com/nats-io/gnatsd
   139ce42..ea9216a  master     -> origin/master
 * [new branch]      temporarily-reverted-INFO-message-format -> origin/temporarily-reverted-INFO-message-format
 * [new tag]         v0.6.8     -> v0.6.8
Updating 139ce42..ea9216a
Fast-forward
 .travis.yml             |  3 ++-
 Dockerfile              |  3 +--
 README.md               | 19 +++++++++++-----
 gnatsd.go               |  6 +++--
 server/client.go        | 18 ++++++++-------
 server/client_test.go   |  2 +-
 server/const.go         |  6 ++++-
 server/monitor.go       |  2 ++
 server/opts.go          |  6 ++++-
 server/opts_test.go     |  1 +
 server/route.go         | 13 ++++++++++-
 server/server.go        | 27 +++++++++++++----------
 test/maxpayload_test.go | 58 +++++++++++++++++++++++++++++++++++++++----------
 test/pedantic_test.go   | 10 +++++----
 test/test.go            |  4 ++--
 15 files changed, 127 insertions(+), 51 deletions(-)
`, `[https://github.com/nats-io/gnatsd (139ce42..ea9216a)] 15*=127+/51-
	139ce42..ea9216a master -> origin/master
	+ [new] temporarily-reverted-INFO-message-format -> origin/temporarily-reverted-INFO-message-format
	+ [new] v0.6.8 -> v0.6.8
changes:
	* .travis.yml 3*=2+/1-
	* Dockerfile 3*=1+/2-
	* README.md 19*=11+/5-
	* gnatsd.go 6*=3+/2-
	* server/client.go 18*=8+/7-
	* server/client_test.go 2*=1+/1-
	* server/const.go 6*=4+/1-
	* server/monitor.go 2*=2+/0-
	* server/opts.go 6*=4+/1-
	* server/opts_test.go 1*=1+/0-
	* server/route.go 13*=10+/1-
	* server/server.go 27*=13+/10-
	* test/maxpayload_test.go 58*=39+/10-
	* test/pedantic_test.go 10*=5+/4-
	* test/test.go 4*=2+/2-
`},

	gitOutputTest{`From https://github.com/mapsme/omim
   6299e3b..1444951  master     -> origin/master
 + 7df9017...35887d4 drape-2.0  -> origin/drape-2.0  (forced update)
 * [new branch]      drape-2.0-map-style-editor -> origin/drape-2.0-map-style-editor
 + 2695f1d...fc1693b map-style-editor-new -> origin/map-style-editor-new  (forced update)
   625a4b6..1d93371  new-search -> origin/new-search
Fetching submodule tools/kothic
From git://github.com/mapsme/kothic
   60c342c..b8d9357  master     -> origin/master
Fetching submodule tools/osmctools
From git://github.com/mapsme/osmctools
   6da67ed..96953b5  master     -> origin/master
Updating 6299e3b..1444951
Fast-forward
 .gitignore                                         |     1 +
 .../server/data_structures/datafacade_base.hpp     |     2 +-
 .../server/data_structures/internal_datafacade.hpp |     2 +-
 .../server/data_structures/shared_datafacade.hpp   |     2 +-
 CONTRIBUTORS                                       |     2 +-
 README.md                                          |     2 +-
 android/UnitTests/build.gradle                     |     2 +-
 android/UnitTests/jni/Android.mk                   |     6 +-
 android/UnitTests/jni/AndroidBeginning.mk          |     4 +-
 android/UnitTests/jni/main.cpp                     |     4 +-
 android/UnitTests/jni/mock.cpp                     |     6 +
 .../maps/unittests/AllTestsActivity.java           |     2 +-
 android/UnitTests/tests_list.sh                    |     2 +-
 android/res/values-ja/strings.xml                  |     2 +-
 .../src/com/mapswithme/maps/sound/TtsPlayer.java   |     2 +
 common.pri                                         |     2 +-
 configure.sh                                       |     7 +-
 data/categories.txt                                |     2 +-
 data/classificator.txt                             |     4 +
 data/drules_proto.bin                              |   Bin 201238 -> 205817 bytes
 data/drules_proto.txt                              |   557 +
 data/drules_proto_clear.bin                        |   Bin 106535 -> 114397 bytes
 data/drules_proto_clear.txt                        | 10930 +++++++++++--------
 data/drules_proto_dark.bin                         |   Bin 104952 -> 112694 bytes
 data/drules_proto_dark.txt                         | 10744 +++++++++++-------
 data/mapcss-dynamic.txt                            |     2 +
 data/mapcss-mapping.csv                            |     4 +
 data/resources-6plus/basic.skn                     |   390 +-
 data/resources-6plus/symbols.png                   |   Bin 165749 -> 181259 bytes
 data/resources-6plus/symbols.sdf                   |   322 +-
 data/resources-6plus_clear/basic.skn               |  1153 +-
 data/resources-6plus_clear/symbols.png             |   Bin 267378 -> 299952 bytes
 data/resources-6plus_clear/symbols.sdf             |   583 +-
 data/resources-6plus_dark/basic.skn                |  1389 +--
 data/resources-6plus_dark/symbols.png              |   Bin 263781 -> 268286 bytes
 data/resources-6plus_dark/symbols.sdf              |   679 +-
 data/resources-hdpi_clear/basic.skn                |  1197 +-
 data/resources-hdpi_clear/symbols.png              |   Bin 132368 -> 138252 bytes
 data/resources-hdpi_clear/symbols.sdf              |   557 +-
 data/resources-hdpi_dark/basic.skn                 |  1389 +--
 data/resources-hdpi_dark/symbols.png               |   Bin 131644 -> 126025 bytes
 data/resources-hdpi_dark/symbols.sdf               |   645 +-
 data/resources-ldpi_clear/basic.skn                |  1147 +-
 data/resources-ldpi_clear/symbols.png              |   Bin 81080 -> 85634 bytes
 data/resources-ldpi_clear/symbols.sdf              |   555 +-
 data/resources-ldpi_dark/basic.skn                 |  1345 +--
 data/resources-ldpi_dark/symbols.png               |   Bin 81886 -> 79597 bytes
 data/resources-ldpi_dark/symbols.sdf               |   647 +-
 data/resources-mdpi_clear/basic.skn                |  1091 +-
 data/resources-mdpi_clear/symbols.png              |   Bin 87583 -> 91399 bytes
 data/resources-mdpi_clear/symbols.sdf              |   479 +-
 data/resources-mdpi_dark/basic.skn                 |  1241 ++-
 data/resources-mdpi_dark/symbols.png               |   Bin 87544 -> 86786 bytes
 data/resources-mdpi_dark/symbols.sdf               |   695 +-
 data/resources-xhdpi_clear/basic.skn               |  1143 +-
 data/resources-xhdpi_clear/symbols.png             |   Bin 231867 -> 241616 bytes
 data/resources-xhdpi_clear/symbols.sdf             |   537 +-
 data/resources-xhdpi_dark/basic.skn                |  1317 +--
 data/resources-xhdpi_dark/symbols.png              |   Bin 229192 -> 225870 bytes
 data/resources-xhdpi_dark/symbols.sdf              |   697 +-
 data/resources-xxhdpi_clear/basic.skn              |   949 +-
 data/resources-xxhdpi_clear/symbols.png            |   Bin 312260 -> 327380 bytes
 data/resources-xxhdpi_clear/symbols.sdf            |   599 +-
 data/resources-xxhdpi_dark/basic.skn               |  1311 +--
 data/resources-xxhdpi_dark/symbols.png             |   Bin 309741 -> 300462 bytes
 data/resources-xxhdpi_dark/symbols.sdf             |   611 +-
 data/styles/clear/include/POI.mapcss               |    71 +-
 data/styles/clear/include/labels.mapcss            |    68 +-
 data/styles/clear/include/landuse.mapcss           |    27 +-
 data/styles/clear/include/natural.mapcss           |    17 +-
 data/styles/clear/include/other.mapcss             |     2 +
 data/styles/clear/include/roads.mapcss             |    53 +-
 data/styles/clear/style-clear/colors.mapcss        |     7 +-
 .../clear/style-clear/symbols/cemetery-l.svg       |    14 +
 .../clear/style-clear/symbols/cemetery-m.svg       |    14 +
 .../clear/style-clear/symbols/cemetery-s.svg       |    14 +
 .../clear/style-clear/symbols/drinking-water-l.svg |     4 +-
 .../clear/style-clear/symbols/drinking-water-m.svg |     4 +-
 .../clear/style-clear/symbols/drinking-water-s.svg |     4 +-
 .../clear/style-clear/symbols/fountain-l.svg       |    18 +-
 .../clear/style-clear/symbols/fountain-m.svg       |    14 +-
 .../clear/style-clear/symbols/fountain-s.svg       |    16 +-
 .../clear/style-clear/symbols/military-l.svg       |    19 +
 .../clear/style-clear/symbols/military-m.svg       |    19 +
 .../clear/style-clear/symbols/military-s.svg       |    19 +
 data/styles/clear/style-clear/symbols/npark-l.svg  |     7 +-
 data/styles/clear/style-clear/symbols/npark-m.svg  |     7 +-
 data/styles/clear/style-clear/symbols/npark-s.svg  |     7 +-
 data/styles/clear/style-clear/symbols/nparkF-l.svg |    16 +
 data/styles/clear/style-clear/symbols/nparkF-m.svg |    16 +
 data/styles/clear/style-clear/symbols/nparkF-s.svg |    16 +
 .../clear/style-clear/symbols/waterfall-l.svg      |    19 +-
 .../clear/style-clear/symbols/waterfall-m.svg      |    21 +-
 .../clear/style-clear/symbols/waterfall-s.svg      |    20 +-
 data/styles/clear/style-night/colors.mapcss        |     3 +-
 .../clear/style-night/symbols/alcohol-l copy.svg   |     4 +-
 .../clear/style-night/symbols/alcohol-m copy.svg   |     4 +-
 .../clear/style-night/symbols/alcohol-s copy.svg   |     4 +-
 .../symbols/america-football-l copy.svg            |     4 +-
 .../symbols/america-football-m copy.svg            |     4 +-
 data/styles/clear/style-night/symbols/atm-l.svg    |     4 +-
 data/styles/clear/style-night/symbols/atm-m.svg    |     4 +-
 .../clear/style-night/symbols/atm-s copy 2.svg     |     4 +-
 data/styles/clear/style-night/symbols/bank-l.svg   |     4 +-
 data/styles/clear/style-night/symbols/bank-m.svg   |     4 +-
 .../clear/style-night/symbols/bank-s copy 5.svg    |     4 +-
 .../clear/style-night/symbols/banknote-l.svg       |     8 +-
 .../clear/style-night/symbols/banknote-m.svg       |     8 +-
 data/styles/clear/style-night/symbols/bar-l.svg    |     4 +-
 data/styles/clear/style-night/symbols/bar-m.svg    |     4 +-
 data/styles/clear/style-night/symbols/bar-s.svg    |     4 +-
 .../clear/style-night/symbols/baseball-l copy.svg  |     4 +-
 .../clear/style-night/symbols/baseball-m copy.svg  |     4 +-
 .../clear/style-night/symbols/baseball-s copy.svg  |     4 +-
 .../style-night/symbols/basketball-l copy.svg      |     4 +-
 .../style-night/symbols/basketball-m copy.svg      |     4 +-
 .../style-night/symbols/basketball-s copy.svg      |     4 +-
 .../styles/clear/style-night/symbols/bicycle-l.svg |     4 +-
 .../styles/clear/style-night/symbols/bicycle-m.svg |     4 +-
 .../styles/clear/style-night/symbols/bicycle-s.svg |     4 +-
 data/styles/clear/style-night/symbols/cafe-l.svg   |     4 +-
 data/styles/clear/style-night/symbols/cafe-m.svg   |     4 +-
 data/styles/clear/style-night/symbols/cafe-s.svg   |     4 +-
 .../clear/style-night/symbols/campsite-l.svg       |     4 +-
 .../clear/style-night/symbols/campsite-m.svg       |     4 +-
 .../clear/style-night/symbols/campsite-s.svg       |     4 +-
 data/styles/clear/style-night/symbols/cave-l.svg   |     4 +-
 data/styles/clear/style-night/symbols/cave-m.svg   |     4 +-
 data/styles/clear/style-night/symbols/cave-s.svg   |     4 +-
 .../clear/style-night/symbols/cemetery-l.svg       |    14 +
 .../clear/style-night/symbols/cemetery-m.svg       |    14 +
 .../clear/style-night/symbols/cemetery-s.svg       |    14 +
 data/styles/clear/style-night/symbols/cinema-l.svg |     4 +-
 data/styles/clear/style-night/symbols/cinema-m.svg |     4 +-
 data/styles/clear/style-night/symbols/cinema-s.svg |     4 +-
 .../styles/clear/style-night/symbols/college-l.svg |     4 +-
 .../styles/clear/style-night/symbols/college-m.svg |     4 +-
 .../styles/clear/style-night/symbols/college-s.svg |     4 +-
 .../clear/style-night/symbols/drinking-water-l.svg |     6 +-
 .../clear/style-night/symbols/drinking-water-m.svg |     6 +-
 .../clear/style-night/symbols/drinking-water-s.svg |     6 +-
 .../styles/clear/style-night/symbols/embassy-l.svg |     4 +-
 .../styles/clear/style-night/symbols/embassy-m.svg |     4 +-
 .../styles/clear/style-night/symbols/embassy-s.svg |     4 +-
 .../clear/style-night/symbols/fastfood-l.svg       |     4 +-
 .../clear/style-night/symbols/fastfood-m.svg       |     4 +-
 .../clear/style-night/symbols/fastfood-s.svg       |     4 +-
 .../clear/style-night/symbols/fountain-l.svg       |     9 +-
 .../clear/style-night/symbols/fountain-m.svg       |     9 +-
 .../clear/style-night/symbols/fountain-s.svg       |     9 +-
 .../clear/style-night/symbols/fuel-l copy.svg      |     4 +-
 .../clear/style-night/symbols/fuel-m copy.svg      |     4 +-
 .../clear/style-night/symbols/fuel-s copy.svg      |     4 +-
 .../styles/clear/style-night/symbols/gallery-l.svg |     4 +-
 .../styles/clear/style-night/symbols/gallery-m.svg |     4 +-
 .../styles/clear/style-night/symbols/gallery-s.svg |     4 +-
 .../clear/style-night/symbols/golf-l copy.svg      |     4 +-
 .../clear/style-night/symbols/golf-m copy.svg      |     4 +-
 .../clear/style-night/symbols/golf-s copy.svg      |     4 +-
 .../styles/clear/style-night/symbols/grocery-l.svg |     4 +-
 .../styles/clear/style-night/symbols/grocery-m.svg |     4 +-
 .../styles/clear/style-night/symbols/grocery-s.svg |     4 +-
 .../clear/style-night/symbols/hospital-l.svg       |     4 +-
 .../clear/style-night/symbols/hospital-m.svg       |     4 +-
 .../clear/style-night/symbols/hospital-s.svg       |     4 +-
 data/styles/clear/style-night/symbols/hotel-l.svg  |     4 +-
 data/styles/clear/style-night/symbols/hotel-m.svg  |     4 +-
 data/styles/clear/style-night/symbols/hotel-s.svg  |     4 +-
 .../style-night/symbols/information-l copy.svg     |     4 +-
 .../clear/style-night/symbols/information-m.svg    |     4 +-
 .../style-night/symbols/information-s copy.svg     |     4 +-
 .../clear/style-night/symbols/kindergarten-l.svg   |     4 +-
 .../clear/style-night/symbols/kindergarten-m.svg   |     4 +-
 .../clear/style-night/symbols/kindergarten-s.svg   |     4 +-
 .../styles/clear/style-night/symbols/library-l.svg |     4 +-
 .../styles/clear/style-night/symbols/library-m.svg |     4 +-
 .../styles/clear/style-night/symbols/library-s.svg |     4 +-
 .../clear/style-night/symbols/lighthouse-l.svg     |     4 +-
 .../clear/style-night/symbols/lighthouse-m.svg     |     4 +-
 .../clear/style-night/symbols/lighthouse-s.svg     |     4 +-
 data/styles/clear/style-night/symbols/mail-l.svg   |     4 +-
 data/styles/clear/style-night/symbols/mail-m.svg   |     4 +-
 data/styles/clear/style-night/symbols/mail-s.svg   |     4 +-
 .../clear/style-night/symbols/military-l.svg       |    14 +
 .../clear/style-night/symbols/military-m.svg       |    14 +
 .../clear/style-night/symbols/military-s.svg       |    14 +
 .../clear/style-night/symbols/monument-l.svg       |     6 +-
 .../clear/style-night/symbols/monument-m.svg       |     6 +-
 .../clear/style-night/symbols/monument-s.svg       |     6 +-
 data/styles/clear/style-night/symbols/museum-l.svg |     4 +-
 data/styles/clear/style-night/symbols/museum-m.svg |     4 +-
 data/styles/clear/style-night/symbols/museum-s.svg |     4 +-
 data/styles/clear/style-night/symbols/nparkF-l.svg |    14 +
 data/styles/clear/style-night/symbols/nparkF-m.svg |    14 +
 data/styles/clear/style-night/symbols/nparkF-s.svg |    14 +
 data/styles/clear/style-night/symbols/peak-l.svg   |     4 +-
 data/styles/clear/style-night/symbols/peak-m.svg   |     4 +-
 data/styles/clear/style-night/symbols/peak-s.svg   |     4 +-
 .../clear/style-night/symbols/pharmacy-l.svg       |     4 +-
 .../clear/style-night/symbols/pharmacy-m.svg       |     4 +-
 data/styles/clear/style-night/symbols/phone-l.svg  |     4 +-
 data/styles/clear/style-night/symbols/phone-m.svg  |     4 +-
 data/styles/clear/style-night/symbols/phone-s.svg  |     4 +-
 .../clear/style-night/symbols/pitch-l copy.svg     |     4 +-
 .../clear/style-night/symbols/pitch-m copy.svg     |     4 +-
 .../clear/style-night/symbols/pitch-s copy.svg     |     4 +-
 .../symbols/place-of-worship-l copy.svg            |     4 +-
 .../symbols/place-of-worship-m copy.svg            |     4 +-
 .../symbols/place-of-worship-s copy.svg            |     4 +-
 data/styles/clear/style-night/symbols/police-l.svg |     4 +-
 data/styles/clear/style-night/symbols/police-m.svg |     4 +-
 data/styles/clear/style-night/symbols/police-s.svg |     4 +-
 data/styles/clear/style-night/symbols/pub-l.svg    |     4 +-
 data/styles/clear/style-night/symbols/pub-m.svg    |     4 +-
 data/styles/clear/style-night/symbols/pub-s.svg    |     4 +-
 .../clear/style-night/symbols/recycling-l copy.svg |     4 +-
 .../clear/style-night/symbols/recycling-m copy.svg |     4 +-
 .../clear/style-night/symbols/recycling-s copy.svg |     4 +-
 .../clear/style-night/symbols/restaurant-l.svg     |     4 +-
 .../clear/style-night/symbols/restaurant-m.svg     |     4 +-
 .../clear/style-night/symbols/restaurant-s.svg     |     4 +-
 data/styles/clear/style-night/symbols/school-l.svg |     4 +-
 data/styles/clear/style-night/symbols/school-m.svg |     4 +-
 data/styles/clear/style-night/symbols/school-s.svg |     4 +-
 .../clear/style-night/symbols/shop-l copy.svg      |     4 +-
 .../clear/style-night/symbols/shop-m copy.svg      |     4 +-
 .../clear/style-night/symbols/shop-s copy.svg      |     4 +-
 .../clear/style-night/symbols/skiing-l copy.svg    |     4 +-
 .../clear/style-night/symbols/skiing-m copy.svg    |     4 +-
 .../clear/style-night/symbols/skiing-s copy.svg    |     4 +-
 .../clear/style-night/symbols/soccer-l copy.svg    |     4 +-
 .../clear/style-night/symbols/soccer-m copy.svg    |     4 +-
 .../clear/style-night/symbols/soccer-s copy.svg    |     4 +-
 .../clear/style-night/symbols/swimming-l copy.svg  |     4 +-
 .../clear/style-night/symbols/swimming-m copy.svg  |     4 +-
 .../clear/style-night/symbols/swimming-s copy.svg  |     4 +-
 .../clear/style-night/symbols/tennis-l copy.svg    |     4 +-
 .../clear/style-night/symbols/tennis-m copy.svg    |     4 +-
 .../clear/style-night/symbols/tennis-s copy.svg    |     4 +-
 .../styles/clear/style-night/symbols/theatre-l.svg |     4 +-
 .../styles/clear/style-night/symbols/theatre-m.svg |     4 +-
 .../styles/clear/style-night/symbols/theatre-s.svg |     4 +-
 .../clear/style-night/symbols/toilets-l copy.svg   |     4 +-
 .../clear/style-night/symbols/toilets-m copy.svg   |     4 +-
 .../clear/style-night/symbols/toilets-s copy.svg   |     4 +-
 .../styles/clear/style-night/symbols/tourism-l.svg |     4 +-
 .../styles/clear/style-night/symbols/tourism-m.svg |     4 +-
 .../styles/clear/style-night/symbols/tourism-s.svg |     4 +-
 .../clear/style-night/symbols/townhall-l.svg       |     4 +-
 .../clear/style-night/symbols/townhall-m.svg       |     4 +-
 .../clear/style-night/symbols/townhall-s.svg       |     4 +-
 .../styles/clear/style-night/symbols/volcano-l.svg |     8 +-
 .../styles/clear/style-night/symbols/volcano-m.svg |     8 +-
 .../styles/clear/style-night/symbols/volcano-s.svg |     8 +-
 .../clear/style-night/symbols/waterfall-l.svg      |    21 +-
 .../clear/style-night/symbols/waterfall-m.svg      |    21 +-
 .../clear/style-night/symbols/waterfall-s.svg      |    20 +-
 data/styles/clear/style-night/symbols/zoo-l.svg    |     4 +-
 data/styles/clear/style-night/symbols/zoo-m.svg    |     4 +-
 data/styles/clear/style-night/symbols/zoo-s.svg    |     4 +-
 data/styles/legacy/include/base_other.mapcss       |    10 +-
 data/styles/legacy/style-yota/style.mapcss         |     4 +-
 data/types.txt                                     |     4 +
 data/visibility.txt                                |     4 +
 docs/INSTALL.md                                    |    21 +-
 docs/MAPS.md                                       |   151 +
 docs/MWM.md                                        |    11 -
 generator/generator_tests/osm_type_test.cpp        |    55 +
 generator/osm2meta.hpp                             |    14 +
 generator/osm2type.cpp                             |    23 +-
 generator/osm_translator.hpp                       |     5 +-
 generator/tesselator.cpp                           |    77 +-
 generator/tesselator.hpp                           |    14 +-
 indexer/drules_selector.cpp                        |    46 +-
 indexer/drules_selector_parser.cpp                 |     4 +-
 indexer/feature_meta.hpp                           |     1 +
 indexer/ftypes_matcher.cpp                         |    13 +
 indexer/ftypes_matcher.hpp                         |     8 +
 .../indexer_tests/drules_selector_parser_test.cpp  |    40 +
 indexer/indexer_tests/jni/Android.mk               |     2 +-
 integration_tests/integration_tests.pro            |    31 -
 integration_tests/jni/Android.mk                   |    15 -
 integration_tests/jni/Application.mk               |     1 -
 integration_tests/jni/test.cpp                     |    15 -
 integration_tests/jni/test.hpp                     |     9 -
 integration_tests/online_cross_tests.cpp           |    49 -
 integration_tests/osrm_route_test.cpp              |   237 -
 integration_tests/osrm_turn_test.cpp               |   350 -
 integration_tests/pedestrian_route_test.cpp        |   509 -
 integration_tests/routing_test_tools.cpp           |   342 -
 .../MWMMapViewControlsCommon.h                     |     0
 .../BottomMenu/MWMBottomMenuCollectionViewCell.h   |    11 +
 .../BottomMenu/MWMBottomMenuCollectionViewCell.mm  |    48 +
 .../MWMBottomMenuCollectionViewLandscapeCell.xib   |    86 +
 .../MWMBottomMenuCollectionViewPortraitCell.xib    |   105 +
 .../BottomMenu/MWMBottomMenuLayout.h               |     6 +
 .../BottomMenu/MWMBottomMenuLayout.mm              |    48 +
 .../MapViewControls/BottomMenu/MWMBottomMenuView.h |    23 +
 .../BottomMenu/MWMBottomMenuView.mm                |   324 +
 .../BottomMenu/MWMBottomMenuViewController.h       |    26 +
 .../BottomMenu/MWMBottomMenuViewController.mm      |   474 +
 .../BottomMenu/MWMBottomMenuViewController.xib     |   242 +
 .../LocationButton/MWMLocationButton.h             |    10 -
 .../LocationButton/MWMLocationButton.mm            |    59 -
 .../LocationButton/MWMLocationButton.xib           |    39 -
 .../LocationButton/MWMLocationButtonView.h         |    12 -
 .../LocationButton/MWMLocationButtonView.mm        |   137 -
 .../MapViewControls/MWMMapViewControlsCommon.h     |    11 +
 .../MapViewControls/MWMMapViewControlsManager.h    |     5 +-
 .../MapViewControls/MWMMapViewControlsManager.mm   |    86 +-
 .../MapViewControls/Search/MWMSearchManager.h      |     1 -
 .../MapViewControls/Search/MWMSearchManager.mm     |    32 +-
 .../MapViewControls/SideMenu/MWMSideMenuButton.h   |    17 -
 .../MapViewControls/SideMenu/MWMSideMenuButton.mm  |   142 -
 .../SideMenu/MWMSideMenuButtonDelegate.h           |    12 -
 .../MapViewControls/SideMenu/MWMSideMenuDelegate.h |     8 -
 .../SideMenu/MWMSideMenuDownloadBadge.h            |    17 -
 .../SideMenu/MWMSideMenuDownloadBadge.m            |    53 -
 .../MapViewControls/SideMenu/MWMSideMenuManager.h  |    21 -
 .../MapViewControls/SideMenu/MWMSideMenuManager.mm |   240 -
 .../SideMenu/MWMSideMenuManagerDelegate.h          |     8 -
 .../MapViewControls/SideMenu/MWMSideMenuView.h     |    15 -
 .../MapViewControls/SideMenu/MWMSideMenuView.mm    |   210 -
 .../MapViewControls/SideMenu/MWMSideMenuViews.xib  |   261 -
 .../MapViewControls/ZoomButtons/MWMZoomButtons.mm  |     3 +
 .../ZoomButtons/MWMZoomButtonsView.m               |    36 +-
 .../MWMNavigationDashboardManager.mm               |     2 +-
 iphone/Maps/Classes/MWMBasePlacePageView.mm        |     7 +-
 iphone/Maps/Classes/MWMExtendedPlacePageView.h     |     4 +-
 iphone/Maps/Classes/MWMPlacePage.h                 |     1 +
 iphone/Maps/Classes/MWMPlacePage.mm                |    41 +-
 iphone/Maps/Classes/MWMPlacePageActionBar.h        |     2 +
 iphone/Maps/Classes/MWMPlacePageActionBar.mm       |     1 -
 iphone/Maps/Classes/MWMPlacePageViewManager.h      |     3 +-
 iphone/Maps/Classes/MWMPlacePageViewManager.mm     |    23 +-
 .../Maps/Classes/MWMPlacePageViewManagerDelegate.h |     2 +-
 iphone/Maps/Classes/MWMiPadPlacePage.mm            |     5 +
 iphone/Maps/Classes/MWMiPhoneLandscapePlacePage.mm |    49 +-
 iphone/Maps/Classes/MWMiPhonePortraitPlacePage.mm  |    34 +-
 iphone/Maps/Classes/MapViewController.mm           |     6 +-
 .../Maps/Images.xcassets/Bottom Menu/Contents.json |     6 +
 .../Bottom Menu/Morphing/Contents.json             |     6 +
 .../ic_follow_mode_light_1.imageset/Contents.json  |    23 +
 .../ic_follow_mode_1_light.png                     |   Bin 0 -> 477 bytes
 .../ic_follow_mode_1_light@2x.png                  |   Bin 0 -> 859 bytes
 .../ic_follow_mode_1_light@3x.png                  |   Bin 0 -> 1188 bytes
 .../ic_follow_mode_light_2.imageset/Contents.json  |    23 +
 .../ic_follow_mode_2_light.png                     |   Bin 0 -> 485 bytes
 .../ic_follow_mode_2_light@2x.png                  |   Bin 0 -> 878 bytes
 .../ic_follow_mode_2_light@3x.png                  |   Bin 0 -> 1223 bytes
 .../ic_follow_mode_light_3.imageset/Contents.json  |    23 +
 .../ic_follow_mode_3_light.png                     |   Bin 0 -> 499 bytes
 .../ic_follow_mode_3_light@2x.png                  |   Bin 0 -> 887 bytes
 .../ic_follow_mode_3_light@3x.png                  |   Bin 0 -> 1319 bytes
 .../ic_follow_mode_light_4.imageset/Contents.json  |    23 +
 .../ic_follow_mode_4_light.png                     |   Bin 0 -> 396 bytes
 .../ic_follow_mode_4_light@2x.png                  |   Bin 0 -> 715 bytes
 .../ic_follow_mode_4_light@3x.png                  |   Bin 0 -> 1052 bytes
 .../ic_follow_mode_light_5.imageset/Contents.json  |    23 +
 .../ic_follow_mode_5_light.png                     |   Bin 0 -> 416 bytes
 .../ic_follow_mode_5_light@2x.png                  |   Bin 0 -> 749 bytes
 .../ic_follow_mode_5_light@3x.png                  |   Bin 0 -> 1096 bytes
 .../ic_follow_mode_light_6.imageset/Contents.json  |    23 +
 .../ic_follow_mode_5_light.png                     |   Bin 0 -> 416 bytes
 .../ic_follow_mode_5_light@2x.png                  |   Bin 0 -> 749 bytes
 .../ic_follow_mode_5_light@3x.png                  |   Bin 0 -> 1096 bytes
 .../Morphing/ic_menu_1.imageset/Contents.json      |    23 +
 .../Morphing/ic_menu_1.imageset/ic_menu_1.png      |   Bin 0 -> 117 bytes
 .../Morphing/ic_menu_1.imageset/ic_menu_1@2x.png   |   Bin 0 -> 145 bytes
 .../Morphing/ic_menu_1.imageset/ic_menu_1@3x.png   |   Bin 0 -> 217 bytes
 .../Morphing/ic_menu_2.imageset/Contents.json      |    23 +
 .../Morphing/ic_menu_2.imageset/ic_menu_2.png      |   Bin 0 -> 116 bytes
 .../Morphing/ic_menu_2.imageset/ic_menu_2@2x.png   |   Bin 0 -> 144 bytes
 .../Morphing/ic_menu_2.imageset/ic_menu_2@3x.png   |   Bin 0 -> 218 bytes
 .../Morphing/ic_menu_3.imageset/Contents.json      |    23 +
 .../Morphing/ic_menu_3.imageset/ic_menu_3.png      |   Bin 0 -> 104 bytes
 .../Morphing/ic_menu_3.imageset/ic_menu_3@2x.png   |   Bin 0 -> 138 bytes
 .../Morphing/ic_menu_3.imageset/ic_menu_3@3x.png   |   Bin 0 -> 191 bytes
 .../Morphing/ic_menu_4.imageset/Contents.json      |    23 +
 .../Morphing/ic_menu_4.imageset/ic_menu_4.png      |   Bin 0 -> 137 bytes
 .../Morphing/ic_menu_4.imageset/ic_menu_4@2x.png   |   Bin 0 -> 198 bytes
 .../Morphing/ic_menu_4.imageset/ic_menu_4@3x.png   |   Bin 0 -> 273 bytes
 .../Morphing/ic_menu_5.imageset/Contents.json      |    23 +
 .../Morphing/ic_menu_5.imageset/ic_menu_5.png      |   Bin 0 -> 136 bytes
 .../Morphing/ic_menu_5.imageset/ic_menu_5@2x.png   |   Bin 0 -> 194 bytes
 .../Morphing/ic_menu_5.imageset/ic_menu_5@3x.png   |   Bin 0 -> 244 bytes
 .../Morphing/ic_menu_6.imageset/Contents.json      |    23 +
 .../Morphing/ic_menu_6.imageset/ic_menu_6.png      |   Bin 0 -> 164 bytes
 .../Morphing/ic_menu_6.imageset/ic_menu_6@2x.png   |   Bin 0 -> 245 bytes
 .../Morphing/ic_menu_6.imageset/ic_menu_6@3x.png   |   Bin 0 -> 323 bytes
 .../ic_menu_rotate_1.imageset/Contents.json        |    23 +
 .../ic_menu_rotate_1.imageset/ic_menu_rotate_1.png |   Bin 0 -> 117 bytes
 .../ic_menu_rotate_1@2x.png                        |   Bin 0 -> 145 bytes
 .../ic_menu_rotate_1@3x.png                        |   Bin 0 -> 217 bytes
 .../ic_menu_rotate_2.imageset/Contents.json        |    23 +
 .../ic_menu_rotate_2.imageset/ic_menu_rotate_2.png |   Bin 0 -> 377 bytes
 .../ic_menu_rotate_2@2x.png                        |   Bin 0 -> 692 bytes
 .../ic_menu_rotate_2@3x.png                        |   Bin 0 -> 936 bytes
 .../ic_menu_rotate_3.imageset/Contents.json        |    23 +
 .../ic_menu_rotate_3.imageset/ic_menu_rotate_3.png |   Bin 0 -> 245 bytes
 .../ic_menu_rotate_3@2x.png                        |   Bin 0 -> 423 bytes
 .../ic_menu_rotate_3@3x.png                        |   Bin 0 -> 598 bytes
 .../ic_menu_rotate_4.imageset/Contents.json        |    23 +
 .../ic_menu_rotate_4.imageset/ic_menu_rotate_4.png |   Bin 0 -> 239 bytes
 .../ic_menu_rotate_4@2x.png                        |   Bin 0 -> 404 bytes
 .../ic_menu_rotate_4@3x.png                        |   Bin 0 -> 558 bytes
 .../ic_menu_rotate_5.imageset/Contents.json        |    23 +
 .../ic_menu_rotate_5.imageset/ic_menu_rotate_5.png |   Bin 0 -> 219 bytes
 .../ic_menu_rotate_5@2x.png                        |   Bin 0 -> 362 bytes
 .../ic_menu_rotate_5@3x.png                        |   Bin 0 -> 497 bytes
 .../ic_menu_rotate_6.imageset/Contents.json        |    23 +
 .../ic_menu_rotate_6.imageset/ic_menu_rotate_6.png |   Bin 0 -> 159 bytes
 .../ic_menu_rotate_6@2x.png                        |   Bin 0 -> 239 bytes
 .../ic_menu_rotate_6@3x.png                        |   Bin 0 -> 325 bytes
 .../Bottom Menu/ic_menu.imageset/Contents.json     |    23 +
 .../Bottom Menu/ic_menu.imageset/ic_menu.png       |   Bin 0 -> 117 bytes
 .../Bottom Menu/ic_menu.imageset/ic_menu@2x.png    |   Bin 0 -> 145 bytes
 .../Bottom Menu/ic_menu.imageset/ic_menu@3x.png    |   Bin 0 -> 217 bytes
 .../Contents.json                                  |    23 +
 .../ic_bookmark_list_light.png                     |   Bin 0 -> 248 bytes
 .../ic_bookmark_list_light@2x.png                  |   Bin 0 -> 414 bytes
 .../ic_bookmark_list_light@3x.png                  |   Bin 0 -> 611 bytes
 .../Contents.json                                  |    23 +
 .../ic_bookmark_list_light_press@1x.png            |   Bin 0 -> 237 bytes
 .../ic_bookmark_list_light_press@2x.png            |   Bin 0 -> 393 bytes
 .../ic_bookmark_list_light_press@3x.png            |   Bin 0 -> 584 bytes
 .../ic_menu_down.imageset/Contents.json            |    23 +
 .../ic_menu_down.imageset/ic_menu_down.png         |   Bin 0 -> 164 bytes
 .../ic_menu_down.imageset/ic_menu_down@2x.png      |   Bin 0 -> 245 bytes
 .../ic_menu_down.imageset/ic_menu_down@3x.png      |   Bin 0 -> 323 bytes
 .../ic_menu_down_press.imageset/Contents.json      |    23 +
 .../ic_menu_down_arrow_light_press.png             |   Bin 0 -> 158 bytes
 .../ic_menu_down_arrow_light_press@2x.png          |   Bin 0 -> 236 bytes
 .../ic_menu_down_arrow_light_press@3x.png          |   Bin 0 -> 302 bytes
 .../ic_menu_download.imageset/Contents.json        |    23 +
 .../ic_menu_download.imageset/ic_download@1x.png   |   Bin 0 -> 346 bytes
 .../ic_menu_download.imageset/ic_download@2x.png   |   Bin 0 -> 647 bytes
 .../ic_menu_download.imageset/ic_download@3x.png   |   Bin 0 -> 961 bytes
 .../ic_menu_download_press.imageset/Contents.json  |    23 +
 .../ic_download_press@1x.png                       |   Bin 0 -> 326 bytes
 .../ic_download_press@2x.png                       |   Bin 0 -> 603 bytes
 .../ic_download_press@3x.png                       |   Bin 0 -> 903 bytes
 .../ic_menu_left.imageset/Contents.json            |    23 +
 .../ic_menu_left.imageset/ic_menu_left_arrow.png   |   Bin 0 -> 159 bytes
 .../ic_menu_left_arrow@2x.png                      |   Bin 0 -> 239 bytes
 .../ic_menu_left_arrow@3x.png                      |   Bin 0 -> 325 bytes
 .../ic_menu_left_press.imageset/Contents.json      |    23 +
 .../ic_menu_left_arrow_light_press.png             |   Bin 0 -> 159 bytes
 .../ic_menu_left_arrow_light_press@2x.png          |   Bin 0 -> 234 bytes
 .../ic_menu_left_arrow_light_press@3x.png          |   Bin 0 -> 304 bytes
 .../ic_menu_location_follow.imageset/Contents.json |    23 +
 .../ic_menu_location_follow.imageset/ic_follow.png |   Bin 0 -> 477 bytes
 .../ic_follow@2x.png                               |   Bin 0 -> 859 bytes
 .../ic_follow@3x.png                               |   Bin 0 -> 1185 bytes
 .../Contents.json                                  |    23 +
 .../ic_follow_and_rotate.png                       |   Bin 0 -> 416 bytes
 .../ic_follow_and_rotate@2x.png                    |   Bin 0 -> 749 bytes
 .../ic_follow_and_rotate@3x.png                    |   Bin 0 -> 1096 bytes
 .../Contents.json                                  |    23 +
 .../ic_get_position.png                            |   Bin 0 -> 272 bytes
 .../ic_get_position@2x.png                         |   Bin 0 -> 508 bytes
 .../ic_get_position@3x.png                         |   Bin 0 -> 740 bytes
 .../Contents.json                                  |    23 +
 .../ic_off_mode_light.png                          |   Bin 0 -> 240 bytes
 .../ic_off_mode_light@2x.png                       |   Bin 0 -> 407 bytes
 .../ic_off_mode_light@3x.png                       |   Bin 0 -> 560 bytes
 .../Contents.json                                  |    23 +
 .../ic_pending_1.png                               |   Bin 0 -> 1197 bytes
 .../ic_pending_1@2x.png                            |   Bin 0 -> 2660 bytes
 .../ic_pending_1@3x.png                            |   Bin 0 -> 4076 bytes
 .../ic_menu_point_to_point.imageset/Contents.json  |    23 +
 .../ic_point_to_point.png                          |   Bin 0 -> 222 bytes
 .../ic_point_to_point@2x.png                       |   Bin 0 -> 380 bytes
 .../ic_point_to_point@3x.png                       |   Bin 0 -> 532 bytes
 .../Contents.json                                  |    23 +
 .../ic_point_to_point_press.png                    |   Bin 0 -> 206 bytes
 .../ic_point_to_point_press@2x.png                 |   Bin 0 -> 349 bytes
 .../ic_point_to_point_press@3x.png                 |   Bin 0 -> 477 bytes
 .../ic_menu_press.imageset/Contents.json           |    23 +
 .../ic_menu_press.imageset/ic_menu_press.png       |   Bin 0 -> 118 bytes
 .../ic_menu_press.imageset/ic_menu_press@2x.png    |   Bin 0 -> 147 bytes
 .../ic_menu_press.imageset/ic_menu_press@3x.png    |   Bin 0 -> 219 bytes
 .../ic_menu_search.imageset/Contents.json          |    23 +
 .../ic_menu_search.imageset/ic_search.png          |   Bin 0 -> 317 bytes
 .../ic_menu_search.imageset/ic_search@2x.png       |   Bin 0 -> 605 bytes
 .../ic_menu_search.imageset/ic_search@3x.png       |   Bin 0 -> 884 bytes
 .../ic_menu_search_press.imageset/Contents.json    |    23 +
 .../ic_search_press.png                            |   Bin 0 -> 296 bytes
 .../ic_search_press@2x.png                         |   Bin 0 -> 554 bytes
 .../ic_search_press@3x.png                         |   Bin 0 -> 812 bytes
 .../ic_menu_settings.imageset/Contents.json        |    23 +
 .../ic_menu_settings.imageset/ic_settings@1x.png   |   Bin 0 -> 318 bytes
 .../ic_menu_settings.imageset/ic_settings@2x.png   |   Bin 0 -> 564 bytes
 .../ic_menu_settings.imageset/ic_settings@3x.png   |   Bin 0 -> 803 bytes
 .../ic_menu_settings_press.imageset/Contents.json  |    23 +
 .../ic_settings_press@1x.png                       |   Bin 0 -> 295 bytes
 .../ic_settings_press@2x.png                       |   Bin 0 -> 508 bytes
 .../ic_settings_press@3x.png                       |   Bin 0 -> 737 bytes
 .../ic_menu_share.imageset/Contents.json           |    23 +
 .../ic_menu_share.imageset/ic_share@1x.png         |   Bin 0 -> 230 bytes
 .../ic_menu_share.imageset/ic_share@2x.png         |   Bin 0 -> 385 bytes
 .../ic_menu_share.imageset/ic_share@3x.png         |   Bin 0 -> 565 bytes
 .../ic_menu_share_press.imageset/Contents.json     |    23 +
 .../ic_share_press@1x.png                          |   Bin 0 -> 216 bytes
 .../ic_share_press@2x.png                          |   Bin 0 -> 363 bytes
 .../ic_share_press@3x.png                          |   Bin 0 -> 529 bytes
 iphone/Maps/Images.xcassets/Contents.json          |     6 +
 .../Contents.json                                  |    23 -
 .../follow_to_followandrotate_1.png                |   Bin 1569 -> 0 bytes
 .../follow_to_followandrotate_1@2x.png             |   Bin 3788 -> 0 bytes
 .../follow_to_followandrotate_1@3x.png             |   Bin 6142 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../follow_to_followandrotate_2.png                |   Bin 1565 -> 0 bytes
 .../follow_to_followandrotate_2@2x.png             |   Bin 3757 -> 0 bytes
 .../follow_to_followandrotate_2@3x.png             |   Bin 6027 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../follow_to_followandrotate_3.png                |   Bin 1573 -> 0 bytes
 .../follow_to_followandrotate_3@2x.png             |   Bin 3788 -> 0 bytes
 .../follow_to_followandrotate_3@3x.png             |   Bin 6024 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../follow_to_followandrotate_4.png                |   Bin 1532 -> 0 bytes
 .../follow_to_followandrotate_4@2x.png             |   Bin 3719 -> 0 bytes
 .../follow_to_followandrotate_4@3x.png             |   Bin 6011 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../follow_to_followandrotate_5.png                |   Bin 1611 -> 0 bytes
 .../follow_to_followandrotate_5@2x.png             |   Bin 3756 -> 0 bytes
 .../follow_to_followandrotate_5@3x.png             |   Bin 6254 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../follow_to_followandrotate_6.png                |   Bin 1564 -> 0 bytes
 .../follow_to_followandrotate_6@2x.png             |   Bin 3753 -> 0 bytes
 .../follow_to_followandrotate_6@3x.png             |   Bin 6108 -> 0 bytes
 .../follow_to_noposition_1.imageset/Contents.json  |    23 -
 .../follow_to_noposition_1.png                     |   Bin 1569 -> 0 bytes
 .../follow_to_noposition_1@2x.png                  |   Bin 3755 -> 0 bytes
 .../follow_to_noposition_1@3x.png                  |   Bin 6008 -> 0 bytes
 .../follow_to_noposition_2.imageset/Contents.json  |    23 -
 .../follow_to_noposition_2.png                     |   Bin 1569 -> 0 bytes
 .../follow_to_noposition_2@2x.png                  |   Bin 3753 -> 0 bytes
 .../follow_to_noposition_2@3x.png                  |   Bin 6028 -> 0 bytes
 .../follow_to_noposition_3.imageset/Contents.json  |    23 -
 .../follow_to_noposition_3.png                     |   Bin 1569 -> 0 bytes
 .../follow_to_noposition_3@2x.png                  |   Bin 3736 -> 0 bytes
 .../follow_to_noposition_3@3x.png                  |   Bin 6000 -> 0 bytes
 .../follow_to_noposition_4.imageset/Contents.json  |    23 -
 .../follow_to_noposition_4.png                     |   Bin 1553 -> 0 bytes
 .../follow_to_noposition_4@2x.png                  |   Bin 3699 -> 0 bytes
 .../follow_to_noposition_4@3x.png                  |   Bin 5924 -> 0 bytes
 .../follow_to_noposition_5.imageset/Contents.json  |    23 -
 .../follow_to_noposition_5.png                     |   Bin 1528 -> 0 bytes
 .../follow_to_noposition_5@2x.png                  |   Bin 3641 -> 0 bytes
 .../follow_to_noposition_5@3x.png                  |   Bin 5843 -> 0 bytes
 .../follow_to_noposition_6.imageset/Contents.json  |    23 -
 .../follow_to_noposition_6.png                     |   Bin 1239 -> 0 bytes
 .../follow_to_noposition_6@2x.png                  |   Bin 2812 -> 0 bytes
 .../follow_to_noposition_6@3x.png                  |   Bin 4741 -> 0 bytes
 .../follow_to_notfollow_1.imageset/Contents.json   |    23 -
 .../follow_to_notfollow_1.png                      |   Bin 1569 -> 0 bytes
 .../follow_to_notfollow_1@2x.png                   |   Bin 3755 -> 0 bytes
 .../follow_to_notfollow_1@3x.png                   |   Bin 6008 -> 0 bytes
 .../follow_to_notfollow_2.imageset/Contents.json   |    23 -
 .../follow_to_notfollow_2.png                      |   Bin 2086 -> 0 bytes
 .../follow_to_notfollow_2@2x.png                   |   Bin 4865 -> 0 bytes
 .../follow_to_notfollow_2@3x.png                   |   Bin 7772 -> 0 bytes
 .../follow_to_notfollow_3.imageset/Contents.json   |    23 -
 .../follow_to_notfollow_3.png                      |   Bin 2048 -> 0 bytes
 .../follow_to_notfollow_3@2x.png                   |   Bin 4862 -> 0 bytes
 .../follow_to_notfollow_3@3x.png                   |   Bin 7776 -> 0 bytes
 .../follow_to_notfollow_4.imageset/Contents.json   |    23 -
 .../follow_to_notfollow_4.png                      |   Bin 1719 -> 0 bytes
 .../follow_to_notfollow_4@2x.png                   |   Bin 4431 -> 0 bytes
 .../follow_to_notfollow_4@3x.png                   |   Bin 7131 -> 0 bytes
 .../follow_to_notfollow_5.imageset/Contents.json   |    23 -
 .../follow_to_notfollow_5.png                      |   Bin 1650 -> 0 bytes
 .../follow_to_notfollow_5@2x.png                   |   Bin 4301 -> 0 bytes
 .../follow_to_notfollow_5@3x.png                   |   Bin 6861 -> 0 bytes
 .../follow_to_notfollow_6.imageset/Contents.json   |    23 -
 .../follow_to_notfollow_6.png                      |   Bin 1362 -> 0 bytes
 .../follow_to_notfollow_6@2x.png                   |   Bin 3072 -> 0 bytes
 .../follow_to_notfollow_6@3x.png                   |   Bin 5191 -> 0 bytes
 .../follow_to_pending_1.imageset/Contents.json     |    23 -
 .../follow_to_pending_1.png                        |   Bin 1564 -> 0 bytes
 .../follow_to_pending_1@2x.png                     |   Bin 3794 -> 0 bytes
 .../follow_to_pending_1@3x.png                     |   Bin 6125 -> 0 bytes
 .../follow_to_pending_2.imageset/Contents.json     |    23 -
 .../follow_to_pending_2.png                        |   Bin 1508 -> 0 bytes
 .../follow_to_pending_2@2x.png                     |   Bin 3657 -> 0 bytes
 .../follow_to_pending_2@3x.png                     |   Bin 5869 -> 0 bytes
 .../follow_to_pending_3.imageset/Contents.json     |    23 -
 .../follow_to_pending_3.png                        |   Bin 1489 -> 0 bytes
 .../follow_to_pending_3@2x.png                     |   Bin 3620 -> 0 bytes
 .../follow_to_pending_3@3x.png                     |   Bin 5802 -> 0 bytes
 .../follow_to_pending_4.imageset/Contents.json     |    23 -
 .../follow_to_pending_4.png                        |   Bin 1316 -> 0 bytes
 .../follow_to_pending_4@2x.png                     |   Bin 2963 -> 0 bytes
 .../follow_to_pending_4@3x.png                     |   Bin 5691 -> 0 bytes
 .../follow_to_pending_5.imageset/Contents.json     |    23 -
 .../follow_to_pending_5.png                        |   Bin 1647 -> 0 bytes
 .../follow_to_pending_5@2x.png                     |   Bin 3864 -> 0 bytes
 .../follow_to_pending_5@3x.png                     |   Bin 6269 -> 0 bytes
 .../follow_to_pending_6.imageset/Contents.json     |    23 -
 .../follow_to_pending_6.png                        |   Bin 1647 -> 0 bytes
 .../follow_to_pending_6@2x.png                     |   Bin 3864 -> 0 bytes
 .../follow_to_pending_6@3x.png                     |   Bin 6269 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_follow_1.png                |   Bin 1642 -> 0 bytes
 .../followandrotate_to_follow_1@2x.png             |   Bin 3802 -> 0 bytes
 .../followandrotate_to_follow_1@3x.png             |   Bin 6409 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_follow_2.png                |   Bin 1555 -> 0 bytes
 .../followandrotate_to_follow_2@2x.png             |   Bin 3738 -> 0 bytes
 .../followandrotate_to_follow_2@3x.png             |   Bin 6083 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_follow_3.png                |   Bin 1532 -> 0 bytes
 .../followandrotate_to_follow_3@2x.png             |   Bin 3719 -> 0 bytes
 .../followandrotate_to_follow_3@3x.png             |   Bin 6011 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_follow_4.png                |   Bin 1573 -> 0 bytes
 .../followandrotate_to_follow_4@2x.png             |   Bin 3788 -> 0 bytes
 .../followandrotate_to_follow_4@3x.png             |   Bin 6024 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_follow_5.png                |   Bin 1565 -> 0 bytes
 .../followandrotate_to_follow_5@2x.png             |   Bin 3757 -> 0 bytes
 .../followandrotate_to_follow_5@3x.png             |   Bin 6027 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_follow_6.png                |   Bin 1569 -> 0 bytes
 .../followandrotate_to_follow_6@2x.png             |   Bin 3755 -> 0 bytes
 .../followandrotate_to_follow_6@3x.png             |   Bin 6008 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_noposition_1.png            |   Bin 1564 -> 0 bytes
 .../followandrotate_to_noposition_1@2x.png         |   Bin 3753 -> 0 bytes
 .../followandrotate_to_noposition_1@3x.png         |   Bin 6108 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_noposition_2.png            |   Bin 1555 -> 0 bytes
 .../followandrotate_to_noposition_2@2x.png         |   Bin 3740 -> 0 bytes
 .../followandrotate_to_noposition_2@3x.png         |   Bin 6076 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_noposition_3.png            |   Bin 1491 -> 0 bytes
 .../followandrotate_to_noposition_3@2x.png         |   Bin 3681 -> 0 bytes
 .../followandrotate_to_noposition_3@3x.png         |   Bin 5931 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_noposition_4.png            |   Bin 1559 -> 0 bytes
 .../followandrotate_to_noposition_4@2x.png         |   Bin 3728 -> 0 bytes
 .../followandrotate_to_noposition_4@3x.png         |   Bin 5937 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_noposition_5.png            |   Bin 1560 -> 0 bytes
 .../followandrotate_to_noposition_5@2x.png         |   Bin 3713 -> 0 bytes
 .../followandrotate_to_noposition_5@3x.png         |   Bin 5932 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_noposition_6.png            |   Bin 1251 -> 0 bytes
 .../followandrotate_to_noposition_6@2x.png         |   Bin 2853 -> 0 bytes
 .../followandrotate_to_noposition_6@3x.png         |   Bin 4773 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_notfollow_1.png             |   Bin 1564 -> 0 bytes
 .../followandrotate_to_notfollow_1@2x.png          |   Bin 3753 -> 0 bytes
 .../followandrotate_to_notfollow_1@3x.png          |   Bin 6108 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_notfollow_2.png             |   Bin 2074 -> 0 bytes
 .../followandrotate_to_notfollow_2@2x.png          |   Bin 4872 -> 0 bytes
 .../followandrotate_to_notfollow_2@3x.png          |   Bin 7809 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_notfollow_3.png             |   Bin 2052 -> 0 bytes
 .../followandrotate_to_notfollow_3@2x.png          |   Bin 4864 -> 0 bytes
 .../followandrotate_to_notfollow_3@3x.png          |   Bin 7838 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_notfollow_4.png             |   Bin 1719 -> 0 bytes
 .../followandrotate_to_notfollow_4@2x.png          |   Bin 4431 -> 0 bytes
 .../followandrotate_to_notfollow_4@3x.png          |   Bin 7131 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_notfollow_5.png             |   Bin 1650 -> 0 bytes
 .../followandrotate_to_notfollow_5@2x.png          |   Bin 4301 -> 0 bytes
 .../followandrotate_to_notfollow_5@3x.png          |   Bin 6861 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_notfollow_6.png             |   Bin 1362 -> 0 bytes
 .../followandrotate_to_notfollow_6@2x.png          |   Bin 3072 -> 0 bytes
 .../followandrotate_to_notfollow_6@3x.png          |   Bin 5191 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_pending_1.png               |   Bin 1642 -> 0 bytes
 .../followandrotate_to_pending_1@2x.png            |   Bin 3748 -> 0 bytes
 .../followandrotate_to_pending_1@3x.png            |   Bin 6291 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_pending_2.png               |   Bin 1555 -> 0 bytes
 .../followandrotate_to_pending_2@2x.png            |   Bin 3738 -> 0 bytes
 .../followandrotate_to_pending_2@3x.png            |   Bin 6083 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_pending_3.png               |   Bin 1509 -> 0 bytes
 .../followandrotate_to_pending_3@2x.png            |   Bin 3666 -> 0 bytes
 .../followandrotate_to_pending_3@3x.png            |   Bin 5894 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_pending_4.png               |   Bin 1485 -> 0 bytes
 .../followandrotate_to_pending_4@2x.png            |   Bin 3623 -> 0 bytes
 .../followandrotate_to_pending_4@3x.png            |   Bin 5835 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_pending_5.png               |   Bin 1316 -> 0 bytes
 .../followandrotate_to_pending_5@2x.png            |   Bin 2963 -> 0 bytes
 .../followandrotate_to_pending_5@3x.png            |   Bin 5689 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../followandrotate_to_pending_6.png               |   Bin 1647 -> 0 bytes
 .../followandrotate_to_pending_6@2x.png            |   Bin 3864 -> 0 bytes
 .../followandrotate_to_pending_6@3x.png            |   Bin 6269 -> 0 bytes
 .../noposition_to_follow_1.imageset/Contents.json  |    23 -
 .../noposition_to_follow_1.png                     |   Bin 1239 -> 0 bytes
 .../noposition_to_follow_1@2x.png                  |   Bin 2812 -> 0 bytes
 .../noposition_to_follow_1@3x.png                  |   Bin 4741 -> 0 bytes
 .../noposition_to_follow_2.imageset/Contents.json  |    23 -
 .../noposition_to_follow_2.png                     |   Bin 1528 -> 0 bytes
 .../noposition_to_follow_2@2x.png                  |   Bin 3641 -> 0 bytes
 .../noposition_to_follow_2@3x.png                  |   Bin 5843 -> 0 bytes
 .../noposition_to_follow_3.imageset/Contents.json  |    23 -
 .../noposition_to_follow_3.png                     |   Bin 1553 -> 0 bytes
 .../noposition_to_follow_3@2x.png                  |   Bin 3699 -> 0 bytes
 .../noposition_to_follow_3@3x.png                  |   Bin 5924 -> 0 bytes
 .../noposition_to_follow_4.imageset/Contents.json  |    23 -
 .../noposition_to_follow_4.png                     |   Bin 1569 -> 0 bytes
 .../noposition_to_follow_4@2x.png                  |   Bin 3736 -> 0 bytes
 .../noposition_to_follow_4@3x.png                  |   Bin 6000 -> 0 bytes
 .../noposition_to_follow_5.imageset/Contents.json  |    23 -
 .../noposition_to_follow_5.png                     |   Bin 1569 -> 0 bytes
 .../noposition_to_follow_5@2x.png                  |   Bin 3753 -> 0 bytes
 .../noposition_to_follow_5@3x.png                  |   Bin 6028 -> 0 bytes
 .../noposition_to_follow_6.imageset/Contents.json  |    23 -
 .../noposition_to_follow_6.png                     |   Bin 1569 -> 0 bytes
 .../noposition_to_follow_6@2x.png                  |   Bin 3755 -> 0 bytes
 .../noposition_to_follow_6@3x.png                  |   Bin 6008 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_followandrotate_1.png            |   Bin 1251 -> 0 bytes
 .../noposition_to_followandrotate_1@2x.png         |   Bin 3052 -> 0 bytes
 .../noposition_to_followandrotate_1@3x.png         |   Bin 5095 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_followandrotate_2.png            |   Bin 1564 -> 0 bytes
 .../noposition_to_followandrotate_2@2x.png         |   Bin 3725 -> 0 bytes
 .../noposition_to_followandrotate_2@3x.png         |   Bin 5928 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_followandrotate_3.png            |   Bin 1563 -> 0 bytes
 .../noposition_to_followandrotate_3@2x.png         |   Bin 3734 -> 0 bytes
 .../noposition_to_followandrotate_3@3x.png         |   Bin 5926 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_followandrotate_4.png            |   Bin 1522 -> 0 bytes
 .../noposition_to_followandrotate_4@2x.png         |   Bin 3714 -> 0 bytes
 .../noposition_to_followandrotate_4@3x.png         |   Bin 5994 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_followandrotate_5.png            |   Bin 1542 -> 0 bytes
 .../noposition_to_followandrotate_5@2x.png         |   Bin 3738 -> 0 bytes
 .../noposition_to_followandrotate_5@3x.png         |   Bin 6090 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_followandrotate_6.png            |   Bin 1564 -> 0 bytes
 .../noposition_to_followandrotate_6@2x.png         |   Bin 3753 -> 0 bytes
 .../noposition_to_followandrotate_6@3x.png         |   Bin 6108 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_notfollow_1.png                  |   Bin 1256 -> 0 bytes
 .../noposition_to_notfollow_1@2x.png               |   Bin 2843 -> 0 bytes
 .../noposition_to_notfollow_1@3x.png               |   Bin 4919 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_notfollow_2.png                  |   Bin 1518 -> 0 bytes
 .../noposition_to_notfollow_2@2x.png               |   Bin 3401 -> 0 bytes
 .../noposition_to_notfollow_2@3x.png               |   Bin 5685 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_notfollow_3.png                  |   Bin 1600 -> 0 bytes
 .../noposition_to_notfollow_3@2x.png               |   Bin 3558 -> 0 bytes
 .../noposition_to_notfollow_3@3x.png               |   Bin 5862 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_notfollow_4.png                  |   Bin 1500 -> 0 bytes
 .../noposition_to_notfollow_4@2x.png               |   Bin 3436 -> 0 bytes
 .../noposition_to_notfollow_4@3x.png               |   Bin 5650 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_notfollow_5.png                  |   Bin 1464 -> 0 bytes
 .../noposition_to_notfollow_5@2x.png               |   Bin 3351 -> 0 bytes
 .../noposition_to_notfollow_5@3x.png               |   Bin 5532 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../noposition_to_notfollow_6.png                  |   Bin 1362 -> 0 bytes
 .../noposition_to_notfollow_6@2x.png               |   Bin 3072 -> 0 bytes
 .../noposition_to_notfollow_6@3x.png               |   Bin 5191 -> 0 bytes
 .../noposition_to_pending_1.imageset/Contents.json |    23 -
 .../noposition_to_pending_1.png                    |   Bin 1251 -> 0 bytes
 .../noposition_to_pending_1@2x.png                 |   Bin 2853 -> 0 bytes
 .../noposition_to_pending_1@3x.png                 |   Bin 4773 -> 0 bytes
 .../noposition_to_pending_2.imageset/Contents.json |    23 -
 .../noposition_to_pending_2.png                    |   Bin 1495 -> 0 bytes
 .../noposition_to_pending_2@2x.png                 |   Bin 3622 -> 0 bytes
 .../noposition_to_pending_2@3x.png                 |   Bin 5809 -> 0 bytes
 .../noposition_to_pending_3.imageset/Contents.json |    23 -
 .../noposition_to_pending_3.png                    |   Bin 1493 -> 0 bytes
 .../noposition_to_pending_3@2x.png                 |   Bin 3606 -> 0 bytes
 .../noposition_to_pending_3@3x.png                 |   Bin 5812 -> 0 bytes
 .../noposition_to_pending_4.imageset/Contents.json |    23 -
 .../noposition_to_pending_4.png                    |   Bin 1316 -> 0 bytes
 .../noposition_to_pending_4@2x.png                 |   Bin 2963 -> 0 bytes
 .../noposition_to_pending_4@3x.png                 |   Bin 5684 -> 0 bytes
 .../noposition_to_pending_5.imageset/Contents.json |    23 -
 .../noposition_to_pending_5.png                    |   Bin 1647 -> 0 bytes
 .../noposition_to_pending_5@2x.png                 |   Bin 3864 -> 0 bytes
 .../noposition_to_pending_5@3x.png                 |   Bin 6269 -> 0 bytes
 .../noposition_to_pending_6.imageset/Contents.json |    23 -
 .../noposition_to_pending_6.png                    |   Bin 1647 -> 0 bytes
 .../noposition_to_pending_6@2x.png                 |   Bin 6606 -> 0 bytes
 .../noposition_to_pending_6@3x.png                 |   Bin 11369 -> 0 bytes
 .../notfollow_to_follow_1.imageset/Contents.json   |    23 -
 .../notfollow_to_follow_1.png                      |   Bin 1362 -> 0 bytes
 .../notfollow_to_follow_1@2x.png                   |   Bin 3072 -> 0 bytes
 .../notfollow_to_follow_1@3x.png                   |   Bin 5191 -> 0 bytes
 .../notfollow_to_follow_2.imageset/Contents.json   |    23 -
 .../notfollow_to_follow_2.png                      |   Bin 1650 -> 0 bytes
 .../notfollow_to_follow_2@2x.png                   |   Bin 4301 -> 0 bytes
 .../notfollow_to_follow_2@3x.png                   |   Bin 6861 -> 0 bytes
 .../notfollow_to_follow_3.imageset/Contents.json   |    23 -
 .../notfollow_to_follow_3.png                      |   Bin 1719 -> 0 bytes
 .../notfollow_to_follow_3@2x.png                   |   Bin 4431 -> 0 bytes
 .../notfollow_to_follow_3@3x.png                   |   Bin 7131 -> 0 bytes
 .../notfollow_to_follow_4.imageset/Contents.json   |    23 -
 .../notfollow_to_follow_4.png                      |   Bin 2051 -> 0 bytes
 .../notfollow_to_follow_4@2x.png                   |   Bin 4837 -> 0 bytes
 .../notfollow_to_follow_4@3x.png                   |   Bin 7793 -> 0 bytes
 .../notfollow_to_follow_5.imageset/Contents.json   |    23 -
 .../notfollow_to_follow_5.png                      |   Bin 2102 -> 0 bytes
 .../notfollow_to_follow_5@2x.png                   |   Bin 4882 -> 0 bytes
 .../notfollow_to_follow_5@3x.png                   |   Bin 7814 -> 0 bytes
 .../notfollow_to_follow_6.imageset/Contents.json   |    23 -
 .../notfollow_to_follow_6.png                      |   Bin 1569 -> 0 bytes
 .../notfollow_to_follow_6@2x.png                   |   Bin 3755 -> 0 bytes
 .../notfollow_to_follow_6@3x.png                   |   Bin 6008 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_followandrotate_1.png             |   Bin 1362 -> 0 bytes
 .../notfollow_to_followandrotate_1@2x.png          |   Bin 3252 -> 0 bytes
 .../notfollow_to_followandrotate_1@3x.png          |   Bin 5514 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_followandrotate_2.png             |   Bin 1664 -> 0 bytes
 .../notfollow_to_followandrotate_2@2x.png          |   Bin 4292 -> 0 bytes
 .../notfollow_to_followandrotate_2@3x.png          |   Bin 6856 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_followandrotate_3.png             |   Bin 1716 -> 0 bytes
 .../notfollow_to_followandrotate_3@2x.png          |   Bin 4407 -> 0 bytes
 .../notfollow_to_followandrotate_3@3x.png          |   Bin 7073 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_followandrotate_4.png             |   Bin 2014 -> 0 bytes
 .../notfollow_to_followandrotate_4@2x.png          |   Bin 4754 -> 0 bytes
 .../notfollow_to_followandrotate_4@3x.png          |   Bin 7617 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_followandrotate_5.png             |   Bin 1542 -> 0 bytes
 .../notfollow_to_followandrotate_5@2x.png          |   Bin 3752 -> 0 bytes
 .../notfollow_to_followandrotate_5@3x.png          |   Bin 6095 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_followandrotate_6.png             |   Bin 1564 -> 0 bytes
 .../notfollow_to_followandrotate_6@2x.png          |   Bin 3753 -> 0 bytes
 .../notfollow_to_followandrotate_6@3x.png          |   Bin 6108 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_noposition_1.png                  |   Bin 1362 -> 0 bytes
 .../notfollow_to_noposition_1@2x.png               |   Bin 3072 -> 0 bytes
 .../notfollow_to_noposition_1@3x.png               |   Bin 5191 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_noposition_2.png                  |   Bin 1471 -> 0 bytes
 .../notfollow_to_noposition_2@2x.png               |   Bin 3365 -> 0 bytes
 .../notfollow_to_noposition_2@3x.png               |   Bin 5573 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_noposition_3.png                  |   Bin 1502 -> 0 bytes
 .../notfollow_to_noposition_3@2x.png               |   Bin 3443 -> 0 bytes
 .../notfollow_to_noposition_3@3x.png               |   Bin 5627 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../nofollow_to_noposition_4.png                   |   Bin 1558 -> 0 bytes
 .../nofollow_to_noposition_4@2x.png                |   Bin 3471 -> 0 bytes
 .../nofollow_to_noposition_4@3x.png                |   Bin 5773 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_noposition_5.png                  |   Bin 1524 -> 0 bytes
 .../notfollow_to_noposition_5@2x.png               |   Bin 3422 -> 0 bytes
 .../notfollow_to_noposition_5@3x.png               |   Bin 5732 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../notfollow_to_noposition_6.png                  |   Bin 1251 -> 0 bytes
 .../notfollow_to_noposition_6@2x.png               |   Bin 2853 -> 0 bytes
 .../notfollow_to_noposition_6@3x.png               |   Bin 4773 -> 0 bytes
 .../notfollow_to_pending_1.imageset/Contents.json  |    23 -
 .../notfollow_to_pending_1.png                     |   Bin 1344 -> 0 bytes
 .../notfollow_to_pending_1@2x.png                  |   Bin 3077 -> 0 bytes
 .../notfollow_to_pending_1@3x.png                  |   Bin 5252 -> 0 bytes
 .../notfollow_to_pending_2.imageset/Contents.json  |    23 -
 .../notfollow_to_pending_2.png                     |   Bin 1362 -> 0 bytes
 .../notfollow_to_pending_2@2x.png                  |   Bin 3072 -> 0 bytes
 .../notfollow_to_pending_2@3x.png                  |   Bin 5191 -> 0 bytes
 .../notfollow_to_pending_3.imageset/Contents.json  |    23 -
 .../notfollow_to_pending_3.png                     |   Bin 1651 -> 0 bytes
 .../notfollow_to_pending_3@2x.png                  |   Bin 4275 -> 0 bytes
 .../notfollow_to_pending_3@3x.png                  |   Bin 6871 -> 0 bytes
 .../notfollow_to_pending_4.imageset/Contents.json  |    23 -
 .../notfollow_to_pending_4.png                     |   Bin 1707 -> 0 bytes
 .../notfollow_to_pending_4@2x.png                  |   Bin 4408 -> 0 bytes
 .../notfollow_to_pending_4@3x.png                  |   Bin 7087 -> 0 bytes
 .../notfollow_to_pending_5.imageset/Contents.json  |    23 -
 .../notfollow_to_pending_5.png                     |   Bin 1423 -> 0 bytes
 .../notfollow_to_pending_5@2x.png                  |   Bin 3125 -> 0 bytes
 .../notfollow_to_pending_5@3x.png                  |   Bin 6210 -> 0 bytes
 .../notfollow_to_pending_6.imageset/Contents.json  |    23 -
 .../notfollow_to_pending_6.png                     |   Bin 1647 -> 0 bytes
 .../notfollow_to_pending_6@2x.png                  |   Bin 3864 -> 0 bytes
 .../notfollow_to_pending_6@3x.png                  |   Bin 6269 -> 0 bytes
 .../pending_to_follow_1.imageset/Contents.json     |    23 -
 .../pending_to_follow_1.png                        |   Bin 1647 -> 0 bytes
 .../pending_to_follow_1@2x.png                     |   Bin 3864 -> 0 bytes
 .../pending_to_follow_1@3x.png                     |   Bin 6269 -> 0 bytes
 .../pending_to_follow_2.imageset/Contents.json     |    23 -
 .../pending_to_follow_2.png                        |   Bin 1316 -> 0 bytes
 .../pending_to_follow_2@2x.png                     |   Bin 2963 -> 0 bytes
 .../pending_to_follow_2@3x.png                     |   Bin 5689 -> 0 bytes
 .../pending_to_follow_3.imageset/Contents.json     |    23 -
 .../pending_to_follow_3.png                        |   Bin 1491 -> 0 bytes
 .../pending_to_follow_3@2x.png                     |   Bin 3605 -> 0 bytes
 .../pending_to_follow_3@3x.png                     |   Bin 5836 -> 0 bytes
 .../pending_to_follow_4.imageset/Contents.json     |    23 -
 .../pending_to_follow_4.png                        |   Bin 1510 -> 0 bytes
 .../pending_to_follow_4@2x.png                     |   Bin 3670 -> 0 bytes
 .../pending_to_follow_4@3x.png                     |   Bin 5872 -> 0 bytes
 .../pending_to_follow_5.imageset/Contents.json     |    23 -
 .../pending_to_follow_5.png                        |   Bin 1564 -> 0 bytes
 .../pending_to_follow_5@2x.png                     |   Bin 3759 -> 0 bytes
 .../pending_to_follow_5@3x.png                     |   Bin 6008 -> 0 bytes
 .../pending_to_follow_6.imageset/Contents.json     |    23 -
 .../pending_to_follow_6.png                        |   Bin 2420 -> 0 bytes
 .../pending_to_follow_6@2x.png                     |   Bin 6471 -> 0 bytes
 .../pending_to_follow_6@3x.png                     |   Bin 10887 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../pending_to_followandrotate_1.png               |   Bin 2613 -> 0 bytes
 .../pending_to_followandrotate_1@2x.png            |   Bin 6606 -> 0 bytes
 .../pending_to_followandrotate_1@3x.png            |   Bin 11369 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../pending_to_followandrotate_2.png               |   Bin 2332 -> 0 bytes
 .../pending_to_followandrotate_2@2x.png            |   Bin 6313 -> 0 bytes
 .../pending_to_followandrotate_2@3x.png            |   Bin 10852 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../pending_to_followandrotate_3.png               |   Bin 2314 -> 0 bytes
 .../pending_to_followandrotate_3@2x.png            |   Bin 6222 -> 0 bytes
 .../pending_to_followandrotate_3@3x.png            |   Bin 10671 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../pending_to_followandrotate_4.png               |   Bin 2334 -> 0 bytes
 .../pending_to_followandrotate_4@2x.png            |   Bin 6252 -> 0 bytes
 .../pending_to_followandrotate_4@3x.png            |   Bin 10745 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../pending_to_followandrotate_5.png               |   Bin 2383 -> 0 bytes
 .../pending_to_followandrotate_5@2x.png            |   Bin 6391 -> 0 bytes
 .../pending_to_followandrotate_5@3x.png            |   Bin 10861 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../pending_to_followandrotate_6.png               |   Bin 2403 -> 0 bytes
 .../pending_to_followandrotate_6@2x.png            |   Bin 6431 -> 0 bytes
 .../pending_to_followandrotate_6@3x.png            |   Bin 10913 -> 0 bytes
 .../pending_to_noposition_1.imageset/Contents.json |    23 -
 .../pending_to_noposition_1.png                    |   Bin 1647 -> 0 bytes
 .../pending_to_noposition_1@2x.png                 |   Bin 3864 -> 0 bytes
 .../pending_to_noposition_1@3x.png                 |   Bin 6269 -> 0 bytes
 .../pending_to_noposition_2.imageset/Contents.json |    23 -
 .../pending_to_noposition_2.png                    |   Bin 1316 -> 0 bytes
 .../pending_to_noposition_2@2x.png                 |   Bin 2963 -> 0 bytes
 .../pending_to_noposition_2@3x.png                 |   Bin 5684 -> 0 bytes
 .../pending_to_noposition_3.imageset/Contents.json |    23 -
 .../pending_to_noposition_3.png                    |   Bin 1498 -> 0 bytes
 .../pending_to_noposition_3@2x.png                 |   Bin 3600 -> 0 bytes
 .../pending_to_noposition_3@3x.png                 |   Bin 5831 -> 0 bytes
 .../pending_to_noposition_4.imageset/Contents.json |    23 -
 .../pending_to_noposition_4.png                    |   Bin 1491 -> 0 bytes
 .../pending_to_noposition_4@2x.png                 |   Bin 3630 -> 0 bytes
 .../pending_to_noposition_4@3x.png                 |   Bin 5809 -> 0 bytes
 .../pending_to_noposition_5.imageset/Contents.json |    23 -
 .../pending_to_noposition_5.png                    |   Bin 1251 -> 0 bytes
 .../pending_to_noposition_5@2x.png                 |   Bin 2853 -> 0 bytes
 .../pending_to_noposition_5@3x.png                 |   Bin 4773 -> 0 bytes
 .../pending_to_noposition_6.imageset/Contents.json |    23 -
 .../pending_to_noposition_6.png                    |   Bin 2278 -> 0 bytes
 .../pending_to_noposition_6@2x.png                 |   Bin 6098 -> 0 bytes
 .../pending_to_noposition_6@3x.png                 |   Bin 10376 -> 0 bytes
 .../pending_to_notfollow_1.imageset/Contents.json  |    23 -
 .../pending_to_nofollow_1.png                      |   Bin 1647 -> 0 bytes
 .../pending_to_nofollow_1@2x.png                   |   Bin 3864 -> 0 bytes
 .../pending_to_nofollow_1@3x.png                   |   Bin 6269 -> 0 bytes
 .../pending_to_notfollow_2.imageset/Contents.json  |    23 -
 .../pending_to_nofollow_2.png                      |   Bin 1423 -> 0 bytes
 .../pending_to_nofollow_2@2x.png                   |   Bin 3125 -> 0 bytes
 .../pending_to_nofollow_2@3x.png                   |   Bin 6210 -> 0 bytes
 .../pending_to_notfollow_3.imageset/Contents.json  |    23 -
 .../pending_to_nofollow_3.png                      |   Bin 1707 -> 0 bytes
 .../pending_to_nofollow_3@2x.png                   |   Bin 4408 -> 0 bytes
 .../pending_to_nofollow_3@3x.png                   |   Bin 7087 -> 0 bytes
 .../pending_to_notfollow_4.imageset/Contents.json  |    23 -
 .../pending_to_nofollow_4.png                      |   Bin 1651 -> 0 bytes
 .../pending_to_nofollow_4@2x.png                   |   Bin 4275 -> 0 bytes
 .../pending_to_nofollow_4@3x.png                   |   Bin 6871 -> 0 bytes
 .../pending_to_notfollow_5.imageset/Contents.json  |    23 -
 .../pending_to_nofollow_5.png                      |   Bin 1362 -> 0 bytes
 .../pending_to_nofollow_5@2x.png                   |   Bin 3072 -> 0 bytes
 .../pending_to_nofollow_5@3x.png                   |   Bin 5191 -> 0 bytes
 .../pending_to_notfollow_6.imageset/Contents.json  |    23 -
 .../pending_to_nofollow_6.png                      |   Bin 3044 -> 0 bytes
 .../pending_to_nofollow_6@2x.png                   |   Bin 7548 -> 0 bytes
 .../pending_to_nofollow_6@3x.png                   |   Bin 12956 -> 0 bytes
 .../btn_white_loading_1.imageset/Contents.json     |    23 -
 .../btn_white_loading_1.png                        |   Bin 1607 -> 0 bytes
 .../btn_white_loading_1@2x.png                     |   Bin 3944 -> 0 bytes
 .../btn_white_loading_1@3x.png                     |   Bin 6343 -> 0 bytes
 .../btn_white_loading_10.imageset/Contents.json    |    23 -
 .../btn_white_loading_10.png                       |   Bin 1659 -> 0 bytes
 .../btn_white_loading_10@2x.png                    |   Bin 4024 -> 0 bytes
 .../btn_white_loading_10@3x.png                    |   Bin 6417 -> 0 bytes
 .../btn_white_loading_11.imageset/Contents.json    |    23 -
 .../btn_white_loading_11.png                       |   Bin 1654 -> 0 bytes
 .../btn_white_loading_11@2x.png                    |   Bin 4004 -> 0 bytes
 .../btn_white_loading_11@3x.png                    |   Bin 6368 -> 0 bytes
 .../btn_white_loading_12.imageset/Contents.json    |    23 -
 .../btn_white_loading_12.png                       |   Bin 1648 -> 0 bytes
 .../btn_white_loading_12@2x.png                    |   Bin 3940 -> 0 bytes
 .../btn_white_loading_12@3x.png                    |   Bin 6328 -> 0 bytes
 .../btn_white_loading_13.imageset/Contents.json    |    23 -
 .../btn_white_loading_13.png                       |   Bin 1607 -> 0 bytes
 .../btn_white_loading_13@2x.png                    |   Bin 3910 -> 0 bytes
 .../btn_white_loading_13@3x.png                    |   Bin 6249 -> 0 bytes
 .../btn_white_loading_14.imageset/Contents.json    |    23 -
 .../btn_white_loading_14.png                       |   Bin 1598 -> 0 bytes
 .../btn_white_loading_14@2x.png                    |   Bin 3887 -> 0 bytes
 .../btn_white_loading_14@3x.png                    |   Bin 6184 -> 0 bytes
 .../btn_white_loading_15.imageset/Contents.json    |    23 -
 .../btn_white_loading_15.png                       |   Bin 1575 -> 0 bytes
 .../btn_white_loading_15@2x.png                    |   Bin 3849 -> 0 bytes
 .../btn_white_loading_15@3x.png                    |   Bin 6144 -> 0 bytes
 .../btn_white_loading_16.imageset/Contents.json    |    23 -
 .../btn_white_loading_16.png                       |   Bin 1558 -> 0 bytes
 .../btn_white_loading_16@2x.png                    |   Bin 3801 -> 0 bytes
 .../btn_white_loading_16@3x.png                    |   Bin 6084 -> 0 bytes
 .../btn_white_loading_17.imageset/Contents.json    |    23 -
 .../btn_white_loading_17.png                       |   Bin 1575 -> 0 bytes
 .../btn_white_loading_17@2x.png                    |   Bin 3849 -> 0 bytes
 .../btn_white_loading_17@3x.png                    |   Bin 6144 -> 0 bytes
 .../btn_white_loading_18.imageset/Contents.json    |    23 -
 .../btn_white_loading_18.png                       |   Bin 1598 -> 0 bytes
 .../btn_white_loading_18@2x.png                    |   Bin 3887 -> 0 bytes
 .../btn_white_loading_18@3x.png                    |   Bin 6184 -> 0 bytes
 .../btn_white_loading_2.imageset/Contents.json     |    23 -
 .../btn_white_loading_2.png                        |   Bin 1648 -> 0 bytes
 .../btn_white_loading_2@2x.png                     |   Bin 3940 -> 0 bytes
 .../btn_white_loading_2@3x.png                     |   Bin 6328 -> 0 bytes
 .../btn_white_loading_3.imageset/Contents.json     |    23 -
 .../btn_white_loading_3.png                        |   Bin 1691 -> 0 bytes
 .../btn_white_loading_3@2x.png                     |   Bin 4007 -> 0 bytes
 .../btn_white_loading_3@3x.png                     |   Bin 6399 -> 0 bytes
 .../btn_white_loading_4.imageset/Contents.json     |    23 -
 .../btn_white_loading_4.png                        |   Bin 1716 -> 0 bytes
 .../btn_white_loading_4@2x.png                     |   Bin 4038 -> 0 bytes
 .../btn_white_loading_4@3x.png                     |   Bin 6452 -> 0 bytes
 .../btn_white_loading_5.imageset/Contents.json     |    23 -
 .../btn_white_loading_5.png                        |   Bin 1709 -> 0 bytes
 .../btn_white_loading_5@2x.png                     |   Bin 4059 -> 0 bytes
 .../btn_white_loading_5@3x.png                     |   Bin 6481 -> 0 bytes
 .../btn_white_loading_6.imageset/Contents.json     |    23 -
 .../btn_white_loading_6.png                        |   Bin 1721 -> 0 bytes
 .../btn_white_loading_6@2x.png                     |   Bin 4092 -> 0 bytes
 .../btn_white_loading_6@3x.png                     |   Bin 6549 -> 0 bytes
 .../btn_white_loading_7.imageset/Contents.json     |    23 -
 .../btn_white_loading_7.png                        |   Bin 1734 -> 0 bytes
 .../btn_white_loading_7@2x.png                     |   Bin 4093 -> 0 bytes
 .../btn_white_loading_7@3x.png                     |   Bin 6594 -> 0 bytes
 .../btn_white_loading_8.imageset/Contents.json     |    23 -
 .../btn_white_loading_8.png                        |   Bin 1721 -> 0 bytes
 .../btn_white_loading_8@2x.png                     |   Bin 4092 -> 0 bytes
 .../btn_white_loading_8@3x.png                     |   Bin 6549 -> 0 bytes
 .../btn_white_loading_9.imageset/Contents.json     |    23 -
 .../btn_white_loading_9.png                        |   Bin 1709 -> 0 bytes
 .../btn_white_loading_9@2x.png                     |   Bin 4059 -> 0 bytes
 .../btn_white_loading_9@3x.png                     |   Bin 6481 -> 0 bytes
 .../btn_white_direction.imageset/Contents.json     |    23 -
 .../btn_white_direction.png                        |   Bin 1564 -> 0 bytes
 .../btn_white_direction@2x.png                     |   Bin 3753 -> 0 bytes
 .../btn_white_direction@3x.png                     |   Bin 6108 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_white_direction_pressed.png                |   Bin 1186 -> 0 bytes
 .../btn_white_direction_pressed@2x.png             |   Bin 3481 -> 0 bytes
 .../btn_white_direction_pressed@3x.png             |   Bin 6126 -> 0 bytes
 .../btn_white_target_off_1.imageset/Contents.json  |    23 -
 .../btn_white_target_off_1.png                     |   Bin 1336 -> 0 bytes
 .../btn_white_target_off_1@2x.png                  |   Bin 3067 -> 0 bytes
 .../btn_white_target_off_1@3x.png                  |   Bin 5198 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_white_target_off_1_pressed.png             |   Bin 961 -> 0 bytes
 .../btn_white_target_off_1_pressed@2x.png          |   Bin 2865 -> 0 bytes
 .../btn_white_target_off_1_pressed@3x.png          |   Bin 5123 -> 0 bytes
 .../btn_white_target_on.imageset/Contents.json     |    23 -
 .../btn_white_target_on.png                        |   Bin 1569 -> 0 bytes
 .../btn_white_target_on@2x.png                     |   Bin 3755 -> 0 bytes
 .../btn_white_target_on@3x.png                     |   Bin 6008 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_white_target_on_pressed.png                |   Bin 1213 -> 0 bytes
 .../btn_white_target_on_pressed@2x.png             |   Bin 3487 -> 0 bytes
 .../btn_white_target_on_pressed@3x.png             |   Bin 5975 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_white_unknow_position.png                  |   Bin 1239 -> 0 bytes
 .../btn_white_unknow_position@2x.png               |   Bin 2812 -> 0 bytes
 .../btn_white_unknow_position@3x.png               |   Bin 4741 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_white_unknow_position_pressed.png          |   Bin 895 -> 0 bytes
 .../btn_white_unknow_position_pressed@2x.png       |   Bin 2718 -> 0 bytes
 .../btn_white_unknow_position_pressed@3x.png       |   Bin 4909 -> 0 bytes
 .../Place Page/ic_back_api.imageset/Contents.json  |    12 +-
 .../ic_back_api.imageset/ic_back_api.png           |   Bin 113 -> 0 bytes
 .../ic_back_api.imageset/ic_back_api@2x.png        |   Bin 178 -> 0 bytes
 .../ic_back_api.imageset/ic_back_api@3x.png        |   Bin 224 -> 0 bytes
 .../Place Page/ic_back_api.imageset/ic_back_pp.png |   Bin 0 -> 206 bytes
 .../ic_back_api.imageset/ic_back_pp@2x.png         |   Bin 0 -> 311 bytes
 .../ic_back_api.imageset/ic_back_pp@3x.png         |   Bin 0 -> 437 bytes
 .../ic_back_api_pressed.imageset/Contents.json     |    12 +-
 .../ic_back_api_pressed.png                        |   Bin 107 -> 0 bytes
 .../ic_back_api_pressed@2x.png                     |   Bin 153 -> 0 bytes
 .../ic_back_api_pressed@3x.png                     |   Bin 189 -> 0 bytes
 .../ic_back_pp_press.png                           |   Bin 0 -> 196 bytes
 .../ic_back_pp_press@2x.png                        |   Bin 0 -> 287 bytes
 .../ic_back_pp_press@3x.png                        |   Bin 0 -> 400 bytes
 .../ic_bookmarks_off.imageset/Contents.json        |    12 +-
 .../ic_bookmark_add_pp.png                         |   Bin 0 -> 272 bytes
 .../ic_bookmark_add_pp@2x.png                      |   Bin 0 -> 465 bytes
 .../ic_bookmark_add_pp@3x.png                      |   Bin 0 -> 669 bytes
 .../ic_bookmarks_off.imageset/ic_bookmarks_off.png |   Bin 254 -> 0 bytes
 .../ic_bookmarks_off@2x.png                        |   Bin 438 -> 0 bytes
 .../ic_bookmarks_off@3x.png                        |   Bin 637 -> 0 bytes
 .../Contents.json                                  |    12 +-
 .../ic_bookmark_add_pp_press.png                   |   Bin 0 -> 253 bytes
 .../ic_bookmark_add_pp_press@2x.png                |   Bin 0 -> 443 bytes
 .../ic_bookmark_add_pp_press@3x.png                |   Bin 0 -> 617 bytes
 .../ic_bookmarks_off_pressed.png                   |   Bin 235 -> 0 bytes
 .../ic_bookmarks_off_pressed@2x.png                |   Bin 382 -> 0 bytes
 .../ic_bookmarks_off_pressed@3x.png                |   Bin 556 -> 0 bytes
 .../ic_bookmarks_on.imageset/Contents.json         |    12 +-
 .../ic_bookmark_remove_pp.png                      |   Bin 0 -> 407 bytes
 .../ic_bookmark_remove_pp@2x.png                   |   Bin 0 -> 723 bytes
 .../ic_bookmark_remove_pp@3x.png                   |   Bin 0 -> 1045 bytes
 .../ic_bookmarks_on.imageset/ic_bookmarks_on.png   |   Bin 512 -> 0 bytes
 .../ic_bookmarks_on@2x.png                         |   Bin 1034 -> 0 bytes
 .../ic_bookmarks_on@3x.png                         |   Bin 1424 -> 0 bytes
 .../ic_bookmarks_on_pressed.imageset/Contents.json |    12 +-
 .../ic_bookmark_remove_pp_press.png                |   Bin 0 -> 401 bytes
 .../ic_bookmark_remove_pp_press@2x.png             |   Bin 0 -> 719 bytes
 .../ic_bookmark_remove_pp_press@3x.png             |   Bin 0 -> 1032 bytes
 .../ic_bookmarks_on_pressed.png                    |   Bin 541 -> 0 bytes
 .../ic_bookmarks_on_pressed@2x.png                 |   Bin 1254 -> 0 bytes
 .../ic_bookmarks_on_pressed@3x.png                 |   Bin 2040 -> 0 bytes
 .../Place Page/ic_route.imageset/Contents.json     |    12 +-
 .../Place Page/ic_route.imageset/ic_route.png      |   Bin 227 -> 0 bytes
 .../Place Page/ic_route.imageset/ic_route@2x.png   |   Bin 362 -> 0 bytes
 .../Place Page/ic_route.imageset/ic_route@3x.png   |   Bin 520 -> 0 bytes
 .../Place Page/ic_route.imageset/ic_route_to.png   |   Bin 0 -> 219 bytes
 .../ic_route.imageset/ic_route_to@2x.png           |   Bin 0 -> 366 bytes
 .../ic_route.imageset/ic_route_to@3x.png           |   Bin 0 -> 513 bytes
 .../ic_route_pressed.imageset/Contents.json        |    12 +-
 .../ic_route_pressed.imageset/ic_route_pressed.png |   Bin 207 -> 0 bytes
 .../ic_route_pressed@2x.png                        |   Bin 334 -> 0 bytes
 .../ic_route_pressed@3x.png                        |   Bin 481 -> 0 bytes
 .../ic_route_to_press.png                          |   Bin 0 -> 195 bytes
 .../ic_route_to_press@2x.png                       |   Bin 0 -> 327 bytes
 .../ic_route_to_press@3x.png                       |   Bin 0 -> 460 bytes
 .../Place Page/ic_share.imageset/Contents.json     |    12 +-
 .../Place Page/ic_share.imageset/ic_share.png      |   Bin 257 -> 0 bytes
 .../Place Page/ic_share.imageset/ic_share@2x.png   |   Bin 454 -> 0 bytes
 .../Place Page/ic_share.imageset/ic_share@3x.png   |   Bin 666 -> 0 bytes
 .../Place Page/ic_share.imageset/ic_share_pp.png   |   Bin 0 -> 230 bytes
 .../ic_share.imageset/ic_share_pp@2x.png           |   Bin 0 -> 385 bytes
 .../ic_share.imageset/ic_share_pp@3x.png           |   Bin 0 -> 565 bytes
 .../ic_share_pressed.imageset/Contents.json        |    12 +-
 .../ic_share_pp_press.png                          |   Bin 0 -> 216 bytes
 .../ic_share_pp_press@2x.png                       |   Bin 0 -> 363 bytes
 .../ic_share_pp_press@3x.png                       |   Bin 0 -> 529 bytes
 .../ic_share_pressed.imageset/ic_share_pressed.png |   Bin 230 -> 0 bytes
 .../ic_share_pressed@2x.png                        |   Bin 406 -> 0 bytes
 .../ic_share_pressed@3x.png                        |   Bin 586 -> 0 bytes
 .../btn_green_bookmarks.imageset/Contents.json     |    23 -
 .../btn_green_bookmarks.png                        |   Bin 2232 -> 0 bytes
 .../btn_green_bookmarks@2x.png                     |   Bin 5332 -> 0 bytes
 .../btn_green_bookmarks@3x.png                     |   Bin 8781 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_green_bookmarks_pressed.png                |   Bin 1491 -> 0 bytes
 .../btn_green_bookmarks_pressed@2x.png             |   Bin 4679 -> 0 bytes
 .../btn_green_bookmarks_pressed@3x.png             |   Bin 7951 -> 0 bytes
 .../btn_green_download.imageset/Contents.json      |    23 -
 .../btn_green_download.png                         |   Bin 1943 -> 0 bytes
 .../btn_green_download@2x.png                      |   Bin 4327 -> 0 bytes
 .../btn_green_download@3x.png                      |   Bin 7277 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_green_download_pressed.png                 |   Bin 1229 -> 0 bytes
 .../btn_green_download_pressed@2x.png              |   Bin 3730 -> 0 bytes
 .../btn_green_download_pressed@3x.png              |   Bin 6462 -> 0 bytes
 .../btn_green_menu_1.imageset/Contents.json        |    23 -
 .../btn_green_menu_1.imageset/btn_green_menu.png   |   Bin 1928 -> 0 bytes
 .../btn_green_menu@2x.png                          |   Bin 4524 -> 0 bytes
 .../btn_green_menu@3x.png                          |   Bin 7503 -> 0 bytes
 .../btn_green_menu_2.imageset/Contents.json        |    23 -
 .../btn_green_menu_2.imageset/btn_green_menu_2.png |   Bin 1905 -> 0 bytes
 .../btn_green_menu_2@2x.png                        |   Bin 4466 -> 0 bytes
 .../btn_green_menu_2@3x.png                        |   Bin 7375 -> 0 bytes
 .../btn_green_menu_3.imageset/Contents.json        |    23 -
 .../btn_green_menu_3.imageset/btn_green_menu_3.png |   Bin 1885 -> 0 bytes
 .../btn_green_menu_3@2x.png                        |   Bin 4466 -> 0 bytes
 .../btn_green_menu_3@3x.png                        |   Bin 7405 -> 0 bytes
 .../btn_green_menu_4.imageset/Contents.json        |    23 -
 .../btn_green_menu_4.imageset/btn_green_menu_4.png |   Bin 1906 -> 0 bytes
 .../btn_green_menu_4@2x.png                        |   Bin 4504 -> 0 bytes
 .../btn_green_menu_4@3x.png                        |   Bin 7394 -> 0 bytes
 .../btn_green_menu_pressed.imageset/Contents.json  |    23 -
 .../btn_green_menu_pressed.png                     |   Bin 1189 -> 0 bytes
 .../btn_green_menu_pressed@2x.png                  |   Bin 3927 -> 0 bytes
 .../btn_green_menu_pressed@3x.png                  |   Bin 6672 -> 0 bytes
 .../btn_green_search.imageset/Contents.json        |    23 -
 .../btn_green_search.imageset/btn_green_search.png |   Bin 1961 -> 0 bytes
 .../btn_green_search@2x.png                        |   Bin 4364 -> 0 bytes
 .../btn_green_search@3x.png                        |   Bin 7487 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_green_search_pressed.png                   |   Bin 1337 -> 0 bytes
 .../btn_green_search_pressed@2x.png                |   Bin 3790 -> 0 bytes
 .../btn_green_search_pressed@3x.png                |   Bin 6571 -> 0 bytes
 .../btn_green_settings.imageset/Contents.json      |    23 -
 .../btn_green_settings.png                         |   Bin 2530 -> 0 bytes
 .../btn_green_settings@2x.png                      |   Bin 5967 -> 0 bytes
 .../btn_green_settings@3x.png                      |   Bin 9748 -> 0 bytes
 .../Contents.json                                  |    23 -
 .../btn_green_settings_pressed.png                 |   Bin 1745 -> 0 bytes
 .../btn_green_settings_pressed@2x.png              |   Bin 5206 -> 0 bytes
 .../btn_green_settings_pressed@3x.png              |   Bin 8654 -> 0 bytes
 .../btn_green_share.imageset/Contents.json         |    23 -
 .../btn_green_share.imageset/btn_green_share.png   |   Bin 2168 -> 0 bytes
 .../btn_green_share@2x.png                         |   Bin 5181 -> 0 bytes
 .../btn_green_share@3x.png                         |   Bin 8447 -> 0 bytes
 .../btn_green_share_pressed.imageset/Contents.json |    23 -
 .../btn_green_share_pressed.png                    |   Bin 1432 -> 0 bytes
 .../btn_green_share_pressed@2x.png                 |   Bin 4458 -> 0 bytes
 .../btn_green_share_pressed@3x.png                 |   Bin 7401 -> 0 bytes
 .../Side Menu/img_counter.imageset/Contents.json   |    23 -
 .../Side Menu/img_counter.imageset/img_counter.png |   Bin 246 -> 0 bytes
 .../img_counter.imageset/img_counter@2x.png        |   Bin 392 -> 0 bytes
 .../img_counter.imageset/img_counter@3x.png        |   Bin 573 -> 0 bytes
 iphone/Maps/Maps.xcodeproj/project.pbxproj         |   114 +-
 iphone/Maps/UIColor+MapsMeColor.h                  |     2 +
 iphone/Maps/UIColor+MapsMeColor.mm                 |     9 +
 iphone/Maps/ja.lproj/Localizable.strings           |     2 +-
 map/address_finder.cpp                             |     4 +-
 map/framework.cpp                                  |   103 +-
 map/framework.hpp                                  |    14 +-
 map/information_display.cpp                        |     2 +-
 omim.pro                                           |     4 +-
 platform/location.hpp                              |     4 +
 qt/res/Info.plist                                  |    23 +
 routing/osrm_data_facade.hpp                       |     2 +-
 routing/osrm_router.cpp                            |   104 +-
 routing/route.cpp                                  |    24 +-
 routing/route.hpp                                  |    14 +-
 routing/routing.pro                                |     2 +
 routing/routing_integration_tests/jni/Android.mk   |    15 +
 .../routing_integration_tests/jni/Application.mk   |     1 +
 routing/routing_integration_tests/jni/test.cpp     |    15 +
 routing/routing_integration_tests/jni/test.hpp     |     9 +
 .../online_cross_tests.cpp                         |    49 +
 .../routing_integration_tests/osrm_route_test.cpp  |   237 +
 .../routing_integration_tests/osrm_turn_test.cpp   |   350 +
 .../pedestrian_route_test.cpp                      |   509 +
 .../routing_integration_tests.pro                  |    36 +
 .../routing_test_tools.cpp                         |   336 +
 .../routing_test_tools.hpp                         |     0
 routing/routing_session.cpp                        |   216 +-
 routing/routing_session.hpp                        |    23 +-
 routing/routing_settings.hpp                       |     7 +-
 routing/routing_tests/route_tests.cpp              |    50 +
 routing/routing_tests/turns_sound_test.cpp         |   195 +-
 routing/routing_tests/turns_tts_text_tests.cpp     |     4 +-
 routing/speed_camera.cpp                           |    65 +
 routing/speed_camera.hpp                           |    14 +
 routing/turns.cpp                                  |     9 +
 routing/turns.hpp                                  |    13 +
 routing/turns_sound.cpp                            |    81 +-
 routing/turns_sound.hpp                            |    24 +-
 routing/turns_sound_settings.hpp                   |     1 +
 routing/turns_tts_text.cpp                         |    23 +-
 routing/turns_tts_text.hpp                         |     3 +-
 search/intermediate_result.cpp                     |    24 +-
 search/intermediate_result.hpp                     |    18 +-
 search/search.pro                                  |     1 +
 search/search_engine.cpp                           |   124 +-
 search/search_engine.hpp                           |    37 +-
 .../test_search_engine.cpp                         |    28 +-
 .../test_search_engine.hpp                         |     3 +
 search/search_query.cpp                            |   109 +-
 search/search_query.hpp                            |    33 +-
 search/search_query_factory.hpp                    |    15 +-
 search/suggest.hpp                                 |    20 +
 sound/tts/sound.txt                                |     6 +-
 storage/country_decl.hpp                           |     4 +-
 storage/country_info.cpp                           |   257 -
 storage/country_info.hpp                           |    84 -
 storage/country_info_getter.cpp                    |   190 +
 storage/country_info_getter.hpp                    |    85 +
 storage/storage.pro                                |     4 +-
 storage/storage_tests/country_info_getter_test.cpp |    81 +
 storage/storage_tests/country_info_test.cpp        |    83 -
 storage/storage_tests/storage_tests.pro            |     2 +-
 strings.txt                                        |    46 +-
 tools/kothic                                       |     2 +-
 tools/osmctools                                    |     2 +-
 tools/publish_to_dropbox.py                        |    91 -
 tools/unix/batch_generate.sh                       |    14 -
 tools/unix/build_omim.sh                           |     4 +-
 tools/unix/clean_indexes.sh                        |    17 +-
 tools/unix/compress_section_files.sh               |     6 -
 tools/unix/create_mac_image_debug.sh               |    10 -
 tools/unix/create_mac_image_release.sh             |    10 -
 tools/unix/ensure_mkspecs.sh                       |     4 -
 tools/unix/external_resources.sh                   |    14 -
 tools/unix/generate_mwm.sh                         |     2 +-
 tools/unix/generate_planet.sh                      |     4 +-
 tools/unix/generate_symbols.sh                     |     6 +-
 tools/unix/make_index_of.sh                        |   108 -
 tools/unix/make_index_of_vng.sh                    |    30 -
 tools/unix/make_routing.sh                         |    12 -
 tools/unix/osrm_generator.sh                       |    77 -
 tools/unix/planet.sh                               |   242 -
 tools/unix/routes_from_mwm_and_osrm.sh             |    20 -
 tools/unix/split_planet_by_polygons.sh             |    23 -
 tools/unix/test_planet.sh                          |     2 +-
 tools/upload_to_dropbox.sh                         |    36 +
 xcode/MapsMe/MapsMe.xcodeproj/project.pbxproj      |     8 +
 xcode/agg/agg.xcodeproj/project.pbxproj            |   765 ++
 xcode/coding/coding.xcodeproj/project.pbxproj      |   356 +-
 xcode/lodepng/lodepng.xcodeproj/project.pbxproj    |   260 +
 xcode/omim.xcworkspace/contents.xcworkspacedata    |    13 +
 xcode/render/render.xcodeproj/project.pbxproj      |     4 -
 xcode/routing/routing.xcodeproj/project.pbxproj    |     8 +
 xcode/storage/storage.xcodeproj/project.pbxproj    |    16 +-
 1320 files changed, 33610 insertions(+), 28746 deletions(-)
 create mode 100644 data/styles/clear/style-clear/symbols/cemetery-l.svg
 create mode 100644 data/styles/clear/style-clear/symbols/cemetery-m.svg
 create mode 100644 data/styles/clear/style-clear/symbols/cemetery-s.svg
 create mode 100644 data/styles/clear/style-clear/symbols/military-l.svg
 create mode 100644 data/styles/clear/style-clear/symbols/military-m.svg
 create mode 100644 data/styles/clear/style-clear/symbols/military-s.svg
 create mode 100644 data/styles/clear/style-clear/symbols/nparkF-l.svg
 create mode 100644 data/styles/clear/style-clear/symbols/nparkF-m.svg
 create mode 100644 data/styles/clear/style-clear/symbols/nparkF-s.svg
 create mode 100644 data/styles/clear/style-night/symbols/cemetery-l.svg
 create mode 100644 data/styles/clear/style-night/symbols/cemetery-m.svg
 create mode 100644 data/styles/clear/style-night/symbols/cemetery-s.svg
 create mode 100644 data/styles/clear/style-night/symbols/military-l.svg
 create mode 100644 data/styles/clear/style-night/symbols/military-m.svg
 create mode 100644 data/styles/clear/style-night/symbols/military-s.svg
 create mode 100644 data/styles/clear/style-night/symbols/nparkF-l.svg
 create mode 100644 data/styles/clear/style-night/symbols/nparkF-m.svg
 create mode 100644 data/styles/clear/style-night/symbols/nparkF-s.svg
 create mode 100644 docs/MAPS.md
 delete mode 100644 docs/MWM.md
 delete mode 100644 integration_tests/integration_tests.pro
 delete mode 100644 integration_tests/jni/Android.mk
 delete mode 120000 integration_tests/jni/Application.mk
 delete mode 100644 integration_tests/jni/test.cpp
 delete mode 100644 integration_tests/jni/test.hpp
 delete mode 100644 integration_tests/online_cross_tests.cpp
 delete mode 100644 integration_tests/osrm_route_test.cpp
 delete mode 100644 integration_tests/osrm_turn_test.cpp
 delete mode 100644 integration_tests/pedestrian_route_test.cpp
 delete mode 100644 integration_tests/routing_test_tools.cpp
 rename iphone/Maps/Classes/CustomViews/MapViewControls/{SideMenu => APIBar}/MWMMapViewControlsCommon.h (100%)
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuCollectionViewCell.h
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuCollectionViewCell.mm
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuCollectionViewLandscapeCell.xib
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuCollectionViewPortraitCell.xib
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuLayout.h
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuLayout.mm
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuView.h
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuView.mm
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuViewController.h
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuViewController.mm
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/BottomMenu/MWMBottomMenuViewController.xib
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/LocationButton/MWMLocationButton.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/LocationButton/MWMLocationButton.mm
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/LocationButton/MWMLocationButton.xib
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/LocationButton/MWMLocationButtonView.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/LocationButton/MWMLocationButtonView.mm
 create mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/MWMMapViewControlsCommon.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuButton.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuButton.mm
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuButtonDelegate.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuDelegate.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuDownloadBadge.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuDownloadBadge.m
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuManager.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuManager.mm
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuManagerDelegate.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuView.h
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuView.mm
 delete mode 100644 iphone/Maps/Classes/CustomViews/MapViewControls/SideMenu/MWMSideMenuViews.xib
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_1.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_1.imageset/ic_follow_mode_1_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_1.imageset/ic_follow_mode_1_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_1.imageset/ic_follow_mode_1_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_2.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_2.imageset/ic_follow_mode_2_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_2.imageset/ic_follow_mode_2_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_2.imageset/ic_follow_mode_2_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_3.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_3.imageset/ic_follow_mode_3_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_3.imageset/ic_follow_mode_3_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_3.imageset/ic_follow_mode_3_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_4.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_4.imageset/ic_follow_mode_4_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_4.imageset/ic_follow_mode_4_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_4.imageset/ic_follow_mode_4_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_5.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_5.imageset/ic_follow_mode_5_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_5.imageset/ic_follow_mode_5_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_5.imageset/ic_follow_mode_5_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_6.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_6.imageset/ic_follow_mode_5_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_6.imageset/ic_follow_mode_5_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_follow_mode_light_6.imageset/ic_follow_mode_5_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_1.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_1.imageset/ic_menu_1.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_1.imageset/ic_menu_1@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_1.imageset/ic_menu_1@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_2.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_2.imageset/ic_menu_2.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_2.imageset/ic_menu_2@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_2.imageset/ic_menu_2@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_3.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_3.imageset/ic_menu_3.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_3.imageset/ic_menu_3@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_3.imageset/ic_menu_3@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_4.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_4.imageset/ic_menu_4.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_4.imageset/ic_menu_4@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_4.imageset/ic_menu_4@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_5.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_5.imageset/ic_menu_5.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_5.imageset/ic_menu_5@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_5.imageset/ic_menu_5@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_6.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_6.imageset/ic_menu_6.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_6.imageset/ic_menu_6@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_6.imageset/ic_menu_6@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_1.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_1.imageset/ic_menu_rotate_1.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_1.imageset/ic_menu_rotate_1@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_1.imageset/ic_menu_rotate_1@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_2.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_2.imageset/ic_menu_rotate_2.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_2.imageset/ic_menu_rotate_2@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_2.imageset/ic_menu_rotate_2@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_3.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_3.imageset/ic_menu_rotate_3.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_3.imageset/ic_menu_rotate_3@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_3.imageset/ic_menu_rotate_3@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_4.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_4.imageset/ic_menu_rotate_4.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_4.imageset/ic_menu_rotate_4@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_4.imageset/ic_menu_rotate_4@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_5.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_5.imageset/ic_menu_rotate_5.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_5.imageset/ic_menu_rotate_5@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_5.imageset/ic_menu_rotate_5@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_6.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_6.imageset/ic_menu_rotate_6.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_6.imageset/ic_menu_rotate_6@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/Morphing/ic_menu_rotate_6.imageset/ic_menu_rotate_6@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu.imageset/ic_menu.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu.imageset/ic_menu@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu.imageset/ic_menu@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light.imageset/ic_bookmark_list_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light.imageset/ic_bookmark_list_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light.imageset/ic_bookmark_list_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light_press.imageset/ic_bookmark_list_light_press@1x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light_press.imageset/ic_bookmark_list_light_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_bookmark_list_light_press.imageset/ic_bookmark_list_light_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down.imageset/ic_menu_down.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down.imageset/ic_menu_down@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down.imageset/ic_menu_down@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down_press.imageset/ic_menu_down_arrow_light_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down_press.imageset/ic_menu_down_arrow_light_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_down_press.imageset/ic_menu_down_arrow_light_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download.imageset/ic_download@1x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download.imageset/ic_download@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download.imageset/ic_download@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download_press.imageset/ic_download_press@1x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download_press.imageset/ic_download_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_download_press.imageset/ic_download_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left.imageset/ic_menu_left_arrow.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left.imageset/ic_menu_left_arrow@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left.imageset/ic_menu_left_arrow@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left_press.imageset/ic_menu_left_arrow_light_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left_press.imageset/ic_menu_left_arrow_light_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_left_press.imageset/ic_menu_left_arrow_light_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow.imageset/ic_follow.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow.imageset/ic_follow@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow.imageset/ic_follow@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow_and_rotate.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow_and_rotate.imageset/ic_follow_and_rotate.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow_and_rotate.imageset/ic_follow_and_rotate@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_follow_and_rotate.imageset/ic_follow_and_rotate@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_get_position.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_get_position.imageset/ic_get_position.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_get_position.imageset/ic_get_position@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_get_position.imageset/ic_get_position@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_off_mode_light.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_off_mode_light.imageset/ic_off_mode_light.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_off_mode_light.imageset/ic_off_mode_light@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_off_mode_light.imageset/ic_off_mode_light@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_pending.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_pending.imageset/ic_pending_1.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_pending.imageset/ic_pending_1@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_location_pending.imageset/ic_pending_1@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point.imageset/ic_point_to_point.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point.imageset/ic_point_to_point@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point.imageset/ic_point_to_point@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point_press.imageset/ic_point_to_point_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point_press.imageset/ic_point_to_point_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_point_to_point_press.imageset/ic_point_to_point_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_press.imageset/ic_menu_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_press.imageset/ic_menu_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_press.imageset/ic_menu_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search.imageset/ic_search.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search.imageset/ic_search@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search.imageset/ic_search@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search_press.imageset/ic_search_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search_press.imageset/ic_search_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_search_press.imageset/ic_search_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings.imageset/ic_settings@1x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings.imageset/ic_settings@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings.imageset/ic_settings@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings_press.imageset/ic_settings_press@1x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings_press.imageset/ic_settings_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_settings_press.imageset/ic_settings_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share.imageset/ic_share@1x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share.imageset/ic_share@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share.imageset/ic_share@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share_press.imageset/Contents.json
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share_press.imageset/ic_share_press@1x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share_press.imageset/ic_share_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Bottom Menu/ic_menu_share_press.imageset/ic_share_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_1.imageset/follow_to_followandrotate_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_1.imageset/follow_to_followandrotate_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_1.imageset/follow_to_followandrotate_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_2.imageset/follow_to_followandrotate_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_2.imageset/follow_to_followandrotate_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_2.imageset/follow_to_followandrotate_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_3.imageset/follow_to_followandrotate_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_3.imageset/follow_to_followandrotate_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_3.imageset/follow_to_followandrotate_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_4.imageset/follow_to_followandrotate_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_4.imageset/follow_to_followandrotate_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_4.imageset/follow_to_followandrotate_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_5.imageset/follow_to_followandrotate_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_5.imageset/follow_to_followandrotate_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_5.imageset/follow_to_followandrotate_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_6.imageset/follow_to_followandrotate_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_6.imageset/follow_to_followandrotate_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_followandrotate_6.imageset/follow_to_followandrotate_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_1.imageset/follow_to_noposition_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_1.imageset/follow_to_noposition_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_1.imageset/follow_to_noposition_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_2.imageset/follow_to_noposition_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_2.imageset/follow_to_noposition_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_2.imageset/follow_to_noposition_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_3.imageset/follow_to_noposition_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_3.imageset/follow_to_noposition_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_3.imageset/follow_to_noposition_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_4.imageset/follow_to_noposition_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_4.imageset/follow_to_noposition_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_4.imageset/follow_to_noposition_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_5.imageset/follow_to_noposition_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_5.imageset/follow_to_noposition_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_5.imageset/follow_to_noposition_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_6.imageset/follow_to_noposition_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_6.imageset/follow_to_noposition_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_noposition_6.imageset/follow_to_noposition_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_1.imageset/follow_to_notfollow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_1.imageset/follow_to_notfollow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_1.imageset/follow_to_notfollow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_2.imageset/follow_to_notfollow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_2.imageset/follow_to_notfollow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_2.imageset/follow_to_notfollow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_3.imageset/follow_to_notfollow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_3.imageset/follow_to_notfollow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_3.imageset/follow_to_notfollow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_4.imageset/follow_to_notfollow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_4.imageset/follow_to_notfollow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_4.imageset/follow_to_notfollow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_5.imageset/follow_to_notfollow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_5.imageset/follow_to_notfollow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_5.imageset/follow_to_notfollow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_6.imageset/follow_to_notfollow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_6.imageset/follow_to_notfollow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_notfollow_6.imageset/follow_to_notfollow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_1.imageset/follow_to_pending_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_1.imageset/follow_to_pending_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_1.imageset/follow_to_pending_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_2.imageset/follow_to_pending_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_2.imageset/follow_to_pending_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_2.imageset/follow_to_pending_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_3.imageset/follow_to_pending_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_3.imageset/follow_to_pending_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_3.imageset/follow_to_pending_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_4.imageset/follow_to_pending_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_4.imageset/follow_to_pending_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_4.imageset/follow_to_pending_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_5.imageset/follow_to_pending_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_5.imageset/follow_to_pending_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_5.imageset/follow_to_pending_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_6.imageset/follow_to_pending_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_6.imageset/follow_to_pending_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/follow_to_pending_6.imageset/follow_to_pending_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_1.imageset/followandrotate_to_follow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_1.imageset/followandrotate_to_follow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_1.imageset/followandrotate_to_follow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_2.imageset/followandrotate_to_follow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_2.imageset/followandrotate_to_follow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_2.imageset/followandrotate_to_follow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_3.imageset/followandrotate_to_follow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_3.imageset/followandrotate_to_follow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_3.imageset/followandrotate_to_follow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_4.imageset/followandrotate_to_follow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_4.imageset/followandrotate_to_follow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_4.imageset/followandrotate_to_follow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_5.imageset/followandrotate_to_follow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_5.imageset/followandrotate_to_follow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_5.imageset/followandrotate_to_follow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_6.imageset/followandrotate_to_follow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_6.imageset/followandrotate_to_follow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_follow_6.imageset/followandrotate_to_follow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_1.imageset/followandrotate_to_noposition_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_1.imageset/followandrotate_to_noposition_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_1.imageset/followandrotate_to_noposition_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_2.imageset/followandrotate_to_noposition_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_2.imageset/followandrotate_to_noposition_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_2.imageset/followandrotate_to_noposition_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_3.imageset/followandrotate_to_noposition_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_3.imageset/followandrotate_to_noposition_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_3.imageset/followandrotate_to_noposition_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_4.imageset/followandrotate_to_noposition_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_4.imageset/followandrotate_to_noposition_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_4.imageset/followandrotate_to_noposition_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_5.imageset/followandrotate_to_noposition_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_5.imageset/followandrotate_to_noposition_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_5.imageset/followandrotate_to_noposition_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_6.imageset/followandrotate_to_noposition_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_6.imageset/followandrotate_to_noposition_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_noposition_6.imageset/followandrotate_to_noposition_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_1.imageset/followandrotate_to_notfollow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_1.imageset/followandrotate_to_notfollow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_1.imageset/followandrotate_to_notfollow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_2.imageset/followandrotate_to_notfollow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_2.imageset/followandrotate_to_notfollow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_2.imageset/followandrotate_to_notfollow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_3.imageset/followandrotate_to_notfollow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_3.imageset/followandrotate_to_notfollow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_3.imageset/followandrotate_to_notfollow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_4.imageset/followandrotate_to_notfollow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_4.imageset/followandrotate_to_notfollow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_4.imageset/followandrotate_to_notfollow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_5.imageset/followandrotate_to_notfollow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_5.imageset/followandrotate_to_notfollow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_5.imageset/followandrotate_to_notfollow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_6.imageset/followandrotate_to_notfollow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_6.imageset/followandrotate_to_notfollow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_notfollow_6.imageset/followandrotate_to_notfollow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_1.imageset/followandrotate_to_pending_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_1.imageset/followandrotate_to_pending_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_1.imageset/followandrotate_to_pending_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_2.imageset/followandrotate_to_pending_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_2.imageset/followandrotate_to_pending_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_2.imageset/followandrotate_to_pending_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_3.imageset/followandrotate_to_pending_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_3.imageset/followandrotate_to_pending_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_3.imageset/followandrotate_to_pending_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_4.imageset/followandrotate_to_pending_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_4.imageset/followandrotate_to_pending_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_4.imageset/followandrotate_to_pending_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_5.imageset/followandrotate_to_pending_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_5.imageset/followandrotate_to_pending_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_5.imageset/followandrotate_to_pending_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_6.imageset/followandrotate_to_pending_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_6.imageset/followandrotate_to_pending_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/followandrotate_to_pending_6.imageset/followandrotate_to_pending_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_1.imageset/noposition_to_follow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_1.imageset/noposition_to_follow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_1.imageset/noposition_to_follow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_2.imageset/noposition_to_follow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_2.imageset/noposition_to_follow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_2.imageset/noposition_to_follow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_3.imageset/noposition_to_follow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_3.imageset/noposition_to_follow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_3.imageset/noposition_to_follow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_4.imageset/noposition_to_follow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_4.imageset/noposition_to_follow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_4.imageset/noposition_to_follow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_5.imageset/noposition_to_follow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_5.imageset/noposition_to_follow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_5.imageset/noposition_to_follow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_6.imageset/noposition_to_follow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_6.imageset/noposition_to_follow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_follow_6.imageset/noposition_to_follow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_1.imageset/noposition_to_followandrotate_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_1.imageset/noposition_to_followandrotate_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_1.imageset/noposition_to_followandrotate_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_2.imageset/noposition_to_followandrotate_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_2.imageset/noposition_to_followandrotate_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_2.imageset/noposition_to_followandrotate_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_3.imageset/noposition_to_followandrotate_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_3.imageset/noposition_to_followandrotate_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_3.imageset/noposition_to_followandrotate_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_4.imageset/noposition_to_followandrotate_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_4.imageset/noposition_to_followandrotate_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_4.imageset/noposition_to_followandrotate_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_5.imageset/noposition_to_followandrotate_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_5.imageset/noposition_to_followandrotate_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_5.imageset/noposition_to_followandrotate_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_6.imageset/noposition_to_followandrotate_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_6.imageset/noposition_to_followandrotate_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_followandrotate_6.imageset/noposition_to_followandrotate_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_1.imageset/noposition_to_notfollow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_1.imageset/noposition_to_notfollow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_1.imageset/noposition_to_notfollow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_2.imageset/noposition_to_notfollow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_2.imageset/noposition_to_notfollow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_2.imageset/noposition_to_notfollow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_3.imageset/noposition_to_notfollow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_3.imageset/noposition_to_notfollow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_3.imageset/noposition_to_notfollow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_4.imageset/noposition_to_notfollow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_4.imageset/noposition_to_notfollow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_4.imageset/noposition_to_notfollow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_5.imageset/noposition_to_notfollow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_5.imageset/noposition_to_notfollow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_5.imageset/noposition_to_notfollow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_6.imageset/noposition_to_notfollow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_6.imageset/noposition_to_notfollow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_notfollow_6.imageset/noposition_to_notfollow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_1.imageset/noposition_to_pending_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_1.imageset/noposition_to_pending_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_1.imageset/noposition_to_pending_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_2.imageset/noposition_to_pending_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_2.imageset/noposition_to_pending_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_2.imageset/noposition_to_pending_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_3.imageset/noposition_to_pending_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_3.imageset/noposition_to_pending_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_3.imageset/noposition_to_pending_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_4.imageset/noposition_to_pending_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_4.imageset/noposition_to_pending_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_4.imageset/noposition_to_pending_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_5.imageset/noposition_to_pending_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_5.imageset/noposition_to_pending_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_5.imageset/noposition_to_pending_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_6.imageset/noposition_to_pending_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_6.imageset/noposition_to_pending_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/noposition_to_pending_6.imageset/noposition_to_pending_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_1.imageset/notfollow_to_follow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_1.imageset/notfollow_to_follow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_1.imageset/notfollow_to_follow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_2.imageset/notfollow_to_follow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_2.imageset/notfollow_to_follow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_2.imageset/notfollow_to_follow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_3.imageset/notfollow_to_follow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_3.imageset/notfollow_to_follow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_3.imageset/notfollow_to_follow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_4.imageset/notfollow_to_follow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_4.imageset/notfollow_to_follow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_4.imageset/notfollow_to_follow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_5.imageset/notfollow_to_follow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_5.imageset/notfollow_to_follow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_5.imageset/notfollow_to_follow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_6.imageset/notfollow_to_follow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_6.imageset/notfollow_to_follow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_follow_6.imageset/notfollow_to_follow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_1.imageset/notfollow_to_followandrotate_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_1.imageset/notfollow_to_followandrotate_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_1.imageset/notfollow_to_followandrotate_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_2.imageset/notfollow_to_followandrotate_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_2.imageset/notfollow_to_followandrotate_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_2.imageset/notfollow_to_followandrotate_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_3.imageset/notfollow_to_followandrotate_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_3.imageset/notfollow_to_followandrotate_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_3.imageset/notfollow_to_followandrotate_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_4.imageset/notfollow_to_followandrotate_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_4.imageset/notfollow_to_followandrotate_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_4.imageset/notfollow_to_followandrotate_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_5.imageset/notfollow_to_followandrotate_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_5.imageset/notfollow_to_followandrotate_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_5.imageset/notfollow_to_followandrotate_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_6.imageset/notfollow_to_followandrotate_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_6.imageset/notfollow_to_followandrotate_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_followandrotate_6.imageset/notfollow_to_followandrotate_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_1.imageset/notfollow_to_noposition_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_1.imageset/notfollow_to_noposition_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_1.imageset/notfollow_to_noposition_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_2.imageset/notfollow_to_noposition_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_2.imageset/notfollow_to_noposition_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_2.imageset/notfollow_to_noposition_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_3.imageset/notfollow_to_noposition_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_3.imageset/notfollow_to_noposition_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_3.imageset/notfollow_to_noposition_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_4.imageset/nofollow_to_noposition_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_4.imageset/nofollow_to_noposition_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_4.imageset/nofollow_to_noposition_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_5.imageset/notfollow_to_noposition_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_5.imageset/notfollow_to_noposition_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_5.imageset/notfollow_to_noposition_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_6.imageset/notfollow_to_noposition_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_6.imageset/notfollow_to_noposition_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_noposition_6.imageset/notfollow_to_noposition_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_1.imageset/notfollow_to_pending_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_1.imageset/notfollow_to_pending_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_1.imageset/notfollow_to_pending_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_2.imageset/notfollow_to_pending_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_2.imageset/notfollow_to_pending_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_2.imageset/notfollow_to_pending_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_3.imageset/notfollow_to_pending_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_3.imageset/notfollow_to_pending_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_3.imageset/notfollow_to_pending_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_4.imageset/notfollow_to_pending_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_4.imageset/notfollow_to_pending_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_4.imageset/notfollow_to_pending_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_5.imageset/notfollow_to_pending_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_5.imageset/notfollow_to_pending_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_5.imageset/notfollow_to_pending_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_6.imageset/notfollow_to_pending_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_6.imageset/notfollow_to_pending_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/notfollow_to_pending_6.imageset/notfollow_to_pending_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_1.imageset/pending_to_follow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_1.imageset/pending_to_follow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_1.imageset/pending_to_follow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_2.imageset/pending_to_follow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_2.imageset/pending_to_follow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_2.imageset/pending_to_follow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_3.imageset/pending_to_follow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_3.imageset/pending_to_follow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_3.imageset/pending_to_follow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_4.imageset/pending_to_follow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_4.imageset/pending_to_follow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_4.imageset/pending_to_follow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_5.imageset/pending_to_follow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_5.imageset/pending_to_follow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_5.imageset/pending_to_follow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_6.imageset/pending_to_follow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_6.imageset/pending_to_follow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_follow_6.imageset/pending_to_follow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_1.imageset/pending_to_followandrotate_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_1.imageset/pending_to_followandrotate_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_1.imageset/pending_to_followandrotate_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_2.imageset/pending_to_followandrotate_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_2.imageset/pending_to_followandrotate_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_2.imageset/pending_to_followandrotate_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_3.imageset/pending_to_followandrotate_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_3.imageset/pending_to_followandrotate_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_3.imageset/pending_to_followandrotate_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_4.imageset/pending_to_followandrotate_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_4.imageset/pending_to_followandrotate_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_4.imageset/pending_to_followandrotate_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_5.imageset/pending_to_followandrotate_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_5.imageset/pending_to_followandrotate_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_5.imageset/pending_to_followandrotate_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_6.imageset/pending_to_followandrotate_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_6.imageset/pending_to_followandrotate_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_followandrotate_6.imageset/pending_to_followandrotate_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_1.imageset/pending_to_noposition_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_1.imageset/pending_to_noposition_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_1.imageset/pending_to_noposition_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_2.imageset/pending_to_noposition_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_2.imageset/pending_to_noposition_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_2.imageset/pending_to_noposition_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_3.imageset/pending_to_noposition_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_3.imageset/pending_to_noposition_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_3.imageset/pending_to_noposition_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_4.imageset/pending_to_noposition_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_4.imageset/pending_to_noposition_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_4.imageset/pending_to_noposition_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_5.imageset/pending_to_noposition_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_5.imageset/pending_to_noposition_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_5.imageset/pending_to_noposition_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_6.imageset/pending_to_noposition_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_6.imageset/pending_to_noposition_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_noposition_6.imageset/pending_to_noposition_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_1.imageset/pending_to_nofollow_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_1.imageset/pending_to_nofollow_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_1.imageset/pending_to_nofollow_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_2.imageset/pending_to_nofollow_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_2.imageset/pending_to_nofollow_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_2.imageset/pending_to_nofollow_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_3.imageset/pending_to_nofollow_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_3.imageset/pending_to_nofollow_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_3.imageset/pending_to_nofollow_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_4.imageset/pending_to_nofollow_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_4.imageset/pending_to_nofollow_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_4.imageset/pending_to_nofollow_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_5.imageset/pending_to_nofollow_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_5.imageset/pending_to_nofollow_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_5.imageset/pending_to_nofollow_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_6.imageset/pending_to_nofollow_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_6.imageset/pending_to_nofollow_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Change State/pending_to_notfollow_6.imageset/pending_to_nofollow_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_1.imageset/btn_white_loading_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_1.imageset/btn_white_loading_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_1.imageset/btn_white_loading_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_10.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_10.imageset/btn_white_loading_10.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_10.imageset/btn_white_loading_10@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_10.imageset/btn_white_loading_10@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_11.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_11.imageset/btn_white_loading_11.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_11.imageset/btn_white_loading_11@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_11.imageset/btn_white_loading_11@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_12.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_12.imageset/btn_white_loading_12.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_12.imageset/btn_white_loading_12@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_12.imageset/btn_white_loading_12@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_13.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_13.imageset/btn_white_loading_13.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_13.imageset/btn_white_loading_13@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_13.imageset/btn_white_loading_13@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_14.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_14.imageset/btn_white_loading_14.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_14.imageset/btn_white_loading_14@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_14.imageset/btn_white_loading_14@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_15.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_15.imageset/btn_white_loading_15.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_15.imageset/btn_white_loading_15@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_15.imageset/btn_white_loading_15@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_16.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_16.imageset/btn_white_loading_16.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_16.imageset/btn_white_loading_16@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_16.imageset/btn_white_loading_16@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_17.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_17.imageset/btn_white_loading_17.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_17.imageset/btn_white_loading_17@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_17.imageset/btn_white_loading_17@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_18.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_18.imageset/btn_white_loading_18.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_18.imageset/btn_white_loading_18@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_18.imageset/btn_white_loading_18@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_2.imageset/btn_white_loading_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_2.imageset/btn_white_loading_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_2.imageset/btn_white_loading_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_3.imageset/btn_white_loading_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_3.imageset/btn_white_loading_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_3.imageset/btn_white_loading_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_4.imageset/btn_white_loading_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_4.imageset/btn_white_loading_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_4.imageset/btn_white_loading_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_5.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_5.imageset/btn_white_loading_5.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_5.imageset/btn_white_loading_5@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_5.imageset/btn_white_loading_5@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_6.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_6.imageset/btn_white_loading_6.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_6.imageset/btn_white_loading_6@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_6.imageset/btn_white_loading_6@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_7.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_7.imageset/btn_white_loading_7.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_7.imageset/btn_white_loading_7@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_7.imageset/btn_white_loading_7@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_8.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_8.imageset/btn_white_loading_8.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_8.imageset/btn_white_loading_8@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_8.imageset/btn_white_loading_8@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_9.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_9.imageset/btn_white_loading_9.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_9.imageset/btn_white_loading_9@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/Pending Position/btn_white_loading_9.imageset/btn_white_loading_9@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction.imageset/btn_white_direction.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction.imageset/btn_white_direction@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction.imageset/btn_white_direction@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction_pressed.imageset/btn_white_direction_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction_pressed.imageset/btn_white_direction_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_direction_pressed.imageset/btn_white_direction_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1.imageset/btn_white_target_off_1.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1.imageset/btn_white_target_off_1@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1.imageset/btn_white_target_off_1@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1_pressed.imageset/btn_white_target_off_1_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1_pressed.imageset/btn_white_target_off_1_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_off_1_pressed.imageset/btn_white_target_off_1_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on.imageset/btn_white_target_on.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on.imageset/btn_white_target_on@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on.imageset/btn_white_target_on@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on_pressed.imageset/btn_white_target_on_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on_pressed.imageset/btn_white_target_on_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_target_on_pressed.imageset/btn_white_target_on_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position.imageset/btn_white_unknow_position.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position.imageset/btn_white_unknow_position@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position.imageset/btn_white_unknow_position@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position_pressed.imageset/btn_white_unknow_position_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position_pressed.imageset/btn_white_unknow_position_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Location Button/btn_white_unknow_position_pressed.imageset/btn_white_unknow_position_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api.imageset/ic_back_api.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api.imageset/ic_back_api@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api.imageset/ic_back_api@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api.imageset/ic_back_pp.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api.imageset/ic_back_pp@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api.imageset/ic_back_pp@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api_pressed.imageset/ic_back_api_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api_pressed.imageset/ic_back_api_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api_pressed.imageset/ic_back_api_pressed@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api_pressed.imageset/ic_back_pp_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api_pressed.imageset/ic_back_pp_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_back_api_pressed.imageset/ic_back_pp_press@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off.imageset/ic_bookmark_add_pp.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off.imageset/ic_bookmark_add_pp@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off.imageset/ic_bookmark_add_pp@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off.imageset/ic_bookmarks_off.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off.imageset/ic_bookmarks_off@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off.imageset/ic_bookmarks_off@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off_pressed.imageset/ic_bookmark_add_pp_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off_pressed.imageset/ic_bookmark_add_pp_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off_pressed.imageset/ic_bookmark_add_pp_press@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off_pressed.imageset/ic_bookmarks_off_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off_pressed.imageset/ic_bookmarks_off_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_off_pressed.imageset/ic_bookmarks_off_pressed@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on.imageset/ic_bookmark_remove_pp.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on.imageset/ic_bookmark_remove_pp@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on.imageset/ic_bookmark_remove_pp@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on.imageset/ic_bookmarks_on.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on.imageset/ic_bookmarks_on@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on.imageset/ic_bookmarks_on@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on_pressed.imageset/ic_bookmark_remove_pp_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on_pressed.imageset/ic_bookmark_remove_pp_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on_pressed.imageset/ic_bookmark_remove_pp_press@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on_pressed.imageset/ic_bookmarks_on_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on_pressed.imageset/ic_bookmarks_on_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_bookmarks_on_pressed.imageset/ic_bookmarks_on_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route.imageset/ic_route.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route.imageset/ic_route@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route.imageset/ic_route@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route.imageset/ic_route_to.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route.imageset/ic_route_to@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route.imageset/ic_route_to@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route_pressed.imageset/ic_route_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route_pressed.imageset/ic_route_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route_pressed.imageset/ic_route_pressed@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route_pressed.imageset/ic_route_to_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route_pressed.imageset/ic_route_to_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_route_pressed.imageset/ic_route_to_press@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share.imageset/ic_share.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share.imageset/ic_share@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share.imageset/ic_share@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share.imageset/ic_share_pp.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share.imageset/ic_share_pp@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share.imageset/ic_share_pp@3x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share_pressed.imageset/ic_share_pp_press.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share_pressed.imageset/ic_share_pp_press@2x.png
 create mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share_pressed.imageset/ic_share_pp_press@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share_pressed.imageset/ic_share_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share_pressed.imageset/ic_share_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Place Page/ic_share_pressed.imageset/ic_share_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks.imageset/btn_green_bookmarks.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks.imageset/btn_green_bookmarks@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks.imageset/btn_green_bookmarks@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks_pressed.imageset/btn_green_bookmarks_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks_pressed.imageset/btn_green_bookmarks_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_bookmarks_pressed.imageset/btn_green_bookmarks_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download.imageset/btn_green_download.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download.imageset/btn_green_download@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download.imageset/btn_green_download@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download_pressed.imageset/btn_green_download_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download_pressed.imageset/btn_green_download_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_download_pressed.imageset/btn_green_download_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_1.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_1.imageset/btn_green_menu.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_1.imageset/btn_green_menu@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_1.imageset/btn_green_menu@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_2.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_2.imageset/btn_green_menu_2.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_2.imageset/btn_green_menu_2@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_2.imageset/btn_green_menu_2@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_3.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_3.imageset/btn_green_menu_3.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_3.imageset/btn_green_menu_3@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_3.imageset/btn_green_menu_3@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_4.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_4.imageset/btn_green_menu_4.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_4.imageset/btn_green_menu_4@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_4.imageset/btn_green_menu_4@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_pressed.imageset/btn_green_menu_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_pressed.imageset/btn_green_menu_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_menu_pressed.imageset/btn_green_menu_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search.imageset/btn_green_search.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search.imageset/btn_green_search@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search.imageset/btn_green_search@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search_pressed.imageset/btn_green_search_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search_pressed.imageset/btn_green_search_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_search_pressed.imageset/btn_green_search_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings.imageset/btn_green_settings.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings.imageset/btn_green_settings@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings.imageset/btn_green_settings@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings_pressed.imageset/btn_green_settings_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings_pressed.imageset/btn_green_settings_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_settings_pressed.imageset/btn_green_settings_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share.imageset/btn_green_share.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share.imageset/btn_green_share@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share.imageset/btn_green_share@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share_pressed.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share_pressed.imageset/btn_green_share_pressed.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share_pressed.imageset/btn_green_share_pressed@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/btn_green_share_pressed.imageset/btn_green_share_pressed@3x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/img_counter.imageset/Contents.json
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/img_counter.imageset/img_counter.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/img_counter.imageset/img_counter@2x.png
 delete mode 100644 iphone/Maps/Images.xcassets/Side Menu/img_counter.imageset/img_counter@3x.png
 create mode 100644 routing/routing_integration_tests/jni/Android.mk
 create mode 120000 routing/routing_integration_tests/jni/Application.mk
 create mode 100644 routing/routing_integration_tests/jni/test.cpp
 create mode 100644 routing/routing_integration_tests/jni/test.hpp
 create mode 100644 routing/routing_integration_tests/online_cross_tests.cpp
 create mode 100644 routing/routing_integration_tests/osrm_route_test.cpp
 create mode 100644 routing/routing_integration_tests/osrm_turn_test.cpp
 create mode 100644 routing/routing_integration_tests/pedestrian_route_test.cpp
 create mode 100644 routing/routing_integration_tests/routing_integration_tests.pro
 create mode 100644 routing/routing_integration_tests/routing_test_tools.cpp
 rename {integration_tests => routing/routing_integration_tests}/routing_test_tools.hpp (100%)
 create mode 100644 routing/speed_camera.cpp
 create mode 100644 routing/speed_camera.hpp
 create mode 100644 search/suggest.hpp
 delete mode 100644 storage/country_info.cpp
 delete mode 100644 storage/country_info.hpp
 create mode 100644 storage/country_info_getter.cpp
 create mode 100644 storage/country_info_getter.hpp
 create mode 100644 storage/storage_tests/country_info_getter_test.cpp
 delete mode 100644 storage/storage_tests/country_info_test.cpp
 delete mode 100755 tools/publish_to_dropbox.py
 delete mode 100755 tools/unix/batch_generate.sh
 delete mode 100755 tools/unix/compress_section_files.sh
 delete mode 100755 tools/unix/create_mac_image_debug.sh
 delete mode 100755 tools/unix/create_mac_image_release.sh
 delete mode 100755 tools/unix/ensure_mkspecs.sh
 delete mode 100755 tools/unix/external_resources.sh
 delete mode 100755 tools/unix/make_index_of.sh
 delete mode 100755 tools/unix/make_index_of_vng.sh
 delete mode 100755 tools/unix/make_routing.sh
 delete mode 100755 tools/unix/osrm_generator.sh
 delete mode 100755 tools/unix/planet.sh
 delete mode 100755 tools/unix/routes_from_mwm_and_osrm.sh
 delete mode 100755 tools/unix/split_planet_by_polygons.sh
 create mode 100755 tools/upload_to_dropbox.sh
 create mode 100644 xcode/agg/agg.xcodeproj/project.pbxproj
 create mode 100644 xcode/lodepng/lodepng.xcodeproj/project.pbxproj
`, ``,
	},
}
