package status

type gitOutputTest struct {
	input  string
	output string
}

var testGitStatusFixtures = []gitOutputTest{
	gitOutputTest{`On branch BE-242
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   types/errors_test.go

no changes added to commit (use "git add" and/or "git commit -a")
`, `[BE-242] == []
  changes not staged for commit:
    * types/errors_test.go
`},
	gitOutputTest{`On branch master
Initial commit
nothing to commit (create/copy files and use "git add" to track)
`, `[master] == [<initial commit>]
`},
	gitOutputTest{`On branch master
Initial commit
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	a.txt
nothing added to commit but untracked files present (use "git add" to track)
`, `[master] == [<initial commit>]
  untracked files:
    ? a.txt
`},
	gitOutputTest{`On branch master
Initial commit
Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
	new file:   a.txt
`, `[master] == [<initial commit>]
  changes to be committed:
    + a.txt
`},
	gitOutputTest{`On branch master
nothing to commit, working directory clean
`, `[master] == []
`},
	gitOutputTest{`On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	b.txt
nothing added to commit but untracked files present (use "git add" to track)
`, `[master] == []
  untracked files:
    ? b.txt
`},
	gitOutputTest{`On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)
	modified:   a.txt
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	b.txt
`, `[master] == []
  changes to be committed:
    * a.txt
  untracked files:
    ? b.txt
`},
	gitOutputTest{`On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)
	modified:   a.txt
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	b.txt
no changes added to commit (use "git add" and/or "git commit -a")
`, `[master] == []
  changes not staged for commit:
    * a.txt
  untracked files:
    ? b.txt
`},
	gitOutputTest{`On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)
	modified:   a.txt
	new file:   b.txt
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	c.txt
`, `[master] == []
  changes to be committed:
    + b.txt
    * a.txt
  untracked files:
    ? c.txt
`},
	gitOutputTest{`On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)
	modified:   a.txt
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)
	modified:   b.txt
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	c.txt
`, `[master] == []
  changes to be committed:
    * a.txt
  changes not staged for commit:
    * b.txt
  untracked files:
    ? c.txt
`},
	gitOutputTest{`On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)
	modified:   a.txt
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)
	modified:   b.txt
	deleted:    c.txt
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	d.txt
`, `[master] == []
  changes to be committed:
    * a.txt
  changes not staged for commit:
    * b.txt
    - c.txt
  untracked files:
    ? d.txt
`},
	gitOutputTest{`On branch master
Your branch and 'origin/master' have diverged,
and have 1 and 2 different commits each, respectively.
  (use "git pull" to merge the remote branch into yours)
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	fa.txt
nothing added to commit but untracked files present (use "git add" to track)
`, `[master+1] ~~ [origin/master+2]
  untracked files:
    ? fa.txt
`},
}
