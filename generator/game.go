package generator

import (
    i "math/rand"
    "time"
    s "crypto/rand"
    "encoding/hex"
    "math"
)

//globall variables for the game and card
    var data = map[string]map[string]bool{
    "p1": map[string]bool{},
    "p2": map[string]bool{},
    "p3": map[string]bool{},
    "p4": map[string]bool{},
    "p5": map[string]bool{},
    "p6": map[string]bool{},
}
var data_card = map[int]string{
    0:"dia-3",
    1:"dia-4",
    2:"dia-5",
    3:"dia-6",
    4:"dia-7",
    5:"dia-8",
    6:"dia-9",
    7:"dia-10",
    8:"dia-j",
    9:"dia-q",
    10:"dia-k",
    11:"dia-a",
    12:"spd-3",
    13:"spd-4",
  14:"spd-5",
  15 :"spd-6",
  16 :"spd-7",
  17 :"spd-8",
  18 :"spd-9",
  19 :"spd-10",
  20 :"spd-j",
  21 :"spd-q",
  22 :"spd-k",
  23 :"spd-a",
  24 :"hrt-3",
  25 :"hrt-4",
  26 :"hrt-5",
  27 :"hrt-6",
  28 :"hrt-7",
  29 :"hrt-8",
  30 :"hrt-9",
  31 :"hrt-10",
  32 :"hrt-j",
  33 :"hrt-q",
  34 :"hrt-k",
  35 :"hrt-a",
  36 :"clb-3",
  37 :"clb-4",
  38 :"clb-5",
  39 :"clb-6",
  40 :"clb-7",
  41 :"clb-8",
  42 :"clb-9",
  43 :"clb-10",
  44 :"clb-j",
  45 :"clb-q",
  46 :"clb-k",
  47 :"clb-a",
}


// struct to write into DB
type game_setup struct {
    Game_id   string
    Players_card map[string]map[string]bool
}

//function to random_card set for 6 players
func random_card()  map[string]map[string]bool {

  random_array := randomizer()
  var count int=0

   for k := range data{
     var player_count int=0
     for player_count<8{
          data[k][data_card[random_array[count]]]=true
         player_count= player_count+1
         count=count+1
     }
  
  }

  return data
}



//Cards shuffle
func randomizer()  []int {
i.Seed(time.Now().Unix())
return i.Perm(48)
}



// random game id generator
func random_id(l int) string {
  buff := make([]byte, int(math.Ceil(float64(l)/2)))
    s.Read(buff)
    str := hex.EncodeToString(buff)
    return str[:l] // strip 1 extra character we get from odd length results
}



//Game generator
func Game() *game_setup {
  random_array_card := random_card()
  random_game_id := random_id(5)
  game := &game_setup{
     Game_id:random_game_id,
    Players_card:random_array_card} 
   return game
}



