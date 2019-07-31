package main

import "fmt"
import "strconv"
import "strings"
import "regexp"
import "os/exec"


const b_end string = "%{B-}"

func main() {
    result := ""
    acpi := exec.Command("acpi")
    acpiOut, _ := acpi.Output()
    acpiData := strings.Split(string(acpiOut), "\n")
    for _, line := range acpiData{
        if strings.Trim(line, " ") == ""{
            continue
        }
        match, _ := regexp.MatchString("^.*unavailable.*$", line)
        if ! match {
            statusReg, _ := regexp.Compile("Charging|Discharging|Full")
            status := statusReg.FindString(line)
            percentage, _ := regexp.Compile("([0-9]+)%")
            percent, _ := strconv.Atoi(percentage.FindStringSubmatch(line)[1])
            if status ==  "Charging" || status == "Full"{
                result = "%{F#f9dc2b} %{F-}" + result
            }
            // if status == "Discharging"{
            // }
            // if status == "Full"{
            //     result = ""
            // }

            bat := ""
            if percent < 20 {
                bat = ""
            } else if percent < 40 {
                bat = ""
            } else if percent < 60 {
                bat = ""
            } else if percent < 80 {
                bat = ""
            } else {
                bat = ""
            }

            max_red := "#c22330"
            max_red_r, _ := strconv.ParseInt(max_red[1:3], 16, 64)
            max_red_g, _ := strconv.ParseInt(max_red[3:5], 16, 64)
            max_red_b, _ := strconv.ParseInt(max_red[5:7], 16, 64)
            max_green := "#19a85b"
            max_green_r, _ := strconv.ParseInt(max_green[1:3], 16, 64)
            max_green_g, _ := strconv.ParseInt(max_green[3:5], 16, 64)
            max_green_b, _ := strconv.ParseInt(max_green[5:7], 16, 64)
            step_r := ( float64(max_green_r) - float64(max_red_r) ) / 100
            step_g := ( float64(max_green_g) - float64(max_red_g) ) / 100
            step_b := ( float64(max_green_b) - float64(max_red_b) ) / 100
            red := int(float64(max_red_r) + float64(percent) * step_r)
            green := int(float64(max_red_g) + float64(percent) * step_g)
            blue := int(float64(max_red_b) + float64(percent) * step_b)
            red_final := strconv.FormatInt(int64(red), 16)
            green_final := strconv.FormatInt(int64(green), 16)
            blue_final := strconv.FormatInt(int64(blue), 16)
            if status ==  "Charging" || status == "Full"{
                result += "%{F#" + red_final + green_final + blue_final + "}" + bat + " " + "%{F#f9dc2b}" + strconv.Itoa(percent) + "%%{F-}" + "%{F-}"
            } else {
                result += "%{F#" + red_final + green_final + blue_final + "}" + bat + " " + strconv.Itoa(percent) + "%%{F-}"
            }

        }
    }
    fmt.Print(result)


}
