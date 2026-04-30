package com.vision.scripter.main.impl.ui

sealed class MainUiCommand {
    object ShowNetworkError : MainUiCommand()
    data class NavigateToScripts(val serial: String) : MainUiCommand()
    data class NavigateToStreaming(val serial: String) : MainUiCommand()
}