package common

import (
  "crypto/md5"
  "encoding/hex"
  "io/ioutil"
  "net"
  "os"
  "os/exec"
  "strconv"
  "strings"
  "syscall"
  "time"
)

func PGrep(cmd string) int {
  _, exitcode := ExecCommand("/usr/bin/pgrep -x '" + cmd +"' > /dev/null")

  return exitcode
}

func ReadFile(path string) (lines []string) {
  if _, err := os.Stat(path); err == nil {
    contents, err := ioutil.ReadFile(path)
    if err != nil {
      return
    }

    lines = strings.Split(string(contents), "\n")
  }
  return
}

func GetUInt64FromFile(path string) uint64 {
  lines := ReadFile(path)
  if len(lines) > 0 {
    return StringToUInt64(lines[0])
  }
  return 0
}

func GetIntFromFile(path string) int {
  lines := ReadFile(path)
  if len(lines) > 0 {
    return StringToInt(lines[0])
  }
  return 0
}

func StringToInt(value string) int {
  i, err := strconv.Atoi(strings.TrimSpace(value))
  if err != nil {
    return 0
  }
  return i
}

func StringToUInt64(value string) uint64 {
  i, err := strconv.ParseUint(strings.TrimSpace(value), 10, 64)
  if err != nil {
    return 0
  }
  return i
}

func KeyInMap(key string, list map[string]string) bool {
  if _, ok := list[key]; ok {
    return true
  }
  return false
}

func StringInArray(key string, list []string) bool {
  for _, l := range list {
    if l == key {
      return true
    }
  }
  return false
}

//func GetEnv(key string, default_value string) string {
//  val, ok := os.LookupEnv(key)
//  if !ok {
//    if len(default_value) > 0 {
//      return default_value
//    }
//  }
//  return val
//}

func Hostname() string {
  host, err := os.Hostname()
  if err != nil {
    return ""
  }

  return host
}

func IpAddress() string {
  addrs, _ := net.InterfaceAddrs()

  for _, a := range addrs {
    if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      if ipnet.IP.To4() != nil {
        return ipnet.IP.String()
      }
    }
  }

  return ""
}

func ToDateTime(timestamp string, layout string) string {
  t, err := time.Parse(layout, timestamp)
  if err != nil {
    return ""
  }
  return t.Format("2006-01-02 15:04:05")
}

func Escape(text string) string {
  return strings.Replace(text, "'", "\\'", -1)
}

func ExecCommand(cmd string) (stdout string, exitcode int) {
  out, err := exec.Command("/bin/bash", "-c", cmd).Output()
  if err != nil {
    if exitError, ok := err.(*exec.ExitError); ok {
      ws := exitError.Sys().(syscall.WaitStatus)
      exitcode = ws.ExitStatus()
    }
  }
  stdout = string(out[:])
  return
}

func MD5(s string) string {
  hash := md5.Sum([]byte(s))
  return hex.EncodeToString(hash[:])
}
