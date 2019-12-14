
+++
title = "Give a damn"
draft = true
date = "2018-11-13T18:43:01.000Z"

+++
In my last post, at the very end, I alluded to a Docker cli gotcha. For those of
you who didn't have a shudder of recognition, I'll tell you something it took me
a while to figure out: docker run koolKontainer:latest /bin/bash  does not
update your local copy of the image with that tag. You must docker pull
koolKontainer:latest  to make sure your image is up to date.

In other words, if you have an old  copy of that image, tagged with latest, the
former command will run that old  local copy. However if the image does not
exists at all locally, it will be retrieved from Docker Hub.

This behavior kind  of makes sense, in retrospect. I guess if you think of a
docker image like a git repo, you wouldn't necessarily assume that your local
master  would be up to date with origin/master  without pulling. I think, for
me, the disconnect comes from the tag name: latest. You sort of assume that
means, well, the most recent version of the image. And you sort of assume docker
cli would reach out to the internet to figure that out. Not so. And not
necessarily wrong as a design choice. The tag convention is really at fault, I
think, but nonetheless I was surprised. And I got bit; I actually spent a good
deal of time trying to figure out why my bug test code was behaving locally
differently than the behavior I was seeing on our CI server. And when I figured
out what was going on there was a bit of frustration.

But fair enough, I didn't know or take the time to research the behavior of
docker run. So here's one piece of advice time-saving advice I'll give you:

Learn the behavior of your tools.That being said, I was a bit miffed when I
learned what was happening. So, I took to the internet. And I found that I
wasn't the only person confused by this behavior. Turns out, there's a
longstanding issue on the project [https://github.com/moby/moby/issues/13331],
with a ton of comments from people who were also surprised by this behavior. And
you know how a good pile-on brings out the folks! I got on that issue, and added
my own little rant and felt pretty good about myself. OSS duty done!

...

After lunch I got back to my desk and looked at the open window with my comment
at the end, and I thought to myself, "You sort of seem like an asshole jumping
in there with all the complaints".

And then I realized something else:

Its always easier to complain than to give a damn. So, why not give a damn?And I
deleted my comment and rolled up my sleeves to pitch in.

Now I want to caveat this with saying that I know contributing to OSS is a
luxury and a privilege. I'm a white, male, childless software engineer, so I
rank relatively high on the advantages scale. I recognize that, for sure. I also
recognize that contributing to a very public project like Docker is super scary,
especially for those who've never contributed to open source. I've a got a post
in the works about making your first OSS contribution, so stay tuned. So, take
what I'm saying with that large rock of salt and not as a finger wagging
directed at anybody but myself (who even reads this blog anyway?)

Caveat caveated, let me just say that I turns out that making this change was
easier than I thought! I hemmed and hawed, intimidated about hacking a huge,
new, codebase. But I browsed the issues on the project and outstanding PRs, and
all the maintainers seemed nice and community minded. So I cracked open VSCode,
and poked around.

As usual, smarter people than me had done most of the work: the requirements
were pretty well hashed out in discussions, as was the actual API of the flag
changes. Also Go is generally such a simple, approachable (some would even say
boring) language that reading the codebase and figuring out where to make the
change, which is usually the hardest part of contributing to a new project, was
pretty straightforward. Go project mostly have a standard structure
[https://github.com/golang-standards/project-layout], so I knew essentially
where I had to look for the code covering the docker run command. And it turns
out that I was basically able to leverage the existing logic , with some slight
modifications and duplications, to achieve my ends. I'm sure the PR
[https://github.com/docker/cli/pull/1498]  needs work, but its gotten a bit of
attention and hopefully it'll land. And I've got the warm, fuzzy feeling that
giving back brings. So remember kids, OSS is free and complaining feels good,
but giving a damn feels better!
