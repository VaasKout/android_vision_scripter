package com.vision.scripter.coroutines.api

import kotlinx.coroutines.CoroutineScope

interface CoroutineScopeFactory {
    fun createForegroundScope(tag: String): CoroutineScope
    fun createBackgroundScope(tag: String): CoroutineScope
    fun createIoScope(tag: String): CoroutineScope
}