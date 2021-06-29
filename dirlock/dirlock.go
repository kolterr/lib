// +build !windows

package dirlock

import (
	"fmt"
	"os"
	"syscall"
)

// locker for dir
type dirLock struct {
	dir string
	f   *os.File
}

// return a instance of dirLock
func New(dir string) (*dirLock, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	l := &dirLock{
		dir: dir,
		f:   f,
	}
	return l, nil
}

// lock the dir
func (l *dirLock) Lock() error {
	err := syscall.Flock(int(l.f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s (possibly in use by another process)", l.dir, err)
	}
	return nil
}

// unlock the dir
func (l *dirLock) UnLock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
