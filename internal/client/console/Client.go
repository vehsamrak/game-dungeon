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
	"sort"
	"strings"
	"time"
)

type Client struct {
	character      *app.Character
	commander      *commands.Commander
	roomRepository app.RoomRepository
	tickDuration   time.Duration
}

func (Client) Create() *Client {
	character := app.Character{}.Create("console")
	pick := app.Item{}.Create("miner pick")
	pick.AddFlag(itemFlag.MineTool)
	fishingPole := app.Item{}.Create("fishing pole")
	fishingPole.AddFlag(itemFlag.FishTool)
	axe := app.Item{}.Create("forester axe")
	axe.AddFlag(itemFlag.CutTreeTool)

	character.AddItems([]*app.Item{pick, fishingPole, axe})

	roomRepository := app.RoomMemoryRepository{}.Create(nil)
	randomizer := random.Randomizer{}.Create()
	commander := commands.Commander{}.Create(roomRepository, randomizer)

	return &Client{
		character:      character,
		commander:      commander,
		roomRepository: roomRepository,
		tickDuration:   5 * time.Second,
	}
}

func (client *Client) Start() {
	client.showEmptyLine()
	client.outputNewline("Game started")
	client.showEmptyLine()

	client.showPrompt()
	client.showEmptyLine()

	client.outputNewline("Enter command:")
	client.showEmptyLine()

	client.startTicker()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		client.showEmptyLine()

		input := scanner.Text()
		if input != "" {
			switch input {
			case "inventory":
				client.outputNewline("Items:")

				items := make(map[string]int)
				for _, item := range client.character.Inventory() {
					itemName := item.Name()

					_, ok := items[itemName]
					if ok {
						items[itemName] += 1
					} else {
						items[itemName] = 1
					}
				}

				// sort inventory
				var itemNames []string
				for name := range items {
					itemNames = append(itemNames, name)
				}
				sort.Strings(itemNames)

				for _, itemName := range itemNames {
					client.outputNewline(fmt.Sprintf("%s: %d", itemName, items[itemName]))
				}

				client.showEmptyLine()
			default:
				client.executeCommand(input)
			}
		}

		client.showPrompt()
		client.showEmptyLine()
	}

	if scanner.Err() != nil {
		client.outputNewline("Error occurred!")
	}
}

func (client *Client) executeCommand(input string) {
	commandWithArguments := strings.Fields(input)
	_, errors := client.commander.Execute(client.character, commandWithArguments)

	for err := range errors {
		client.outputNewline("Error: " + err.Error() + "\n")
	}
}

func (client *Client) showPrompt() {
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

	var exitsString string
	if len(exits) > 0 {
		exitsString = strings.Join(exits, ", ")
	} else {
		exitsString = "not explored. (type \"explore [north/south/east/west]\")"
	}

	client.outputNewline("Exits: " + exitsString)
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
func (client *Client) startTicker() {
	ticker := time.NewTicker(client.tickDuration)

	go func() {
		restCommand := commands.RestCommand{}.Create()

		for range ticker.C {
			commandResult := restCommand.Execute(client.character)

			if !commandResult.HasErrors() {
				client.outputNewline("HP increased")
				client.showEmptyLine()
				client.showPrompt()
				client.showEmptyLine()
			}
		}
	}()
}
