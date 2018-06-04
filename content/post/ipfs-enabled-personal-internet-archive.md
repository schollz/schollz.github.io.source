---
title: "IPFS enabled personal internet archive"
date: 2018-06-04T07:16:21-07:00
tags: [ipfs]
slug: ipfs-enabled-personal-internet-archive
written: ["2018-06-04","2018-06","2018"]
---

There are some great ways to archive sites on the web. You can use [archive.is](http://archive.is/) or the [Internet Archive](http://archive.org/web/). But, using the IPFS, its really easy to make your own hosted internet archive.

Say you found a cool website, like `www.bzarg.com/p/how-a-kalman-filter-works-in-pictures/` which shows how a Kalman filter works. All you need to do is recursively download the site, with cross-domain assets, and then pin it to IPFS. 

```
$ mkdir site1 && cd site1
$ wget --restrict-file-names=windows -k --adjust-extension --span-hosts --convert-links --backup-converted --page-requisites http://www.bzarg.com/p/how-a-kalman-filter-works-in-pictures/ 
$ cd ../ && ipfs add -r site1 -Q
<some hash>
$ ipfs pin add -r <some hash>
```

That's it! IPFS will automatically resuse assets if they have the same hash, so common files (CSS, JS) will only ever be added once. You can then view your site with the IPFS gateway. For example, the site above is available at https://ipfs.io/ipfs/QmdLSyPnQLiVk1wy91tZvhVfqbFbD7zSCuRtC51Z9yeY8E/www.bzarg.com/p/how-a-kalman-filter-works-in-pictures/.
