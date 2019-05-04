
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

Sample config

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

First 3 lines define positions of blocks and their order

```
left:music|packages
center:workspaces
right:torrent|volume|battery|brightness|redshift|wifi|layout|date
```

Set separator character (if not using powerline mode)

```
separator=" %{F#FF0000}#%{F-} "
```

Turns on powerline mode, needs specified color for each block
```
powerline=true
```

Sets script that is ran after bar is created (to set your xorg layers properly). Features example used in BSPWM.
```
after_run=/home/infiniter/.scripts/lego_fix_layers
```

Sets background color, format #aarrggbb (allows transparency)

```
background=#ff0b0b0b
```

Accepts multiple fonts, uses them all (one for text, one icons)

```
font="Inconsolata for Powerline-19"
font="FontAwesome-19"
```

Settings for stalonetray

```
stalonetray--geometry=1x1+700+0
stalonetray--grow-gravity=E
stalonetray--icon-gravity=W
stalonetray-bg=#0b0b0b
```


Example block with name "music", no refresh interval, color #FF282A2E and enabled subscribe script. This script will be ran to refresh block on events.
```
[music
    interval=0
    color=#FF282A2E
    subscribe=/usr/bin/subscribe_music
]
```

Comments start with "#"

```
# Comment line
```

## License

This project is licensed under the DO_WTF_U_WANT_WITH_IT License.

## Acknowledgments

Project is under initial development.
