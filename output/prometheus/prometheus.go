package prometheus

import (
  "fmt"
  "strings"
  "gitlab.com/swapbyt3s/zenit/accumulator"
)

func Run() {
  var a = accumulator.Load()

  for _, m := range *a {
    switch m.Values.(type) {
    case int, uint, uint64, float64:
      fmt.Printf("%s{%s} %s\n", m.Key, getTags(m.Tags), getValue(m.Values))
    case []accumulator.Value:
      for _, i := range m.Values.([]accumulator.Value) {
        fmt.Printf("%s{%s,type=\"%s\"} %s\n", m.Key, getTags(m.Tags), i.Key, getValue(i.Value))
      }
    }
  }
}

func getTags(tags []accumulator.Tag) string {
  s := []string{}
  for t := range(tags) {
    k := tags[t].Name
    v := strings.ToLower(tags[t].Value)
    s = append(s, fmt.Sprintf("%s=\"%s\"", k, v))
  }
  return strings.Join(s,",")
}

func getValue(value interface{}) string {
  switch v := value.(type) {
  case int, uint, uint64:
    return fmt.Sprintf("%d", v)
  case float64:
    return fmt.Sprintf("%.2f", v)
  }

  return "0"
}