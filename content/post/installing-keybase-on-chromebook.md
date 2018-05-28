---
title: "Installing Keybase on Chromebook"
date: 2018-05-28T10:28:00-07:00
tags: [coding]
slug: installing-keybase-on-chromebook
written: ["2018-05-28","2018-05","2018"]
---

Keybase is a great system for mounting a end-to-end encrypted filesystem. It also has an encrypted git feature which is great for storing secure information in repositories. While keybase has releases for most systems, that don't have any releases for `armh`, of which Chromebooks fall. 

With [help from the keybase developers](https://github.com/keybase/client/issues/8815#issuecomment-334397388%29) I have determined a way to get keybase up and running on a Chromebook in no time.

First install the pre-requisites:

```
$ sudo apt-get install fuse
```

Then install keybase:

```
$ go get -u -v github.com/keybase/client/go/keybase
$ go install -v -tags production github.com/keybase/client/go/keybase
```

Login with your keybase name (or add the computer if it isn't already added.

```
$ keybase login
```

Now install the keybase filesystem, *kbfsfuse*:

```
$ go get -v github.com/keybase/kbfs/...
$ cd $GOPATH/src/github.com/keybase/kbfs/kbfsfuse
$ go install
$ sudo mkdir -p /keybase
$ sudo chown $USER /keybase
$ KEYBASE_RUN_MODE=prod kbfsfuse /keybase & 
```

And finally, install the *git-remote-keybase* for using git with keybase.

```
$ cd $GOPATH/src/github.com/keybase/kbfs/kbfsgit/git-remote-keybase 
$ go install
```
