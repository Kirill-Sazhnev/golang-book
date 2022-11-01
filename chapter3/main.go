package main

import (
    "fmt"
    "math/rand"
)

var era = "AD"

func main() {

	for i := 0; i < 10; i++ {
    year := rand.Intn(2022)+1
    month := rand.Intn(12) + 1
    daysInMonth := 31

    switch month {
    case 2:
        daysInMonth = 28
				if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
					daysInMonth += 1
				}
    case 4, 6, 9, 11:
        daysInMonth = 30
    }

    day := rand.Intn(daysInMonth) + 1
    fmt.Println(era, year, month, day)
	}
}
  /*
  // func HasPrefix(s, prefix string) bool
  fmt.Println(strings.HasPrefix("test", "te"))

  // func Count(s, sep string) int
  fmt.Println(strings.Count("test", "t"))
  // => 2

  // Contains(s, substr string) bool
  fmt.Println(strings.Contains("test", "es"))
  // => true

  // func HasSuffix(s, suffix string) bool
  fmt.Println(strings.HasSuffix("test", "st"))
  // => true

  // func Join(a []string, sep string) string
  fmt.Println(strings.Join([]string{"a","b"}, "-"))
  // => "a-b"

  // func Index(s, sep string) int
  fmt.Println(strings.Index("test", "e"))
  // => 1

  //func Repeat(s string, count int) string
  fmt.Println(strings.Repeat("a", 5))
  // => "aaaaa"

  //func Replace(s, old, new string, n int) string
  fmt.Println(strings.Replace("aaaa", "a", "b", 2))
  // => "bbaa"

  // func Split(s, sep string) []string
  fmt.Println(strings.Split("a-b-c-d-e", "-"))
  // => []string{"a","b","c","d","e"}

  var buf bytes.Buffer
  buf.Write([]byte("test"))
  fmt.Println(buf)
  */
