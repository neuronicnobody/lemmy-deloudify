

A plugin for [Lemmy](https://join-lemmy.org/) that modifies community posts by detecting excessive use of caps lock and converting the text to lowercase.
This helps normalize the text, making posts more readable and suitable for certain communities.
The plugin was developed to test out this [proof of concept for adding a plugin system to Lemmy](https://github.com/LemmyNet/lemmy/pull/4695) using 
[Extism](https://extism.org/)

## Installation
Place the `deloudify.wasm` file in the plugins directory of your Lemmy instance. See [here](https://github.com/LemmyNet/lemmy-docs/blob/plugins/src/contributors/08-plugins.md) for more details

## Modifying the Plugin
- install [tinygo](https://tinygo.org/)
- modify the main.go file to your liking
- run `make build`

## Extism Go Plugin
This plugin was written in Go using the Extism Go PDK.

See more documentation at https://github.com/extism/go-pdk and
[join us on Discord](https://extism.org/discord) for more help.


