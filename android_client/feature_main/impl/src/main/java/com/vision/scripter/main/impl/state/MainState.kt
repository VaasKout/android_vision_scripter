package com.vision.scripter.main.impl.state

import com.vision.scripter.data.api.models.AdbDevice
import com.vision.scripter.ui.states.LoadingState

data class MainState(
    val loadingState: LoadingState = LoadingState.LoadingOnStart,
    val devices: List<AdbDevice> = listOf()
)
