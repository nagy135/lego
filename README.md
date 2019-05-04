
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

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
