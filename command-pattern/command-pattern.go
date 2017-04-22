package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type Light struct {
	isOn bool
}

func (l Light) On() {
	l.isOn = true
	fmt.Println("Light is on")
}

func (l Light) Off() {
	l.isOn = false
	fmt.Println("Light is off")
}

type GarageDoor struct {
	isOpen bool
}

func (g GarageDoor) Open() {
	g.isOpen = true
	fmt.Println("Garage Door is Open")
}

func (g GarageDoor) Close() {
	g.isOpen = false
	fmt.Println("Garage Door is Closed")
}

type NoCommand struct {
}

func (loc NoCommand) Execute() {
	fmt.Println("No Command")
}

func (loc NoCommand) Undo() {
	fmt.Println("No Command")
}

type LightOnCommand struct {
	light Light
}

func (loc LightOnCommand) Execute() {
	loc.light.On()
}

func (loc LightOnCommand) Undo() {
	loc.light.Off()
}

type LightOffCommand struct {
	light Light
}

func (loc LightOffCommand) Execute() {
	loc.light.Off()
}

func (loc LightOffCommand) Undo() {
	loc.light.On()
}

type GarageDoorOpenCommand struct {
	door GarageDoor
}

func (loc GarageDoorOpenCommand) Execute() {
	loc.door.Open()
}

func (loc GarageDoorOpenCommand) Undo() {
	loc.door.Close()
}

type GarageDoorCloseCommand struct {
	door GarageDoor
}

func (loc GarageDoorCloseCommand) Execute() {
	loc.door.Close()
}

func (loc GarageDoorCloseCommand) Undo() {
	loc.door.Open()
}

type SimpleRemoteControl struct {
	offCommands [2]Command
	onCommands  [2]Command
	undoCommand Command
}

func (src *SimpleRemoteControl) SetOnCommand(slot int, command Command) {
	src.onCommands[slot] = command
}

func (src *SimpleRemoteControl) SetOffCommand(slot int, command Command) {
	src.offCommands[slot] = command
}

func (src *SimpleRemoteControl) SetUndoCommand(command Command) {
	src.undoCommand = command
}

func (src *SimpleRemoteControl) PressOnButton(slot int) {
	cmd := src.onCommands[slot]
	cmd.Execute()
	src.undoCommand = cmd
}

func (src *SimpleRemoteControl) PressOffButton(slot int) {
	cmd := src.offCommands[slot]
	cmd.Execute()
	src.undoCommand = cmd
}

func (src *SimpleRemoteControl) PressUndoButton() {
	src.undoCommand.Undo()
}

func main() {
	remote := &SimpleRemoteControl{}
	garageDoor := GarageDoor{}
	light := Light{}

	lightOn := LightOnCommand{light}
	lightOff := LightOffCommand{light}
	garageDoorOpen := GarageDoorOpenCommand{garageDoor}
	garageDoorClose := GarageDoorCloseCommand{garageDoor}
	noCommand := NoCommand{}

	remote.SetOnCommand(0, lightOn)
	remote.SetOffCommand(0, lightOff)
	remote.SetOnCommand(1, garageDoorOpen)
	remote.SetOffCommand(1, garageDoorClose)
	remote.SetUndoCommand(noCommand)

	remote.PressUndoButton()
	remote.PressOnButton(0)
	remote.PressUndoButton()
	remote.PressOnButton(1)
	remote.PressUndoButton()
	remote.PressOffButton(0)
	remote.PressUndoButton()
	remote.PressOffButton(1)
	remote.PressUndoButton()

}
