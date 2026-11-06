# First implementation. Logic was gotten from AI. 1ms runtime and 19.49MB memory use.
class Solution:
    def calPoints(self, operations: list[str]) -> int:
        records = []
        for val in operations:
            try:
                val = int(val)
                records.append(val)
            except ValueError:
                if val == "C":
                    records.pop()
                elif val == "D":
                    v = records[len(records)-1] * 2
                    records.append(v)
                elif val == "+":
                    v = sum(records[len(records)-2:])
                    records.append(v)

        return sum(records)
        
