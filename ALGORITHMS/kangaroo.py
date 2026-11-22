# THIS IS A HACKERRANK TASK.
def kangaroo(x1, v1, x2, v2):
  """Return 'YES' if two kangaroos land on the same spot on the same jump, else 'NO'."""
    if v1 <= v2:
        return "NO"

    j = (x2 - x1) // (v1 - v2)
    if x1 + (j * v1) == x2 + (j * v2):
        return "YES"

    return "NO"
