class Quack {
  execute(){
    console.log("Quack");
  }
}

class Squeak {
  execute(){
    console.log("Squeak");
  }
}

class MuteQuack {
  execute(){
    console.log("<< Silence >>");
  }
}

class FlyWithWings {
  execute(){
    console.log("I'm flying");
  }
}


class FlyNoWay {
  execute(){
    console.log("I can't fly");
  }
}

class Duck {
  display() {
    console.log("I don't know what type of duck I am");
  }

  swim() {
    console.log("All ducks swim");
  }

  quack() {
    this.quackBehavior.execute();
  }

  fly() {
    this.flyBehavior.execute();
  }
}

class RubberDuck extends Duck {
  constructor() {
    super();
    this.quackBehavior = new Squeak();
    this.flyBehavior = new FlyNoWay();
  }

  display() {
    console.log("I'm a rubber duck");
  }
}


class MallardDuck extends Duck {
  constructor() {
    super();
    this.quackBehavior = new Quack();
    this.flyBehavior = new FlyWithWings();
  }

  display() {
    console.log("I'm a mallard duck");
  }
}


class ModelDuck extends Duck {
  constructor() {
    super();
    this.quackBehavior = new MuteQuack();
    this.flyBehavior = new FlyNoWay();
  }
  display() {
    console.log("I'm a model duck");
  }
}

ducks = [new RubberDuck(), new MallardDuck(), new ModelDuck()]

ducks.forEach(function(duck) {
  console.log("------------------");
  duck.display();
  duck.swim();
  duck.quack();
  duck.fly();
})


