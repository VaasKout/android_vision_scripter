package com.vision.scripter.ui

import androidx.compose.runtime.Stable
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.Job
import kotlinx.coroutines.channels.Channel
import kotlinx.coroutines.flow.AbstractFlow
import kotlinx.coroutines.flow.FlowCollector
import kotlinx.coroutines.flow.emitAll
import kotlinx.coroutines.flow.receiveAsFlow

@Stable
@OptIn(ExperimentalCoroutinesApi::class)
class CommandFlow<T>(scope: CoroutineScope) : AbstractFlow<T>() {

    private val channel = Channel<T>(Channel.UNLIMITED)

    init {
        scope.coroutineContext[Job]?.invokeOnCompletion {
            channel.cancel()
        }
    }

    fun tryEmit(value: T): Boolean {
        return channel.trySend(value).isSuccess
    }

    override suspend fun collectSafely(collector: FlowCollector<T>) {
        collector.emitAll(channel.receiveAsFlow())
    }
}

infix fun <T> CommandFlow<T>.emit(value: T) {
    this.tryEmit(value)
}