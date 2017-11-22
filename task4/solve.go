package main
import "fmt"
import "unicode"

func RemoveEven(s []int) (ret []int) {
  for _, i := range s {
    if i % 2 != 0 {
      ret = append(ret, i)
    }
  }
  return ret
}

func PowerGenerator(x int) func() int {
  a := x
  res := 1
  gen := func () int {
    res = res * a
    return res;
  }
  return gen
}

func DifferentWordsCount(s string) int {
  dict := make(map[string]int)
  i := 0
  for i < len(s) {
    var l []string
    var cur string
    flag := false
    for i < len(s) && unicode.IsLetter(rune(s[i])) {
      k := unicode.ToLower(rune(s[i]))
      c := string(k)
      cur = cur + c
      l = append(l, c)
      i++
      flag = true
    }
    for i < len(s) && !unicode.IsLetter(rune(s[i])) {
      i++
    }
    if flag {
      dict[cur] = 1
    }
  }
  return len(dict)
}

// func main() {
//   s := []int{7, 56, 3, 55, 66, 7}
//   fmt.Println(RemoveEven(s))
//   f := PowerGenerator(2)
//   fmt.Println(f())
//   fmt.Println(f())
//   fmt.Println(f())
//   fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
// }
