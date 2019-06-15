package console

import (
	"bufio"
	"fmt"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"os"
	"strings"
)

type Client struct {
	character      *app.Character
	commander      *commands.Commander
	roomRepository app.RoomRepository
}

func (Client) Create() *Client {
	roomRepository := app.RoomMemoryRepository{}.Create(nil)
	character := app.Character{}.Create("console")
	pick := app.Item{}.Create()
	pick.AddFlag(itemFlag.MineTool)
	character.AddItem(pick)
	fishingPole := app.Item{}.Create()
	fishingPole.AddFlag(itemFlag.FishTool)
	character.AddItem(fishingPole)
	randomizer := random.Random{}.Create()
	commander := commands.Commander{}.Create(roomRepository, randomizer)

	return &Client{character: character, commander: commander, roomRepository: roomRepository}
}

func (client *Client) Start() {
	client.showEmptyLine()
	client.outputNewline("Game started")
	client.showEmptyLine()
	client.ShowPrompt()
	client.showEmptyLine()
	client.outputNewline("Enter command:")
	client.showEmptyLine()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		client.showEmptyLine()

		input := scanner.Text()
		if input != "" {
			client.ExecuteCommand(input)
		}

		client.ShowPrompt()
		client.showEmptyLine()
	}

	if scanner.Err() != nil {
		client.outputNewline("Error occurred!")
	}
}

func (client *Client) ExecuteCommand(input string) {
	commandWithArguments := strings.Fields(input)
	_, errors := client.commander.Execute(client.character, commandWithArguments)

	for err := range errors {
		client.outputNewline("Error: " + err.Error() + "\n")
	}
}

func (client *Client) ShowPrompt() {
	room := client.roomRepository.FindByXYandZ(
		client.character.X(),
		client.character.Y(),
		client.character.Z(),
	)

	client.outputInline(fmt.Sprintf("%d/100 HP | ", client.character.Health()))

	var biom string
	if room != nil {
		biom = room.Biom().String()
	} else {
		biom = "N/A"
	}

	client.outputInline(
		fmt.Sprintf(
			"%s [%d/%d/%d] | ",
			biom,
			client.character.X(),
			client.character.Y(),
			client.character.Z(),
		),
	)

	var roomFlags []string
	for roomFlag := range room.Flags() {
		roomFlags = append(roomFlags, roomFlag.String())
	}

	client.outputNewline("Room flags: " + strings.Join(roomFlags, ", "))

	client.outputNewline("Items:")
	for _, item := range client.character.Inventory() {
		client.outputNewline(item)
	}

	var exits []string
	directions := []direction.Direction{
		direction.North,
		direction.South,
		direction.East,
		direction.West,
		direction.Up,
		direction.Down,
	}

	for _, exitDirection := range directions {
		x, y, z := exitDirection.DiffXYZ()
		exitRoom := client.roomRepository.FindByXYandZ(
			client.character.X()+x,
			client.character.Y()+y,
			client.character.Z()+z,
		)
		if exitRoom != nil {
			exits = append(exits, exitDirection.String())
		}
	}

	client.outputNewline("Exits: " + strings.Join(exits, ", "))
}

func (client *Client) showEmptyLine() {
	client.outputNewline("----------------")
}

func (client *Client) outputNewline(message interface{}) {
	client.output(message, true)
}

func (client *Client) outputInline(message interface{}) {
	client.output(message, false)
}

func (client *Client) output(message interface{}, isNewLine bool) {
	newLine := ""
	if isNewLine {
		newLine = "\n"
	}

	fmt.Printf("%v%s", message, newLine)
}
