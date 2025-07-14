package reboot

import (
	"fmt"

	"github.com/Dorrrke/spacectl/internal/reboot"
	"github.com/spf13/cobra"
)

func rebootAgentLauncher() *cobra.Command {
	return &cobra.Command{
		Use:   "agent",
		Short: "Перезапустить Space Agent PC",
		Run: func(cmd *cobra.Command, args []string) {
			if err := reboot.RebootGlintLauncher(); err != nil {
				fmt.Println("Ошибка перезапуска Space Agent PC:", err)
			} else {
				fmt.Println("Space Agent PC перезапущен")
			}
		},
	}
}
