package console

import (
	"bufio"
	"fmt"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
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

	scanner := bufio.NewScanner(os.Stdin)
	client.output("Enter command:")

	for scanner.Scan() {
		client.outputEmptyLine()
		client.ExecuteCommand(scanner.Text())
		client.ShowPrompt()
		client.outputEmptyLine()
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
			"Biom: %s Coordinates: [%d/%d/%d]",
			biom,
			client.character.X(),
			client.character.Y(),
			client.character.Z(),
		),
	)

	client.output("Room flags:")
	for roomFlag := range room.Flags() {
		client.output(roomFlag.String())
	}

	client.output("Items:")
	for _, item := range client.character.Inventory() {
		client.output(item)
	}
}

func (client *Client) output(message interface{}) {
	fmt.Printf("%v\n", message)
}

func (client *Client) outputEmptyLine() {
	fmt.Printf("----------------\n")
}
