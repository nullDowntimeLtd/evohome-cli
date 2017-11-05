package main

import (
    "strings"
    "github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest {
    { Text: "cancel", Description: "Cancel configuration" },
    { Text: "set", Description: "Change configuration" },
    { Text: "show", Description: "Retrieve configuration/status" },
    { Text: "exit", Description: "Exit application" },
}

func authCompleter(d prompt.Document) []prompt.Suggest {
    return []prompt.Suggest {}
}

func mainCompleter(d prompt.Document) []prompt.Suggest {
    args := strings.Split(d.TextBeforeCursor(), " ")
    return argumentsCompleter(args)
}

func argumentsCompleter(args []string) []prompt.Suggest {
    if len(args) <= 1 {
        return prompt.FilterHasPrefix(commands, args[0], true)
    }

    if !clientInitialized() {
        error("Evohome client not initialized")
        return []prompt.Suggest {}
    }
    t := client.TemperatureControlSystem()
    first := args[0]
    second := args[1]
    switch first {
        case "cancel":
            if len(args) == 2 {
                subCommands := []prompt.Suggest {
                    { Text: "zone", Description: "Cancel zone configuration" },
                }
                return prompt.FilterHasPrefix(subCommands, second, true)
            }

            third := args[2]
            if len(args) == 3 {
                var subCommands []prompt.Suggest
                switch second {
                    case "zone":
                        subCommands = listToSubCommands(t.ZoneNamesWithOverride())
                }
                return prompt.FilterHasPrefix(subCommands, third, true)
            }

            fourth := args[3]
            if len(args) == 4 {
                subCommands := []prompt.Suggest {
                    { Text: "temperature", Description: "Cancel zone temperature configuration" },
                }
                return prompt.FilterHasPrefix(subCommands, fourth, true)
            }

            fifth := args[4]
            if len(args) == 5 {
                subCommands := []prompt.Suggest {
                    { Text: "override", Description: "Cancel zone temperature override" },
                }
                return prompt.FilterHasPrefix(subCommands, fifth, true)
            }
        case "set":
            if len(args) == 2 {
                subCommands := []prompt.Suggest {
                    { Text: "zone", Description: "Change zone configuration" },
                }
                return prompt.FilterHasPrefix(subCommands, second, true)
            }

            third := args[2]
            if len(args) == 3 {
                var subCommands []prompt.Suggest
                switch second {
                    case "zone":
                        subCommands = listToSubCommands(t.ZoneNames())
                }
                return prompt.FilterHasPrefix(subCommands, third, true)
            }

            fourth := args[3]
            if len(args) == 4 {
                subCommands := []prompt.Suggest {
                    { Text: "temperature", Description: "Change zone temperature" },
                }
                return prompt.FilterHasPrefix(subCommands, fourth, true)
            }

            if len(args) == 6 {
                sixth := args[5]
                subCommands := []prompt.Suggest {
                    { Text: "until", Description: "Change zone temperature until point in time" },
                }
                return prompt.FilterHasPrefix(subCommands, sixth, true)
            }

            if len(args) == 7 {
                sixth := args[6]
                subCommands := []prompt.Suggest {
                    { Text: "yyyy/mm/dd hh:mm" },
                }
                return prompt.FilterHasPrefix(subCommands, sixth, true)
            }
        case "show":
            if len(args) == 2 {
                subCommands := []prompt.Suggest {
                    { Text: "zone", Description: "Show zone state" },
                    { Text: "zones", Description: "List all zones" },
                }
                return prompt.FilterHasPrefix(subCommands, second, true)
            }

            third := args[2]
            if len(args) == 3 {
                switch second {
                    case "zone":
                        subCommands := listToSubCommands(t.ZoneNames())
                        return prompt.FilterHasPrefix(subCommands, third, true)
                }
            }
    }

    return []prompt.Suggest {}
}

func listToSubCommands(list []string) ([]prompt.Suggest) {
    subCommands := []prompt.Suggest {}
    for _, zone := range list {
        subCommands = append(subCommands, prompt.Suggest{ Text: zone })
    }
    return subCommands
}
