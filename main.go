package main

import "fmt"
import "strings"
import "log"
import "time"
import "os"
import "io"
import "bytes"
import "os/signal"
import "syscall"
import "io/ioutil"
import "os/exec"
import "strconv"

var order map[string][]string
var settings map[string]map[string]string

const separator_left string = ""
const separator_right string = ""

type timestamp struct {
    t time.Time
}

func main(){
    order, settings = ParseConfig()
    holder := make(map[string]string)
    init_holder(holder)
    history := make(map[string]time.Time)
    init_history(history)

    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)
    timer := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
    for key,_ := range settings {
        if _, ok := settings[key]["subscribe"]; ok {
            sub_file := settings[key]["subscribe"]

            sub := exec.Command("bash", sub_file)
            sub.Start()
        }
    }

    if bg == "" {
        bg = "FF0b0b0b"
    }

    lemonbar_args := []string{"-p",
        "-B", bg,
        "-F", "#CCCCCC",
        "-g", "1920x25+0+0",
    }
    if len(fonts) > 0 {
        for _,val := range fonts {
            back := make([]string, len(lemonbar_args))
            copy(back, lemonbar_args[1:])
            lemonbar_args = append(lemonbar_args[:1], "-f")
            lemonbar_args = append(lemonbar_args, val)
            lemonbar_args = append(lemonbar_args, back...)
        }
    }

    fmt.Println("lemonbar_args", lemonbar_args)
    go func() {
        fmt.Println("starting lemonbar")
        lemon := exec.Command("lemonbar", lemonbar_args...)
        lemon_in, _ := lemon.StdinPipe()
        lemon.Start()

        stalone_args := make([]string, 0)

        for key,val := range stalonetray_settings {
            stalone_args = append(stalone_args, key, val )

        }
        fmt.Println("stalonetray", strings.Join(stalone_args, " "))
        stalonetray := exec.Command("stalonetray", stalone_args...)
        stalonetray.Start()

        ticker := time.NewTicker(time.Second)
        go func() {
            for _ = range ticker.C {
                now := time.Now()
                for key,_ := range history {
                    then := history[key]
                    interval, _ := strconv.ParseFloat(settings[key]["interval"], 64)
                    if now.Sub(then).Seconds() >= interval && interval > 0{
                        history[key] = time.Now()
                        holder[key] = get_block(key)
                        timer <- true
                    }
                }
            }
        }()

        if after_run != "" {
            go func() {
                var stdBuffer bytes.Buffer
                mw := io.MultiWriter(os.Stdout, &stdBuffer)

                time.Sleep(time.Millisecond * 1000)
                fmt.Println("running after_run script")

                after_cmd := exec.Command("bash", after_run)
                after_cmd.Stdout = mw
                after_cmd.Stderr = mw
                if err := after_cmd.Run(); err != nil {
                    log.Println(stdBuffer.String())
                    log.Panic(err)
                }


            }()
        }

        for {
            lemon_in.Write([]byte(lemonize(holder)))
            select{
                case sig := <-sigs:
                    switch sig {
                    case syscall.SIGINT:
                        fmt.Println("interrupted")
                        done <- true
                    case syscall.SIGTERM:
                        fmt.Println("terminated")
                        done <- true
                    case syscall.SIGUSR1:
                        refresh_targets, _ := ioutil.ReadFile("/tmp/lego_refresh")
                        refresh_targets_slices := strings.Split(strings.TrimSuffix(string(refresh_targets), "\n"), "|")
                        fmt.Println("user signal: refreshing blocks", refresh_targets_slices)
                        for _, val := range refresh_targets_slices {
                            holder[val] = get_block(val)
                            lemon_in.Write([]byte(lemonize(holder)))
                        }
                    }
                case _ = <-timer:
                    continue
            }

        }

        done <- true
        lemon_in.Close()
    }()
    <-done
}

func init_holder(holder map[string]string){
    for key,_ := range settings {
        holder[key] = get_block(key)
    }
}

func init_history(history map[string]time.Time){
    for key, _ := range settings {
        history[key] = time.Now()
    }
}

func get_block(key string) string{
    script := exec.Command(config_path + "/lego/bricks/" + key)
    script_out, _ := script.Output()
    return string(script_out)
}
func lemonize(holder map[string]string) string{
    var left, center, right []string
    for _,val := range order["left"]{
        if holder[val] == "" {
            continue
        }
        left = append(left, wrap(holder[val], "left", val, holder))
    }
    for _,val := range order["center"]{
        if holder[val] == "" {
            continue
        }
        center = append(center, wrap(holder[val], "center", val, holder))
    }
    for _,val := range order["right"]{
        if holder[val] == "" {
            continue
        }
        right = append(right, wrap(holder[val], "right", val, holder))
    }
    return "%{l}" + strings.Join(left, "") + "%{c}" + strings.Join(center, "") + "%{r}" + strings.Join(right, "")
}

func wrap(value string, position string, key string, holder map[string]string) string {

    if cmd, ok := settings[key]["left_click"]; ok && settings[key]["left_click"] != "" {
        value = "%{A1:" + cmd + ":}" + value + "%{A}"
    }
    if powerline {
        value = " " + value + " "
        var bg_open string = "%{B" + settings[key]["color"] + "}"
        var fg_open string = "%{F" + settings[key]["color"] + "}"
        var fg_close string = "%{F-}"
        var bg_close string = "%{B-}"
        if position == "left" {
            for i,str := range order[position]{
                if str == key {
                    if i+1 < len(order[position]) && holder[order[position][i+1]] != "" {
                        next := order[position][i+1]
                        var next_bg_open string = "%{B" + settings[next]["color"] + "}"
                        return bg_open + value + bg_close + fg_open + next_bg_open + separator_right + fg_close + bg_close
                    } else {
                        return bg_open + value + bg_close + fg_open + separator_right + fg_close
                    }
                }
            }
            return value
        } else if position == "right" {
            for i,str := range order[position]{
                if str == key {
                    if i-1 >= 0 && holder[order[position][i-1]] != "" {
                        next := order[position][i-1]
                        var next_bg_open string = "%{B" + settings[next]["color"] + "}"
                        return fg_open + next_bg_open + separator_left + bg_close + fg_close + bg_open + value + bg_close
                    } else {
                        return fg_open + separator_left + fg_close + bg_open + value + bg_close
                    }
                }
            }
            return value
        } else { //center
            return value
        }
    } else {
        if position == "left" {
            for i,str := range order[position]{
                if str == key {
                    if i+1 < len(order[position]) && holder[order[position][i+1]] != "" {
                        return separator + value
                    } else {
                        return value
                    }
                }
            }
            return value
        } else if position == "right" {
            for i,str := range order[position]{
                if str == key {
                    if i > 0 && holder[order[position][i-1]] != "" {
                        return separator + value
                    } else {
                        return value
                    }
                }
            }
            return value

        } else { // center
            return value
        }
    }
}
