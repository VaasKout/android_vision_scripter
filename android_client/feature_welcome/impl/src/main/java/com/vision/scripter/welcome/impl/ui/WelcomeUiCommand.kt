package com.vision.scripter.welcome.impl.ui

sealed class WelcomeUiCommand {
    object NavigateToMain : WelcomeUiCommand()
    object ShowAddressError : WelcomeUiCommand()
}