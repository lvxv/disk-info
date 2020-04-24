
#### DEP
```
yum install -y golang
```

#### BUILD
`make`

#### RUN
```
./build/disk-info -f [mount point of disk]
```


```
output:
Type |Disk |Total   |Used    |Free   |Root Reserve |Lamb Reserve
df   |/tmp |234 GiB |202 GiB |16 GiB |15 GiB       |0 B
lamb |/tmp |218 GiB |204 GiB |15 GiB |15 GiB       |1.2 GiB

```
