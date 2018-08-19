---
title: "Project graveyard"
date: 2018-08-19T20:57:15-07:00
tags: [coding]
slug: graveyard
written: ["2018-08-19","2018-08","2018"]
---

I like to write code and always have - my [earliest project is a band name generator](https://github.com/schollz/BandGenerator) from 2003. On Github  I have [over 230 repos](https://github.com/schollz?utf8=%E2%9C%93&tab=repositories&q=&type=source&language=), but I also have some on [Bitbucket](https://bitbucket.org/schollz/), [Gitlab](https://gitlab.com/schollz), and my own [self-hosted fossil](https://fossil.schollz.com/). 

Like [others](https://dev.to/t/graveyard), I've had projects that I stopped working on and stopped using altogether - thus relegating them to the project graveyard. Here is a sliver of projects that ended up not meeting my expectations and have been buried in their grave of 1's and 0's.



## 🌅 [sundial](https://github.com/schollz/sundial)

*Get the next sunset and sunrise time based on latitude and longitude.*

I write most programs in Python and Go but I specifically wrote this one in C because I was planning on using it on an Arduino. I wanted an Arduino to know the time of sunrise and sunsets so it could be used to open/close a chicken coop door to let the chickens out at the right times.

The code here works great, however I when I put it on an Arduino it was always off by hours. I learned the error occurred because the [Uno boards use a double with 32-bit precision, not 64-bit](https://www.arduino.cc/reference/en/language/variables/data-types/double/) so there wasn't enough precision to actually make the mathematical calculations in the code. A hardware issue which I didn't see anyway around - oh well, to the grave you go.

**What I learned**: 

- Equations for calculating sunrise/sunset times.
- Someone, somewhere, may find your code useful. I got a [huge PR](https://github.com/schollz/sundial/pull/1) out of the blue from someone who found this code immensely useful even though it had no use to me anymore.





## 🥚 [Food identicons](https://github.com/schollz/food-identicon)

*A identicon made of foods.*

I thought since people recognize food really well and are used to seeing recipes, that you could make an identicon that could uniquely identify some bytes with 9 pictures of food arranged in a square.

The program basically works, but it turned out to be hard to get nice pictures of food that are easily discernible so the identicons looked strange and indecipherable.

**What I learned**:

- How hard it is to find and collect high quality images.
- How to work with images in Go.






## 🎶 [musicsaur](https://github.com/schollz/musicsaur)

*Music synchronization across browsers.*

This is the one that saddens me the most. I was living in a two story house and I wanted to easily play music on both stories with two computers, so I wrote this program to synchronize music *playing through your browser*! It could could also play it headless on Raspberry Pis which was great fun.

I was never really satisfied that it worked *well enough*. Though it worked, every once and awhile you'd get a song that had >30 ms latency between the start times and then it would sound wonky. I used Javascript to control the skipping, but that is always problematic because Javascript is slow and can take ~0-10 ms to accurately skip to the position and it was near impossible to figure out exactly how much time is needed to account for the skip.

If you want to try, it still works, but you just need to use commit `b775910` from the `tcolgate/mp3` library since they made a breaking change since I've played with it.

**What I learned**:

- How hard it is to work with streaming audio.
- How to use HTML5 audio in Javascript
- Time synchronization routines






## 🎲 [No dice](https://github.com/schollz/no-dice)

*Generating random numbers from squiggles.*

I had an idea of generating random numbers just by drawing quick squiggles and counting the intersections and then doing some modulus. Of course its slower than dice, but its better than nothing if you have no dice. I did some coding to figure out how "random" this can be, but realized that since humans are doing the drawing I'd need to sample from actual humans instead of simulating human drawing with this code.

**What I learned**:

- How hard it is to simulate human actions.
- How to do bezier curves






## 📘 [bol](https://github.com/schollz/bol)

*A distributed, synchronized journal.*

All my life I've wanted to create the perfect distributed, synchronized journal. Nothing has ever satisfied. This is one attempt that I abandoned. It had a nice idea and worked across platforms, but I just didn't like using that much because of the UI and didn't have the will to make improve UX. I stopped using it so it became abandoned.

Since [bol](https://github.com/schollz/bol) I've worked on a totally CLI version, [gojot](https://github.com/schollz/gojot). However I've abandoned and [rewritten gojot 5 times](https://github.com/schollz/gojot) and I still don't use it now. Lately I've been pretty happy with a web version of another program I wrote, [cowyo](https://github.com/schollz/cowyo) but I've actually found myself recently rewriting this one into [rwtxt](https://rwtxt.com). 

**What I learned:**

- How hard UI/UX design is and how hard I am to please
- How hard it is to make distributed programs






## 👥 [kiki](https://github.com/schollz/kiki)

*A distributed, offline-friendly social network.*

There are all sorts of reasons to have alternatives to Facebook, so I decided to try to make one myself. I took a lot of inspiration from [scuttlebutt](https://www.scuttlebutt.nz/) but wanted to make something that was so easy to use that my mom could use it and wanted something that could be made private if I wanted to.

I had a lot of fun writing the code, but in the end I realized I wouldn't use it. To use it would encourage other people to use it and then I'd be stuck with all the problems facing other social networks - how to regulate bad behavior while encouraging openness, etc. I had no desire to write code to regulate content but that what you have to do nowadays if you have a social network.

**What I learned:**

- How hard social networks are to maintain
- Creating distributed message passing system







## 🎹 [pypiano](https://github.com/schollz/pyplayerpiano)

*Let your computer play a piano duet with you.*

Here is one happy story. I wanted to create an AI to play music with me at the piano. The obvious way to do this at the time was to use Python - they have a great gaming ([pygame](https://www.pygame.org/wiki/GettingStarted)) library which handles realtime MIDI out-of-the-box. It turned out, though, that development in Windows was super problematic and I found pygame would work with some functions but not with others (see my entire complaints here: https://rpiai.com/piano). So I gave up and abandoned it.

But, this is a happy story because the project came back from the grave. I had an inclination it'd be easier to write this in Go because it requires so many little async processes, easily done with goroutines. After a few more non-programming obstacles I was happily using the code: [PianoAI](https://github.com/schollz/PIanoAI). Maybe its my favorite project I've ever made.

**What I learned:**

- How hard it is to do realtime applications, especially with MIDI
- How MIDI works
- Heuristically learning synthetic piano lines

---

I could go on and on and on. At least I can say that I've genuinely learned a lot from writing *and abandoning* code - trying (and maybe failing) is my favorite way to learn. I will continue to abandon projects, I think because I just like the journey more than the destination, even if that destination is the graveyard.
