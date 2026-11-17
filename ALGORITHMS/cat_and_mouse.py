# THIS IS A HACKERRANK TASK.
def catAndMouse(x, y, z):
  """
  If cat  catches the mouse first, print Cat A.
If cat  catches the mouse first, print Cat B.
If both cats reach the mouse at the same time, print Mouse C as the two cats fight and mouse escapes.
"""
    catA = max(x, z) - min(x, z)
    catB = max(y, z) - min(y, z)
    if catA > catB:
        return "Cat B"
    elif catB > catA:
        return "Cat A"
    else:
        return "Mouse C"
