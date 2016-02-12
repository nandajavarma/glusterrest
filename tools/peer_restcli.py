#!/usr/bin/env python
"""

"""
from __future__ import print_function
from argparse import ArgumentParser, RawDescriptionHelpFormatter
import json
import sys
import subprocess

# enable, disable, cert-gen, start, stop, reload, add, reset, delete, config


APPS_FILE = "/var/lib/glusterd/rest/apps.json"
CONFIG_FILE = "/etc/glusterfs/glusterrest.json"
CONFIG_KEYS = ["port", "https", "enabled", "events_enabled"]
DEFAULT_CONFIG = {
    "port": 443,
    "https": True,
    "enabled": False,
    "events_enabled": False
}


def execute(cmd, fail_msg):
    p = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    out, err = p.communicate()
    if p.returncode != 0:
        output_error(fail_msg + "\n" + err.strip())


def output_error(message):
    print (message, file=sys.stderr)
    sys.exit(1)


def handle_enable(args):
    # Check the current state, if enabled return success
    # Event is dependent on REST Server
    if args.name == "events":
        execute(["systemctl", "enable", "glustereventsd"],
                fail_msg="Failed to enable glustereventsd")

    execute(["systemctl", "enable", "glusterrestd"],
            fail_msg="Failed to enable glusterrestd")


def handle_disable(args):
    if args.name == "rest":
        execute(["systemctl", "disable", "glusterrestd"],
                fail_msg="Failed to enable glusterrestd")

    execute(["systemctl", "disable", "glustereventsd"],
            fail_msg="Failed to enable glustereventsd")


def handle_start(args):
    execute(["systemctl", "start", "gluster{0}d".format(args.name)],
            fail_msg="Failed to start gluster{0}d".format(args.name))


def handle_stop(args):
    execute(["systemctl", "stop", "gluster{0}d".format(args.name)],
            fail_msg="Failed to stop gluster{0}d".format(args.name))


def handle_reload(args):
    execute(["systemctl", "reload", "gluster{0}d".format(args.name)],
            fail_msg="Failed to reload gluster{0}d".format(args.name))


def boolify(value):
    val = False
    if value.lower() in ["enabled", "true", "on", "yes"]:
        val = True
    return val


def handle_config(args):
    data = json.load(open(CONFIG_FILE))

    if args.key is None:
        for k, v in data.items():
            print ("{0}  => {1}".format(k, v))

        return

    if args.key not in CONFIG_KEYS:
        output_error("Invalid Config item")

    if args.reset:
        data[args.key] = DEFAULT_CONFIG[args.key]
    elif args.value is not None:
        # TODO: Validate
        v = args.value
        if args.key == "port":
            try:
                v = int(args.value)
            except ValueError:
                v = DEFAULT_CONFIG[args.key]
        elif args.key in ["https", "enabled", "events_enabled"]:
            v = boolify(args.value)

        data[args.key] = v
    else:
        print (data[args.key])
        return

    with open(CONFIG_FILE, "w") as f:
        f.write(json.dumps(data))


def handle_app_add(args):
    data = json.load(open(APPS_FILE))
    if data.get(args.app_id, None) is not None:
        output_error("Application already exists")

    data[args.app_id] = args.secret

    with open(APPS_FILE, "w") as f:
        f.write(json.dumps(data))


def handle_app_reset(args):
    data = json.load(open(APPS_FILE))
    if data.get(args.app_id, None) is None:
        output_error("Application does not exists")

    data[args.app_id] = args.secret

    with open(APPS_FILE, "w") as f:
        f.write(json.dumps(data))


def handle_app_delete(args):
    data = json.load(open(APPS_FILE))
    if data.get(args.app_id, None) is None:
        output_error("Application does not exists")

    del data[args.app_id]

    with open(APPS_FILE, "w") as f:
        f.write(json.dumps(data))


def main():
    parser = ArgumentParser(formatter_class=RawDescriptionHelpFormatter,
                            description=__doc__)
    parser.add_argument("--prefix", help="Install Prefix",
                        default="/usr/local")
    subparsers = parser.add_subparsers(dest="mode")
    p = subparsers.add_parser("enable", help="Enable REST Server/Eventing")
    p.add_argument("name", help="rest|events")

    p = subparsers.add_parser("disable", help="Disable REST Server/Eventing")
    p.add_argument("name", help="rest|events")

    p = subparsers.add_parser("config",
                              help="REST Server/Eventing Configuration")
    p.add_argument("-k", "--key", help="Config Key")
    p.add_argument("-v", "--value", help="Config Value")
    p.add_argument("--reset", help="Config Reset", action="store_true")

    p = subparsers.add_parser("start", help="Start REST Server")
    p.add_argument("name", help="rest|events", choices=["rest", "events"])

    p = subparsers.add_parser("stop", help="Stop REST Server")
    p.add_argument("name", help="rest|events", choices=["rest", "events"])

    p = subparsers.add_parser("reload", help="Reload REST Server/Eventing")
    p.add_argument("name", help="rest|events", choices=["rest", "events"])

    subparsers.add_parser("cert-gen", help="Generate Certificate for Https")

    p = subparsers.add_parser("app-add", help="Add application")
    p.add_argument("app_id", help="Application ID")
    p.add_argument("secret", help="Application Secret")

    p = subparsers.add_parser("app-reset", help="Reset application")
    p.add_argument("app_id", help="Application ID")
    p.add_argument("secret", help="Application Secret")

    p = subparsers.add_parser("app-delete", help="Delete application")
    p.add_argument("app_id", help="Application ID")

    args = parser.parse_args()
    globals()["handle_" + args.mode.replace("-", "_")](args)


if __name__ == "__main__":
    main()
