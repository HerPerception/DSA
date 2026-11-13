def breakingRecords(scores):
"""Return a list containing the number of times Maria breaks her season high and low scoring records."""
    min_score = scores[0]
    max_score = scores[0]
    min_breaks = 0
    max_breaks = 0
    for i in range(1, len(scores)):
        if scores[i] > max_score:
            max_score = scores[i]
            max_breaks += 1
        elif scores[i] < min_score:
            min_score = scores[i]
            min_breaks += 1

    return [max_breaks, min_breaks]
