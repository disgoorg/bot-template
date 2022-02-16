package test

import (
	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/discord"
	"github.com/YourName/YourRepo/internal/types"
)

var (
	_ types.Module         = (*Module)(nil)
	_ types.CommandsModule = (*Module)(nil)
	_ types.ListenerModule = (*Module)(nil)
)

type Module struct{}

func (Module) Commands() []types.Command {
	return []types.Command{
		{
			Create: discord.SlashCommandCreate{
				Name:              "test",
				Description:       "Test command",
				DefaultPermission: true,
			},
			Handler: testHandler,
		},
	}
}

func (Module) OnEvent(b *types.Bot, event core.Event) {

}