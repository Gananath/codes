import numpy as np
import mygrad as mg


def Uniform(*shape, lower_bound=0, upper_bound=1):
    mg.Tensor(np.random.uniform(lower_bound,upper_bound,shape))

def He_normal(*shape,gain=1):
    if len(shape) < 2:
        raise ValueError("He Normal initialization requires at least two dimensions")
    
    tensor = np.empty(shape)
    std = gain / np.sqrt(shape[1] * tensor[0, 0].size)
    return Tensor(np.random.normal(0, std, shape))
