# 附录1-Ubuntu22.04添加新分区

## PART1. 创建新分区

### 1.1 进入`fdisk`

```
root@es-node-2:~# fdisk /dev/sda

Welcome to fdisk (util-linux 2.37.2).
Changes will remain in memory only, until you decide to write them.
Be careful before using the write command.

GPT PMBR size mismatch (41943039 != 167772159) will be corrected by write.
This disk is currently in use - repartitioning is probably a bad idea.
It's recommended to umount all file systems, and swapoff all swap
partitions on this disk.
```

### 1.2 创建新分区

```
Command (m for help): n
Partition number (4-128, default 4): 
First sector (41940992-167772126, default 41940992): 
Last sector, +/-sectors or +/-size{K,M,G,T,P} (41940992-167772126, default 167772126): 

Created a new partition 4 of type 'Linux filesystem' and of size 60 GiB.
```

### 1.3 写入更改并退出

```
Command (m for help): w
The partition table has been altered.
Syncing disks.
```

## PART2. 格式化新分区

```
root@es-node-2:~# mkfs.ext4 /dev/sda4
mke2fs 1.46.5 (30-Dec-2021)
Creating filesystem with 15728640 4k blocks and 3939840 inodes
Filesystem UUID: 3e483007-6ca0-4a94-adde-6c70779f1257
Superblock backups stored on blocks: 
	32768, 98304, 163840, 229376, 294912, 819200, 884736, 1605632, 2654208, 
	4096000, 7962624, 11239424

Allocating group tables: done                            
Writing inode tables: done                            
Creating journal (65536 blocks): done
Writing superblocks and filesystem accounting information: done   
```

## PART3. 挂载新分区

### 3.1 创建挂载点

```
root@es-node-2:~# mkdir /mnt/new_partition
```

### 3.2 挂载新分区

```
root@es-node-2:~# mount /dev/sda4 /mnt/new_partition
```

### 3.3 确认挂载结果

```
root@es-node-2:~# df -h
Filesystem                         Size  Used Avail Use% Mounted on
tmpfs                              1.6G  1.6M  1.6G   1% /run
/dev/mapper/ubuntu--vg-ubuntu--lv   18G  7.2G  9.8G  43% /
tmpfs                              7.8G     0  7.8G   0% /dev/shm
tmpfs                              5.0M     0  5.0M   0% /run/lock
/dev/sda2                          1.8G  129M  1.5G   8% /boot
tmpfs                              1.6G  4.0K  1.6G   1% /run/user/0
/dev/sda4                           59G   24K   56G   1% /mnt/new_partition
```

## PART4. 实现自动挂载

```
root@es-node-2:~# vim /etc/fstab 
root@es-node-2:~# cat /etc/fstab
```

```
# /etc/fstab: static file system information.
#
# Use 'blkid' to print the universally unique identifier for a
# device; this may be used with UUID= as a more robust way to name devices
# that works even if disks are added and removed. See fstab(5).
#
# <file system> <mount point>   <type>  <options>       <dump>  <pass>
# / was on /dev/ubuntu-vg/ubuntu-lv during curtin installation
/dev/disk/by-id/dm-uuid-LVM-R0i9ICndgK5GX5UpNxjiTKQlPfg6ByqDdSq2RW50dJZhT6mKcsNF9fqD8ea14RPs / ext4 defaults 0 1
# /boot was on /dev/sda2 during curtin installation
/dev/disk/by-uuid/1acdc704-2ad7-46c5-8e4b-84d757795de7 /boot ext4 defaults 0 1
/swap.img	none	swap	sw	0	0
/dev/sda4    /mnt/new_partition    ext4    defaults    0    2
```

其中:

`/dev/sda4    /mnt/new_partition    ext4    defaults    0    2`

这行内容是新加的