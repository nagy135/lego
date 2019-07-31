package main

import (
    "fmt"
    "strings"
    "os/exec"
    "regexp"
    // "strconv"
)

func main(){
    var result string
    mpcCommand := exec.Command("mpc")
    mpcOut, _ := mpcCommand.Output()
    mpcArray := strings.Split(string(mpcOut), "\n")
    if len(mpcArray) < 3 {
        result = "%{F#555555}  playlist empty%{F-}"
    } else {
        name := mpcArray[0]
        data := mpcArray[1]
        // flags := mpcArray[2]

        if len(name) >= 50 {
            name = name[:46] + "..." + name[len(name)-3:]
        }

        // PLAYING
        r_playing := regexp.MustCompile("\\[([a-z]+)\\]")
        capture_playing := r_playing.FindStringSubmatch(data)
        playing := capture_playing[1]

        if playing == "playing" {
            result = " %{F#b2d3d9} " + name + "%{F-}"
        } else {
            result = " %{F#555555} " + name + "%{F-}"
        }

    }
    fmt.Print("%{A:mpc toggle:}" + result + "%{A}")
}
