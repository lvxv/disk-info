package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/lvxv/disk-info/disk"
	"github.com/lvxv/disk-info/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"text/tabwriter"
)

const (
	flagFolder = "folder"
)


func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disk-info",
		Short: "show disk-info",
		RunE:  runInfoCmd,
	}

	cmd.Flags().StringP(flagFolder, "f", "", "show usage of the folder")
	return cmd
}

func runInfoCmd(cmd *cobra.Command, args []string) error {
	dst := viper.GetString(flagFolder)
	if dst == "" {
		return cmd.Help()
	}

	di, err := disk.GetInfo(dst)
	if err != nil {
		return fmt.Errorf("get disk info with err %v", err)
	}

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintln(w,"Type\tDisk\tTotal\tUsed\tFree\tRoot Reserve\tLamb Reserve")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"df",
		dst,
		humanize.IBytes(di.Total + di.Reserved),
		humanize.IBytes(di.Total-di.Free),
		humanize.IBytes(di.Free),
		humanize.IBytes(di.Reserved),
		humanize.IBytes(0),
	)

	// exclude preserved space for every device
	// here we just ignore if disk is under same device
	free := uint64(0)
	_free := int64(float64(di.Free)*0.95 - utils.DiskMinFreeSpace)
	if _free > 0 {
		free = uint64(_free)
	}
	reserved := uint64(float64(di.Free)*0.05 + utils.DiskMinFreeSpace)

	fmt.Fprintf(w,"%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"lamb",
		dst,
		humanize.IBytes(di.Total),
		humanize.IBytes(di.Total-free),
		humanize.IBytes(free),
		humanize.IBytes(di.Reserved),
		humanize.IBytes(reserved),
	)

	w.Flush()
	return nil
}

func preRun(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		return nil
	}
	return cmd
}

func main() {
	rootCmd := RootCmd()
	rootCmd.SilenceUsage = true
	_ = preRun(rootCmd).Execute()

}
