package com.vision.scripter.welcome.impl

import androidx.compose.material3.SnackbarHostState
import androidx.compose.runtime.remember
import androidx.hilt.lifecycle.viewmodel.compose.hiltViewModel
import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder
import androidx.navigation.compose.composable
import com.vision.scripter.welcome.api.FeatureWelcome
import com.vision.scripter.welcome.api.WelcomeRoute
import com.vision.scripter.welcome.impl.commandobservers.WelcomeUiCommandObserver
import com.vision.scripter.welcome.impl.state.WelcomeViewModel
import com.vision.scripter.welcome.impl.ui.WelcomeScreen
import dagger.hilt.android.scopes.ActivityScoped
import javax.inject.Inject

@ActivityScoped
class FeatureWelcomeImpl @Inject constructor() : FeatureWelcome {

    override fun register(
        navGraphBuilder: NavGraphBuilder,
        navController: NavController,
    ) {
        navGraphBuilder.composable(route = WelcomeRoute) {
            val snackbarHostState = remember { SnackbarHostState() }
            val welcomeViewModel = hiltViewModel<WelcomeViewModel>()

            WelcomeUiCommandObserver(
                uiStateHolder = welcomeViewModel,
                navController = navController,
                snackbarHostState = snackbarHostState,
            )

            WelcomeScreen(
                uiStateHolder = welcomeViewModel,
                snackbarHostState = snackbarHostState,
            )
        }
    }
}
