package main

import (
	"fmt"

	"github.com/StackExchange/wmi"
)

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

func getStorageInfo() {
	var storageinfo []storageInfo
	var loaclStorages []Storage
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
	if err != nil {
		return
	}

	for _, storage := range storageinfo {
		info := Storage{
			Name:       storage.Name,
			FileSystem: storage.FileSystem,
			Total:      storage.Size / 1024 / 1024 / 1024,  // 字节转为G
			Free:       storage.FreeSpace / 1024 / 1024 / 1024,
		}
		loaclStorages = append(loaclStorages, info)
	}
	fmt.Println("localStorage=", loaclStorages)
	// localStorage= [{C: NTFS 99 14} {D: NTFS 132 85} {E: NTFS 1277 332} {F: NTFS 585 42} {H:  0 0}]
	// localStorage= [{C: NTFS 共99G 剩余可用14G}
}

func main() {
	getStorageInfo()
}
