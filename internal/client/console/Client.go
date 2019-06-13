package console

import (
	"bufio"
	"fmt"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
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
	randomizer := random.Random{}.Create()
	commander := commands.Commander{}.Create(roomRepository, randomizer)

	return &Client{character: character, commander: commander, roomRepository: roomRepository}
}

func (client *Client) Start() {
	client.output("Game started")
	client.ShowPrompt()
	client.showEmptyLine()
	client.output("Enter command:")
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
		client.output("Error occurred!")
	}
}

func (client *Client) ExecuteCommand(input string) {
	commandWithArguments := strings.Fields(input)
	command, err := client.commander.Command(commandWithArguments[0])

	errors := make(map[gameError.Error]bool)
	if err == "" {
		commandResult := command.Execute(client.character, strings.Join(commandWithArguments[1:], " "))

		if commandResult.HasErrors() {
			for err := range commandResult.Errors() {
				errors[err] = true
			}
		}
	} else {
		errors[err] = true
	}

	for err := range errors {
		client.output("Error: " + err.Error() + "\n")
	}

}

func (client *Client) ShowPrompt() {
	room := client.roomRepository.FindByXYandZ(
		client.character.X(),
		client.character.Y(),
		client.character.Z(),
	)

	var biom string
	if room != nil {
		biom = room.Biom().String()
	} else {
		biom = "N/A"
	}

	client.output(
		fmt.Sprintf(
			"Biom: %s [%d/%d/%d]",
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

	client.output("Room flags: " + strings.Join(roomFlags, ", "))

	client.output("Items:")
	for _, item := range client.character.Inventory() {
		client.output(item)
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

	client.output("Exits: " + strings.Join(exits, ", "))
}

func (client *Client) output(message interface{}) {
	fmt.Printf("%v\n", message)
}

func (client *Client) showEmptyLine() {
	fmt.Printf("----------------\n")
}
