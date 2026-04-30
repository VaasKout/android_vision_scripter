package com.vision.scripter.scripts.api

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder

const val ScriptsRoute = "scripts"
const val ScriptsRouteWithArg = "scripts/{serial}"
const val SerialArg = "serial"

interface FeatureScripts {

    fun register(
        navGraphBuilder: NavGraphBuilder,
        navController: NavController,
    )
}