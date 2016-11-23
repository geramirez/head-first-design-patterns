class Quack:
    def execute(self):
        print("Quack")


class Squeak:
    def execute(self):
        print("Squeak")


class MuteQuack:
    def execute(self):
        print("<< Silence >>")


class FlyWithWings:
    def execute(self):
        print("I'm flying")


class FlyNoWay:
    def execute(self):
        print("I can't fly")


class Duck:
    def __init__(self, quackBehavior=None, flyBehavior=None):
        """ Making __init__ to take both the quackBehavior and
            the flyBehavior allows us to create new ducks on the fly """ 
        self.quackBehavior = quackBehavior()
        self.flyBehavior = flyBehavior()

    def display(self):
        print("I don't know what type of duck I am")

    def swim(self):
        print("All ducks swim")

    def quack(self):
        self.quackBehavior.execute()

    def fly(self):
        self.flyBehavior.execute()


class Rubber(Duck):
    def __init__(self):
        self.quackBehavior = Squeak()
        self.flyBehavior = FlyNoWay()

    def display(self):
        print("I am a rubber duck")


class Mallard(Duck):
    def __init__(self):
        self.quackBehavior = Quack()
        self.flyBehavior = FlyWithWings() 

    def display(self):
        print("I am a mallard duck")


if __name__ == "__main__":

    ducks = [
        Rubber(),
        Mallard(),
        Duck(quackBehavior=MuteQuack, flyBehavior=FlyNoWay)
    ]
    
    for idx, duck in enumerate(ducks):
        print("This is duck #%s" % idx)
        duck.display()
        duck.swim()
        duck.quack()
        duck.fly()
        print("_______")
