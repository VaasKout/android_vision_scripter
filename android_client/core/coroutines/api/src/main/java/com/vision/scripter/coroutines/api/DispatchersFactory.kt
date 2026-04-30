package com.vision.scripter.coroutines.api

import kotlinx.coroutines.CoroutineDispatcher

interface DispatchersFactory {

    val main: CoroutineDispatcher

    val immediate: CoroutineDispatcher

    val default: CoroutineDispatcher

    val io: CoroutineDispatcher

    val unconfined: CoroutineDispatcher
}