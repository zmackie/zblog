---
title: "H101ctf Microcms 2"
date: 2020-08-23T20:56:21-04:00
draft: true
---

### H101ctf Micro-CMS 2
Second set of challenges. These were a bit trickier. One thing I'm starting to get is just ho much these involve "try all the things". I bet you can be systematic about it, but I'm still just throwing shit at the wall and trying to get it to stick.

#### Flag 0
This was SQLi. It took a bit of fiddling. I was able to get the error quickly...but I a lot of time reading the stacktrace to realize the sql was just retrieving a password. Which meant that a simple union with "1" would let me supply that same password: `' union select "1"#`.

#### Flag 1
This was a "just try shit". Realizing that apps are complicated, I tried a lot of unauthenticated actions and ultimately was able to edit a page without being logged in. Found this before 0.

#### Flag 2
I had to look this up. I had a suspicion based on the message supplied when I bypassed login that I'd have to find an actual user and password. Manually enumerating this field was going to be hell; lucky somebody wrote `sqlmap` for these situations! Found the cred pair with `sqlmap -u http://35.227.24.107/b6ddb620aa/login --data "username=&password=" --dump` (which also, incidentally, turned up flag0 through a different avenue).

