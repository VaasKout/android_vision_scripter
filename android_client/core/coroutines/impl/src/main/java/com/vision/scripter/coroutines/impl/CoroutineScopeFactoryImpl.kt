package com.vision.scripter.coroutines.impl

import com.vision.scripter.coroutines.api.CoroutineScopeFactory
import com.vision.scripter.coroutines.api.DispatchersFactory
import kotlinx.coroutines.CoroutineName
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.SupervisorJob
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class CoroutineScopeFactoryImpl @Inject constructor(
    private val dispatchersFactory: DispatchersFactory,
) : CoroutineScopeFactory {
    override fun createForegroundScope(tag: String) = CoroutineScope(
        dispatchersFactory.immediate + SupervisorJob() + CoroutineName(tag)
    )

    override fun createBackgroundScope(tag: String) = CoroutineScope(
        dispatchersFactory.default + SupervisorJob() + CoroutineName(tag)
    )

    override fun createIoScope(tag: String): CoroutineScope = CoroutineScope(
        dispatchersFactory.io + SupervisorJob() + CoroutineName(tag)
    )
}