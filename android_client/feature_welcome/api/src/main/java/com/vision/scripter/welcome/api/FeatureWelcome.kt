package com.vision.scripter.welcome.api

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder

const val WelcomeRoute = "welcome"

interface FeatureWelcome {

    fun register(
        navGraphBuilder: NavGraphBuilder,
        navController: NavController,
    )
}