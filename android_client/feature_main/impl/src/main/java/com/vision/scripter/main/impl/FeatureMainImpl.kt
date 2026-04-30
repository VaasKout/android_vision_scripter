package com.vision.scripter.main.impl

import androidx.compose.material3.SnackbarHostState
import androidx.compose.runtime.remember
import androidx.hilt.lifecycle.viewmodel.compose.hiltViewModel
import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder
import androidx.navigation.compose.composable
import com.vision.scripter.main.api.FeatureMain
import com.vision.scripter.main.api.MainRoute
import com.vision.scripter.main.impl.commandobservers.MainUiCommandObserver
import com.vision.scripter.main.impl.state.MainViewModel
import com.vision.scripter.main.impl.ui.MainUiScreen
import dagger.hilt.android.scopes.ActivityScoped
import javax.inject.Inject

@ActivityScoped
class FeatureMainImpl @Inject constructor() : FeatureMain {
    override fun register(
        navGraphBuilder: NavGraphBuilder,
        navController: NavController,
    ) {
        navGraphBuilder.composable(route = MainRoute) {
            val snackbarHostState = remember { SnackbarHostState() }
            val mainViewModel = hiltViewModel<MainViewModel>()
            MainUiCommandObserver(
                uiStateHolder = mainViewModel,
                navController = navController,
                snackbarHostState = snackbarHostState,
            )

            MainUiScreen(
                uiStateHolder = mainViewModel,
                snackbarHostState = snackbarHostState,
            )
        }
    }
}