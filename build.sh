#!/bin/sh

go build -o dnstiming main.go
container=$(buildah from fedora:33)
mnt=$(buildah mount $container)
cp dnstiming $mnt
chmod 755 $mnt/dnstiming
buildah config --cmd /dnstiming $container
buildah commit $container dnstiming
buildah unmount $container
