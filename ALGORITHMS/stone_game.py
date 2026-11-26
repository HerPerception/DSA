# First implementation. 0ms runtime and 19.25MB memory use.
class Solution:
    def canAliceWin(self, n: int) -> bool:
        turn = 10
        Alice = 0
        Bob = 0
        while n >= turn and turn > 0:
            n -= turn
            if turn % 2 == 0:
                Alice += 1
            else:
                Bob += 1
            turn -= 1
            
        return Alice > Bob
