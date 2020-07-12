package nfs

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
)

type NFS struct {
	Host       string
	Dir        string
	MountDir   string
	DeployFlag string
}

var Nfs = NFS{}

func init()  {
	Nfs.MountDir = "/Users/leemike/testnfs/test"
	Nfs.Host = "192.168.9.82"
	Nfs.Dir = "/home/centos/testnfs/test"
	Nfs.mountNfsLocal()
}


func (nfs *NFS) mountNfsLocal() error {
	shellFile, err := os.OpenFile("mount.sh", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("MountNfsLocal: open mount.sh failed", err)
	}
	defer shellFile.Close()
	var cmd string
	switch runtime.GOOS {
	case "darwin":
		cmd = fmt.Sprintf("echo Password | sudo mount -o nolock -o resvport %s:%s  %s", nfs.Host, nfs.Dir, nfs.MountDir)
	case "linux":
		cmd = fmt.Sprintf("echo Password | sudo mount -o nolock -t nfs %s:%s  %s", nfs.Host, nfs.Dir, nfs.MountDir)
	}
	_, err = shellFile.Write([]byte(cmd))
	if err != nil {
		return fmt.Errorf("MountNfsLocal: Write mount.sh failed", err)
	}
	command := exec.Command("/bin/bash", "mount.sh")
	_, err = command.Output()
	if err != nil {
		return fmt.Errorf("MountNfsLocal: Execute Shell failed", err)
	}
	return nil
}

func (nfs *NFS) CreateDir(dirName string) error {
	serviceDir := fmt.Sprintf("%s/%s", nfs.MountDir, dirName)
	err := os.MkdirAll(serviceDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("CreateServiceDir: create service dir failed", err)
	}
	return nil
}

//在nfs服务器创建对应的文件夹
func (nfs *NFS) CreateDirInNFS() error {

	//在nfs服务器创建对应的文件夹
		if err := nfs.CreateDir("aaa"); err != nil {
			return err
		}

	return nil
}