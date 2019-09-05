// Package cachestore provide a storage backend based on the memory, only cache the bytes.
package cachestore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kyl2016/tusd"
	"github.com/kyl2016/tusd/uid"
	"gopkg.in/Acconut/lockfile.v1"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var defaultFilePerm = os.FileMode(0664)

type CacheFile struct {
	Len  int
	Data *[]byte
}

type CacheStore struct {
	Path  string
	files map[string]*CacheFile // id and file info (len and data)
}

func New(path string) *CacheStore {
	return &CacheStore{path, map[string]*CacheFile{}}
}

func (store *CacheStore) UseIn(composer *tusd.StoreComposer) {
	composer.UseCore(store)
	composer.UseGetReader(store)
	composer.UseTerminater(store)
	composer.UseLocker(store)
	composer.UseConcater(store)
}

func (store *CacheStore) NewUpload(info tusd.FileInfo) (id string, err error) {
	id = uid.Uid()
	info.ID = id

	store.files[id] = &CacheFile{Len: 0, Data: &[]byte{}}

	err = store.writeInfo(id, info)
	return
}

func (store *CacheStore) WriteChunk(id string, offset int64, src io.Reader) (int64, error) {
	fileInfo, ok := store.files[id]
	if !ok {
		return -1, fmt.Errorf("Not found file id %s", id)
	}

	buffer, err := ioutil.ReadAll(src)
	if err != nil {
		return -1, err
	}

	*fileInfo.Data = append(*fileInfo.Data, buffer...)
	fileInfo.Len += len(buffer)

	return int64(len(buffer)), nil
}

func (store *CacheStore) GetInfo(id string) (tusd.FileInfo, error) {
	info := tusd.FileInfo{}

	_, ok := store.files[id]
	if !ok {
		return info, fmt.Errorf("Not found file id %s", id)
	}

	data, err := ioutil.ReadFile(store.infoPath(id))
	if err != nil {
		return info, err
	}
	if err := json.Unmarshal(data, &info); err != nil {
		return info, err
	}

	fileInfo, ok := store.files[id]
	if !ok {
		return info, fmt.Errorf("Not found file id %s", id)
	}

	info.Offset = int64(fileInfo.Len)

	return info, nil
}

func (store *CacheStore) GetReader(id string) (io.Reader, error) {
	file, ok := store.files[id]
	if !ok {
		return nil, fmt.Errorf("Not found file id %s", id)
	}

	return bytes.NewReader(*file.Data), nil
}

func (store *CacheStore) Terminate(id string) error {
	if err := os.Remove(store.infoPath(id)); err != nil {
		return err
	}
	delete(store.files, id)
	return nil
}

func (store *CacheStore) ConcatUploads(dest string, uploads []string) (err error) {
	file, ok := store.files[dest]
	if !ok {
		return fmt.Errorf("Not found file id %s", dest)
	}

	for _, id := range uploads {
		src, err := store.GetReader(id)
		if err != nil {
			return err
		}
		buffer, err := ioutil.ReadAll(src)
		if err != nil {
			return err
		}
		*file.Data = append(*file.Data, buffer...)
	}
	return
}

func (store *CacheStore) LockUpload(id string) error {
	lock, err := store.newLock(id)
	if err != nil {
		return err
	}
	err = lock.TryLock()
	if err == lockfile.ErrBusy {
		return tusd.ErrFileLocked
	}
	return err
}

func (store *CacheStore) UnlockUpload(id string) error {
	lock, err := store.newLock(id)
	if err != nil {
		return err
	}
	err = lock.Unlock()
	if os.IsNotExist(err) {
		err = nil
	}
	return err
}

func (store *CacheStore) newLock(id string) (lockfile.Lockfile, error) {
	path, err := filepath.Abs(filepath.Join(store.Path, id+".lock"))
	if err != nil {
		return lockfile.Lockfile(""), err
	}
	return lockfile.Lockfile(path), nil
}

func (store *CacheStore) infoPath(id string) string {
	return filepath.Join(store.Path, id+".info")
}

func (store *CacheStore) writeInfo(id string, info tusd.FileInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(store.infoPath(id), data, defaultFilePerm)
}
