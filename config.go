package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var configFile = getCurrentPath() + "/config.json"

type MajsoulExInfoDialogAttr struct {
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Alpha    float64 `json:"alpha"`
	FontSize float64 `json:"font_size"`
}

type MajsoulExRGBA struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

type MajsoulExChongColorAttr struct {
	Normal               MajsoulExRGBA `json:"normal"`
	LessThan5Percent     MajsoulExRGBA `json:"less_than_5_percent"`
	LessThan10Percent    MajsoulExRGBA `json:"less_than_10_percent"`
	LessThan15Percent    MajsoulExRGBA `json:"less_than_15_percent"`
	GreaterThan15Percent MajsoulExRGBA `json:"greater_than_15_percent"`
}

type MajsoulExInGamePromptAttr struct {
	MultipleInfo       bool                    `json:"multiple_info"`
	MultipleInfoAttr   MajsoulExInfoDialogAttr `json:"multiple_info_attr"`
	MouseEnterInfo     bool                    `json:"mouse_enter_info"`
	MouseEnterInfoAttr MajsoulExInfoDialogAttr `json:"mouse_enter_info_attr"`
	ChongTileColor     bool                    `json:"chong_tile_color"`
	ChongTileColorAttr MajsoulExChongColorAttr `json:"chong_tile_color_attr"`
	MoqieColor         bool                    `json:"moqie_color"`
}

type MajsoulExProxyAttr struct {
	IsUseProxy bool   `json:"is_use_proxy"`
	ProxyType  string `json:"proxy_type"`
	ProxyAddr  string `json:"proxy_addr"`
	ProxyPort  int    `json:"proxy_port"`
}

type GameConfig struct {
	MajsoulAccountIDs []int `json:"majsoul_account_ids"`

	CurrentActiveMajsoulAccountID int    `json:"-"`
	CurrentActiveTenhouUsername   string `json:"-"`

	MajsoulExDeploy       bool                      `json:"majsoulex_deploy"`
	MajsoulExServer       int                       `json:"majsoulex_server"`
	MajsoulExAutoDiscard  bool                      `json:"majsoulex_auto_discard"`
	MajsoulExAutoConfirm  bool                      `json:"majsoulex_auto_confirm"`
	MajsoulExFallback     bool                      `json:"majsoulex_fallbak"`
	MajsoulExImprove      bool                      `json:"majsoulex_improve"`
	MajsoulExBugQDY       bool                      `json:"majsoulex_bugqdy"`
	MajsoulExChiitoi02    bool                      `json:"majsoulex_chiitoi_02"`
	MajsoulExInGamePrompt MajsoulExInGamePromptAttr `json:"majsoulex_in_game_prompt"`
	MajsoulExProxy        MajsoulExProxyAttr        `json:"majsoulex_proxy"`
}

var GameConf = &GameConfig{
	MajsoulAccountIDs:             []int{},
	CurrentActiveMajsoulAccountID: -1,
	MajsoulExInGamePrompt: MajsoulExInGamePromptAttr{
		MultipleInfo: false,
		MultipleInfoAttr: MajsoulExInfoDialogAttr{
			Width:    1250.0,
			Height:   730.0,
			X:        200.0,
			Y:        200.0,
			Alpha:    0.8,
			FontSize: 30.0,
		},
		MouseEnterInfo: false,
		MouseEnterInfoAttr: MajsoulExInfoDialogAttr{
			Width:    1250.0,
			Height:   210.0,
			X:        200.0,
			Y:        550.0,
			Alpha:    0.8,
			FontSize: 50.0,
		},
		ChongTileColor: false,
		ChongTileColorAttr: MajsoulExChongColorAttr{
			Normal:               MajsoulExRGBA{R: 255 / 255, G: 255 / 255, B: 255 / 255, A: 1.0},
			LessThan5Percent:     MajsoulExRGBA{R: 0 / 255, G: 255 / 255, B: 255 / 255, A: 1.0},
			LessThan10Percent:    MajsoulExRGBA{R: 255 / 255, G: 255 / 255, B: 0, A: 1.0},
			LessThan15Percent:    MajsoulExRGBA{R: 255 / 255, G: 127.0 / 255.0, B: 80.0 / 255.0, A: 1.0},
			GreaterThan15Percent: MajsoulExRGBA{R: 255 / 255, G: 199.0 / 255.0, B: 199.0 / 255.0, A: 1.0},
		},
		MoqieColor: false,
	},
	MajsoulExProxy: MajsoulExProxyAttr{
		IsUseProxy: false,
		ProxyType:  "http",
		ProxyAddr:  "127.0.0.1",
		ProxyPort:  0,
	},
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}

func init() {
	ReadNewConfig()
}

func ReadNewConfig() {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		// if debugMode {
		// 	panic(err)
		// }
		return
	}

	if err := json.NewDecoder(bytes.NewReader(data)).Decode(GameConf); err != nil {
		// if debugMode {
		// 	panic(err)
		// }
		return
	}
}

func (c *GameConfig) SaveConfigToFile() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "\t"); err != nil {
		return err
	}
	if err := ioutil.WriteFile(configFile, out.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (c *GameConfig) IsIDExist(majsoulAccountID int) bool {
	for _, id := range c.MajsoulAccountIDs {
		if id == majsoulAccountID {
			return true
		}
	}
	return false
}

func (c *GameConfig) AddMajsoulAccountID(majsoulAccountID int) error {
	if c.IsIDExist(majsoulAccountID) {
		return nil
	}
	GameConf.MajsoulAccountIDs = append(GameConf.MajsoulAccountIDs, majsoulAccountID)
	return c.SaveConfigToFile()
}

func (c *GameConfig) SetMajsoulAccountID(majsoulAccountID int) {
	c.CurrentActiveMajsoulAccountID = majsoulAccountID
}
