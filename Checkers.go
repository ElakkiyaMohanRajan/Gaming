package main
import "fmt"
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
var Directions =[2]int{7,9}
var convertTo100 = [64]int{
  11,12,13,14,15,16,17,18,
  21,22,23,24,25,26,27,28,
  31,32,33,34,35,36,37,38,
  41,42,43,44,45,46,47,48,
  51,52,53,54,55,56,57,58,
  61,62,63,64,65,66,67,68,
  71,72,73,74,75,76,77,78,
  81,82,83,84,85,86,87,88}
  var ply int=0
  var positions int=0
  var maxPly int=0
  var humanMove int=0
  var humanScore int=0
  func GetNumForDir(index1 int,index int,dir int,board []int,us int) (int,int,int){
    fmt.Printf("Inside GetNumForDir index %4d \n",index)
    fmt.Printf("Inside GetNumForDir index1 %4d \n",index1)
    fmt.Printf("Inside GetNumForDir us%4d \n",us)
    found:=0
    border:=0
    dir1:=0
     if(index% 8 == 7 && (index - dir)% 8 == 0){

         fmt.Printf("Inside If BORDER\n")
          border=1
     }else{
       if(index% 8 == 0 && (index - dir)% 8 == 7){
           fmt.Printf("Inside If BORDER\n")
            border=1
       }else{
         fmt.Printf("Inside Else BORDER \n")
           fmt.Printf("board[convertTo100[index]] %4d n",board[convertTo100[index]])
    if(board[convertTo100[index]] == us ){
      fmt.Printf("=dir %4d \n",dir)
      if(index-dir < 8){
        fmt.Printf("index-dir %4d \n",index-dir)
          if(board[convertTo100[index+dir]] == EMPTY){
            found=1
            dir1=-dir
          }else{
            if(board[convertTo100[index+dir]] == NOUGHTS){
              found=2
              dir1=-(dir+dir)
            }
          }
           }else{

      fmt.Printf("board[convertTo100[index-dir]] %4d \n",board[convertTo100[index-dir]])
     if(board[convertTo100[index-dir]] == EMPTY ){
                found=1
                dir1=dir
  }else{
    if(board[convertTo100[index-dir]] == NOUGHTS ){
      if((index-dir)-dir < 0){

      }else{
  fmt.Printf("(index-dir)-dir--1%4d \n",(index-dir)-dir)
      if((index-dir)-dir% 8 == 0 ){
        }else{
fmt.Printf("board[convertTo100[index-(dir+dir)]]%4d \n",board[convertTo100[index-(dir+dir)]])
        if(board[convertTo100[index-(dir+dir)]] == EMPTY){
  fmt.Printf("board[convertTo100[index-(dir+dir)]]...1%4d \n",board[convertTo100[index-(dir+dir)]])
              found=2
              dir1=dir+dir
            }
   }
  }
}
}
}
}
}
}
fmt.Printf("found returning%4d \n",found)
  fmt.Printf("board returning%4d \n",border)
   fmt.Printf("dir1 returning%4d \n",dir1)
    return found,border,dir1
  }

  func FindNoCoinsofOpponent(board []int,us int) int{
  index:=0
  winvalue:=0
  for index =0 ; index < 64;index++{
    if(board[convertTo100[index]]== us){
        winvalue=0
        break
    }else{
       winvalue=1
       continue
    }
  }

  return winvalue
  }
func EvalForWin(board []int,us int) int{

    if(us == 0){
      us = 1
    }
    if(us == 1){
      us = 0
    }


    if(FindNoCoinsofOpponent(board,us) != 0){
      return 1
    }
    return 0
  }
  func MinMax(board []int,side int) (int,int,int,int){
    DirIndex:=0
    dir:=0
    var score int= -2
    var index int
    var move1= 0
    var move2=0
    var move3=0
    var found =0
    var border =0
    var dirValue =0

      score= EvalForWin(board,side)

  for index=0;index < 64; index=index+1 {

    for DirIndex =0 ; DirIndex < 2;DirIndex++{

      dir=Directions[DirIndex]

      found,border,dirValue =GetNumForDir(convertTo100[index],index,dir,board,1)

    if(found >0 && border==0){

        move1 =index
        move2= index-dirValue
        break
      }else{


      }
  }
  if(found > 0 && border==0){

      move1 =index
      move2= index-dirValue
      break
    }else{

        continue

    }
}
  if(score != 0){
    return move1,move2,move3,score
  }

if(dirValue == dir*2){
     move3=move1-dir
}
return move1,move2,move3,score
}


  func initializeBoard(board []int){
var index1 int = 0
for index1 = 0; index1< 100 ; index1++ {
    board[index1] = BORDER
}
for index1 = 0; index1 < 64; index1++ {
  board[convertTo100[index1]] = EMPTY
}
for index1 = 0; index1 < 8; index1=index1+2 {
      board[convertTo100[index1]] = NOUGHTS
  }
for index1 = 9; index1 < 16; index1=index1+2 {
        board[convertTo100[index1]] = NOUGHTS
    }
for index1 = 16; index1 < 24; index1=index1+2 {
          board[convertTo100[index1]] = NOUGHTS
      }
for index1 = 41; index1 < 48; index1=index1+2 {
            board[convertTo100[index1]] = CROSSES
        }
for index1 = 48; index1 < 55; index1=index1+2 {
              board[convertTo100[index1]] = CROSSES
          }
for index1 = 57; index1 < 64; index1=index1+2 {
                board[convertTo100[index1]] = CROSSES
            }
}

func PrintBoard(board []int){
  var index1 int = 0
  const pceChars = "OX|-"
fmt.Printf("\n Board: \n")

for index1 = 0; index1< 64; index1++ {
    if(index1 !=0 && index1 % 8 ==0){
      fmt.Printf("\n\n")
    }
    fmt.Printf("%4c",pceChars[board[convertTo100[index1]]])
  }
  fmt.Printf("\n")
}
func HasEmpty(board []int) bool{
  index:=0
  for index =0;index<9;index++{
    if(board[convertTo100[index]] == EMPTY){
      return true
    }
  }
  return false
}

func MakeComputerMove(board []int,sq int ,side int){

    board[convertTo100[sq]]=side
}
func DeleteAlreadyComputerMove(board []int,sq int){

  board[convertTo100[sq]]=EMPTY
}
func GetHumanMove(board []int) (int,int,int){
//var userInput [4]byte
moveOk1:=0
moveOk2:=0
move1 :=-1
move2 :=-1
move3 :=-1
for ( moveOk1==0){
  fmt.Printf("Please Enter From Which position You want to move the  NOUGHTS: ")
  userInput1, err:=fmt.Scanf("%d\n", &move1)
  if err != nil {
           fmt.Println(userInput1, err)
       }
  if(userInput1 != 1){
    move1= -1
    fmt.Printf("Invalid scanf() \n")
    continue
  }
  if(move1 <1 || move1 >64){
    move1=-1
    fmt.Printf("Invalid Range \n")
    continue
  }
  move1=move1-1
  if(board[convertTo100[move1]] == EMPTY){
    move1=-1;
    fmt.Printf("NOUGHTS Not Available \n")
    continue
  }
  moveOk1 = 1
    move1=move1+1
}
for ( moveOk2==0){
  fmt.Printf("Please Enter To Which position You want to move the  NOUGHTS: ")
  userInput2, err:=fmt.Scanf("%d\n", &move2)

  if err != nil {
           fmt.Println(userInput2, err)
       }
  if(userInput2 != 1){
    move2= -1
    fmt.Printf("Invalid scanf() \n")
    continue
  }
  if(move2 <1 || move2 >64){
    move2=-1
    fmt.Printf("Invalid Range \n")
    continue
  }
  if(move2 < move1){
    move2=-1
    fmt.Printf("The position you are moving should be Diagnol wise forward \n")
    continue
  }
  if(((move2 != move1+7) && (move2 != move1+9))){

    if(board[convertTo100[move1+6]] == CROSSES && board[convertTo100[move2-1]] == EMPTY){
      move3 = move1+7

      }else{

        if(board[convertTo100[move1+8]] == CROSSES && board[convertTo100[move2-1]] == EMPTY){

          move3 = move1+9
        }else{
    fmt.Printf("The position you are moving should be Diagnol wise forward \n")
    continue
  }
  }
  }else{

  /*  if(board[convertTo100[move2]] == CROSSES && board[convertTo100[move2+7]] == EMPTY){
      board[move2+7] = NOUGHTS
      board[move2] = EMPTY
    }else{if(board[convertTo100[move2]] == CROSSES && board[convertTo100[move2+9]] == EMPTY){
        board[move2+9] = NOUGHTS
        board[move2] = EMPTY
      }
    } */
  }

/*  if(board[convertTo100[move2]] != EMPTY){
    move2=-1;
    fmt.Printf("Square Not Available \n")
    continue
  } */
  move2=move2-1
  move1=move1-1
  move3=move3-1
  moveOk2 = 1
  humanMove=humanMove+1
  humanScore=humanScore+1

}

return move1,move2,move3
}
func GetComputerMove(board []int,side1 int) (int,int,int,int){
  ply=0
  positions=0
  maxPly=0

  move1,move2,move3,best :=MinMax(board,side1)
   return move1,move2,move3,best
}
func RunGame(){
  GameOver:=false
  side:=NOUGHTS
  LastMoveMadeFrom,LastMoveMadeTo,AnotherLastMoveMadeFrom := 0,0,0
  var board =make([]int,100,100)
  score:=0
  initializeBoard(board)
  PrintBoard(board)
  for (!(GameOver)){
      if(side == NOUGHTS ){
      LastMoveMadeFrom,LastMoveMadeTo,AnotherLastMoveMadeFrom  = GetHumanMove(board)
       MakeComputerMove(board,LastMoveMadeTo,side)
       DeleteAlreadyComputerMove(board,LastMoveMadeFrom)

       if(AnotherLastMoveMadeFrom > 0){
           DeleteAlreadyComputerMove(board,AnotherLastMoveMadeFrom)
       }
       side=CROSSES
        PrintBoard(board)
       }else{

           LastMoveMadeFrom,LastMoveMadeTo,AnotherLastMoveMadeFrom,score= GetComputerMove(board,side)

           MakeComputerMove(board,LastMoveMadeTo,side)
           DeleteAlreadyComputerMove(board,LastMoveMadeFrom)
           if(AnotherLastMoveMadeFrom !=0){
               DeleteAlreadyComputerMove(board,AnotherLastMoveMadeFrom)
           }
           side=NOUGHTS
           PrintBoard(board)
       }
       winScore:=FindNoCoinsofOpponent(board,side)
       if(winScore == 1){
         fmt.Printf("Game Over! \n")
         GameOver=true
           if(side == NOUGHTS ){
               fmt.Printf("Computer Wins! \n")
               fmt.Println("Score :",score)
           }else{
             fmt.Printf("Human Wins! \n")
           }
       }

      }

  }


func main(){
  RunGame()
var board =make([]int,100,100)
initializeBoard(board)
//PrintBoard(board)
}
