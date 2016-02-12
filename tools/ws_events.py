from websocket import create_connection, WebSocketConnectionClosedException
import hmac
import base64
import hashlib
from datetime import datetime
import time


def get_sign(**kwargs):
    # TODO: Add Data to Sign
    message = bytes("{method}\n{content_type}\n{http_date}\n{url}".format(
        **kwargs)).encode('utf-8')
    secret = bytes(kwargs.get("app_secret")).encode('utf-8')
    return base64.b64encode(hmac.new(secret, message,
                                     digestmod=hashlib.sha256).digest())


def main():
    content_type = "application/json"
    http_date = datetime.strftime(datetime.now(),
                                  "%a, %d %b %Y %H:%M:%S +0000")
    signature = get_sign(url="/v1/events",
                         app_secret="sumne",
                         app_id="hello",
                         method="GET",
                         http_date=http_date,
                         content_type=content_type)
    ws = create_connection("ws://fvm1:8080/v1/events",
                           header={
                               "Content-Type": content_type,
                               "Date": http_date,
                               "Authorization": "HMAC_SHA256 {0}:{1}".format(
                                   "hello", signature)
                           })
    print "Sending 'Hello, World'..."
    ws.send("Hello, World")
    print "Sent"
    print "Receiving..."
    while True:
        result = ws.recv()
        print "Received '%s'" % result
        time.sleep(1)

    ws.close()

if __name__ == "__main__":
    try:
        main()
    except WebSocketConnectionClosedException:
        pass
