package main

// First implementation. 0ms runtime and 4.69MB memory use.
func canAliceWin(n int) bool {
    turn := 10
    Alice := 0
    Bob := 0
    for n >= turn && turn > 0{
        n -= turn
        if turn % 2 == 0{
            Alice += 1
        }else{
            Bob += 1
        }
        turn -= 1
        }
        return Alice > Bob

}
