package com.vision.scripter.main.impl.ui

import com.vision.scripter.main.impl.R
import com.vision.scripter.ui.CommandFlow
import kotlinx.collections.immutable.persistentListOf
import kotlinx.collections.immutable.persistentMapOf
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.SharedFlow

internal val mainUiStateLoadingPreview = MainUiState()
internal val mainUiStateNoDataPreview = MainUiState(isLoading = false)
internal val mainUiStateWithDataPreview = MainUiState(
    isLoading = false,
    devices = persistentListOf(
        UiDevice(
            serial = "c88893qa",
            deviceParams = persistentMapOf(
                R.string.serial_key to "c9858321",
                R.string.model_key to "Samsung",
            ),
        ),
    )
)

internal class MainScreenUiStateHolderPreview(state: MainUiState) : MainUiStateHolder {

    override val uiStateFlow: SharedFlow<MainUiState> = MutableStateFlow(state)
    override val uiCommandsFlow: CommandFlow<MainUiCommand>
        get() = throw UnsupportedOperationException()

    override fun onLoadData(onStart: Boolean) {}

}