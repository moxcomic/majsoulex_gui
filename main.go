package main

import (
	"os"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {
	initUI()
}

func initUI() {
	ReadNewConfig()
	gameConf := GameConf
	c := gameConf.MajsoulExInGamePrompt

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQWidget(nil, 0)
	window.SetFixedSize2(800, 600)
	window.SetWindowTitle("雀魂Ex功能可视化配置器")

	mainLayout := widgets.NewQGridLayout2()
	window.SetLayout(mainLayout)

	mainVLayout := widgets.NewQVBoxLayout()
	mainLayout.AddLayout(mainVLayout, 0, 0, 0)

	// ====
	serverHLayout := widgets.NewQHBoxLayout()

	serverLabel := widgets.NewQLabel2("游戏区服:", nil, 0)
	serverLabel.SetFixedSize2(100, 25)
	serverHLayout.Layout().AddWidget(serverLabel)
	serverComboBox := widgets.NewQComboBox(nil)
	serverComboBox.SetFixedSize2(200, 25)
	serverComboBox.AddItems([]string{"国服/国际服/港澳台", "日服"})
	if gameConf.MajsoulExServer > -1 && gameConf.MajsoulExServer < 2 {
		serverComboBox.SetCurrentIndex(gameConf.MajsoulExServer)
	}
	serverComboBox.ConnectCurrentIndexChanged(func(index int) {
		gameConf.MajsoulExServer = index
	})
	serverHLayout.Layout().AddWidget(serverComboBox)
	serverHLayout.AddStretch(0)
	serverHLayout.AddStretch(0)
	saveBtn := widgets.NewQPushButton2("保存配置", nil)
	saveBtn.ConnectClicked(func(bool) {
		gameConf.SaveConfigToFile()
		widgets.QMessageBox_Information(nil, "提示", "保存配置成功, 注意部分配置需重启生效", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	serverHLayout.Layout().AddWidget(saveBtn)
	serverHLayout.SetAlignment(saveBtn, core.Qt__AlignRight)
	serverHLayout.AddStretch(0)
	mainVLayout.AddLayout(serverHLayout, 0)

	bugqdyCheckBox := widgets.NewQCheckBox2("千万不要开启此功能", nil)
	if gameConf.MajsoulExBugQDY {
		bugqdyCheckBox.SetCheckState(core.Qt__Checked)
	}
	bugqdyCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExBugQDY = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExBugQDY = false
		}
	})
	mainVLayout.Layout().AddWidget(bugqdyCheckBox)

	// ====
	proxyGroupBox := widgets.NewQGroupBox2("代理设置", nil)
	proxyGroupBoxMainVLayout := widgets.NewQVBoxLayout()
	proxyGroupBoxMainVLayout.SetContentsMargins(0, 0, 0, 0)
	proxyGroupBoxMainVLayout.SetSpacing(0)
	proxyGroupBoxMainVLayout.Layout().SetSpacing(0)
	proxyGroupBox.SetLayout(proxyGroupBoxMainVLayout)
	mainVLayout.Layout().AddWidget(proxyGroupBox)

	// ====
	isUseProxyCheckBox := widgets.NewQCheckBox2("启用代理", nil)
	if gameConf.MajsoulExProxy.IsUseProxy {
		isUseProxyCheckBox.SetCheckState(core.Qt__Checked)
	}
	isUseProxyCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExProxy.IsUseProxy = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExProxy.IsUseProxy = false
		}
	})
	proxyGroupBoxMainVLayout.Layout().AddWidget(isUseProxyCheckBox)

	proxyTypeHLayout := widgets.NewQHBoxLayout()
	proxyTypeLabel := widgets.NewQLabel2("代理模式:", nil, 0)
	proxyTypeLabel.SetFixedSize2(100, 25)
	proxyTypeHLayout.Layout().AddWidget(proxyTypeLabel)
	proxyTypeComboBox := widgets.NewQComboBox(nil)
	proxyTypeComboBox.SetFixedSize2(100, 25)
	proxyTypeComboBox.AddItems([]string{"http", "https", "socks5"})
	proxyTypeComboBox.SetCurrentText(gameConf.MajsoulExProxy.ProxyType)
	proxyTypeComboBox.ConnectCurrentTextChanged(func(text string) {
		gameConf.MajsoulExProxy.ProxyType = text
	})
	proxyTypeHLayout.Layout().AddWidget(proxyTypeComboBox)
	proxyTypeHLayout.AddStretch(0)
	proxyTypeHLayout.AddStretch(0)
	proxyGroupBoxMainVLayout.AddLayout(proxyTypeHLayout, 0)

	proxyAddrHLayout := widgets.NewQHBoxLayout()
	proxyAddrLabel := widgets.NewQLabel2("代理地址:", nil, 0)
	proxyAddrLabel.SetFixedSize2(100, 25)
	proxyAddrHLayout.Layout().AddWidget(proxyAddrLabel)
	proxyAddrValue := widgets.NewQLineEdit(nil)
	proxyAddrValue.SetText(gameConf.MajsoulExProxy.ProxyAddr)
	proxyAddrValue.ConnectTextChanged(func(text string) {
		gameConf.MajsoulExProxy.ProxyAddr = text
	})
	proxyAddrHLayout.Layout().AddWidget(proxyAddrValue)
	proxyTypeHLayout.AddStretch(0)
	proxyTypeHLayout.AddStretch(0)
	proxyGroupBoxMainVLayout.AddLayout(proxyAddrHLayout, 0)

	proxyPortHLayout := widgets.NewQHBoxLayout()
	proxyPortLabel := widgets.NewQLabel2("代理端口:", nil, 0)
	proxyPortLabel.SetFixedSize2(100, 25)
	proxyPortHLayout.Layout().AddWidget(proxyPortLabel)
	proxyPortValue := widgets.NewQLineEdit(nil)
	proxyPortValue.SetText(strconv.Itoa(gameConf.MajsoulExProxy.ProxyPort))
	proxyPortValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseInt(text, 10, 64); e == nil {
			gameConf.MajsoulExProxy.ProxyPort = int(f)
		}
	})
	proxyPortHLayout.Layout().AddWidget(proxyPortValue)
	proxyTypeHLayout.AddStretch(0)
	proxyTypeHLayout.AddStretch(0)
	proxyGroupBoxMainVLayout.AddLayout(proxyPortHLayout, 0)

	// ====
	autoDiscardGroupBox := widgets.NewQGroupBox2("自动出牌", nil)
	autoDiscardGroupBoxMainVLayout := widgets.NewQVBoxLayout()
	autoDiscardGroupBoxMainVLayout.SetContentsMargins(0, 0, 0, 0)
	autoDiscardGroupBoxMainVLayout.SetSpacing(0)
	autoDiscardGroupBoxMainVLayout.Layout().SetSpacing(0)
	autoDiscardGroupBox.SetLayout(autoDiscardGroupBoxMainVLayout)
	mainVLayout.Layout().AddWidget(autoDiscardGroupBox)

	// ====

	autoDiscardCheckBox := widgets.NewQCheckBox2("自动出牌 (该功能不再支持段位战内使用)", nil)
	if gameConf.MajsoulExAutoDiscard {
		autoDiscardCheckBox.SetCheckState(core.Qt__Checked)
	}
	autoDiscardCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExAutoDiscard = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExAutoDiscard = false
		}
	})
	autoDiscardGroupBoxMainVLayout.Layout().AddWidget(autoDiscardCheckBox)

	autoConfirmCheckBox := widgets.NewQCheckBox2("自动确认 (开启该功能后会跳过对局结束结算页面, 不影响好感度、礼物等物品获取)", nil)
	if gameConf.MajsoulExAutoConfirm {
		autoConfirmCheckBox.SetCheckState(core.Qt__Checked)
	}
	autoConfirmCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExAutoConfirm = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExAutoConfirm = false
		}
	})
	autoDiscardGroupBoxMainVLayout.Layout().AddWidget(autoConfirmCheckBox)

	fallBackCheckBox := widgets.NewQCheckBox2("早巡退向 (开启该功能后在早巡时如果牌型较差会进行向听倒退)", nil)
	if gameConf.MajsoulExFallback {
		fallBackCheckBox.SetCheckState(core.Qt__Checked)
	}
	fallBackCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExFallback = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExFallback = false
		}
	})
	autoDiscardGroupBoxMainVLayout.Layout().AddWidget(fallBackCheckBox)

	improveCheckBox := widgets.NewQCheckBox2("立直改良 (开启该功能后在立直时如果听砍张将会尝试进行两轮的摸牌改良)", nil)
	if gameConf.MajsoulExImprove {
		improveCheckBox.SetCheckState(core.Qt__Checked)
	}
	improveCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExImprove = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExImprove = false
		}
	})
	autoDiscardGroupBoxMainVLayout.Layout().AddWidget(improveCheckBox)

	// chiitoiCheckBox1 := widgets.NewQCheckBox2("七对判断 (判断1)", nil)
	// if gameConf.MajsoulExChiitoi01 {
	// 	chiitoiCheckBox1.SetCheckState(core.Qt__Checked)
	// }
	// chiitoiCheckBox1.ConnectStateChanged(func(state int) {
	// 	switch state {
	// 	case int(core.Qt__Checked):
	// 		gameConf.MajsoulExChiitoi01 = true
	// 	case int(core.Qt__Unchecked):
	// 		gameConf.MajsoulExChiitoi01 = false
	// 	}
	// })
	chiitoiCheckBox2 := widgets.NewQCheckBox2("七对判断 (关闭后将不会优先做七对, 会出现拆刻子、对子等行为)", nil)
	if gameConf.MajsoulExChiitoi02 {
		chiitoiCheckBox2.SetCheckState(core.Qt__Checked)
	}
	chiitoiCheckBox2.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExChiitoi02 = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExChiitoi02 = false
		}
	})
	// mainVLayout.Layout().AddWidget(chiitoiCheckBox1)
	autoDiscardGroupBoxMainVLayout.Layout().AddWidget(chiitoiCheckBox2)

	// ====
	inGamePromptGroupBox := widgets.NewQGroupBox2("游戏内输出", nil)
	inGamePromptGroupBoxMainVLayout := widgets.NewQVBoxLayout()
	inGamePromptGroupBoxMainVLayout.SetContentsMargins(0, 0, 0, 0)
	inGamePromptGroupBoxMainVLayout.SetSpacing(0)
	inGamePromptGroupBoxMainVLayout.Layout().SetSpacing(0)
	inGamePromptGroupBox.SetLayout(inGamePromptGroupBoxMainVLayout)
	mainVLayout.Layout().AddWidget(inGamePromptGroupBox)

	// ====
	multipleInfoCheckBox := widgets.NewQCheckBox2("多行输出 (游戏内右下角点击感叹号后输出窗口)", nil)
	if gameConf.MajsoulExInGamePrompt.MultipleInfo {
		multipleInfoCheckBox.SetCheckState(core.Qt__Checked)
	}
	multipleInfoCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExInGamePrompt.MultipleInfo = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExInGamePrompt.MultipleInfo = false
		}
	})
	inGamePromptGroupBoxMainVLayout.Layout().AddWidget(multipleInfoCheckBox)

	// ====
	multipleInfoHLayout := widgets.NewQHBoxLayout()

	multipleInfoWidthLabel := widgets.NewQLabel2("宽度：", nil, 0)
	multipleInfoHLayout.Layout().AddWidget(multipleInfoWidthLabel)

	multipleInfoWidthValue := widgets.NewQLineEdit(nil)
	multipleInfoWidthValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MultipleInfoAttr.Width > -1 && c.MultipleInfoAttr.Width < 10000 {
		multipleInfoWidthValue.SetText(strconv.FormatFloat(c.MultipleInfoAttr.Width, 'g', -1, 64))
	}
	multipleInfoWidthValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MultipleInfoAttr.Width = f
			}
		}
	})
	multipleInfoHLayout.Layout().AddWidget(multipleInfoWidthValue)

	multipleInfoHeightLabel := widgets.NewQLabel2("高度：", nil, 0)
	multipleInfoHLayout.Layout().AddWidget(multipleInfoHeightLabel)

	multipleInfoHeightValue := widgets.NewQLineEdit(nil)
	multipleInfoHeightValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MultipleInfoAttr.Height > -1 && c.MultipleInfoAttr.Height < 10000 {
		multipleInfoHeightValue.SetText(strconv.FormatFloat(c.MultipleInfoAttr.Height, 'g', -1, 64))
	}
	multipleInfoHeightValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MultipleInfoAttr.Height = f
			}
		}
	})
	multipleInfoHLayout.Layout().AddWidget(multipleInfoHeightValue)

	multipleInfoXLabel := widgets.NewQLabel2("X：", nil, 0)
	multipleInfoHLayout.Layout().AddWidget(multipleInfoXLabel)

	multipleInfoXValue := widgets.NewQLineEdit(nil)
	multipleInfoXValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MultipleInfoAttr.X > -1 && c.MultipleInfoAttr.X < 10000 {
		multipleInfoXValue.SetText(strconv.FormatFloat(c.MultipleInfoAttr.X, 'g', -1, 64))
	}
	multipleInfoXValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MultipleInfoAttr.X = f
			}
		}
	})
	multipleInfoHLayout.Layout().AddWidget(multipleInfoXValue)

	multipleInfoYLabel := widgets.NewQLabel2("Y：", nil, 0)
	multipleInfoHLayout.Layout().AddWidget(multipleInfoYLabel)

	multipleInfoYValue := widgets.NewQLineEdit(nil)
	multipleInfoYValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MultipleInfoAttr.Y > -1 && c.MultipleInfoAttr.Y < 10000 {
		multipleInfoYValue.SetText(strconv.FormatFloat(c.MultipleInfoAttr.Y, 'g', -1, 64))
	}
	multipleInfoYValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MultipleInfoAttr.Y = f
			}
		}
	})
	multipleInfoHLayout.Layout().AddWidget(multipleInfoYValue)

	multipleInfoAlphaLabel := widgets.NewQLabel2("透明度：", nil, 0)
	multipleInfoHLayout.Layout().AddWidget(multipleInfoAlphaLabel)

	multipleInfoAlphaValue := widgets.NewQLineEdit(nil)
	multipleInfoAlphaValue.SetValidator(gui.NewQDoubleValidator2(0, 1, 1, nil))
	if c.MultipleInfoAttr.Alpha > -1 && c.MultipleInfoAttr.Alpha <= 1 {
		multipleInfoAlphaValue.SetText(strconv.FormatFloat(c.MultipleInfoAttr.Alpha, 'g', -1, 64))
	}
	multipleInfoAlphaValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f <= 1 {
				gameConf.MajsoulExInGamePrompt.MultipleInfoAttr.Alpha = f
			}
		}
	})
	multipleInfoHLayout.Layout().AddWidget(multipleInfoAlphaValue)

	multipleInfoFontSizeLabel := widgets.NewQLabel2("字体大小：", nil, 0)
	multipleInfoHLayout.Layout().AddWidget(multipleInfoFontSizeLabel)

	multipleInfoFontSizeValue := widgets.NewQLineEdit(nil)
	multipleInfoFontSizeValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MultipleInfoAttr.FontSize > -1 && c.MultipleInfoAttr.FontSize < 10000 {
		multipleInfoFontSizeValue.SetText(strconv.FormatFloat(c.MultipleInfoAttr.FontSize, 'g', -1, 64))
	}
	multipleInfoFontSizeValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MultipleInfoAttr.FontSize = f
			}
		}
	})
	multipleInfoHLayout.Layout().AddWidget(multipleInfoFontSizeValue)

	inGamePromptGroupBoxMainVLayout.AddLayout(multipleInfoHLayout, 0)

	// ====
	mouseEnterInfoCheckBox := widgets.NewQCheckBox2("单行输出 (游戏内鼠标移动到手牌上后的输出窗口)", nil)
	if gameConf.MajsoulExInGamePrompt.MouseEnterInfo {
		mouseEnterInfoCheckBox.SetCheckState(core.Qt__Checked)
	}
	mouseEnterInfoCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExInGamePrompt.MouseEnterInfo = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExInGamePrompt.MouseEnterInfo = false
		}
	})
	inGamePromptGroupBoxMainVLayout.Layout().AddWidget(mouseEnterInfoCheckBox)

	// ====
	mouseEnterInfoInfoHLayout := widgets.NewQHBoxLayout()

	mouseEnterInfoInfoWidthLabel := widgets.NewQLabel2("宽度：", nil, 0)
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoWidthLabel)

	mouseEnterInfoInfoWidthValue := widgets.NewQLineEdit(nil)
	mouseEnterInfoInfoWidthValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MouseEnterInfoAttr.Width > -1 && c.MouseEnterInfoAttr.Width < 10000 {
		mouseEnterInfoInfoWidthValue.SetText(strconv.FormatFloat(c.MouseEnterInfoAttr.Width, 'g', -1, 64))
	}
	mouseEnterInfoInfoWidthValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MouseEnterInfoAttr.Width = f
			}
		}
	})
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoWidthValue)

	mouseEnterInfoInfoHeightLabel := widgets.NewQLabel2("高度：", nil, 0)
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoHeightLabel)

	mouseEnterInfoInfoHeightValue := widgets.NewQLineEdit(nil)
	mouseEnterInfoInfoHeightValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MouseEnterInfoAttr.Height > -1 && c.MouseEnterInfoAttr.Height < 10000 {
		mouseEnterInfoInfoHeightValue.SetText(strconv.FormatFloat(c.MouseEnterInfoAttr.Height, 'g', -1, 64))
	}
	mouseEnterInfoInfoHeightValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MouseEnterInfoAttr.Height = f
			}
		}
	})
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoHeightValue)

	mouseEnterInfoInfoXLabel := widgets.NewQLabel2("X：", nil, 0)
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoXLabel)

	mouseEnterInfoInfoXValue := widgets.NewQLineEdit(nil)
	mouseEnterInfoInfoXValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MouseEnterInfoAttr.X > -1 && c.MouseEnterInfoAttr.X < 10000 {
		mouseEnterInfoInfoXValue.SetText(strconv.FormatFloat(c.MouseEnterInfoAttr.X, 'g', -1, 64))
	}
	mouseEnterInfoInfoXValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MouseEnterInfoAttr.X = f
			}
		}
	})
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoXValue)

	mouseEnterInfoInfoYLabel := widgets.NewQLabel2("Y：", nil, 0)
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoYLabel)

	mouseEnterInfoInfoYValue := widgets.NewQLineEdit(nil)
	mouseEnterInfoInfoYValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MouseEnterInfoAttr.Y > -1 && c.MouseEnterInfoAttr.Y < 10000 {
		mouseEnterInfoInfoYValue.SetText(strconv.FormatFloat(c.MouseEnterInfoAttr.Y, 'g', -1, 64))
	}
	mouseEnterInfoInfoYValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MouseEnterInfoAttr.Y = f
			}
		}
	})
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoYValue)

	mouseEnterInfoInfoAlphaLabel := widgets.NewQLabel2("透明度：", nil, 0)
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoAlphaLabel)

	mouseEnterInfoInfoAlphaValue := widgets.NewQLineEdit(nil)
	mouseEnterInfoInfoAlphaValue.SetValidator(gui.NewQDoubleValidator2(0, 1, 1, nil))
	if c.MouseEnterInfoAttr.Alpha > -1 && c.MouseEnterInfoAttr.Alpha <= 1 {
		mouseEnterInfoInfoAlphaValue.SetText(strconv.FormatFloat(c.MouseEnterInfoAttr.Alpha, 'g', -1, 64))
	}
	mouseEnterInfoInfoAlphaValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f <= 1 {
				gameConf.MajsoulExInGamePrompt.MouseEnterInfoAttr.Alpha = f
			}
		}
	})
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoAlphaValue)

	mouseEnterInfoInfoFontSizeLabel := widgets.NewQLabel2("字体大小：", nil, 0)
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoFontSizeLabel)

	mouseEnterInfoInfoFontSizeValue := widgets.NewQLineEdit(nil)
	mouseEnterInfoInfoFontSizeValue.SetValidator(gui.NewQDoubleValidator2(0, 9999, 2, nil))
	if c.MouseEnterInfoAttr.FontSize > -1 && c.MouseEnterInfoAttr.FontSize < 10000 {
		mouseEnterInfoInfoFontSizeValue.SetText(strconv.FormatFloat(c.MouseEnterInfoAttr.FontSize, 'g', -1, 64))
	}
	mouseEnterInfoInfoFontSizeValue.ConnectTextChanged(func(text string) {
		if f, e := strconv.ParseFloat(text, 64); e == nil {
			if f > -1 && f < 10000 {
				gameConf.MajsoulExInGamePrompt.MouseEnterInfoAttr.FontSize = f
			}
		}
	})
	mouseEnterInfoInfoHLayout.Layout().AddWidget(mouseEnterInfoInfoFontSizeValue)

	inGamePromptGroupBoxMainVLayout.AddLayout(mouseEnterInfoInfoHLayout, 0)
	// ====
	chongTileColorCheckBox := widgets.NewQCheckBox2("手牌染色 (游戏内根据计算的铳率对手牌进行染色, 染色时机为自己摸牌时、有鸣牌操作时)", nil)
	if gameConf.MajsoulExInGamePrompt.ChongTileColor {
		chongTileColorCheckBox.SetCheckState(core.Qt__Checked)
	}
	chongTileColorCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExInGamePrompt.ChongTileColor = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExInGamePrompt.ChongTileColor = false
		}
	})
	inGamePromptGroupBoxMainVLayout.Layout().AddWidget(chongTileColorCheckBox)

	// ====
	chongTileColorAttrHLayout := widgets.NewQHBoxLayout()
	// ====
	chongTileColorNormalAttrBtn := widgets.NewQPushButton2("设置0%铳率颜色", nil)
	chongTileColorNormalAttrBtn.ConnectClicked(func(bool) {
		n := gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.Normal
		color := widgets.QColorDialog_GetColor(gui.QColor_FromRgbF(n.R, n.G, n.B, n.A), window, "选择颜色", widgets.QColorDialog__ShowAlphaChannel)
		if color.IsValid() {
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.Normal.R = color.RedF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.Normal.G = color.GreenF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.Normal.B = color.BlueF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.Normal.A = color.AlphaF()
		}
	})
	chongTileColorAttrHLayout.Layout().AddWidget(chongTileColorNormalAttrBtn)
	// ====
	chongTileColor5AttrBtn := widgets.NewQPushButton2("设置低于5%铳率颜色", nil)
	chongTileColor5AttrBtn.ConnectClicked(func(bool) {
		n := gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan5Percent
		color := widgets.QColorDialog_GetColor(gui.QColor_FromRgbF(n.R, n.G, n.B, n.A), window, "选择颜色", widgets.QColorDialog__ShowAlphaChannel)
		if color.IsValid() {
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan5Percent.R = color.RedF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan5Percent.G = color.GreenF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan5Percent.B = color.BlueF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan5Percent.A = color.AlphaF()
		}
	})
	chongTileColorAttrHLayout.Layout().AddWidget(chongTileColor5AttrBtn)
	// ====
	chongTileColor10AttrBtn := widgets.NewQPushButton2("设置低于10%铳率颜色", nil)
	chongTileColor10AttrBtn.ConnectClicked(func(bool) {
		n := gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan10Percent
		color := widgets.QColorDialog_GetColor(gui.QColor_FromRgbF(n.R, n.G, n.B, n.A), window, "选择颜色", widgets.QColorDialog__ShowAlphaChannel)
		if color.IsValid() {
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan10Percent.R = color.RedF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan10Percent.G = color.GreenF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan10Percent.B = color.BlueF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan10Percent.A = color.AlphaF()
		}
	})
	chongTileColorAttrHLayout.Layout().AddWidget(chongTileColor10AttrBtn)
	// ====
	chongTileColor15AttrBtn := widgets.NewQPushButton2("设置低于15%铳率颜色", nil)
	chongTileColor15AttrBtn.ConnectClicked(func(bool) {
		n := gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan15Percent
		color := widgets.QColorDialog_GetColor(gui.QColor_FromRgbF(n.R, n.G, n.B, n.A), window, "选择颜色", widgets.QColorDialog__ShowAlphaChannel)
		if color.IsValid() {
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan15Percent.R = color.RedF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan15Percent.G = color.GreenF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan15Percent.B = color.BlueF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.LessThan15Percent.A = color.AlphaF()
		}
	})
	chongTileColorAttrHLayout.Layout().AddWidget(chongTileColor15AttrBtn)
	// ====
	chongTileColor15GAttrBtn := widgets.NewQPushButton2("设置高于15%铳率颜色", nil)
	chongTileColor15GAttrBtn.ConnectClicked(func(bool) {
		n := gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.GreaterThan15Percent
		color := widgets.QColorDialog_GetColor(gui.QColor_FromRgbF(n.R, n.G, n.B, n.A), window, "选择颜色", widgets.QColorDialog__ShowAlphaChannel)
		if color.IsValid() {
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.GreaterThan15Percent.R = color.RedF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.GreaterThan15Percent.G = color.GreenF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.GreaterThan15Percent.B = color.BlueF()
			gameConf.MajsoulExInGamePrompt.ChongTileColorAttr.GreaterThan15Percent.A = color.AlphaF()
		}
	})
	chongTileColorAttrHLayout.Layout().AddWidget(chongTileColor15GAttrBtn)
	// ====
	inGamePromptGroupBoxMainVLayout.AddLayout(chongTileColorAttrHLayout, 0)
	// ====
	miqieColorCheckBox := widgets.NewQCheckBox2("摸切染色 (游戏内对牌河摸切进行染色)", nil)
	if gameConf.MajsoulExInGamePrompt.MoqieColor {
		miqieColorCheckBox.SetCheckState(core.Qt__Checked)
	}
	miqieColorCheckBox.ConnectStateChanged(func(state int) {
		switch state {
		case int(core.Qt__Checked):
			gameConf.MajsoulExInGamePrompt.MoqieColor = true
		case int(core.Qt__Unchecked):
			gameConf.MajsoulExInGamePrompt.MoqieColor = false
		}
	})
	inGamePromptGroupBoxMainVLayout.Layout().AddWidget(miqieColorCheckBox)

	tipLabel := widgets.NewQLabel2("部分功能修改后重启软件才会生效, 最低支持版本0.0.15", nil, 0)
	tipLabel.SetStyleSheet(`color:red`)
	mainVLayout.Layout().AddWidget(tipLabel)
	mainVLayout.SetAlignment(tipLabel, core.Qt__AlignCenter)

	window.Show()
	app.Exec()
}
