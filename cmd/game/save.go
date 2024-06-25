package game

import (
	"fmt"

	"github.com/DMXMax/cli-test/util/db"
	gdb "github.com/DMXMax/cli-test/util/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save a game to a file",
	Long:  `Save a game to a given file name. Currently it will overwrite any existing file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gdb.Current == nil {
			return fmt.Errorf("no game selected")
		}
		g := gdb.Current
		/*//filename := g.Name + ".gob"
		//file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		//defer file.Close()
		//enc := gob.NewEncoder(file)
		//err = enc.Encode(g)
		if err != nil {
			return err
		}
		fmt.Printf("Game saved to %s\n", filename)*/

		// lets do all the db here:
		/*DB, err := gorm.Open(sqlite.Open("data/games.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		err = DB.AutoMigrate(&gdb.Game{})
		if err != nil {
			panic("failed to migrate Game")
		}*/
		if g.ID == 0 {
			db.GamesDB.Create(g)
		} else {
			db.GamesDB.Save(g)
		}
		//fmt.Printf("Game saved as %d", g.ID)

		return nil
	},
}

func init() {

}
