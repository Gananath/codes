import mygrad as mg


class Loss:
    '''    https://towardsdatascience.com/common-loss-functions-in-machine-learning-46af0ffc4d23
    '''
    def __init__(self):
        pass
    def SquaredError(self,y_real,y_pred):
        return (y_real-y_pred)**2
    def AbsoluteError(self,y_real,y_pred):
        return mg.absolute(y_real - y_pred)
    def HuberLoss(self,y_real,y_pred,delta=1):
        if mg.absolute(y_real - y_pred)<=delta:
            return 0.5*(y_real - y_pred)**2
        else:
            return delta*mg.absolute(y_real - y_pred)-0.5*delta**2
    def MeanSquaredError(self,y_real,y_pred):
        # Learn outlier
        return ((y_real - y_pred)**2).mean()
    def MeanAbsoluteError(self,y_real,y_pred):
        # Ignores outlier
        return (mg.absolute(y_real - y_pred)).mean()
    def MeanBiasError(self,y_real,y_pred):
        return (y_real - y_pred).mean()
    def CrossEntropy(self,y_real,y_pred,eps=1e-10):
        y_pred = mg.clip(y_pred, eps, 1. - eps)
        N = y_pred.shape[0]
        return -mg.sum(y_real*mg.log(y_pred+1e-9))/N
