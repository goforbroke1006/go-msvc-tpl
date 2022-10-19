package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

func NewDaemonPeriodicPrintCmd() *cobra.Command {
	return &cobra.Command{
		Use: "periodic-print",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()

			go func() {
			RunLoop:
				for {
					select {
					case <-ctx.Done():
						break RunLoop
					case <-time.After(5 * time.Second):
						fmt.Println("hello...")
					}
				}
			}()

			<-ctx.Done()
		},
	}
}

func init() {
	daemonCmd.AddCommand(NewDaemonPeriodicPrintCmd())
}
