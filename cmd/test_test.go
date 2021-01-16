package cmd

import "testing"

func TestTest(t *testing.T) {

	rootCmd.SetArgs([]string{"failed"})
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	if err := rootCmd.Execute(); err != nil {
		t.Errorf(err.Error())
	}

}
