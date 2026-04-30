package com.vision.scripter.welcome.impl.state

import com.vision.scripter.welcome.impl.ui.WelcomeUiState
import dagger.hilt.android.scopes.ViewModelScoped
import javax.inject.Inject

@ViewModelScoped
class WelcomeStateUiMapper @Inject constructor() {
    fun map(state: WelcomeState): WelcomeUiState {
        return WelcomeUiState(
            url = state.url,
            port = state.port,
        )
    }
}