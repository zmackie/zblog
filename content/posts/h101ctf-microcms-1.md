---
title: "Writeup 📝"
date: 2020-08-21T16:31:47-04:00
draft: false
---

### H101ctf Micro-cms 1 
First set of actual challenges. These were good to get warmed up.

#### Flag 0
This one relies on bypassing some sort of access control, probably a path filter.
Pages are accessed by ID, so I poked around and found one that was blocked. At first I tried messing with `Referer:` in the headers, but that didn't work. Then I tried to create a redirect to that page via a link, which fails for the same reason. But then I realized, "maybe I can edit that page?" 

💣 We're in!

#### Flag 1

What can we do if we add weird content in the markdown? Well script tags don't work...but I looked up XSS in markdown and tried a div payload:  `<div style="padding: 20px; opacity: 0;height: 20px;" onmouseout="alert('Gotcha!')"></div>`. Which pops...but nothing...huh. Turns out the flag gets added to the element in the dom when it gets rendered.

💣 We're in!

##### Flag 2

This one was similair to the above, but I discovered it by accident adding a script tag into the title. I didn't realize it had worked until I went back to the root page and saw my xss fire.

💣 We're in!

##### Flag 3

This took a bit of doing. I tried random inputs in everything but it was only after a couple hours of get requests that I realized I hadn't messed with the path in the edit page. SQLi it was `/page/edit/2'`!

💣 We're in!
