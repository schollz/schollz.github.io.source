---
title: "How to post a site to IPFS easily"
date: 2018-06-03T14:13:59-07:00
tags: [ipfs,coding]
slug: how-to-post-a-site-to-ipfs
written: ["2018-06-03","2018-06","2018"]
---

There's a bit of a chicken-and-egg problem with IPFS to get pages to link to each other that [is frusturating to IPFS users](https://macwright.org/2017/08/09/decentralize-ipfs.html). To have the IPFS hash for the link you first need to get the IPFS hash of the other link. One way around this is to make all the URLs relative. This is a lot easier to do than you think.

First make a blog however you want and then host it locally, e.g. at `localhost:1313` (for Hugo people).

Then open another terminal and do the following `wget` command:

```
$ wget -k --recursive --no-clobber --page-requisites --adjust-extension --span-hosts --convert-links --restrict-file-names=windows --domains localhost --no-parent http://localhost:1313/
```

This will generate a offline version of your website that has everything with relative links. Now create an IPFS version using the downloaded version with

```
$ ipfs add -r localhost+1313 -Q
<your hash>
```

Now you can view your site on IPFS at `https://ipfs.io/ipfs/<your hash>`.

