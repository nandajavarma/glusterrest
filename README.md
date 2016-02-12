# Gluster REST API Server and Eventing


## Install

    git clone https://github.com/aravindavk/glusterrestd
    make build
    sudo make install

Prefix can be specified during install, by default installs to `/usr/local`

    sudo make install PREFIX=/usr

## Usage

Gluster CLI work is in progress, meanwhile we can use `gluster system:: execute` command to enable/disable REST Server/Eventing

    gluster system:: execute restcli.py rest enable|disable

To enable Eventing,

    gluster system:: execute restcli.py eventing enable|disable

## Secured Server
To run glusterrestd in secured mode using https, run the following commands to generate the key and certificate

    gluster system:: execute restcli.py cert-gen

By default rest server runs in port 443, can be changed using config command,

    gluster system:: execute restcli.py config port 4443
    gluster system:: execute restcli.py config https enabled|disabled
    
Finally start the REST server,

    gluster system:: execute restcli.py start|stop|reload

If port or https config is changed, then we need to restart the REST service. For other configs or registering apps reload is sufficient.

## API Documentation
API documentation is available [here](docs/API.md).
