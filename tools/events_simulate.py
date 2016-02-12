import socket
import sys

client = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
client.connect("/var/run/gluster/events.sock")

key = sys.argv[1]
value = sys.argv[2]

client.sendall("{0}={1}".format(key, value))
print client.recv(1)
