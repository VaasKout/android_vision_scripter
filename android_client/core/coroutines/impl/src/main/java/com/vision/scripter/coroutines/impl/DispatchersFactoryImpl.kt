package com.vision.scripter.coroutines.impl

import com.vision.scripter.coroutines.api.DispatchersFactory
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.coroutines.Dispatchers
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class DispatchersFactoryImpl @Inject constructor() : DispatchersFactory {
    override val main: CoroutineDispatcher = Dispatchers.Main
    override val immediate: CoroutineDispatcher = Dispatchers.Main.immediate
    override val default: CoroutineDispatcher = Dispatchers.Default
    override val io: CoroutineDispatcher = Dispatchers.IO
    override val unconfined: CoroutineDispatcher = Dispatchers.Unconfined
}