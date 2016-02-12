#!/usr/bin/env python
"""
A tool to generate Curl commands for Gluster based on the inputs
"""
import hashlib
import hmac
import base64
from datetime import datetime
from urlparse import urlparse

from argparse import ArgumentParser, RawDescriptionHelpFormatter


def get_sign(**kwargs):
    # TODO: Add Data to Sign
    message = bytes("{method}\n{content_type}\n{http_date}\n{url}".format(
        **kwargs)).encode('utf-8')
    secret = bytes(kwargs.get("app_secret")).encode('utf-8')
    return base64.b64encode(hmac.new(secret, message,
                                     digestmod=hashlib.sha256).digest())


def get_curl_cmd(args, space=""):
    op = ["{0}curl".format(space)]
    if args.show_headers:
        op[-1] += " -i"

    op[-1] += " -X {0}".format(args.method)
    http_date = datetime.strftime(datetime.now(),
                                  "%a, %d %b %Y %H:%M:%S +0000")
    content_type = "application/json"
    op.append('-H "Content-Type: {0}"'.format(content_type))
    op.append('-H "Date: {0}"'.format(http_date))
    if args.id:
        url_data = urlparse(args.url)
        signature = get_sign(url=url_data.path,
                             app_secret=args.secret,
                             app_id=args.id,
                             method=args.method,
                             http_date=http_date,
                             content_type=content_type)
        op.append('-H "Authorization: HMAC_SHA256 {id}:{signature}"'.format(
            id=args.id, signature=signature))

    op.append("{url}".format(url=args.url))

    join_str = " \\\n{0}    ".format(space)
    return join_str.join(op)


def main():
    parser = ArgumentParser(formatter_class=RawDescriptionHelpFormatter,
                            description=__doc__)
    parser.add_argument("--id", help="Application ID")
    parser.add_argument("--secret", help="Application Secret")
    parser.add_argument("--method", help="HTTP Method", default="GET")
    parser.add_argument("--data", help="JSON encoded Data")
    parser.add_argument("--show-headers", help="Show HTTP Response headers",
                        action="store_true")
    parser.add_argument("url", help="URL")
    args = parser.parse_args()

    print get_curl_cmd(args)


if __name__ == "__main__":
    main()
