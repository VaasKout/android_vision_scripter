package com.vision.scripter.welcome.impl.commandobservers

import androidx.compose.material3.SnackbarHostState
import androidx.compose.runtime.Composable
import androidx.compose.ui.res.stringResource
import androidx.navigation.NavController
import com.vision.scripter.main.api.MainRoute
import com.vision.scripter.ui.observe
import com.vision.scripter.welcome.api.WelcomeRoute
import com.vision.scripter.welcome.impl.R
import com.vision.scripter.welcome.impl.ui.WelcomeUiCommand
import com.vision.scripter.welcome.impl.ui.WelcomeUiStateHolder

@Composable
internal fun WelcomeUiCommandObserver(
    uiStateHolder: WelcomeUiStateHolder,
    navController: NavController,
    snackbarHostState: SnackbarHostState,
) {
    val addressErrorMsg = stringResource(R.string.address_error)

    uiStateHolder.uiCommandsFlow.observe {
        when (it) {
            is WelcomeUiCommand.NavigateToMain -> navController.navigate(MainRoute) {
                popUpTo(WelcomeRoute) {
                    inclusive = true
                }
            }

            is WelcomeUiCommand.ShowAddressError -> {
                snackbarHostState.showSnackbar(message = addressErrorMsg)
            }
        }
    }
}