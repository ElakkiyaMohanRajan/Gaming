package main
import (
  "fmt"
)
const (
  NOUGHTS = 0
  CROSSES = 1
  BORDER  = 2
  EMPTY   = 3)
const (
  HUMANWIN = 0
  COMPWIN  = 1
  DRAW     = 2
)
var Directions =[4]int{1,5,4,6}
var convertTo25 = [9]int{
  6,7,8,
  11,12,13,
  16,17,18}
  const InMiddle int=4
  var Corners = [4]int{0,2,6,8}
  var ply int=0
  var positions int=0
  var maxPly int=0
func GetNumForDir(startsq int,dir int,board []int,us int) int{
  found:=0
  for(board[startsq] != BORDER){
    if(board[startsq] != us){
      break
    }
    found=found+1
    startsq = startsq+dir
  }
  return found
}
func FindThreeInARow(board []int,ourindex int,us int) int{
DirIndex:=0
dir:=0
threeCount :=1
for DirIndex =0 ; DirIndex < 4;DirIndex++{
  dir=Directions[DirIndex]
  threeCount = threeCount +GetNumForDir(ourindex+dir,dir,board,us)
  threeCount = threeCount +GetNumForDir(ourindex+dir * -1 ,dir * -1,board,us)
  if(threeCount == 3){
    break
  }
  threeCount =1
}
return threeCount
}
func FindThreeInARowAllBoard(board []int,us int) int{
threeFound :=0
var index int=0
for index =0 ; index < 9;index++{
  if(board[convertTo25[index]] == us){
    if(FindThreeInARow(board,convertTo25[index],us) == 3){
      threeFound = 1
      break
    }
  }
}
return threeFound
}
func EvalForWin(board []int,us int) int{
  if(FindThreeInARowAllBoard(board,us) != 0){
    return 1
  }
  if(FindThreeInARowAllBoard(board,us ^ 1) != 0){
    return -1
  }
  return 0
}

func MinMax(board []int,side int) int{
  var MoveList [10]int
  var MoveCount int= 0
  var bestScore int= -2
  var score int= -2
  var bestMove int =-1
  var Move int
  var index int
  if(ply > maxPly){
    maxPly=ply
    positions=positions+1
  }
  if(ply >0){
    score= EvalForWin(board,side)
    if(score != 0){
      return score
    }
  }

for index=0;index<9;index++{
  if(board[convertTo25[index]] == EMPTY){
    MoveList[MoveCount]=convertTo25[index]
    MoveCount=MoveCount+1
  }
}

for index=0;index<MoveCount;index++{
  Move=MoveList[index]
  board[Move]=side

  ply++
  score= -MinMax(board,side^1)
  if(score > bestScore){
    bestScore = score
    bestMove = Move
  }
  board[Move] = EMPTY
  ply=ply-1
}

if(MoveCount == 0){
  bestScore=FindThreeInARowAllBoard(board,side)
}
if(ply != 0){
  return bestScore
}else{
  return bestMove
}
}
func initializeBoard(board []int){
var index int = 0
for index = 0; index< 25 ; index++ {
  board[index] = BORDER
}
for index = 0; index < 9; index++{
  board[convertTo25[index]] = EMPTY
}
}
func PrintBoard(board []int){
var index int =0
const pceChars = "OX|-"
fmt.Printf("\n\n Board: \n\n")
for index = 0; index < 9 ; index++ {
  if(index !=0 && index % 3 ==0){
    fmt.Printf("\n\n")
  }
  fmt.Printf("%4c",pceChars[board[convertTo25[index]]])
}
fmt.Printf("\n")
}
func HasEmpty(board []int) bool{
  index:=0
  for index =0;index<9;index++{
    if(board[convertTo25[index]] == EMPTY){
      return true
    }
  }
  return false
}
func MakeMove(board []int,sq int ,side int){
  board[sq]=side
}
func GetComputerMove(board []int,side1 int) int{
  ply=0
  positions=0
  maxPly=0
  best := MinMax(board,side1)
    return best

}

func GetHumanMove(board []int) int{
moveOk:=0
var move int=-1
for ( moveOk==0){
  fmt.Printf("Please Enter a move from 1 to 9: ")
  userInput, err:=fmt.Scanf("%d\n", &move)
  if err != nil {
           fmt.Println(userInput, err)
       }
  if(userInput != 1){
    move= -1
    fmt.Printf("Invalid scanf() \n")
    continue
  }
  if(move <1 || move >9){
    move=-1
    fmt.Printf("Invalid Range \n")
    continue
  }
  move=move-1
  if(board[convertTo25[move]] != EMPTY){
    move=-1;
    fmt.Printf("Square Not Available \n")
    continue
  }
  moveOk = 1

}
return convertTo25[move]
}
func RunGame(){
  GameOver:=false
  side:=NOUGHTS
  LastMoveMade := 0
  var board =make([]int,25,25)
  initializeBoard(board)
  PrintBoard(board)
  for (!(GameOver)){
    if(side == NOUGHTS ){
       LastMoveMade = GetHumanMove(board)
       MakeMove(board,LastMoveMade,side)
       side=CROSSES
    }else{
      LastMoveMade = GetComputerMove(board,side)
      MakeMove(board,LastMoveMade,side)
      side=NOUGHTS
      PrintBoard(board)
    }
    if(FindThreeInARow(board,LastMoveMade,side ^ 1) == 3){
        fmt.Printf("Game Over! \n")
        GameOver=true
          if(side == NOUGHTS ){
              fmt.Printf("Computer Wins! \n")
          }else{
            fmt.Printf("Human Wins! \n")
          }
    }
    if(!HasEmpty(board)){
      fmt.Printf("Game Over! \n")
      GameOver=true
      fmt.Printf("It's a Draw! \n")
    }
  }
 PrintBoard(board)
}

func main(){
  RunGame()
var board =make([]int,25,25)
initializeBoard(board)
}
