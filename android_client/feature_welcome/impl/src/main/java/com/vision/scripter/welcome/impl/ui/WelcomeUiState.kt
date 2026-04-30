package com.vision.scripter.welcome.impl.ui

import androidx.compose.runtime.Immutable

@Immutable
data class WelcomeUiState(
    val url: String = "",
    val port: String = "",
)