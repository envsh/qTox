
### 清除现在css中的color行
go run reset_theme.go

### 代码上的theme相关清理改动
GenericChatroomWidget::reloadTheme(): return

清除mainwindow.ui中的palette设置，styleSheet值

widget.cpp:227 theme调用

去掉：reloadTheme(), applyTheme(),
去掉：->setStyleSheet(), SET_STYLESHEET()

去掉：genericchatroomwidget.cpp: setBackgroundRole(), setForegroundRole(), setPalette()

图标：

tooltips:

### 抽象提取的颜色值：
* background-color
* foreground-color 
* hover-background-color
* selected-background-color
* scollbar-background-color
* scollbar-foreground-color
* self-icon-area-background-color
* avatar-background-color

