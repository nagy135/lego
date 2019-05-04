
# LemonbarGo(LEGO)

Simple configurable lemonbar wrapper written in go. Supports single block refresh, timers, user signals, stalonetray. Contains multiple examples and preconfigured config.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Needs [lemonbar](https://github.com/LemonBoy/bar),preconfigured to work with [stalonetray](https://wiki.archlinux.org/index.php/Stalonetray) if installed.

### Installing

Installing process is very simple. After prerequisities are installed run

```
git clone "https://github.com/nagy135/lego"
```

and then

```
cd lego
make
sudo make install
./lego
```

### Config

Firstly check config

```
left:music|packages
center:workspaces
right:torrent|volume|battery|brightness|redshift|wifi|layout|date

separator=" %{F#FF0000}#%{F-} "
powerline=true
after_run=/home/infiniter/.scripts/lego_fix_layers

# format #aarrggbb
background=#ff0b0b0b

font="Inconsolata for Powerline-19"
font="FontAwesome-19"

stalonetray--geometry=1x1+700+0
stalonetray--grow-gravity=E
stalonetray--icon-gravity=W
stalonetray-bg=#0b0b0b

[music
    interval=0
    color=#FF282A2E
    subscribe=/usr/bin/subscribe_music
]
```

## License

This project is licensed under the DO_WTF_U_WANT_WITH_IT License.

## Acknowledgments

Project is under initial development.
