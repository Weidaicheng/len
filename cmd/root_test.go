package cmd

import (
	"testing"

	"github.com/Weidaicheng/len/pkg/version"
	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	wantVersion := version.Version
	version := rootCmd.Version

	if wantVersion != version {
		t.Fatalf(`rootCmd.Version = %v, want equals to %v`, version, wantVersion)
	}

	wantSilenceErrors := true
	silenceErrors := rootCmd.SilenceErrors
	if wantSilenceErrors != silenceErrors {
		t.Fatalf(`rootCmd.SilenceErrors = %v, want to be %v`, silenceErrors, wantSilenceErrors)
	}

	wantDisableFlagsInUseLine := true
	disableFlagsInUseLine := rootCmd.DisableFlagsInUseLine
	if wantDisableFlagsInUseLine != disableFlagsInUseLine {
		t.Fatalf(`rootCmd.DisableFlagsInUseLine = %v, want to be %v`, disableFlagsInUseLine, wantDisableFlagsInUseLine)
	}

}

func TestRootCmdArgs(t *testing.T) {
	wantArgs := cobra.MinimumNArgs(1)
	args := rootCmd.Args

	arrWithLength1 := []string{
		"abc",
	}
	if args(nil, arrWithLength1) != nil || wantArgs(nil, arrWithLength1) != nil {
		t.Fatal("rootCmd.Args is not completable with cobra.MinimumNArgs(1) for 1 arg")
	}

	arrWithLength0 := []string{}
	if args(nil, arrWithLength0) == nil || wantArgs(nil, arrWithLength0) == nil {
		t.Fatal("rootCmd.Args is not completable with cobra.MinimumNArgs(1) for 0 arg")
	}
}
