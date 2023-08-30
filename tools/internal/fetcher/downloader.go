/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(path string, url string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Printf("error creating %s", dir)
		log.Fatal(err)
	}
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err = out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	resp, err := http.Get(rawUrl(url))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func rawUrl(url string) string {
	return gitlabRawUrl(githubRawUrl(url))
}
