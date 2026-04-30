package com.vision.scripter.main.impl.ui

import androidx.compose.runtime.Stable
import com.vision.scripter.ui.CommandFlow
import kotlinx.coroutines.flow.SharedFlow

@Stable
interface MainUiStateHolder {
    val uiStateFlow: SharedFlow<MainUiState>
    val uiCommandsFlow: CommandFlow<MainUiCommand>

    fun onLoadData(onStart: Boolean)
}