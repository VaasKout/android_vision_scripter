package com.vision.scripter.main.impl.state

import com.vision.scripter.data.api.models.AdbDevice
import com.vision.scripter.main.impl.R
import com.vision.scripter.main.impl.ui.MainUiState
import com.vision.scripter.main.impl.ui.UiDevice
import com.vision.scripter.ui.states.LoadingState
import dagger.hilt.android.scopes.ViewModelScoped
import kotlinx.collections.immutable.ImmutableMap
import kotlinx.collections.immutable.toImmutableList
import kotlinx.collections.immutable.toImmutableMap
import javax.inject.Inject

@ViewModelScoped
class MainUiStateMapper @Inject constructor() {
    fun map(state: MainState): MainUiState {
        return MainUiState(
            isLoading = state.loadingState == LoadingState.LoadingOnStart,
            isRefreshing = state.loadingState == LoadingState.RefreshLoading,
            devices = state.devices.map {
                UiDevice(
                    serial = it.serial,
                    deviceParams = it.toMap(),
                    preview = null,
                )
            }.toImmutableList(),
        )
    }
}

fun AdbDevice.toMap(): ImmutableMap<Int, String> {
    return buildMap {
        if (serial != "") put(
            R.string.serial_key, serial.replace(" ", "\u00A0")
        )
        if (brand != "") put(
            R.string.brand_key, brand.replace(" ", "\u00A0")
        )
        if (device != "") put(
            R.string.device_key, device.replace(" ", "\u00A0")
        )
        if (model != "") put(
            R.string.model_key, model.replace(" ", "\u00A0")
        )
        if (osVersion != "") put(
            R.string.os_version_key, osVersion.replace(" ", "\u00A0")
        )
    }.toImmutableMap()
}