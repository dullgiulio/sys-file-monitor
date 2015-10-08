package main

import (
	"time"
)

//type rdir string

type entry struct {
	dir   string
	mtime time.Time
}

func (e *entry) equal(o *entry) bool {
	return e.dir == o.dir
}

func (e *entry) after(o *entry) bool {
	return e.mtime.After(o.mtime)
}

type lru struct {
	entries []*entry
	n       int
	oldest  *entry
}

func newLRU(n int) *lru {
	return &lru{
		entries: make([]*entry, n),
	}
}

func (l *lru) add(e *entry) bool {
	if l.n < cap(l.entries) {
		return l.addUpdate(e)
	}
	if e.after(l.oldest) {
		return l.addUpdate(e)
	}
	return false
}

func (l *lru) addUpdate(e *entry) bool {
	if n, ok := l.find(e); ok {
		// Element already in list, update mtime
		l.entries[n].mtime = e.mtime
	} else {
		if l.n < cap(l.entries) {
			// Not filled, append
			l.entries[l.n] = e
			l.n++
		} else {
			// Substitute this entry to the oldest
			for i := range l.entries {
				if l.entries[i] == l.oldest {
					l.entries[i] = e
				}
			}
		}
	}
	l.updateOldest()
	return true
}

func (l *lru) find(e *entry) (int, bool) {
	for i, el := range l.entries {
		if el.equal(e) {
			return i, true
		}
	}
	return -1, false
}

func (l *lru) updateOldest() {
	var oldest *entry
	for _, e := range l.entries {
		if oldest == nil || oldest.after(e) {
			oldest = e
		}
	}
	l.oldest = oldest
}
