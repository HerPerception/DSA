# First implementation. Suggested by AI. 124ms runtime and 33.42MB memory use.
class ProductOfNumbers:

    def __init__(self):
        self.prefixes = [1]

    def add(self, num: int) -> None:
        if num == 0:
            self.prefixes = [1]
        else:
            self.prefixes.append((self.prefixes[-1] * num))

    def getProduct(self, k: int) -> int:
        if k >= len(self.prefixes):
            return 0
        return self.prefixes[-1]//self.prefixes[-1-k]


# Your ProductOfNumbers object will be instantiated and called as such:
# obj = ProductOfNumbers()
# obj.add(num)
# param_2 = obj.getProduct(k)
