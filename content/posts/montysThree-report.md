---
title: "MontysThree Report"
date: 2020-10-14T15:56:51-04:00
draft: true
---

## ThreesMonty and what you need to know

###  Summary
On October 8, 2020, Kaspersky analysts published a report on a highly-targeted malware toolset they dubbed "MontysThree". MontysThree targets the industrial space and appears to target victims in Cyrillic speaking locales. Additionally, MT3 appears to have been compiled in 2018 and does not appear to be currently active in the wild. The malware contains a number of novel techniques, but is not to be considered a widespread outside of Russian language companies in the ICS space, due to its selective nature.

### Key points
 - An email program executing an self-extracting archive that executes a DLL masquerading as anything eg. a document (`.pdf` or `.doc`) should be viewed as suspicious.
 - C2 traffic will be hard to distinguish from legitimate traffic because it uses popular cloud storage providers. However, if a process is communicating with cloud services in addition to possible MT3 infrastructure, this would be a strong indicator of malicious activity.
 - Processes modifying host `.lnk` files should be viewed as suspicious.
 - Processes that execute RDP, WebDAV, or Citrix client applications as child processes sand subsequently execute clipboard actions (Ctrl-V, Ctrl-p, etc) should be viewed as suspicious 
 - MT3 targets Cyrillic language hosts.
 
### Details

MontysThree (MT3) is delivered via spearphishing emails containing references to medical results, corporate phone records, and various other topics. The malware is delivered via an email attachment in the form of a self-extracting archive (`SFX`) that extracts and opens the loader module and then deletes it from disk. The loader module (now running in memory) then unpacks MT3 core functionality from a specially prepared image (steganography).

Ultimately the actions of this malware are set by a configuration file embedded within the executable. Configuration specifies tasks for the malware to run. MT3 looks for recent files run by Microsoft office (using the Russian words for "Recent Files") and may take screenshots and 

MT3 is equipped to upload and download data from GDrive and Dropbox and know the RDP, WebDAV, Citrix, and HTTP protocols, but uses client applcations for communication. During remote connections MT3 will literally send commands to move data in and out of the system clipboard. It seems to reach out to C2 infrastructure hosted on Digital Ocean.

MT3 achieves persistence by modifying `.lnk` file in the Windows Quick Launch panel to run the loader alongside legitimate applications when the user executes them with the modified link.

### Known indicators
* Caveat * Indicators come from single report. No sample was available for analysis.

Dropper sample MD5: 3afa43e1bc578460be002eb58fa7c2de 

Domains and IPs:
- autosport-club.tekcities[.]com
- dl10-web-stock[.]ru
- dl16-web-eticket[.]ru
- dl166-web-eticket[.]ru
- dl55-web-yachtbooking[.]xyz

MT3 uses legitimate cloud services for C2

### Sources:

- https://securelist.com/montysthree-industrial-espionage/98972/
- https://otx.alienvault.com/pulse/5f7f1fbb47295b54fbf69b6a
- https://www.virustotal.com/gui/file/81ff04b2ce933c7064c3aee78aa97d521752d966738c4e02dfba5755da7d3af9/detection


### Analyst notes/ outstanding questions
- Why the long time-lag between compilation and initial report? Some sandboxes and samples seem to go back to 2018, but they only get followup in 2020.
- Further samples for more in-depth followup if interested: https://twitter.com/craiu/status/1314147971160735745?s=20
