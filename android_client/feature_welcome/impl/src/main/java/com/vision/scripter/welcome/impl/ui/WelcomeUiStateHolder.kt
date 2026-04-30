package com.vision.scripter.welcome.impl.ui

import androidx.compose.runtime.Stable
import com.vision.scripter.ui.CommandFlow
import kotlinx.coroutines.flow.SharedFlow

@Stable
interface WelcomeUiStateHolder {
    val uiStateFlow: SharedFlow<WelcomeUiState>
    val uiCommandsFlow: CommandFlow<WelcomeUiCommand>
    fun onInitData()

    fun editUrl(url: String)
    fun editPort(port: String)
    fun onApplyData()
}
