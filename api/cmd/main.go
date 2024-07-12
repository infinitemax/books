package main

import (
	"context"
	"fmt"
	"github.com/infinitemax/books/internal/server"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hey dude! This program is staring...")

	// start the server
	rootCmd := RootCmd()

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RootCmd() *cobra.Command {
	fmt.Println("rootcmd doing its thing...")

	cmd := &cobra.Command{
		Use:   "books-api",
		Short: "This is the root command for the application",
	}

	cmd.AddCommand(ServeCmd())
	return cmd
}

func ServeCmd() *cobra.Command {
	fmt.Println("Servecmd doing its thing...")
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			srv := server.NewServer()
			if err := srv.Init(); err != nil {
				return fmt.Errorf("failer to initialise server: %w", err)
			}
			return srv.Run(ctx)
		},
	}
}
