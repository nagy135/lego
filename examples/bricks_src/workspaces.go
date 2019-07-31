package main

import "fmt"
// import "strconv"
import "strings"
// import "reflect"
// import "time"
// import "os"
// import "os/signal"
// import "syscall"
import "io/ioutil"
// import "os/exec"


const f_black string = "%{F#0b0b0b}"
const b_black string = "%{B#0b0b0b}"
const f_red string = "%{F#c22330}"
const b_red string = "%{B#c22330}"
const f_green string = "%{F#19a85b}"
const b_green string = "%{B#19a85b}"
const f_yellow string = "%{F#f9dc2b}"
const b_yellow string = "%{B#f9dc2b}"
const f_blue string = "%{F#304bcc}"
const b_blue string = "%{B#304bcc}"
const f_magenta string = "%{F#d13273}"
const b_magenta string = "%{B#d13273}"
const f_cyan string = "%{F#2dc189}"
const b_cyan string = "%{B#2dc189}"
const f_white string = "%{F#b2d3d9}"
const b_white string = "%{B#b2d3d9}"
const f_end string = "%{F-}"
const b_end string = "%{B-}"

func main() {
    dat, _ := ioutil.ReadFile("/tmp/workspaces")
    str_dat := strings.TrimSuffix(string(dat), "\n")
    arr := strings.Split(str_dat, ":")
    // monitor := arr[0]
    workspaces := arr[1:11]
    layout := ""
    if len(arr) >= 12{
        layout = arr[11]
    }
    state := ""
    if len(arr) >= 13{
        state = arr[12]
    }
    flags := ""
    if len(arr) >= 14{
        flags = arr[13]
    }

    // LAYOUT:  Monocle vs Tiled
    layout_wrap:= ""
    if layout == "LM" && 1 == 2 {
        layout_wrap = " " + f_red + "" + f_end + " "
    }

    // FLAGS: Sticky vs Private vs Locked vs Marked
    flags_sep_start := ""
    flags_sep_end := ""
    if strings.Contains(flags, "S"){
        flags_sep_end = flags_sep_end + " " + b_cyan + " " + b_end
        flags_sep_start = b_cyan + " " + b_end + " " + flags_sep_start
    }
    if strings.Contains(flags, "P"){
        flags_sep_end = flags_sep_end + " " + b_yellow + " " + b_end
        flags_sep_start = b_yellow + " " + b_end + " " + flags_sep_start
    }
    if strings.Contains(flags, "L"){
        flags_sep_end = flags_sep_end + " " + b_red + " " + b_end
        flags_sep_start = b_red + " " + b_end + " " + flags_sep_start
    }
    if strings.Contains(flags, "M"){
        flags_sep_end = flags_sep_end + " " + b_blue + " " + b_end
        flags_sep_start = b_blue + " " + b_end + " " + flags_sep_start
    }

    // WORKSPACES
    workspaces_res := ""
    for _,val := range workspaces {
        name:= " " + string(val[1:]) + " "
        occupancy_focus:= string(val[0])

        var res string
        if strings.ToLower(occupancy_focus) == "o" { // occupied
            if occupancy_focus == "O" { // focused occupied
                res = b_green + f_black + name + f_end + b_end
                if state == "TF" {
                    res = b_blue + f_yellow + name + f_end + b_end
                }
                res = flags_sep_start + res + flags_sep_end
            } else { // not focused occupied
                res = b_black + f_green + name + f_end + b_end
            }
        } else if strings.ToLower(occupancy_focus) == "u" { // urgent
            if occupancy_focus == "U" { // urgent free
                res = b_red + f_yellow + name + f_end + b_end
                res = flags_sep_start + res + flags_sep_end
            } else { // not focused free
                res = b_yellow + f_red + name + f_end + b_end
            }
        } else if strings.ToLower(occupancy_focus) == "f" { // free
            if occupancy_focus == "F" { // focused free
                res = b_white + f_black + name + f_end + b_end
                res = flags_sep_start + res + flags_sep_end
            } else { // not focused free
                res = b_black + f_white + name + f_end + b_end
            }
        }
        workspaces_res += res
    }

    result := layout_wrap + workspaces_res + layout_wrap

    fmt.Print(result)





}
