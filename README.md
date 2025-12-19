# Flatpak Raw Executor

A simple command line application to let you run commands from within a Flatpak context 

## Why?

Allows you to test execution within a Flatpak environment to test permission levels.

## Security

Obviously it's insecure. Don't ship this in your project.

## Install steps

```shell
sudo apt install flatpak-builder
flatpak remote-add --user --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
flatpak install flathub org.freedesktop.Sdk//23.08 org.freedesktop.Platform//23.08
flatpak install flathub org.freedesktop.Sdk.Extension.golang//23.08
```