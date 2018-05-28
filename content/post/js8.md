---
title: "Downloading sounds from the JS-8"
date: 2015-04-10T13:54:40-06:00
tags: [coding]
slug: js8
written: ["2015","2015-04","2015-04-10"]
---

Boss makes a really cool Audio player with guitar effects, the JS-8. It turns out that some of the cool sounds that is has can be downloaded directly via a little hacking.
​
![The JS8 audio player](/img/js8.jpg)
​
You'll need to create a fake JS-8 device. To do this you basically download the backup data and then format a thumbdrive so that it *looks* like a JS-8 device.


1. First [download eBand JS-8 factory installed data in the included SD Card](http://roland.com/support/article/?q=downloads&p=JS-8&id=62153672).

2. Insert a thumbdrive, format it and name it "JS-8".

3. Unzip the "eBand JS-8 factory installed data in the included SD Card" and drag the "ROLAND" folder onto the newly formatted thumbdrive.

4. [Download the eBand Song List Editor Ver.1.01 for Windows](http://www.boss.info/support/by_product/eband_js-10/updates_drivers/5410)

5. Unzip, then install "eBand Song List Editor"

6. Open "eBand Song List Editor", it will automatically detect your faked JS-8 device and allow you to export all the songs from it.