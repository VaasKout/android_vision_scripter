package com.vision.scripter.welcome.impl.state

import androidx.core.net.toUri
import com.vision.scripter.coroutines.api.CoroutineScopeFactory
import com.vision.scripter.data.api.ScripterRepository
import com.vision.scripter.prefs.api.DataStoreRepository
import com.vision.scripter.ui.CommandFlow
import com.vision.scripter.welcome.impl.ui.WelcomeUiCommand
import com.vision.scripter.welcome.impl.ui.WelcomeUiState
import com.vision.scripter.welcome.impl.ui.WelcomeUiStateHolder
import dagger.hilt.android.scopes.ViewModelScoped
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.SharedFlow
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.shareIn
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import javax.inject.Inject

@ViewModelScoped
class WelcomeInteractor @Inject constructor(
    coroutineScopeFactory: CoroutineScopeFactory,
    private val uiStateMapper: WelcomeStateUiMapper,
    private val dataStoreRepository: DataStoreRepository,
    private val scripterRepository: ScripterRepository,
) : WelcomeUiStateHolder {

    private val _stateFlow = MutableStateFlow(WelcomeState())
    private val stateFlow: StateFlow<WelcomeState> = _stateFlow.asStateFlow()
    private val coroutineScope: CoroutineScope =
        coroutineScopeFactory.createBackgroundScope("welcome_interactor")

    private val currentState: WelcomeState
        get() = _stateFlow.value

    override val uiStateFlow: SharedFlow<WelcomeUiState>
        get() = stateFlow.map(uiStateMapper::map)
            .shareIn(coroutineScope, SharingStarted.WhileSubscribed(), replay = 1)

    override val uiCommandsFlow: CommandFlow<WelcomeUiCommand> = CommandFlow(coroutineScope)

    override fun onInitData() {
        coroutineScope.launch {
            val fullUrl = dataStoreRepository.getServerUrl()
            if (fullUrl.isEmpty()) return@launch

            val serverAvailable = scripterRepository.pingServer()
            if (!serverAvailable) {
                val uri = fullUrl.toUri()
                val port = uri.port
                val url = "${uri.scheme}://${uri.host}"
                _stateFlow.update {
                    it.copy(
                        url = url,
                        port = if (port <= 0) "" else port.toString(),
                    )
                }
                return@launch
            }
            uiCommandsFlow.tryEmit(WelcomeUiCommand.NavigateToMain)
        }
    }

    override fun editUrl(url: String) {
        _stateFlow.update { it.copy(url = url) }
    }

    override fun editPort(port: String) {
        if (port.length > 5) return
        _stateFlow.update { it.copy(port = port) }
    }

    override fun onApplyData() {
        coroutineScope.launch {
            if (currentState.url.isEmpty()) {
                uiCommandsFlow.tryEmit(WelcomeUiCommand.ShowAddressError)
                return@launch
            }
            val uri = currentState.url.toUri()
            if (uri.scheme == null || uri.host == null) {
                uiCommandsFlow.tryEmit(WelcomeUiCommand.ShowAddressError)
                return@launch
            }

            val fullUrl = buildString {
                append(uri.scheme)
                append("://")
                append(uri.host)
                if (currentState.port != "") append(":${currentState.port}")
            }

            dataStoreRepository.setServerUrl(fullUrl)
            val serverAvailable = scripterRepository.pingServer()
            if (!serverAvailable) {
                uiCommandsFlow.tryEmit(WelcomeUiCommand.ShowAddressError)
                return@launch
            }

            uiCommandsFlow.tryEmit(WelcomeUiCommand.NavigateToMain)
        }
    }
}
