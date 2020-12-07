#!/bin/sh

podman run --rm --name fedora -d fedora sleep 60
podman run --rm --name dnstiming -t dnstiming /dnstiming fedora
