package worker

import (
	"../helper"
	"fmt"
	"net/url"
	"path"
)

type SnortConfig struct {
	SourceUrl      string
	DestinationUrl string
	Hash           string
	DestPrepared   bool
}

func (cfg *SnortConfig) SourcePath() string {
	src_url, err := url.Parse(cfg.SourceUrl)
	if err != nil {
		return cfg.SourceUrl
	}
	return src_url.Path
}

func (cfg *SnortConfig) DestinationPath() string {
	dst_url, err := url.Parse(cfg.DestinationUrl)

	if err != nil {
		panic(err)
	}

	base := path.Base(cfg.SourcePath())
	dest := path.Join(dst_url.Path, base)

	if cfg.DestPrepared != true {
		_, err = helper.PrepareDestPath(dest)

		if err != nil {
			fmt.Println("destination path error: " + err.Error())
			panic(err)
		}
		cfg.DestPrepared = true
	}
	return dest
}

func (cfg *SnortConfig) HasMd5() bool {
	if cfg.Hash == "MD5" {
		return true
	} else {
		return false
	}
}

func (cfg *SnortConfig) HasCrc() bool {
	if cfg.Hash == "CRC32" {
		return true
	} else {
		return false
	}
}
