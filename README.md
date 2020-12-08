# haco

Simple cli s3 uploader

# Description

pronounced *ha-co*.

# Installation

1. Build with `go build`.

2. Set environment variables like this.

```sh
export AWS_ACCESS_KEY_ID=C99F5C7EE00F1EXAMPLE
export AWS_SECRET_ACCESS_KEY=a63xWEj9ZFbigxqA7wI3Nuwj3mte3RDBdEXAMPLE
export HACO_AWS_REGION=us-west-2
export HACO_AWS_BUCKET=examplebucket
```
# Usage

Upload a file to bucket root directory.

```
$ ./haco testimg.jpg
```
Upload a file as other name.

```
$ ./haco foo.txt othername.txt
```

Upload a file to specified directory.

```
$ haco foo.txt bar/baz/foo.txt
```

# TODO

* Validation
* Error handling
* Upload directory
* Upload multiple files
* Flags
