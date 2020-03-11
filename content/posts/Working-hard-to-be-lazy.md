
+++
title = "Working hard to be lazy"
draft = false
date = "2018-06-24T12:50:45.000Z"

+++

![Ursula k Le Guin](../../data/Phillips-Ursula-K-LeGuin.jpg)

So that's a picture of a great writer, UKLG. She's completely not lazy, so
apologies to her for inclusion. She's here, though, because she invented
something great: The ansible 🌌. And why is that relevant? We'll get there...

As any smart human person in charge of configuring and managing computers knows,
I'm bad at my job. Or rather, my job is nearly impossible to be good at, given
my human limitations. Why?

There are way too many things to keep track of to be good wrangling a computer,
let alone more than one of them 🖥 𝔛 ⧝!

That's why smart humans invented ✨⚡️AUTOMATION⚡️✨.

Automation means: you keep track of all the random changes your make to your
myriad .dotfiles  in a central, version controlled place. You can run a
convergent install.sh  script that puts every alias, .ignore, .gdbinit, package,
etc. you use on your computer, sets it all up, clones the repos you work on and
curls and installs the things not easily package-managed, symlinks applicable
files (.bashrc, etc.), and prints out instructions (that you keep up to date) of
manual steps you must execute. You make the effort to improve this script and
configuration by re-imaging your computer often and re-running it. You re-run it
when its not fresh to test it's convergence. This way you get to be lazy  when
you set up a new laptop or workstation; you simply git clone dotfiles && cd $_
&& ./install.sh. Nothing to remember, not much work to do.

Automation also means, and here we return to our opening image, making the
effort to perform and document server setup via code. I happen to like Ansible
for this, for a few reasons: its relatively lightweight, the manifest format is
fairly straightforward, it doesn't require an agent to be installed on the
machine being managed, and it works fairly hard to be idempotent[1]

Security, for example: There's an article kicking around about the first few
minutes on a server, which is a great reference for a simple security baseline:
lock things down, turn on updates, etc. I figured, why not do one better and
automate this, since I run these same steps on basically every server I boot up
(and should do on all of them, which automation makes easy!). So now I have a
manifest on my github that can do this again, and I don't have to remember the
exact steps. Again, I get to be lazy. I don't really have to look things up. I
just know that I did at one point, codified that knowledge in some code that
does the thing, and know where that code lives. Some  documentation might be
nice, because things break, but in general I expect that using the script over
time will shake out bugs and make it run smoother every time.

And in general this is true. Documenting an obscure process that doesn't often
run or a setup is complicated to get right doesn't have much as much value as
the ability to achieve the desired end without really know how it happened.
Sure, sometimes how it happened or how it works matters, such as when a
dependency breaks you (but you're protecting yourself from that, right?) or you
have to add something. In general though, I've been learning, you've always got
more work to do. And reducing some of the toil[2]  involved in your work frees
you up to actually think at a more synthesizing level. That's the good stuff! 🜚

Computers in general function because we limn their absurd complexity via
abstractions. Do you need to know about pipelines and interrupts and L2 caches
to write CSS? No! That'd be a terrible waste of time. Do you need to know what
to do to lock down a server, if that's your job. Yes! But make your job a little
easier by adding in automation, and you'll start to operate strategically rather
than always being caught up in the minutia.

So here's an exhortation: work hard to be lazy. Future you will be glad.


--------------------------------------------------------------------------------

 1. basically you can keep running the steps as you're trying to get the whole
    thing to work and nothing break. Good for fiddlers... ↩︎


 2. https://landing.google.com/sre/book/chapters/eliminating-toil.html ↩︎
