import numpy as np
import mygrad as mg

def softmax(x, axis=None):
    x = x - x.max(axis=axis, keepdims=True)
    y = mg.exp(x)
    return y / y.sum(axis=axis, keepdims=True)

def sigmoid(x):
    return 1 / (1 + mg.exp(-x))

def tanh(z):
    x = mg.exp(z)
    y = mg.exp(-z)
    return (x-y)/(x+y)

def relu(x):
    return mg.maximum(x,0)