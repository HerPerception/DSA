import numpy as np

def softmax(logits):
    logits = np.array(logits)
    shifted_logits = logits-np.max(logits)
    exp_values = np.exp(shifted_logits)
    probabilities = exp_values / np.sum(exp_values)
    return probabilities
print(softmax([1000, 1001, 1002]))

