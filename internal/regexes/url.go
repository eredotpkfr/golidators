package regexes

import rgx "regexp"

const (
  // IP octets
  ipMiddleOctet string = "(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5]))"
  ipLastOctet string = "(?:\\.(?:0|[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-5]))"

  urlPattern string = "^" +
  // protocol identifier
  "(?:(?:https?|ftp)://)" +
  // user:pass authentication
  "(?:[-a-z\u00a1-\uffff0-9._~%!$&'()*+,;=:]+" +
  "(?::[-a-z0-9._~%!$&'()*+,;=:]*)?@)?" +
  "(?:" +
  "(?P<private_ip>" +
  // IP address exclusion
  // private & local networks
  "(?:(?:10|127)" + ipMiddleOctet + "{2}" + ipLastOctet + ")|" +
  "(?:(?:169\\.254|192\\.168)" + ipMiddleOctet + ipLastOctet + ")|" +
  "(?:172\\.(?:1[6-9]|2\\d|3[0-1])" + ipMiddleOctet + ipLastOctet + "))" +
  "|" +
  // private & local hosts
  "(?P<private_host>" +
  "(?:localhost))" +
  "|" +
  // IP address dotted notation octets
  // excludes loopback network 0.0.0.0
  // excludes reserved space >= 224.0.0.0
  // excludes network & broadcast addresses
  // (first & last IP address of each class)
  "(?P<public_ip>" +
  "(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])" +
  "" + ipMiddleOctet + "{2}" +
  "" + ipLastOctet + ")" +
  "|" +
  // IPv6 RegEx from https://stackoverflow.com/a/17871737
  "\\[(" +
  // 1:2:3:4:5:6:7:8
  "([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|" +
  // 1::                              1:2:3:4:5:6:7::
  "([0-9a-fA-F]{1,4}:){1,7}:|" +
  // 1::8             1:2:3:4:5:6::8  1:2:3:4:5:6::8
  "([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|" +
  // 1::7:8           1:2:3:4:5::7:8  1:2:3:4:5::8
  "([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|" +
  // 1::6:7:8         1:2:3:4::6:7:8  1:2:3:4::8
  "([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|" +
  // 1::5:6:7:8       1:2:3::5:6:7:8  1:2:3::8
  "([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|" +
  // 1::4:5:6:7:8     1:2::4:5:6:7:8  1:2::8
  "([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|" +
  // 1::3:4:5:6:7:8   1::3:4:5:6:7:8  1::8
  "[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|" +
  // ::2:3:4:5:6:7:8  ::2:3:4:5:6:7:8 ::8       ::
  ":((:[0-9a-fA-F]{1,4}){1,7}|:)|" +
  // fe80::7:8%eth0   fe80::7:8%1
  // (link-local IPv6 addresses with zone index)
  "fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|" +
  "::(ffff(:0{1,4}){0,1}:){0,1}" +
  "((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}" +
  // ::255.255.255.255   ::ffff:255.255.255.255  ::ffff:0:255.255.255.255
  // (IPv4-mapped IPv6 addresses and IPv4-translated addresses)
  "(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|" +
  "([0-9a-fA-F]{1,4}:){1,4}:" +
  "((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}" +
  // 2001:db8:3:4::192.0.2.33  64:ff9b::192.0.2.33
  // (IPv4-Embedded IPv6 Address)
  "(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])" +
  ")\\]|" +
  // host name
  "(?:(?:(?:xn--)|[a-z\u00a1-\uffff\U00010000-\U0010ffff0-9]-?)*" +
  "[a-z\u00a1-\uffff\U00010000-\U0010ffff0-9]+)" +
  // domain name
  "(?:\\.(?:(?:xn--)|[a-z\u00a1-\uffff\U00010000-\U0010ffff0-9]-?)*" +
  "[a-z\u00a1-\uffff\U00010000-\U0010ffff0-9]+)*" +
  // TLD identifier
  "(?:\\.(?:(?:xn--[a-z\u00a1-\uffff\U00010000-\U0010ffff0-9]{2,})|" +
  "[a-z\u00a1-\uffff\U00010000-\U0010ffff]{2,}))" +
  ")" +
  // port number
  "(?::\\d{2,5})?" +
  // resource path
  "(?:/[-a-z\u00a1-\uffff\U00010000-\U0010ffff0-9._~%!$&'()*+,;=:@/]*)?" +
  // query string
  "(?:\\?\\S*)?" +
  // fragment
  "(?:#\\S*)?" +
  "$"
)

// URLRegex for validating URL data
var URLRegex = rgx.MustCompile(makeCaseInsensitive(urlPattern))
