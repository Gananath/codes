import mygrad as mg
import numpy as np


class Optimizers():
    '''    https://hackernoon.com/implementing-different-variants-of-gradient-descent-optimization-algorithm-in-python-using-numpy-809e7ab3bab4
    '''
    def __init__(self,params = None,learning_rate=0.01,beta=[0.9,0.999],gamma=0.1,epsilon=1e-07):
        self.lr = float(learning_rate)
        self.params = params
        self.beta_1,self.beta_2 = beta
        self.gamma = float(gamma)
        self.eps = float(epsilon)
        self.new_updates = 1
        self.value = 0
    @property
    def _get_parameters_from_memory(self):
        params = dict()
        for i in list(globals()):
            if "__" in i:
                continue
            try:
                if eval(i).grad is not None:
                    params[i]=eval(i)
                    
            except:
                pass
        return params
    @property
    def _get_params(self):
        if self.params==None:
            params = self._get_parameters_from_memory.values()
        else:
            params = self.params
        return params
    @property
    def _get_value(self):
        params = self._get_params
        val = []
        for i in params:
            val.append(mg.zeros(i.shape))
        return val
    def ZeroGrad(self):
        params = self._get_params
        for p in params:
            p.null_gradients()
    def GradientDescent(self):
        params = self._get_params
        for p in params:
            p.data -= self.lr * p.grad
            
            
    def Momentum(self):
        params = self._get_params
        if self.value == 0:
            self.value = self._get_value
        for p,v in zip(params,self.value):
            v.data = self.gamma * v.data + self.lr * p.grad
            p.data -= v.data
            
        
    def NesterovAcceleratedGradientDescent(self):
        # currently not works
        params = self._get_params
        if self.value == 0:
            self.value = self._get_value
        for p,v in zip(params,self.value):
            p_temp = p.data - self.gamma * v.data
            p.data = p_temp - self.lr * p_temp.grad
            
    def AdaGrad(self):
        params = self._get_params
        if self.value == 0:
            self.value = self._get_value
        for p,v in zip(params,self.value):
            v.data = v.data + (p.grad)**2
            p.data -= (self.lr/(np.sqrt(v.data)+self.eps))*p.grad
            
    def RMSProp(self):
        params = self._get_params
        if self.value == 0:
            self.value = self._get_value
        for p,v in zip(params,self.value):
            v.data = self.beta_1*v.data + (1 - self.beta_1) * (p.grad)**2
            p.data -= (self.lr/(np.sqrt(v.data)+self.eps))*p.grad
        
    def Adam(self):
        params = self._get_params
        if self.value == 0:
            self.value = self._get_value
        for p,v in zip(params,self.value):
            m = self.beta_1*v.data + (1 - self.beta_1)*p.grad
            v.data = self.beta_2*v.data + (1 - self.beta_2)*(p.grad**2)
            # bias correction
            m = m/(1-np.power(self.beta_1,self.new_updates))
            v.data = v.data/(1-np.power(self.beta_2,self.new_updates))
            p.data = p.data.astype('float') - (self.lr/(np.sqrt(v.data)+self.eps))*m
            self.new_updates += 1
