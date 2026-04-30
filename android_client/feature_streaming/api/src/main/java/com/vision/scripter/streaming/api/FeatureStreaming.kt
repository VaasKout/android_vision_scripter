package com.vision.scripter.streaming.api

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder

const val StreamingRoute = "streaming"
const val StreamingRouteWithArgs = "streaming/{serial}"
const val StreamingSerialArg = "serial"

interface FeatureStreaming {

    fun register(
        navGraphBuilder: NavGraphBuilder,
        navController: NavController,
    )
}