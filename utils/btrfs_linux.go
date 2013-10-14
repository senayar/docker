package utils

/*
#include <sys/ioctl.h>
#include <linux/fs.h>
#include <errno.h>

// See linux.git/fs/btrfs/ioctl.h
#define BTRFS_IOCTL_MAGIC 0x94
#define BTRFS_IOC_CLONE _IOW(BTRFS_IOCTL_MAGIC, 9, int)

int	btrfs_reflink(int fd_out, int fd_in)
{
  int	res;

  res = ioctl(fd_out, BTRFS_IOC_CLONE, fd_in);
  if (res < 0)
    return errno;
  return 0;
}

*/
import "C"

import (
	"syscall"
)

func BtrfsReflink(fd_out, fd_in uintptr) error {
	if res := C.btrfs_reflink(C.int(fd_out), C.int(fd_in)); res != 0 {
		return syscall.Errno(res)
	}
	return nil
}