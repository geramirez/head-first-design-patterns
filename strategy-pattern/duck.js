function Quack() {
  this.execute = function() {
    console.log("Quack");
  }
}

function Squeak() {
  this.execute = function() {
    console.log("Squeak");
  }
}

function MuteQuack() {
  this.execute = function() {
    console.log("<< Silence >>");
  }
}

var FlyWithWings = {
  execute: function() {
    console.log("I'm flying");
  },
}


var FlyNoWay = {
  execute: function() {
    console.log("I can't fly");
  },
}

function Duck() {}

Duck.prototype.display = function() {
  console.log("I don't know what type of duck I am");
}

Duck.prototype.swim = function() {
  console.log("All ducks swim");
}

Duck.prototype.quack = function() {
  this.quackBehavior.execute();
}

Duck.prototype.fly = function() {
  this.flyBehavior.execute();
}

function RubberDuck() {
  this.display = function() {
    console.log("I'm a rubber duck");
  }
  this.quackBehavior = new Squeak();
  this.flyBehavior = FlyNoWay;
}
RubberDuck.prototype = Object.create(Duck.prototype);
RubberDuck.prototype.constructor = RubberDuck;


function MallardDuck() {
  this.display = function() {
    console.log("I'm a mallard duck");
  }
  this.quackBehavior = new Quack();
  this.flyBehavior = FlyWithWings;
}
MallardDuck.prototype = Object.create(Duck.prototype);
MallardDuck.prototype.constructor = RubberDuck;


function ModelDuck() {
  this.display = function() {
    console.log("I'm a model duck");
  }
  this.quackBehavior = new MuteQuack();
  this.flyBehavior = FlyNoWay;
}
ModelDuck.prototype = Object.create(Duck.prototype);
ModelDuck.prototype.constructor = RubberDuck;



ducks = [new RubberDuck(), new MallardDuck(), new ModelDuck()]

ducks.forEach(function(duck) {
  console.log("------------------");
  duck.display();
  duck.swim();
  duck.quack();
  duck.fly();
})

