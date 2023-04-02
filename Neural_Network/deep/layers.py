import numpy as np
import mygrad as mg
from .initializers import Uniform
from .activations import softmax, sigmoid, tanh


class Linear:
    '''
    Linear/Dense layer
    '''
    def __init__(self,in_dim,out_dim,initializer=Uniform,bias=True):
        self.in_dim = in_dim
        self.out_dim = out_dim
        self.weights = mg.Tensor(np.random.rand(out_dim,in_dim))
        if bias:
            self.bias = mg.Tensor(np.random.rand(out_dim))
        else:
            self.bias = None
    def __call__(self,inp):
        _,y = inp.shape
        if y != self.in_dim:
            print(f'Wrong Input Features. Please use tensor with {self.in_dim} Input Features')
        output = mg.matmul(inp,self.weights.T)
        if self.bias is not None:
            output += self.bias
        return output
    @property
    def parameters(self):
        if self.bias is not None:
            return [self.weights,self.bias]
        else:
            return [self.weights]


class BatchNorm:
    '''
    Batch Normalization: In batch normalization, the statistics are computed across the batch 
    '''
    def __init__(self,in_dim,eps=1e-05):
        self.gamma = mg.Tensor(np.random.randn(1,in_dim))
        self.beta = mg.Tensor(np.random.randn(1,in_dim))
        self.eps = eps
    def __call__(self,x):
        N, D = x.shape
        # mini-batch mean
        mu = (1./N)*mg.sum(x,axis =0)
        # mini batch variance
        sqr_mu = (x - mu)**2
        var = (1./N)*mg.sum(x,axis =0)
        # normalize
        xhat = (x - mu)/(mg.sqrt(var + self.eps))
        return mg.matmul(xhat,self.gamma.T) + self.beta
    @property
    def parameters(self):
        return [self.gamma,self.beta]


class LayerNorm:
    '''
    Layer Normalization: in layer normalization, the statistics are computed across each feature and are independent of other examples.
    '''
    def __init__(self,in_dim,eps=1e-05):
        self.gamma = mg.Tensor(np.random.randn(1,in_dim))
        self.beta = mg.Tensor(np.random.randn(1,in_dim))
        self.eps = eps
    def __call__(self,x):
        N = x.shape[0]
        # mini-batch mean
        mu = mg.expand_dims(x.mean(axis=1),axis=1)
        # mini batch variance
        sqr_mu = (x - mu)**2
        var = mg.expand_dims((1./N)*mg.sum(x,axis =1),axis=1)
        # normalize
        xhat = (x - mu)/(mg.sqrt(var + self.eps))
        return mg.matmul(xhat,self.gamma.T) + self.beta
    @property
    def parameters(self):
        return [self.gamma,self.beta]


class Dropout:
    '''
    Dropout    
    '''
    def __init__(self,p=0.5):
        self.p = p
    def __call__(self,x):
        mask = np.random.binomial(1, self.p, size=x.shape)
        return (mask * x)/self.p
    @property
    def parameters(self):
        # Dropout has no parameters to return
        return []


#https://sjmielke.com/jax-purify.htm
class LSTMCell:
    '''
    https://mlexplained.com/2019/02/15/building-an-lstm-from-scratch-in-pytorch-lstms-in-depth-part-1/
    '''
    def __init__(self,in_dim,out_dim, bias=True):
        self.W_ih = mg.Tensor(np.random.randn(4*out_dim,in_dim))
        self.W_hh = mg.Tensor(np.random.randn(4*out_dim,out_dim))
        if bias:
            self.bias = mg.Tensor(np.zeros(4*out_dim))
        else:
            self.bias = 0
    def __call__(self,x, h=None, c=None):
        # reshaping input of shape (batch, time, channels)
        # to (time, batch, channels)
#         B, T, C = x.shape
#         x = x.reshape(T,B,C)
        if h.any()==None:
            h = mg.Tensor(np.random.randn(1,B,self.W_hh.shape[1]))
        if c.any()==None:
            c = mg.Tensor(np.random.randn(1,B,self.W_hh.shape[1]))
        ifgo  = (x @ self.W_ih.T) + (h @ self.W_hh.T) + self.bias
        i, f, g, o = np.split(ifgo, 4, axis=-1)
        i = sigmoid(i)
        f = sigmoid(f)
        g = tanh(g)
        o = sigmoid(o)
        new_c = f * c + i * g
        new_h = o * tanh(new_c)
        return new_h, new_c
    @property
    def parameters(self):
        return [self.W_ih,self.W_hh,self.bias]
    
class LSTM:
    def __init__(self,in_dim,out_dim, layers=1, bias=True):
        self.lstm = LSTMCell(in_dim,out_dim)
    def __call__(self,x,h,c):
        T, B, C = x.shape
        H = []
        for i in range(T):
            h,c = lstm(X[i],h,c)
            H.append(h.squeeze())
        return mg.Tensor(H),c
    @property
    def parameters(self):
        return [self.lstm.parameters]

class RNNCell:
    '''
    https://medium.com/dair-ai/building-rnns-is-fun-with-pytorch-and-google-colab-3903ea9a3a79
    pic->https://hackernoon.com/hn-images/1*uubYiUNDmhmR5KOPdJKYtQ.png
    '''
    def __init__(self,n_inputs,n_neurons):
        self.weight_ih = mg.Tensor(np.random.randn(n_inputs,n_neurons))
        self.weight_hh = mg.Tensor(np.random.randn(n_neurons,n_neurons))
        self.bias = mg.zeros((1,n_neurons))
    def __call__(self,x,h):
        Y0 = mg.nnet.tanh(mg.matmul(x,self.weight_ih)) - self.bias
        Y1 = mg.nnet.tanh(mg.matmul(Y0,self.weight_hh) - mg.matmul(h,self.weight_ih)+self.bias)
        return Y0,Y1
    @property
    def parameters(self):
        return [self.weight_ih,self.weight_hh,self.bias]


class Attention:
    '''
    Simple Attention
    '''
    def __init__(self, embed_dim):
        self.d = embed_dim
        self.W_kqv = mg.Tensor(np.random.randn(3*self.d,self.d))
        self.W_out = mg.Tensor(np.random.randn(self.d,self.d))
    def __call__(self,X):
        B,T,_ = X.shape
        K, Q, V = np.split(X@self.W_kqv.T,3,axis=-1)
        attn = softmax((K@Q.swapaxes(1,2))/mg.sqrt(self.d))
        out = (attn @ V) @ W_out
        return out, attn    
    @property
    def parameters(self):
        return [self.W_kqv,self.W_out]
        
class MultiHeadAttention:
    def __init__(self, embed_dim, heads):
        self.d = embed_dim
        self.h = heads
        self.W_kqv = mg.Tensor(np.random.randn(3*self.d,self.d))
        self.W_out = mg.Tensor(np.random.randn(self.d,self.d))
    def __call__(self,X):
        B, T, _ = X.shape
        K, Q, V =  np.split(X@W_kqv.T,3,axis=-1)
        # multi headed attention
        K, Q, V = [i.reshape(B,self.h,T,self.d//self.h) for i in [K,Q,V]]
        attn = softmax((K@Q.swapaxes(-1,-2))/mg.sqrt(self.d//self.h))
        out = (attn @ V).reshape(B,T,self.d) @ self.W_out
        return out, attn    
    @property
    def parameters(self):
        return [self.W_kqv,self.W_out]