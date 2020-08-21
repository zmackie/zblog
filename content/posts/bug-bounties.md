---
title: "Bug Bounties 🗺"
date: 2020-08-21T10:24:27-04:00
draft: False
---

I've been dancing around web hacking for a while. I think my deep-seated hatred of web programming (mainly css) makes this fun 😈. The web is just a tangled complicated mess. It can be satisfy to rip things apart a bit...anyway this is a post about my process for getting into doing bug bounties.

#### Watch this first 🔑:
This video is a bit long, but its so compendious and worth your time. Why? She talks about learning theory, vets resources, goes against the common (bad) advice, and debunks lots of myths. And then she gives you bunch of sane paths based on where people are at (not everyone is starting from zero):
https://youtu.be/hDYqWZ11njU


#### What to focus on?

I should note that I'm not starting from zero. I've been programming for a while and gone through some of this material before. But I've learned a lot about, well, *learning* in the past couple years and I'm going to try and apply a ["Whole Game"](https://www.gse.harvard.edu/news/uk/09/01/education-bat-seven-principles-educators) approach: quickly get the lay of the land, do some practical challenges first, figure out where the high value is, figure out what to focus on (harder things, $$$ bugs).

_I've been doing my usual learning resource recon. Here are some huge ones:_

 This [guide to a great interview](https://twitter.com/MarcoFigueroa/status/1293264102421192704?s=20) with Jason Haddix has so many nuggets.
> 1st month: "You should not expect to find anything great for your first three months." That crap will get you to private invites. You'll get used to it, used to bounty. 3 months in: standardize and undertand recon automation. Find fresh targets, get private invites. My bread and butter - recon finds old shit. "Most of my great bugs have been pretty generic. SQLi, auth bypass, idor. Then get into something newer (like mobile). Get at mobile web services in a file. You're gonna be ahead with that because nobobody wants to do that."

> Nugget:
> 90% of my great vulns are from content discover. Recon, finding shit other people don't, and then finding endpoints in Javascript. Learn to parse JS!!
- Jhaddix [Syllabus](https://twitter.com/Jhaddix/status/1292969859937005568?s=20)

* [This twitter AMA with Agarri is great](https://www.agarri.fr/blog/archives/2020/06/19/a_recap_of_the_q_ampa_session_on_twitter/index.html):

> For bug bounty, I don't look for XSS and CSRF bugs. Mostly injections, RCE, XXE, SSRF and business-related vulnerabilities.

>« Any tips for testing for RCE? In my experience a lot of RCE reports are not disclosed, and payloads aren't that largely available, making it hard to get a feeling for what might be vulnerable. 

> » RCE is an impact, not a vulnerability class. Vulnerabilities you're looking for: command injections (of course), SQLi (f.e. with xp_cmdshell), leaked or weak credentials (CMS admin can change Jinja templates), file creation (overwrite a script), SSRF to an internal unprotected admin interface, and of course, everything with a CVE (ImageTragick, ShellShock, ...). And I disagree which the point that RCE reports and/or payloads aren't commonly available. Relevant Google dork: "site:hackerone.com/reports/ rce".

### The raw materials

To be executed concurrently:    `go`
- Reading materials: 
    * "Real world bug hunting". 
    * ("Web app hacker's handbook" is usually recommended but)
        * Its too long (play the game!)
        * WebSecAcademy is basically a third edition of that book (with labs! and videos!)
- Practicals: Juice shop, Pentesterlabs, WebSecAcademy
- Videos: https://www.youtube.com/playlist?list=PLbyncTkpno5FAC0DJYuJrEqHSMdudEffw
- Course: https://web.stanford.edu/class/cs253/ (For supplemental material)
- Tool: Burp 😜

### Can haz order?
    
There's a huge bunch of vulnerabilities everywhere and it can be overwhelming to figure out which to learn about first. Tool is burp. Always burp. And in order of vulnerability, start with the easy/classic ones and progress from there: 
```shell
injections/SQLis/xss
lfi, rfi (local/remote file inclusion)
RCE
XXE
SSRF
IDOR
business-related vulnerabilities (logic)

```

### First up:
* https://ctf.hacker101.com/ctf
* https://portswigger.net/web-security/sql-injection

