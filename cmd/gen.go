/*
Copyright © 2023 Aino L. Spring <info@aino-spring.com>
*/
package cmd

import (
	"fmt"
	"mazeCli/maze"
	"time"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates the maze and outputs it into a txt file.",
	Long: `It generates the maze for the given arguments and saves it
to the specified output file.`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("size")
		seed, _ := cmd.Flags().GetInt64("seed")
		out, _ := cmd.Flags().GetString("out")

		fmt.Println("Generating maze...")

		var mazeObj = maze.NewMaze(maze.NewVector(2).Fill(float64(size)), seed)

		var display = maze.DisplayMaze(mazeObj)
		display.SetValue(mazeObj.Start, 'S')
		display.SetValue(mazeObj.Finish, 'F')
		display.Save(out)

		fmt.Printf("Saved maze at %v.\n", out)

	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	genCmd.Flags().IntP("size", "z", 10, "Sets the size of the maze.")
	genCmd.Flags().Int64P("seed", "s", int64(time.Now().Nanosecond()), "Sets the seed of the maze.")
	genCmd.Flags().StringP("out", "o", "maze.txt", "Sets the output file.")

}
