
+++
title: "The Joys of debugging"
draft: true
date: "2018-11-01T21:31:21.000Z"

+++
Today at work I was trying to merging a PR, and I noticed something very strange
 during the rebase. I work on a team that, among other things, manages
dependencies for applications [https://github.com/buildpack].  We define the
versions of a runtime available to an application through a manifest, which is
in YAML form. And part of the manifest includes SHAs of the runtimes, which we
use to verify their integrity.

A select few of SHAs now had single quotes around them. Wat?This wasn't
necessarily a problem, but it was a confusing change, at the very least.  Spooky
👻. But was it a bug? And why was it happening? I had to find out.

I love  a good debug.  For the curious and patient, its a wonderfully rewarding
exercise in problem solving. You're guaranteed to learn things about your
codebase and services that you run. And if you're deliberate in your approach,
you can learn things about debugging! Which makes you a better debugger!
Feedback loops!

Okay so I always start with questions. What's the behavior I'm seeing? Do we
have time to understand this or is this an emergency that we can triage later?
And most importantly,

What's changed recently?So we checked the commit of the YAML change. Automation.
Where are the code paths that actually write this YAML, that are invoked by our
automation? Hmm...nothing changed there recently. Some 2 year old Ruby code.

Okay let's try and correlate this YAML change with something else. This
technique roughly looks like finding when the new behavior started, and
backtracking to any pertinent changes that happened just before the behavior
started exhibiting. Its a bit like being a detective (at least the masterpiece
theater type 🧐) -  you're trying to develop a timeline of the event in
question, building up a picture in your mind. Don't jump on causes just yet.
Steep yourself in information and let diffuse thinking be your guide.

You'll notice that I said pertinent  changes. In the ideal case, every aspect of
the environment in question is version controlled in some way, even at the
system level. Luckily, our CI tasks run in docker containers. Unfortunately, we
use the latest tag (some shame on us), so its a bit harder to track what exactly
was in the container that ran the code in question, but in general the point
stands. And we're much better off than if we have some IT provisioned snowflake

Git, of course, is a great source for this information...its actually basically
the only source of information that's reliable. Even the release notes of your
dependencies could be wrong, but git don't lie (at least I hope not). Git it is
basically the entire magic sauce to doing debugging at the unit and integration
level (IE, within one codebase and then between system boundaries). If there was
a git-like thing for reality, detectives would have a much easier time. I don't
even really know what that means; hopefully I didn't just describe the
blockchain ▇ ⛓.

Anyway, so we looked through our git log and found something that happened
around that time: the CI image was updated from ruby:2.3.1-slim -> ruby:2.3-slim
. Okay so that was probably it. So, next step:

Develop a hypothesis and test itThe change had something to do with YAML and
Ruby, so lets look into that. Hypothesis: some version of either libyaml  or the
Ruby YAML library (basically a thin wrapper around libyaml) got bumped with that
CI change and its now dumping YAML differently. Not exactly a brilliant insight,
but its something to go on.

Create an isolated example of the bug to test your hypothesisIn my mind, you
don't have a bug if you can't reproduce it; you have a Heisenbug or some
distributed systems Gremlin. Or you're drunk. Anway, try and catch that Gremlin
and put it in a bottle 👹. We wrote a little Ruby program that should reproduce
the behavior:
```ruby
require 'yaml'

put YAML::VERSION

testCase = {sha:
'0911c3aeb2c25dd0a41f0225e0c0f2baaa404ffb9cd772166133572f5fb91112'}

puts YAML.dump(testCase)

So now, bug-in-a-bottle in hand, we can test our hypothesis. Here's how we
tested it:

○ → docker run -v $(pwd):/tmp/test -it cfbuildpacks/ci:latest bash -c "cd /tmp/test && ruby test.rb

2.0.17
sha: 0911c3aeb2c25dd0a41f0225e0c0f2baaa404ffb9cd772166133572f5fb91112


○ → ruby test.rb

2.1.0
sha: '0911c3aeb2c25dd0a41f0225e0c0f2baaa404ffb9cd772166133572f5fb91112'

```

AHA⚠️ So there it is...problem isolated. But...wait that's not what we
expected...that's the opposite! I mean, clearly there's something to our
hypothesis, but we were just confused at that point! Hmmm. So leaning back, I
let my mind relax. After I woke up from my nap, I vaguely remembered seeing
something like this before...something to do with floats and YAML...lets check
the psych  commits to see if anything stands out.

Again, git to the rescue. Luckily, Github has a great compare ui, which I
utilized like so to [compare the tags](https://github.com/ruby/psych/compare/v2.0.17...v2.1.0). And now that  looks
promising

> "Support YAML 1.2 Core Schema."Looking through that commit, and the issue
linked, I found myself in the YAML spec. Oh lord.

Well it actually wasn't all that bad. I learned that YAML is crazy! But more
importantly, I learned that a new part of the spec changes the behavior for
recognizing floats, making strings starting with a zero optionally resolved as
floats via some implicit typecasting. And sure enough, with [this commit](https://github.com/ruby/psych/commit/b737f0811a9687cf86f44f0a35f61cbde9eac673),
psych was wrapping value starting in zeros in quotes.


And that was more spec compliant, thought it threw us off. It was wasn't a bug,
just a surprise. So we must be pulling in the new code. So the lesson here:  if
you want to really learn why things behave as they do

Read specs and read the code of your dependenciesThat latter piece of advice,
reading the code of your dependencies, is the most tedious but most rewarding
and useful part of debugging a complex issue like this.

We jumped into one of our jobs running on CI, and ran the test code to confirm
the new behavior. Mystery solved.

---

Or was it? Remember the output we saw in the ci:latest  docker container? That
seemed to indicate something else was going on. Spoiler alert: it was, but it
was a bit of a red herring. And yes, I went down that rabbit-hole and learned
something again. Check out the [follow up](/give-a-damn/).
