# haco

Simple cli s3 uploader

# Description

pronounced *ha-co*.

# Installation

1. Download binary from [release page](https://github.com/hikitest/haco/releases),

2. Put it somewhere(like `/usr/local/bin` in Linux),

3. Set environment variables like this.

```sh
export AWS_ACCESS_KEY_ID=C99F5C7EE00F1EXAMPLE
export AWS_SECRET_ACCESS_KEY=a63xWEj9ZFbigxqA7wI3Nuwj3mte3RDBdEXAMPLE
export HACO_AWS_REGION=ap-northeast-1
export HACO_AWS_BUCKET=examplebucket
```
# Usage

Upload a file to bucket root directory.

```
$ haco helloworld.java
$ haco ~/Pictures/gopher.jpg
```
Upload a file as other name.

```
$ haco foo.txt othername.txt
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
