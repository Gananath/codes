from deep.optimizers import Optimizers
from deep.layers import Linear,BatchNorm,Dropout
from deep.losses import Loss
import numpy as np
import mygrad as mg

class Net:
    def __init__(self):
        self.dense1 = Linear(2,3,bias=True)
        self.dense2 = Linear(3,1,bias=True)
        self.bn = BatchNorm(3)
        self.dropout = Dropout()
    def __call__(self,x):
        x = mg.nnet.relu(self.dense1(x))
        x = self.bn(x)
        x = self.dropout(x)
        x = mg.nnet.sigmoid(self.dense2(x))
        return x
    @property
    def parameters(self):
        return self.dense1.parameters + self.dense2.parameters + self.bn.parameters 

model = Net()
opt = Optimizers(model.parameters)
criterion = Loss()

# xor problem
X = np.array([[0,0],
              [1,0],
              [0,1],
              [1,1]])

y = np.array([[0],
              [1],
              [1],
              [0]])

for i in range(500):
    out = model(X)
    loss = criterion.CrossEntropy(y,out)
    loss.backward()
    opt.Momentum()
    loss.null_gradients()
    if i%10==0:
        #print(np.where(model(X)>0.5,1,0))
        print(model(X))
        print("Loss",loss.data.tolist())
        print("\n")
