# dudir

## Introduction

This is an example program that can calculate directory sizes and print them in json format.

Then you are able to set the suid bit on it, and monitor directory sizes with unprivileged users.

Quick start:

```bash
# clone and build
git clone git@github.com:nagylzs/dudir.git
cd dudir
go build dudir.go

# set SUID and change owner to root
sudo chmod 4555 dudir
sudo chown root: dudir

# get directory size, executed by any user
./dudir /tmp /etc
# prints:
# [{"path":"/tmp","size":71558},{"path":"/etc","size":8411501}]
```

## Use constant directories instead of command line arguments

Just replace `err := doMain(os.Args[1:])` with `err := doMain([]string{"/tmp", "/etc"})` in `func main()`.

## Use it from telegraf

Example, using a telegraf module to monitor directory sizes:

```
[[inputs.exec]]
commands = [ "/usr/bin/dudir /var/lib/influxdb/data /var/lib/postgresql/pg_wal" ]
timeout = "5s"
interval = "5m"
name_override = "dudir"
name_suffix = ""
data_format = "json"
tag_keys = [ "dudir" ]
```
