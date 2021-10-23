package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	src := "./input"
	dest := "./output"

	err := entry(src, dest, "")
	if err != nil {
		os.Exit(1)
	}
}

func entry(src, dest, path string) error {
	sp := filepath.Join(src, path)
	fs, err := os.Stat(sp)
	if err != nil {
		return err
	}

	if fs.IsDir() { // ディレクトリの場合
		if err := addDir(src, dest, path); err != nil {
			return err
		}
	} else { // ファイルの場合
		if err := addFile(src, dest, path); err != nil {
			return err
		}
	}

	return nil
}

func addDir(src, dest, path string) error {
	dp := filepath.Join(dest, path)
	if err := os.Mkdir(dp, 0700); err != nil {
		return err
	}

	sp := filepath.Join(src, path)
	fi, err := os.ReadDir(sp)
	if err != nil {
		return err
	}
	for _, f := range fi {
		err := entry(sp, dp, f.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func addFile(src, dest, path string) error {
	d, err := ioutil.ReadFile(filepath.Join(src, path))
	if err != nil {
		return err
	}

	df, err := os.Create(filepath.Join(dest, path))
	if err != nil {
		return err
	}
	defer df.Close()

	if _, err = df.Write(d); err != nil {
		return err
	}

	return nil
}
