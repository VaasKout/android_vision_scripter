package com.vision.scripter.welcome.impl.ui

import androidx.compose.runtime.Immutable

@Immutable
data class WelcomeUiState(
    val oldUrl: String = "",
    val oldPort: String = "",
)