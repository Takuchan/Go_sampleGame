package main

import "fmt"
import "math/rand"
import "time"
// わんちゃん、import ()できそう？

func main(){
  rand.Seed(time.Now().UnixNano())
  maxAge := 100
  putAge := rand.Intn(maxAge)
  guess := 0

  fmt.Println("AI年齢当てゲームをしよう")
  for {
    fmt.Println("推測で私の年齢を答えて。何歳やと思う?: ")
    fmt.Scanf("%d",&guess)

    if guess < putAge {
      fmt.Println("若く見られてて嬉しいわ。もっと上よ。")
    }else if guess < putAge{
      fmt.Println("うーん。もっと下よ。")
    }else{
      fmt.Println("正解")
      break
    }
  }
}
