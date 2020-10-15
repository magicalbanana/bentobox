# Bento Box Challenge

This repository contains the code implementation for the requirements given by
Bento Box.

Instead of splitting the application into two applications, I instead opted to
write one CLI app that has subcommands that address the two exercises.

## Requirements

### Bare Metal

[golang](golang.org) needs to be installed in your machine to run the
application.

You can opt to run the application in a container but running the `dirls`
command will require mounting a volume to the directory that needs to be
traversed.

### Container

The CLI app can be executed inside a container. To do so, you need to build
the image first. Run the following command:

```
╰─$ docker build -t bentobox/bentobox-cli -f Dockerfile .
```

## CLI App

To view the commands in the CLI App simply run:

__bare metal__
```
╰─$ go run main.go
NAME:
   main - A new cli application

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   dirls    Recursively walks a directory and lists the contents with the file size
   strcomp  Compress a string
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

__container__
```
╰─$ docker run -it --rm bentobox/bentobox-cli
NAME:
   bentobox-cli - A new cli application

USAGE:
   bentobox-cli [global options] command [command options] [arguments...]

COMMANDS:
   dirls    Recursively walks a directory and lists the contents with the file size
   strcomp  Compress a string
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Excerise 1

To display the help message for running the first command run:

__bare metal__
```
╰─$ go run main.go dirls
NAME:
   main dirls - Recursively walks a directory and lists the contents with the file size

USAGE:
   main dirls [command options] <dir>

OPTIONS:
   --sort-by value  sort by either asc or desc (default: "asc")
   --output value   Print either as a flat directory or tree directory, vlaues are: simple | tree (default: "simple")

```

__container__
```
╰─$ docker run -it --rm --mount src="$(pwd)",target="/app",type=bind bentobox/bentobox-cli dirls
NAME:
   bentobox-cli dirls - Recursively walks a directory and lists the contents with the file size

USAGE:
   bentobox-cli dirls [command options] <dir>

OPTIONS:
   --sort-by value  sort by either asc or desc (default: "asc")
   --output value   Print either as a flat directory or tree directory, vlaues are: simple | tree (default: "simple")
```
__NOTES__:
- in order to list the directory in your host machine, we should mount a
directory to the container using the `--mount` flag. Make sure the target is
`/app` so that it doesn't conflict with any existing direcotries in the
container. You can read more about it here, [docker mounts](https://docs.docker.com/storage/bind-mounts/).

*example*:
```
╰─$ docker run -it --rm --mount src="$(pwd)",target="/app",type=bind bentobox/bentobox-cli dirls /app
/app/.git/MERGE_RR 0
/app/.git/COMMIT_EDITMSG 11
/app/.git/HEAD 23
/app/.git/refs/remotes/origin/master 41
/app/.git/ORIG_HEAD 41
/app/.git/refs/heads/master 41
/app/.git/objects/18/41e42a2aebf388beec3e9e4acd1fe9f56cef6c 54
/app/.git/description 73
/app/.git/objects/1e/c1f22e243a31fb713e33968aeff827630201a5 82
/app/.git/objects/80/0edea2ca60805e40f641a930efad57b22ea7b3 83
/app/.git/objects/0e/3adca1795837beadb1d7c9e3c90c080304f8fa 83
/app/.git/objects/40/40d77c8d0990c037cfdb252d5d38f3fa551b3c 85
/app/.git/objects/65/d5132db54bff7717895f56bf3e838000b9c3d7 87
/app/.git/objects/84/aff1055e4763cb33c52ce73a9b7bde29172973 117
/app/.git/objects/fa/41cbf5a05a6ec47c06cba694b16ac3fbc5141a 118
/app/.git/objects/71/6cf521c8d2dc057f6eb107da8fbb2424ed72e5 132
/app/.git/logs/refs/remotes/origin/master 145
/app/go.mod 154
/app/.git/objects/9d/76e07c0ef649dfea35c1c193ca972c2f5d0523 175
/app/.git/hooks/post-update.sample 189
...
```
- passing in `.` for the directory doesn't work. You'll need to pass a full in
    the path.

__Bad__:
```
╰─$ docker run -it --rm --mount src="$(pwd)",target="/app",type=bind bentobox/bentobox-cli dirls .
2020/10/15 12:39:38 lstat proc/1/fd/3: no such file or directory
```
__Good__:
```
╰─$ docker run -it --rm --mount src="$(pwd)",target="/app",type=bind bentobox/bentobox-cli dirls /app
/app/.git/MERGE_RR 0
/app/.git/COMMIT_EDITMSG 11
...
```


## Excerise 2

To display the help message for running the second command run:

__bare metal__
```
╰─$ go run main.go strcomp
NAME:
   main strcomp - Compress a string

USAGE:
   main strcomp <word to compress>
```

__container__
```
╰─$ docker run -it --rm bentobox/bentobox-cli strcomp
NAME:
   bentobox-cli strcomp - Compress a string

USAGE:
   bentobox-cli strcomp <word to compress>
```
