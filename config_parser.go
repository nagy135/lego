package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "regexp"
    "strconv"
)

var config_path = os.Getenv("XDG_CONFIG_HOME")
var separator string
var bg string = ""
var left_click string
var powerline bool = false
var after_run string
var fonts []string
var stalonetray_settings = make(map[string]string)

func ParseConfig() (map[string][]string, map[string]map[string]string) {
    legorc, err := os.Open(config_path + "/lego/legorc")
    if err != nil {
        fmt.Println(err)
    }
    defer legorc.Close()

    scanner := bufio.NewScanner(legorc)
    var i int
    order := make(map[string][]string)
    settings := make(map[string]map[string]string)
    var setting_focused string


    for scanner.Scan() {
        line := scanner.Text()
        line = strings.TrimSuffix(line, "\n")
        if  line == "" || line == "]"{
            continue
        }
        if match, _ := regexp.MatchString("^\\s*#.*", line); match {
            continue
        }
        if i < 3{
            line_split := strings.Split(line, ":")
            order[line_split[0]] = strings.Split(line_split[1], "|")
        } else if strings.Contains(line, "separator"){
            separator_pair := strings.Split(line,"=")
            separator = strings.Replace(separator_pair[1], "\"", "", 2)
        } else if strings.Contains(line, "powerline"){
            powerline_pair := strings.Split(line,"=")
            powerline, _ = strconv.ParseBool(powerline_pair[1])
        } else if strings.Contains(line, "after_run"){
            after_pair := strings.Split(line,"=")
            after_run = string(after_pair[1])
        } else if strings.Contains(line, "background"){
            bg_pair := strings.Split(line,"=")
            bg = string(bg_pair[1])
        } else if strings.Contains(line, "font"){
            fonts_pair := strings.Split(line,"=")
            fonts = append(fonts, strings.Replace(string(fonts_pair[1]),"\"", "", 2))
        } else if strings.Contains(line, "stalonetray"){
            stalone_pair := strings.Split(line,"=")
            stalone_key := strings.Replace(string(stalone_pair[0]), "stalonetray", "", -1)
            stalonetray_settings[stalone_key] = strings.Replace(string(stalone_pair[1]),"\"", "", 2)
        } else {
            if strings.ContainsAny(line, "["){
                setting_focused = strings.Replace(line, "[", "", 1)
                settings[setting_focused] = make(map[string]string)
            } else {
                line := strings.Replace(line, " ", "", -1)
                setting_pair := strings.Split(line, "=")
                settings[setting_focused][setting_pair[0]] = setting_pair[1]
            }
        }
        i += 1
    }
    fmt.Println("settings", settings)
    return order, settings
}
