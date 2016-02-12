import requests
import hashlib
import hmac
import base64

APP_NAME = "hello"
APP_SECRET = "sumne"
ROOT = "http://192.168.122.188:8080"
VOLUME_NAME = "gv1"


def auth_header(**kwargs):
    message = bytes("{method}\n{url}".format(**kwargs)).encode('utf-8')
    secret = bytes(APP_SECRET).encode('utf-8')
    signature = base64.b64encode(hmac.new(secret, message,
                                          digestmod=hashlib.sha256).digest())

    return "{0} {1}".format(APP_NAME, signature)


def test_volume_start():
    v = requests.post(ROOT + "/v1/volumes/gv1/start")
    assert v.status_code == 401
    assert (v.json() ==
            {"message": "Missing 'Authorization' header"})


def test_volume_start_with_login():
    url = "/v1/volumes/gv1/start"
    headers = {"Authorization": auth_header(method="POST", url=url)}
    v = requests.post(ROOT + url, headers=headers)
    assert v.status_code == 200
