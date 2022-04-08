package auths

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"zeus/pkg/api/dao"
	"zeus/pkg/api/domain/perm"
	"zeus/pkg/api/dto"
	"zeus/pkg/api/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	datacfg string
	config  string
	AuthCmd = &cobra.Command{
		Use:     "authcommand",
		Short:   "import batch commands from json.file",
		Example: "zeus authcommand -a config/auth.json",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	AuthCmd.PersistentFlags().StringVarP(&config, "config", "c", "./config/in-local.yaml", "Start server with provided configuration file")
	AuthCmd.PersistentFlags().StringVarP(&datacfg, "datacfg", "d", "./data/auth.json", "start server with provided configuration file")

}

func usage() {
	fmt.Println("insert auth command")
}

func setup() {
	fmt.Println("setup")

	zerolog.SetGlobalLevel(zerolog.Level(0))

	viper.SetConfigFile(config)
	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	//3.Set up run mode
	mode := viper.GetString("mode")
	gin.SetMode(mode)

	dao.Setup()
	perm.SetUp(false)
}

func run() error {
	command_list, err := importAuth()

	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	for _, v := range command_list {

		menuModel := model.Menu{
			Name:     v.Name,
			ParentId: v.ParentId,
			DomainId: v.DomainId,
			Url:      v.Url,
			Perms:    v.Perms,
			Alias:    v.Alias,
			MenuType: v.MenuType,
			Icon:     v.Icon,
			OrderNum: 1,
		}

		db := dao.GetDb()
		var m model.Menu

		result := db.Table("menu").Where("domain_id = ? and name = ? and perms = ? and parent_id = ?", v.DomainId, v.Name, v.Perms, v.ParentId).First(&m)

		if result.RowsAffected > 0 {
			continue
		}

		db.Save(&menuModel)

	}

	return nil
}

func importAuth() ([]*dto.MenuCreateDto, error) {
	data, err := ioutil.ReadFile(datacfg)
	if err != nil {
		return nil, err
	}

	var auth_list []*dto.MenuCreateDto
	err = json.Unmarshal(data, &auth_list)

	if err != nil {
		return nil, err
	}

	return auth_list, nil

}
