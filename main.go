package main

import (
	"time"
)

/*

	- lru: add: record the oldest entry and oldest date
	- In each directory get the newest file (mtime)
	- If len(lru) < nlru, add to lru
	- else add if newer than oldest(lru)

*/

type rdir string

type entry struct {
	dir   string
	mtime time.Time
}

type lru struct {
	entries    map[rdir]time.Time
	n          int
	oldestRdir rdir
	oldestTime time.Time
}

func newLRU(n int) *lru {
	return &lru{
		entries: make(map[rdir]time.Time),
		n:       n,
	}
}

func (l *lru) add(dir rdir, mtime time.Time) bool {
	if len(l.entries) < n {
		return l.addUpdate(dir, mtime)
	}
	if mtime.After(l.oldestTime) {
		return l.addUpdate(dir, mtime)
	}
	return false
}

func (l *lru) addUpdate(dir rdir, mtime time.Time) bool {
	if _, ok := l.entries[dir]; ok {
		l.entries[dir] = mtime
		if l.oldestRdir == dir {
			l.olderstTime = mtime
		}
	} else {
		delete(l.entries, l.oldest)
		l.entries[dir] = mtime
	}
	l.setOldest()
	return true
}

func (l *lru) setOldest() {
	var (
		oldestRdir string
		oldestTime time.Time
	)
	for dir, mtime := range l.entries {
		if oldestRdir == "" || mtime.Before(oldestTime) {
			oldestRdir = dir
			oldestTime = mtime
		}
	}
	l.oldestRdir = oldestRdir
	l.oldestTime = oldestTime
}
