import numpy as np
import mygrad as mg
from .initializers import Uniform


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
        self.gamma = mg.Tensor(np.random.rand(1,in_dim))
        self.beta = mg.Tensor(np.random.rand(1,in_dim))
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
        self.gamma = mg.Tensor(np.random.rand(1,in_dim))
        self.beta = mg.Tensor(np.random.rand(1,in_dim))
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
    def __init__(self,in_dim,out_dim,h,c):
        self.weight_ih = mg.Tensor(np.random.rand(4*out_dim,in_dim))
        self.weight_hh = mg.Tensor(np.random.rand(4*out_dim,in_dim))    
        self.bias = mg.Tensor(np.zeros(4*out_dim))
    def __call__(self,x):
        ifgo = mg.matmul(self.weight_ih,x) + mg.matmul(self.weight_hh,h)
        i, f, g, o = np.split(ifgo,4)
        i = mg.nnet.sigmoid(i)
        f = mg.nnet.sigmoid(f)
        g = mg.nnet.tanh(g)
        o = mg.nnet.sigmoid(o)
        new_c = f * c + i * g
        new_h = o * mg.nnet.tanh(new_c)
        return new_h, new_c

class RNNCell:
    '''
    https://medium.com/dair-ai/building-rnns-is-fun-with-pytorch-and-google-colab-3903ea9a3a79
    pic->https://hackernoon.com/hn-images/1*uubYiUNDmhmR5KOPdJKYtQ.png
    '''
    def __init__(self,n_inputs,n_neurons):
        self.weight_ih = mg.Tensor(np.random.rand(n_inputs,n_neurons))
        self.weight_hh = mg.Tensor(np.random.rand(n_neurons,n_neurons))
        self.bias = mg.zeros((1,n_neurons))
    def __call__(self,x,h):
        Y0 = mg.nnet.tanh(mg.matmul(x,self.weight_ih)) - self.bias
        Y1 = mg.nnet.tanh(mg.matmul(Y0,self.weight_hh) - mg.matmul(h,self.weight_ih)+self.bias)
        return Y0,Y1
    @property
    def parameters(self):
        return [self.weight_ih,self.weight_hh,self.bias]



