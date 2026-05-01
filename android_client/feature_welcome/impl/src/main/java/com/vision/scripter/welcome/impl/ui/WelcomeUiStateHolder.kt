package com.vision.scripter.welcome.impl.ui

import androidx.compose.runtime.Stable
import com.vision.scripter.ui.CommandFlow
import kotlinx.coroutines.flow.StateFlow

@Stable
interface WelcomeUiStateHolder {
    val uiStateFlow: StateFlow<WelcomeUiState?>
    val uiCommandsFlow: CommandFlow<WelcomeUiCommand>
    fun onInitData()

    fun onApplyData(url: String, port: String)
}
