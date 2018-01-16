

go run reset_theme.go

GenericChatroomWidget::reloadTheme(): return

清除mainwindow.ui中的palette设置，styleSheet值

widget.cpp:227 theme调用

去掉：reloadTheme(), applyTheme(),
去掉：->setStyleSheet(), SET_STYLESHEET()

去掉：genericchatroomwidget.cpp: setBackgroundRole(), setForegroundRole(), setPalette()

图标：

tooltips:

