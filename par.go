package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/sync/errgroup"
)

const (
	shell = "/bin/bash"
)

var (
	dryRun  = flag.Bool("n", false, "Dry run - print commands instead of executing them")
	verbose = flag.Bool("v", false, "Verbose - print commands as they are executed")
)

func runCommand(command string) error {
	cmd := exec.Command(shell, "-c", command)
	if cmd.Err != nil {
		return cmd.Err
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if *dryRun || *verbose {
		fmt.Fprintf(os.Stderr, "%s\n", cmd.String())
	}

	if *dryRun {
		return nil
	}

	return cmd.Run()
}

func run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(runtime.NumCPU())

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		group.Go(func() error { return runCommand(line) })
	}

	if err := group.Wait(); err != nil {
		return err
	}

	return scanner.Err()
}

func main() {
	flag.Parse()

	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
