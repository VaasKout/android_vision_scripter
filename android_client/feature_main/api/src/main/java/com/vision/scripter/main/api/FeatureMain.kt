package com.vision.scripter.main.api

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder

const val MainRoute = "main"

interface FeatureMain {

    fun register(
        navGraphBuilder: NavGraphBuilder,
        navController: NavController,
    )
}